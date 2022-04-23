package gateway

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/binchencoder/janus-gateway/examples/internal/proto/examplepb"
	gwruntime "github.com/binchencoder/janus-gateway/gateway/runtime"
	"google.golang.org/grpc"
)

// newGateway returns a new gateway server which translates HTTP into gRPC.
func newGateway(ctx context.Context, conn *grpc.ClientConn, opts []gwruntime.ServeMuxOption) (http.Handler, error) {
	sgs := gwruntime.GetServicGroups()
	fmt.Printf("gwruntime.GetServicGroups: %v \n", sgs)
	for _, sg := range sgs {
		go sg.Enable()
		spec := sg.Spec
		fmt.Printf("[serviceName:%s | namespace:%s | portName:%s]\n",
			spec.ServiceName, spec.Namespace, spec.PortName)
	}

	mux := gwruntime.NewServeMux(opts...)

	// examplepb.Enable_CUSTOM_EASE_GATEWAY_TEST__default__grpc_ServiceGroup()
	for _, f := range []func(context.Context, *gwruntime.ServeMux, *grpc.ClientConn) error{
		examplepb.RegisterEchoServiceHandler,
	} {
		if err := f(ctx, mux, conn); err != nil {
			return nil, err
		}
	}
	return mux, nil
}

func dial(ctx context.Context, network, addr string) (*grpc.ClientConn, error) {
	switch network {
	case "tcp":
		return dialTCP(ctx, addr)
	case "unix":
		return dialUnix(ctx, addr)
	default:
		return nil, fmt.Errorf("unsupported network type %q", network)
	}
}

// dialTCP creates a client connection via TCP.
// "addr" must be a valid TCP address with a port number.
func dialTCP(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	return grpc.DialContext(ctx, addr, grpc.WithInsecure())
}

// dialUnix creates a client connection via a unix domain socket.
// "addr" must be a valid path to the socket.
func dialUnix(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	d := func(ctx context.Context, addr string) (net.Conn, error) {
		return (&net.Dialer{}).DialContext(ctx, "unix", addr)
	}
	return grpc.DialContext(ctx, addr, grpc.WithInsecure(), grpc.WithContextDialer(d))
}
