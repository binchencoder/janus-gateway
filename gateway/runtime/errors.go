package runtime

import (
	"context"
	"io"
	"net/http"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"

	fpb "github.com/binchencoder/gateway-proto/frontend"
)

// HTTPStatusFromCode converts a gRPC error code into the corresponding HTTP response status.
// See: https://github.com/googleapis/googleapis/blob/master/google/rpc/code.proto
func HTTPStatusFromCode(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		// Note, this deliberately doesn't translate to the similarly named '412 Precondition Failed' HTTP response status.
		return http.StatusBadRequest
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusBadRequest
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	}

	grpclog.Infof("Unknown gRPC error code: %v", code)
	return http.StatusInternalServerError
}

var (
	// HTTPError replies to the request with an error.
	//
	// HTTPError is called:
	//  - From generated per-endpoint gateway handler code, when calling the backend results in an error.
	//  - From gateway runtime code, when forwarding the response message results in an error.
	//
	// The default value for HTTPError calls the custom error handler configured on the ServeMux via the
	// WithProtoErrorHandler serve option if that option was used, calling GlobalHTTPErrorHandler otherwise.
	//
	// To customize the error handling of a particular ServeMux instance, use the WithProtoErrorHandler
	// serve option.
	//
	// To customize the error format for all ServeMux instances not using the WithProtoErrorHandler serve
	// option, set GlobalHTTPErrorHandler to a custom function.
	//
	// Setting this variable directly to customize error format is deprecated.
	HTTPError = DefaultHTTPError

	// GlobalHTTPErrorHandler is the HTTPError handler for all ServeMux instances not using the
	// WithProtoErrorHandler serve option.
	//
	// You can set a custom function to this variable to customize error format.
	GlobalHTTPErrorHandler = DefaultHTTPError

	// OtherErrorHandler handles gateway errors from parsing and routing client requests for all
	// ServeMux instances not using the WithProtoErrorHandler serve option.
	//
	// It returns the following error codes: StatusMethodNotAllowed StatusNotFound StatusBadRequest
	//
	// To customize parsing and routing error handling of a particular ServeMux instance, use the
	// WithProtoErrorHandler serve option.
	//
	// To customize parsing and routing error handling of all ServeMux instances not using the
	// WithProtoErrorHandler serve option, set a custom function to this variable.
	OtherErrorHandler = DefaultOtherErrorHandler
)

// MuxOrGlobalHTTPError uses the mux-configured error handler, falling back to GlobalErrorHandler.
func MuxOrGlobalHTTPError(ctx context.Context, mux *ServeMux, marshaler Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	if mux.protoErrorHandler != nil {
		mux.protoErrorHandler(ctx, mux, marshaler, w, r, err)
	} else {
		GlobalHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
	}
}

type errorBody struct {
	Error *fpb.Error `protobuf:"bytes,1,name=error" json:"error"`
	// This is to make the error more compatible with users that expect errors to be Status objects:
	// https://github.com/grpc/grpc/blob/master/src/proto/grpc/status/status.proto
	// It should be the exact same message as the Error field.
	Code    int32      `protobuf:"varint,1,name=code" json:"code"`
	Message string     `protobuf:"bytes,2,name=message" json:"message"`
	Details []*any.Any `protobuf:"bytes,3,rep,name=details" json:"details,omitempty"`
}

// Make this also conform to proto.Message for builtin JSONPb Marshaler
func (e *errorBody) Reset()         { *e = errorBody{} }
func (e *errorBody) String() string { return proto.CompactTextString(e) }
func (*errorBody) ProtoMessage()    {}

// DefaultHTTPError is the default implementation of HTTPError.
// If "err" is an error from gRPC system, the function replies with the status code mapped by HTTPStatusFromCode.
// If otherwise, it replies with http.StatusInternalServerError.
//
// The response body returned by this function is a JSON object,
// which contains a member whose key is "error" and whose value is err.Error().
func DefaultHTTPError(ctx context.Context, mux *ServeMux, marshaler Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	w.Header().Del("Trailer")
	w.Header().Del("Transfer-Encoding")

	contentType := marshaler.ContentType()
	// Check marshaler on run time in order to keep backwards compatibility
	// An interface param needs to be added to the ContentType() function on
	// the Marshal interface to be able to remove this check
	if typeMarshaler, ok := marshaler.(contentTypeMarshaler); ok {
		pb := s.Proto()
		contentType = typeMarshaler.ContentTypeFromMessage(pb)
	}
	w.Header().Set("Content-Type", contentType)

	e := fpb.Error{}
	desc := grpc.ErrorDesc(err)
	if erru := jsonpb.UnmarshalString(desc, &e); erru != nil {
		e.Code = fpb.ErrorCode_UNDEFINED
		e.Params = []string{desc}
	}
	body := &errorBody{
		Error:   &e,
		Message: s.Message(),
		Code:    int32(s.Code()),
		Details: s.Proto().GetDetails(),
	}

	buf, merr := marshaler.Marshal(body)
	if merr != nil {
		grpclog.Infof("Failed to marshal error message %q: %v", body, merr)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallback); err != nil {
			grpclog.Infof("Failed to write response: %v", err)
		}
		return
	}

	md, ok := ServerMetadataFromContext(ctx)
	if !ok {
		grpclog.Infof("Failed to extract ServerMetadata from context")
	}

	handleForwardResponseServerMetadata(w, mux, md)

	// RFC 7230 https://tools.ietf.org/html/rfc7230#section-4.1.2
	// Unless the request includes a TE header field indicating "trailers"
	// is acceptable, as described in Section 4.3, a server SHOULD NOT
	// generate trailer fields that it believes are necessary for the user
	// agent to receive.
	var wantsTrailers bool

	if te := r.Header.Get("TE"); strings.Contains(strings.ToLower(te), "trailers") {
		wantsTrailers = true
		handleForwardResponseTrailerHeader(w, md)
		w.Header().Set("Transfer-Encoding", "chunked")
	}

	st := HTTPStatusFromCode(s.Code())
	w.WriteHeader(st)
	if _, err := w.Write(buf); err != nil {
		grpclog.Infof("Failed to write response: %v", err)
	}

	if wantsTrailers {
		handleForwardResponseTrailer(w, md)
	}
}

// DefaultOtherErrorHandler is the default implementation of OtherErrorHandler.
// It simply writes a string representation of the given error into "w".
func DefaultOtherErrorHandler(w http.ResponseWriter, _ *http.Request, msg string, code int) {
	http.Error(w, msg, code)
}
