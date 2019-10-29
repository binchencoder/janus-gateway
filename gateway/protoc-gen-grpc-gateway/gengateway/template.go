package gengateway

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"unicode"

	"github.com/golang/glog"
	generator2 "github.com/golang/protobuf/protoc-gen-go/generator"

	// "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/descriptor"
	"binchencoder.com/ease-gateway/gateway/protoc-gen-grpc-gateway/descriptor"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
)

type param struct {
	*descriptor.File
	Imports            []descriptor.GoPackage
	UseRequestContext  bool
	RegisterFuncSuffix string
	AllowPatchFeature  bool
}

type binding struct {
	*descriptor.Binding
	Registry          *descriptor.Registry
	AllowPatchFeature bool
}

// GetBodyFieldPath returns the binding body's fieldpath.
func (b binding) GetBodyFieldPath() string {
	if b.Body != nil && len(b.Body.FieldPath) != 0 {
		return b.Body.FieldPath.String()
	}
	return "*"
}

// badToUnderscore is the mapping function used to generate Go names from
// package names, which can be dotted in the input .proto file.  It
// replaces non-identifier characters such as dot or dash with underscore.
func badToUnderscore(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
		return r
	}
	return '_'
}

// GetUniqueName return the unique valid go variable name
// based on package name.
func (p param) GetUniqueName() string {
	s := strings.Split(p.GetName(), ".")
	return strings.Map(badToUnderscore, s[0])
}

// HasQueryParam determines if the binding needs parameters in query string.
//
// It sometimes returns true even though actually the binding does not need.
// But it is not serious because it just results in a small amount of extra codes generated.
func (b binding) HasQueryParam() bool {
	if b.Body != nil && len(b.Body.FieldPath) == 0 {
		return false
	}
	fields := make(map[string]bool)
	for _, f := range b.Method.RequestType.Fields {
		fields[f.GetName()] = true
	}
	if b.Body != nil {
		delete(fields, b.Body.FieldPath.String())
	}
	for _, p := range b.PathParams {
		delete(fields, p.FieldPath.String())
	}
	return len(fields) > 0
}

func (b binding) QueryParamFilter() queryParamFilter {
	var seqs [][]string
	if b.Body != nil {
		seqs = append(seqs, strings.Split(b.Body.FieldPath.String(), "."))
	}
	for _, p := range b.PathParams {
		seqs = append(seqs, strings.Split(p.FieldPath.String(), "."))
	}
	return queryParamFilter{utilities.NewDoubleArray(seqs)}
}

// HasEnumPathParam returns true if the path parameter slice contains a parameter
// that maps to an enum proto field that is not repeated, if not false is returned.
func (b binding) HasEnumPathParam() bool {
	return b.hasEnumPathParam(false)
}

// HasRepeatedEnumPathParam returns true if the path parameter slice contains a parameter
// that maps to a repeated enum proto field, if not false is returned.
func (b binding) HasRepeatedEnumPathParam() bool {
	return b.hasEnumPathParam(true)
}

// hasEnumPathParam returns true if the path parameter slice contains a parameter
// that maps to a enum proto field and that the enum proto field is or isn't repeated
// based on the provided 'repeated' parameter.
func (b binding) hasEnumPathParam(repeated bool) bool {
	for _, p := range b.PathParams {
		if p.IsEnum() && p.IsRepeated() == repeated {
			return true
		}
	}
	return false
}

// LookupEnum looks up a enum type by path parameter.
func (b binding) LookupEnum(p descriptor.Parameter) *descriptor.Enum {
	e, err := b.Registry.LookupEnum("", p.Target.GetTypeName())
	if err != nil {
		return nil
	}
	return e
}

// FieldMaskField returns the golang-style name of the variable for a FieldMask, if there is exactly one of that type in
// the message. Otherwise, it returns an empty string.
func (b binding) FieldMaskField() string {
	var fieldMaskField *descriptor.Field
	for _, f := range b.Method.RequestType.Fields {
		if f.GetTypeName() == ".google.protobuf.FieldMask" {
			// if there is more than 1 FieldMask for this request, then return none
			if fieldMaskField != nil {
				return ""
			}
			fieldMaskField = f
		}
	}
	if fieldMaskField != nil {
		return generator2.CamelCase(fieldMaskField.GetName())
	}
	return ""
}

// queryParamFilter is a wrapper of utilities.DoubleArray which provides String() to output DoubleArray.Encoding in a stable and predictable format.
type queryParamFilter struct {
	*utilities.DoubleArray
}

func (f queryParamFilter) String() string {
	encodings := make([]string, len(f.Encoding))
	for str, enc := range f.Encoding {
		encodings[enc] = fmt.Sprintf("%q: %d", str, enc)
	}
	e := strings.Join(encodings, ", ")
	return fmt.Sprintf("&utilities.DoubleArray{Encoding: map[string]int{%s}, Base: %#v, Check: %#v}", e, f.Base, f.Check)
}

type trailerParams struct {
	Services           []*descriptor.Service
	UseRequestContext  bool
	RegisterFuncSuffix string
	AssumeColonVerb    bool
}

func applyTemplate(p param, reg *descriptor.Registry) (string, error) {
	w := bytes.NewBuffer(nil)
	if err := headerTemplate.Execute(w, p); err != nil {
		return "", err
	}

	if err := validatorTemplate.Execute(w, p); err != nil {
		return "", err
	}

	var targetServices []*descriptor.Service

	for _, msg := range p.Messages {
		msgName := generator2.CamelCase(*msg.Name)
		msg.Name = &msgName
	}
	for _, svc := range p.Services {
		var methodWithBindingsSeen bool
		svcName := generator2.CamelCase(*svc.Name)
		svc.Name = &svcName
		for _, meth := range svc.Methods {
			glog.V(2).Infof("Processing %s.%s", svc.GetName(), meth.GetName())
			methName := generator2.CamelCase(*meth.Name)
			meth.Name = &methName
			for _, b := range meth.Bindings {
				methodWithBindingsSeen = true
				if err := handlerTemplate.Execute(w, binding{
					Binding:           b,
					Registry:          reg,
					AllowPatchFeature: p.AllowPatchFeature,
				}); err != nil {
					return "", err
				}
			}
		}
		if methodWithBindingsSeen {
			targetServices = append(targetServices, svc)
		}
	}
	if len(targetServices) == 0 {
		return "", errNoTargetService
	}

	assumeColonVerb := true
	if reg != nil {
		assumeColonVerb = !reg.GetAllowColonFinalSegments()
	}
	tp := trailerParams{
		Services:           targetServices,
		UseRequestContext:  p.UseRequestContext,
		RegisterFuncSuffix: p.RegisterFuncSuffix,
		AssumeColonVerb:    assumeColonVerb,
	}
	if err := trailerTemplate.Execute(w, tp); err != nil {
		return "", err
	}
	return w.String(), nil
}

var (
	headerTemplate = template.Must(template.New("header").Parse(`
// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: {{.GetName}}

/*
Package {{.GoPkg.Name}} is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package {{.GoPkg.Name}}
import (
	{{range $i := .Imports}}{{if $i.Standard}}{{$i | printf "%s\n"}}{{end}}{{end}}

	{{range $i := .Imports}}{{if not $i.Standard}}{{$i | printf "%s\n"}}{{end}}{{end}}
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ sync.RWMutex
var _ proto.Message
var _ context.Context
var _ grpc.ClientConn
var _ client.ServiceCli
var _ vexpb.ServiceId
var _ = http.MethodGet
var _ regexp.Regexp
var _ = balancer.ConsistentHashing
var _ option.BalancerCreator
var _ naming.Resolver
var _ strings.Reader
var _ = utf8.UTFMax

// TODO (jiezmo): check if there is any rule before create this var.
var {{.GetUniqueName}}_error = lgr.ToGrpcError(codes.InvalidArgument, &fpb.Error{Code: fpb.ErrorCode_BADPARAM_ERROR, Params: []string{"Validation error"},})
`))

	validatorTemplate = template.Must(template.New("validator").Parse(`
{{$fileUniqueName := .GetUniqueName}}
// Validation methods start
{{range $, $message := .Messages}}
	{{if $message.HasRule}}
	func {{$message.GetValidationMethodName}}(v *{{$message.File.GoPkg.Path  | $message.GoType }}) error{
	if v == nil {
		return nil
	}
	// Validation for each Fields
	{{range $,$f := $message.Fields}}  {{/*0*/}}

		{{if $f.HasRule}}
			{{if $f.IsRepeated}}
				for _, vv2 := range v.{{$f.GoName}} {
			{{else}}
				{
				vv2 := v.{{if $f.IsOneOf}}Get{{$f.GoName}}(){{else}}{{$f.GoName}}{{end}}
				{{if $f.IsOneOf}}
					if _, ok := v.{{$f.OneOfDeclGoName}}.(*{{$message.File.GoPkg.Path  | $message.GoType }}_{{$f.GoName}}); ok {
				{{end}}
			{{end}}
		{{end}}

		// Validation Field: {{if $f.IsOneOf}}Get{{$f.GoName}}(){{else}}{{$f.GoName}}{{end}}
		{{range $,$r := $f.Rules}}   {{/*1*/}}
			{{if $r.IsOpMatch}}   {{/*2*/}}
				{{if $r.IsTypeString}}
					// Match pattern {{if $f.IsOneOf}}Get{{$f.GoName}}(){{else}}{{$f.GoName}}{{end}}
					if matched, err := regexp.MatchString("{{$r.Value}}", vv2); matched == false || err != nil{
						return {{$fileUniqueName}}_error
					}
				{{else}}
					// TODO(jiezmo): fail the build
					//Err, only string can have pattern match
				{{end}}
			{{else if $r.IsOpEq}}
				{{if $r.IsTypeString}}
					if "{{$r.Value}}" != vv2{
						return {{$fileUniqueName}}_error
					}
				{{else}}
					if {{$r.Value}} != vv2{
						return {{$fileUniqueName}}_error
					}
				{{end}}
			{{else if $r.IsOpGt}}
				{{if $r.IsTypeNumber}}
					if vv2 <= {{$r.Value}} {
						return {{$fileUniqueName}}_error
					}
				{{else}}
					// TODO(jiezmo): fail the build
					// Err
				{{end}}
			{{else if $r.IsOpLt}}
				{{if $r.IsTypeNumber}}
					if vv2 >= {{$r.Value}} {
						return {{$fileUniqueName}}_error
					}
				{{else}}
					// TODO(jiezmo): fail the build
					// Err
				{{end}}
			{{else if $r.IsOpNotNil}}
				{{if $r.IsTypeObj}}
					if vv2 == nil{
						return {{$fileUniqueName}}_error
					}
				{{else}}
					// TODO(jiezmo): fail the build
					// Err
				{{end}}
			{{else if $r.IsLenEq}}
				{{if $r.IsTypeString}}
					{{if $r.NeedTrim}}
					if utf8.RuneCountInString(strings.TrimSpace(vv2)) != {{$r.Value}}{
						return {{$fileUniqueName}}_error
					}
					{{else}}
					if {{$r.Value}} != utf8.RuneCountInString(vv2){
						return {{$fileUniqueName}}_error
					}
					{{end}}
				{{else}}
					// TODO(jiezmo): fail the build
					// Err
				{{end}}
			{{else if $r.IsLenGt}}
				{{if $r.IsTypeString}}
					{{if $r.NeedTrim}}
					if utf8.RuneCountInString(strings.TrimSpace(vv2)) <= {{$r.Value}}{
						return {{$fileUniqueName}}_error
					}
					{{else}}
					if utf8.RuneCountInString(vv2) <= {{$r.Value}}{
						return {{$fileUniqueName}}_error
					}
					{{end}}
				{{else}}
					// TODO(jiezmo): fail the build
					// Err
				{{end}}
			{{else if $r.IsLenLt}}
				{{if $r.IsTypeString}}
					{{if $r.NeedTrim}}
					if utf8.RuneCountInString(strings.TrimSpace(vv2)) >= {{$r.Value}}{
						return {{$fileUniqueName}}_error
					}
					{{else}}
					if utf8.RuneCountInString(vv2) >= {{$r.Value}}{
						return {{$fileUniqueName}}_error
					}
					{{end}}
				{{else}}
					// TODO(jiezmo): fail the build
					// Err
				{{end}}
			{{else}}
				// TODO(jiezmo): fail the build
				// Error
			{{end}}  {{/*2*/}}
		{{end}}  {{/*1*/}}

		{{if $f.HasRule}}
			{{if $f.IsOneOf}}
				}
			{{end}}
			}
		{{end}}

		{{if $f.FieldMessage}}
			{{if and $f.FieldMessage.HasRule}}
				{{if $f.IsRepeated}}
				for _, vv := range v.{{$f.GoName}} {
					if err := {{$message.File.GoPkg.Path | $f.FieldMessage.GetValidationMethodQualifiedName}}(vv); err != nil {
						return err
					}
				}
				{{else}}
				if err := {{$message.File.GoPkg.Path | $f.FieldMessage.GetValidationMethodQualifiedName}}(v.{{if $f.IsOneOf}}Get{{$f.GoName}}(){{else}}{{$f.GoName}}{{end}}); err != nil {
					return err
				}
				{{end}}
			{{end}}
		{{end}}
	{{end}}
	return nil
	}
	{{end}}

{{end}}
// Validation methods done
`))

	handlerTemplate = template.Must(template.New("handler").Parse(`
{{if and .Method.GetClientStreaming .Method.GetServerStreaming}}
{{template "bidi-streaming-request-func" .}}
{{else if .Method.GetClientStreaming}}
{{template "client-streaming-request-func" .}}
{{else}}
{{template "client-rpc-request-func" .}}
{{end}}
`))

	_ = template.Must(handlerTemplate.New("request-func-signature").Parse(strings.Replace(`
{{if .Method.GetServerStreaming}}
func request_{{.Method.Service.GetName}}_{{.Method.GetName}}_{{.Index}}(ctx context.Context, marshaler runtime.Marshaler, client {{.Method.Service.GetName}}Client, req *http.Request, pathParams map[string]string) ({{.Method.Service.GetName}}_{{.Method.GetName}}Client, runtime.ServerMetadata, error)
{{else}}
func request_{{.Method.Service.GetName}}_{{.Method.GetName}}_{{.Index}}(ctx context.Context, marshaler runtime.Marshaler, client {{.Method.Service.GetName}}Client, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error)
{{end}}`, "\n", "", -1)))

	_ = template.Must(handlerTemplate.New("client-streaming-request-func").Parse(`
{{template "request-func-signature" .}} {
	var metadata runtime.ServerMetadata
	stream, err := client.{{.Method.GetName}}(ctx)
	if err != nil {
		grpclog.Infof("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	for {
		var protoReq {{.Method.RequestType.GoType .Method.Service.File.GoPkg.Path}}
		err = dec.Decode(&protoReq)
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Infof("Failed to decode request: %v", err)
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
		if err = stream.Send(&protoReq); err != nil {
			if err == io.EOF {
				break
			}
			grpclog.Infof("Failed to send request: %v", err)
			return nil, metadata, err
		}
	}

	if err := stream.CloseSend(); err != nil {
		grpclog.Infof("Failed to terminate client stream: %v", err)
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		grpclog.Infof("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header
{{if .Method.GetServerStreaming}}
	return stream, metadata, nil
{{else}}
	msg, err := stream.CloseAndRecv()
	metadata.TrailerMD = stream.Trailer()
	return msg, metadata, err
{{end}}
}
`))

	_ = template.Must(handlerTemplate.New("client-rpc-request-func").Parse(`
{{$AllowPatchFeature := .AllowPatchFeature}}
{{if .HasQueryParam}}
var (
	filter_{{.Method.Service.GetName}}_{{.Method.GetName}}_{{.Index}} = {{.QueryParamFilter}}
)
{{end}}
{{template "request-func-signature" .}} {
	var protoReq {{.Method.RequestType.GoType .Method.Service.File.GoPkg.Path}}
	var metadata runtime.ServerMetadata
	spec := internal_{{.Method.Service.GetName}}_{{.Method.Service.ServiceId}}_spec
{{if .Body}}
	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&{{.Body.AssignableExpr "protoReq"}}); err != nil && err != io.EOF  {
		runtime.RequestHandled(ctx, spec, "{{.Method.Service.GetName}}", "{{.Method.GetName}}", nil, &metadata, err)
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	{{- if and $AllowPatchFeature (and (eq (.HTTPMethod) "PATCH") (.FieldMaskField))}}
	if protoReq.{{.FieldMaskField}} != nil && len(protoReq.{{.FieldMaskField}}.GetPaths()) > 0 {
		runtime.CamelCaseFieldMask(protoReq.{{.FieldMaskField}})
	} {{if not (eq "*" .GetBodyFieldPath)}} else {
			if fieldMask, err := runtime.FieldMaskFromRequestBody(newReader()); err != nil {
				return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
			} else {
				protoReq.{{.FieldMaskField}} = fieldMask
			}
	} {{end}}
	{{end}}
{{end}}
{{if .PathParams}}
	var (
		val string
{{- if .HasEnumPathParam}}
		e int32
{{- end}}
{{- if .HasRepeatedEnumPathParam}}
		es []int32
{{- end}}
		ok bool
		err error
		_ = err
	)
	{{$binding := .}}
	{{range $param := .PathParams}}
	{{$enum := $binding.LookupEnum $param}}
	val, ok = pathParams[{{$param | printf "%q"}}]
	if !ok {
		runtime.RequestHandled(ctx, spec, "{{.Method.Service.GetName}}", "{{.Method.GetName}}", nil, &metadata, err)
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", {{$param | printf "%q"}})
	}
{{if $param.IsNestedProto3}}
	err = runtime.PopulateFieldFromPath(&protoReq, {{$param | printf "%q"}}, val)
{{else if $enum}}
	e{{if $param.IsRepeated}}s{{end}}, err = {{$param.ConvertFuncExpr}}(val{{if $param.IsRepeated}}, {{$binding.Registry.GetRepeatedPathParamSeparator | printf "%c" | printf "%q"}}{{end}}, {{$enum.GoType $param.Target.Message.File.GoPkg.Path}}_value)
{{else}}
	{{$param.AssignableExpr "protoReq"}}, err = {{$param.ConvertFuncExpr}}(val{{if $param.IsRepeated}}, {{$binding.Registry.GetRepeatedPathParamSeparator | printf "%c" | printf "%q"}}{{end}})
{{end}}
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", {{$param | printf "%q"}}, err)
	}
{{if and $enum $param.IsRepeated}}
	s := make([]{{$enum.GoType $param.Target.Message.File.GoPkg.Path}}, len(es))
	for i, v := range es {
		s[i] = {{$enum.GoType $param.Target.Message.File.GoPkg.Path}}(v)
	}
	{{$param.AssignableExpr "protoReq"}} = s
{{else if $enum}}
	{{$param.AssignableExpr "protoReq"}} = {{$enum.GoType $param.Target.Message.File.GoPkg.Path}}(e)
{{end}}
	{{end}}
{{end}}
{{if .HasQueryParam}}
	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_{{.Method.Service.GetName}}_{{.Method.GetName}}_{{.Index}}); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
{{end}}
{{if .Method.GetServerStreaming}}
	stream, err := client.{{.Method.GetName}}(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
{{else}}
	// Only hook up for non-stream call for now.

	{{if and .Method.RequestType.HasRule}}
	// Validate
	// {{.Method.RequestType.GoType .Method.Service.File.GoPkg.Path}}
	if err :={{.Method.RequestType.GetValidationMethodName}}(&protoReq); err != nil {
		runtime.RequestHandled(ctx, spec, "{{.Method.Service.GetName}}", "{{.Method.GetName}}", nil, &metadata, err)
		return nil, metadata, err
	}
	{{end}}
	runtime.RequestParsed(ctx, spec, "{{.Method.Service.GetName}}", "{{.Method.GetName}}", &protoReq, &metadata)
	ctx = runtime.PreLoadBalance(ctx, "{{$.Method.Service.Balancer.String}}", "{{.Method.HashKey}}", &protoReq)
	msg, err := client.{{.Method.GetName}}(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	runtime.RequestHandled(ctx, spec, "{{.Method.Service.GetName}}", "{{.Method.GetName}}", msg, &metadata, err)
	return msg, metadata, err
{{end}}
}`))

	_ = template.Must(handlerTemplate.New("bidi-streaming-request-func").Parse(`
{{template "request-func-signature" .}} {
	var metadata runtime.ServerMetadata
	stream, err := client.{{.Method.GetName}}(ctx)
	if err != nil {
		grpclog.Infof("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	handleSend := func() error {
		var protoReq {{.Method.RequestType.GoType .Method.Service.File.GoPkg.Path}}
		err := dec.Decode(&protoReq)
		if err == io.EOF {
			return err
		}
		if err != nil {
			grpclog.Infof("Failed to decode request: %v", err)
			return err
		}
		if err := stream.Send(&protoReq); err != nil {
			grpclog.Infof("Failed to send request: %v", err)
			return err
		}
		return nil
	}
	if err := handleSend(); err != nil {
		if cerr := stream.CloseSend(); cerr != nil {
			grpclog.Infof("Failed to terminate client stream: %v", cerr)
		}
		if err == io.EOF {
			return stream, metadata, nil
		}
		return nil, metadata, err
	}
	go func() {
		for {
			if err := handleSend(); err != nil {
				break
			}
		}
		if err := stream.CloseSend(); err != nil {
			grpclog.Infof("Failed to terminate client stream: %v", err)
		}
	}()
	header, err := stream.Header()
	if err != nil {
		grpclog.Infof("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}
`))

	trailerTemplate = template.Must(template.New("trailer").Parse(`
// Register itself to runtime.
func init() {
	var s *runtime.Service
	var spec *skypb.ServiceSpec

	_ = s
	_ = spec
{{range $svc := .Services}}
	spec = internal_{{$svc.GetName}}_{{$svc.ServiceId}}_spec
	s = &runtime.Service {
		Spec    : *spec,
		Name    : "{{$svc.GetName}}",
		Register: Register{{$svc.GetName}}{{$.RegisterFuncSuffix}}FromEndpoint,
		Enable  : Enable{{$svc.GetName}}_Service,
		Disable : Disable{{$svc.GetName}}_Service,
	}

	{{if $svc.GenController}}
		runtime.AddService(s, Enable_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_ServiceGroup, Disable_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_ServiceGroup)
	{{else}}
		runtime.AddService(s, nil, nil)
	{{end}}
{{end}}
}

{{$UseRequestContext := .UseRequestContext}}
{{range $svc := .Services}}
// Register{{$svc.GetName}}{{$.RegisterFuncSuffix}}FromEndpoint is same as Register{{$svc.GetName}}{{$.RegisterFuncSuffix}} but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func Register{{$svc.GetName}}{{$.RegisterFuncSuffix}}FromEndpoint(mux *runtime.ServeMux) (err error) {
	// conn, err := grpc.Dial(endpoint, opts...)
	// if err != nil {
	// 	return err
	// }
	// defer func() {
	// 	if err != nil {
	// 		if cerr := conn.Close(); cerr != nil {
	// 			grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
	// 		}
	// 		return
	// 	}
	// 	go func() {
	// 		<-ctx.Done()
	// 		if cerr := conn.Close(); cerr != nil {
	// 			grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
	// 		}
	// 	}()
	// }()

	// return Register{{$svc.GetName}}{{$.RegisterFuncSuffix}}(ctx, mux, conn)
	return Register{{$svc.GetName}}{{$.RegisterFuncSuffix}}(nil, mux, nil)
}

// Register{{$svc.GetName}}{{$.RegisterFuncSuffix}} registers the http handlers for service {{$svc.GetName}} to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func Register{{$svc.GetName}}{{$.RegisterFuncSuffix}}(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return Register{{$svc.GetName}}{{$.RegisterFuncSuffix}}Client(ctx, mux, New{{$svc.GetName}}Client(conn))
}

// Register{{$svc.GetName}}{{$.RegisterFuncSuffix}}Client registers the http handlers for service {{$svc.GetName}}
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "{{$svc.GetName}}Client".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "{{$svc.GetName}}Client"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "{{$svc.GetName}}Client" to call the correct interceptors.
func Register{{$svc.GetName}}{{$.RegisterFuncSuffix}}Client(ctx context.Context, mux *runtime.ServeMux, client {{$svc.GetName}}Client) error {
	spec := internal_{{$svc.GetName}}_{{$svc.ServiceId}}_spec

	{{range $m := $svc.Methods}}
	{{range $b := $m.Bindings}}
	runtime.AddMethod(spec, "{{$svc.GetName}}", "{{$m.GetName}}", "{{$b.PathTmpl.Template}}", {{$b.HTTPMethod | printf "%q"}}, {{$m.LoginRequired}}, {{$m.ClientSignRequired}}, {{$m.IsThirdParty}}, "{{$m.SpecSourceType}}", "{{$m.ApiSource}}", "{{$m.TokenType}}", "{{$m.Timeout}}")
	mux.Handle({{$b.HTTPMethod | printf "%q"}}, pattern_{{$svc.GetName}}_{{$m.GetName}}_{{$b.Index}}, vexpb.ServiceId_{{$svc.ServiceId}}, func(inctx context.Context, w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		// TODO(mojz): review all locking/unlocking logic.
		// internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_lock.RLock()
		// defer internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_lock.RUnlock()
		cli := internal_{{$svc.GetName}}_{{$svc.ServiceId}}_client
		if cli == nil {
			runtime.DefaultOtherErrorHandler(w, req, "service disabled", http.StatusInternalServerError)
			return
		}

		ctx, err := runtime.RequestAccepted(inctx, internal_{{$svc.GetName}}_{{$svc.ServiceId}}_spec, "{{$svc.GetName}}", "{{$m.GetName}}", w, req)
		if err != nil {
			grpclog.Errorf("runtime.HTTPError error: %v", err)
			runtime.HTTPError(ctx, nil, &runtime.JSONBuiltin{}, w, req, err)
			return
		}

		{{ if $UseRequestContext }}
			ctx, cancel := context.WithCancel(req.Context())
		{{- else -}}
			ctx, cancel := context.WithCancel(ctx)
		{{- end }}
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			grpclog.Errorf("runtime.HTTPError error: %v", err)
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_{{$svc.GetName}}_{{$m.GetName}}_{{$b.Index}}(rctx, inboundMarshaler, cli, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			grpclog.Errorf("runtime.HTTPError error: %v", err)
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		{{if $m.GetServerStreaming}}
		forward_{{$svc.GetName}}_{{$m.GetName}}_{{$b.Index}}(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)
		{{else}}
		{{ if $b.ResponseBody }}
		forward_{{$svc.GetName}}_{{$m.GetName}}_{{$b.Index}}(ctx, mux, outboundMarshaler, w, req, response_{{$svc.GetName}}_{{$m.GetName}}_{{$b.Index}}{resp}, mux.GetForwardResponseOptions()...)
		{{ else }}
		forward_{{$svc.GetName}}_{{$m.GetName}}_{{$b.Index}}(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)
		{{end}}
		{{end}}
	})
	{{end}}
	{{end}}
	return nil
}

{{range $m := $svc.Methods}}
{{range $b := $m.Bindings}}
{{if $b.ResponseBody}}
type response_{{$svc.GetName}}_{{$m.GetName}}_{{$b.Index}} struct {
	proto.Message
}

func (m response_{{$svc.GetName}}_{{$m.GetName}}_{{$b.Index}}) XXX_ResponseBody() interface{} {
	response := m.Message.(*{{$m.ResponseType.GoType $m.Service.File.GoPkg.Path}})
	return {{$b.ResponseBody.AssignableExpr "response"}}
}
{{end}}
{{end}}
{{end}}

{{if $svc.GenController}}
func Disable_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_ServiceGroup() {
	// internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_lock.Lock()
	// defer internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_lock.Unlock()
	if internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_skycli != nil {
		spec := internal_{{$svc.GetName}}_{{$svc.ServiceId}}_spec
		sg := runtime.GetServiceGroup(spec)
		for _, svc := range sg.Services {
			svc.Disable()
		}

		internal_{{$svc.GetName}}_{{$svc.ServiceId}}_client = nil
		internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_skycli.Shutdown()
		internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_skycli = nil
	}
}

func Enable_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_ServiceGroup() {
	// internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_lock.Lock()
	// defer internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_lock.Unlock()

	internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_skycli = client.NewServiceCli(runtime.CallerServiceId)

	// Resolve service
	spec := internal_{{$svc.GetName}}_{{$svc.ServiceId}}_spec

	{{if eq $svc.Balancer.String "ROUND_ROBIN"}}
	internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_skycli.Resolve(spec)
	{{else if eq $svc.Balancer.String "CONSISTENT"}}
	internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_skycli.Resolve(spec,
		option.WithBalancerCreator(func(r naming.Resolver) grpc.Balancer {
			return balancer.ConsistentHashing(r)
		}),
	)
	{{end}}
	internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_skycli.EnableResolveFullEps()
	internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_skycli.Start(func(spec *skypb.ServiceSpec, conn *grpc.ClientConn) {
		sg := runtime.GetServiceGroup(spec)
		for _, svc := range sg.Services {
			svc.Enable(spec, conn)
		}
	})
}
{{end}}

func Enable{{$svc.GetName}}_Service(spec *skypb.ServiceSpec, conn *grpc.ClientConn) {
	internal_{{$svc.GetName}}_{{$svc.ServiceId}}_client = New{{$svc.GetName}}Client(conn)
}

func Disable{{$svc.GetName}}_Service() {
	internal_{{$svc.GetName}}_{{$svc.ServiceId}}_client = nil
}

var (
	{{range $m := $svc.Methods}}
	{{range $b := $m.Bindings}}
	pattern_{{$svc.GetName}}_{{$m.GetName}}_{{$b.Index}} = runtime.MustPattern(runtime.NewPattern({{$b.PathTmpl.Version}}, {{$b.PathTmpl.OpCodes | printf "%#v"}}, {{$b.PathTmpl.Pool | printf "%#v"}}, {{$b.PathTmpl.Verb | printf "%q"}}, runtime.AssumeColonVerbOpt({{$.AssumeColonVerb}})))
	{{end}}
	{{end}}
)

var (
	{{range $m := $svc.Methods}}
	{{range $b := $m.Bindings}}
	forward_{{$svc.GetName}}_{{$m.GetName}}_{{$b.Index}} = {{if $m.GetServerStreaming}}runtime.ForwardResponseStream{{else}}runtime.ForwardResponseMessage{{end}}
	{{end}}
	{{end}}
)

var (
	internal_{{$svc.GetName}}_{{$svc.ServiceId}}_spec = client.NewServiceSpec("{{$svc.Namespace}}", vexpb.ServiceId_{{$svc.ServiceId}}, "{{$svc.PortName}}")
	internal_{{$svc.GetName}}_{{$svc.ServiceId}}_client {{$svc.GetName}}Client
	{{if $svc.GenController}}
	internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_skycli client.ServiceCli

	internal_{{$svc.ServiceId}}__{{$svc.Namespace}}__{{$svc.PortName}}_lock =sync.RWMutex{}
	{{end}}
)

{{end}}`))
)
