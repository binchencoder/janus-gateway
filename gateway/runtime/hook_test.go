package runtime

import (
	"net/http"
	"testing"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"

	pb "github.com/binchencoder/skylb-api/proto"
	ggr "github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type GatewayServiceHookFake struct {
}

func (g *GatewayServiceHookFake) Bootstrap(sgs map[pb.ServiceSpec]*ServiceGroup) error {
	return nil
}
func (g *GatewayServiceHookFake) RequestReceived(w http.ResponseWriter, r *http.Request) (ctx context.Context, err error) {
	return nil, nil
}

func (g *GatewayServiceHookFake) RequestAccepted(ctx context.Context, svc *Service, m *Method, w http.ResponseWriter, r *http.Request) (ctxret context.Context, err error) {
	return nil, nil
}

func (g *GatewayServiceHookFake) RequestParsed(ctx context.Context, svc *Service, m *Method, in proto.Message, meta *ggr.ServerMetadata) error {
	return nil
}

func (g *GatewayServiceHookFake) RequestHandled(ctx context.Context, svc *Service, m *Method, out proto.Message, meta *ggr.ServerMetadata, err error) {
}

func TestHook(t *testing.T) {
	fakeHook := GatewayServiceHookFake{}
	err := SetGatewayServiceHook(&fakeHook)
	if err != nil {
		t.Errorf("Failed to set hook err=%+v", err)
	}
}
