package runtime

import (
	"net/http"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"

	pb "github.com/binchencoder/skylb-api/proto"
	ggr "github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// GatewayServiceHook collects the injection points with which gateway runtime
// calls back on different stages of request processing. It allows us to
// intercept the flow and inject our business logic, such as authentication,
// etc.
//
// At runtime we only need one instance of GatewayServiceHook since it's
// designed as goroutine-safe.
type GatewayServiceHook interface {
	// Bootstrap is a callback function which will be called when
	// SetGatewayServiceHook() is executed. The call is guaranteed to be
	// after all gateway services are registered and before any request
	// arrives.
	//
	// The passed-in service map contains all the service information known
	// by the gateway at compile time. The interface implementation can
	// safely hold this map because it will not be changed as soon as
	// Bootstrap() is called.
	//
	// The typical jobs which can be done in Bootstrap() include
	// initialization or starting service maintenance workers.
	Bootstrap(svcs map[string]*ServiceGroup) error

	// RequestReceived is called after the request arrives at the gateway
	// but before the routing decision is made.
	// Parameters:
	//     w: the raw HTTP response writer of current request
	//     r: the raw HTTP request
	//
	// Returns:
	//     err: the error returned to HTTP handler; when it's non-nil, the
	// request fails with an internal server error (500)
	RequestReceived(w http.ResponseWriter, r *http.Request) (ctx context.Context, err error)

	// RequestAccepted is called when a request is accepted and routed at
	// the gateway but before the protocol buffer is parsed.
	//
	// Parameters:
	//     svc: the service object to which the gateway routes
	//     m: the method object to which the gateway routes
	//     w:   the raw HTTP response writer of current request
	//     r:    the raw HTTP request
	// Returns:
	//     ctxVals: the values to be put into context and passed along the
	//              chain
	//     err:     the error returned to HTTP handler; when it's non-nil,
	//              the request fails with an internal server error (500)
	RequestAccepted(ctx context.Context, svc *Service, m *Method, w http.ResponseWriter, r *http.Request) (ctxret context.Context, err error)

	// RequestParsed is called after the request payload is unmarshaled and
	// before the gRPC call is invoked.
	//
	// Parameters:
	//    ctx:      the context
	//    svc:      the service object to which the gateway routes
	//    m: the method object to which the gateway routes
	//    reqProto: the request proto message
	//    meta:     the server meta data
	RequestParsed(ctx context.Context, svc *Service, m *Method, reqProto proto.Message, meta *ggr.ServerMetadata) error

	// RequestHandled is called after a request is completely handled, either
	// succeeded or failed.
	//
	// Parameters:
	//    ctx:  is the context.
	//    svc:  is the service object under current context.
	//    m: the method object to which the gateway routes
	//    out:  is the response message from grpc server.
	//    meta: the meta data.
	//    err:  is the err which returned from grpc server
	RequestHandled(ctx context.Context, svc *Service, m *Method, responseProto proto.Message, meta *ggr.ServerMetadata, err error)
}

var (
	hook           GatewayServiceHook
	defaultContext = context.Background()
)

// SetGatewayServiceHook sets a GatewayServiceHook. It should be called exactly
// once, after all init() functions are called (so that all gateway handlers
// are properly registered). That said, do not call it in function init().
func SetGatewayServiceHook(h GatewayServiceHook) error {
	hook = h
	if err := hook.Bootstrap(availableServiceGroups); err != nil {
		return err
	}
	return nil
}

// RequestReceived will forward call to the hook if set; otherwise no-op.
func RequestReceived(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	if hook == nil {
		return defaultContext, nil
	}

	return hook.RequestReceived(w, r)
}

// RequestAccepted will forward call to the hook if been set, otherwise no-op.
func RequestAccepted(ctx context.Context, spec *pb.ServiceSpec, name string, methodName string, w http.ResponseWriter, r *http.Request) (context.Context, error) {
	if hook == nil {
		return nil, nil
	}

	sg := GetServiceGroup(spec)
	s := sg.Services[name]
	return hook.RequestAccepted(ctx, s, getMethod(s, methodName), w, r)
}

// RequestParsed forwards the call to the RequestParsed method of
// GatewayServiceHook.
func RequestParsed(ctx context.Context, spec *pb.ServiceSpec, name string, methodName string, reqProto proto.Message, meta *ggr.ServerMetadata) error {
	if hook == nil {
		return nil
	}

	sg := GetServiceGroup(spec)
	s := sg.Services[name]
	return hook.RequestParsed(ctx, s, getMethod(s, methodName), reqProto, meta)
}

// RequestHandled will forward call to the hook if been set otherwise noop.
func RequestHandled(ctx context.Context, spec *pb.ServiceSpec, name string, methodName string, out proto.Message, meta *ggr.ServerMetadata, err error) {
	if hook != nil {
		sg := GetServiceGroup(spec)
		s := sg.Services[name]
		hook.RequestHandled(ctx, s, getMethod(s, methodName), out, meta, err)
	}
}

func getMethod(s *Service, methodName string) *Method {
	for _, m := range s.Methods {
		if m.Name == methodName {
			return m
		}
	}
	return nil
}
