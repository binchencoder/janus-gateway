package integrate

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	gr "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	options "github.com/binchencoder/ease-gateway/gateway/options"
	"github.com/binchencoder/ease-gateway/gateway/runtime"
	"github.com/binchencoder/ease-gateway/integrate/metrics"
	"github.com/binchencoder/ease-gateway/util"
	"github.com/binchencoder/letsgo/grpc"
	"github.com/binchencoder/letsgo/trace"

	// vexpb "github.com/binchencoder/ease-gateway/proto/data"
	fpb "github.com/binchencoder/ease-gateway/proto/frontend"
)

var (
	debugMode    = flag.Bool("debug-mode", false, "If debug mode,not use etcd config and redis.")
	debugUid     = flag.String("debug-uid", "179", "Fake user ID, used only in debug mode.")
	debugCid     = flag.String("debug-cid", "10", "Fake company ID, used only in debug mode.")
	debugAid     = flag.String("debug-aid", "", "Fake personal account ID, used only in debug mode.")
	debugService = flag.String("debug-service", "", "Specifies service which will be started, be used only in debug mode.")

	// For docker-compose command line override only.
	ectdEndpointsStr     = flag.String("etcd-endpoints", "", "ETCD endpoints")
	mysqlHostsStr        = flag.String("mysql-hosts", "", "Mysql hosts")
	mysqlUserStr         = flag.String("mysql-user", "", "Mysql user")
	mysqlPwdStr          = flag.String("mysql-password", "", "Mysql password")
	redisAddressesStr    = flag.String("redis-addresses", "", "Redis addresses")
	redisMgwAddressesStr = flag.String("redis-mgw-addresses", "", "Redis mgw addresses")
)

// gatewayHook implements interface GatewayServiceHook in package
// github.com/binchencoder/ease-gateway/gateway/runtime.
type gatewayHook struct {
	mux  *runtime.ServeMux
	host string
}

// Bootstrap starts the gateway and sets up the housekeeping goroutine.
func (gh *gatewayHook) Bootstrap(sgs map[string]*runtime.ServiceGroup) error {
	util.Logf(util.DefaultLogger, "*****debug-mode=%t.All program services size=%v", *debugMode, len(sgs))

	if len(sgs) == 0 {
		panic("No program service was found.")
	}

	return gh.bootstrap(sgs)
}

func (gh *gatewayHook) RequestReceived(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	ctx := context.Background()
	// traceid deal.
	if xt := r.Header.Get(XRequestId); xt == "" {
		tid := trace.GenerateTraceId()
		ctx = trace.WithTraceId(ctx, tid)
		r.Header.Add(XRequestId, tid)
		w.Header().Set(XRequestId, tid)
	} else {
		w.Header().Set(XRequestId, xt)
		ctx = trace.WithTraceId(ctx, xt)
	}
	// ldap-gateway set header.
	// TODO(chenbin) 2019/08/07 notes
	// if runtime.CallerServiceId == vexpb.ServiceId_LDAP_GATEWAY {
	// 	r.Header.Set(XSource, ResourceLdap)
	// } else if runtime.CallerServiceId == vexpb.ServiceId_OPEN_GATEWAY {
	// 	r.Header.Set(XSource, ResourceOpenPlatform)
	// }
	// Set request start time for calculating latency.
	ctx = context.WithValue(ctx, RequestReceivedTime, time.Now())
	return ctx, nil
}

func (gh *gatewayHook) RequestAccepted(ctx context.Context, svc *runtime.Service, m *runtime.Method, w http.ResponseWriter,
	r *http.Request) (context.Context, error) {
	if m.IsThirdParty {
		r.Header.Set(XSource, ResourceThird)
	}
	// 对未设置x-source头的请求进行处理.
	if r.Header.Get(XSource) == "" {
		// api指定补充x-source为web的处理.
		if m.SpecifiedSource == options.SpecSourceType_WEB {
			r.Header.Set(XSource, ResourceWeb)
		}
	}

	tid := trace.GetTraceIdOrEmpty(ctx)
	util.Logf(util.RestLogger, util.RequestRestFormat, tid, svc.Spec.GetServiceName(), m.HttpMethod, m.Path, getReqInfo(r))

	ctx, err := gh.requestAccepted(ctx, svc, m, w, r)

	// 转发全部http header
	h := make(map[string]string)
	for k, v := range r.Header {
		h[k] = v[0]
	}
	ctxret := metadata.NewOutgoingContext(ctx, metadata.New(h))
	return ctxret, err
}

func (gh *gatewayHook) RequestParsed(ctx context.Context, svc *runtime.Service, m *runtime.Method,
	req proto.Message, meta *runtime.ServerMetadata) error {

	return nil
}

func (gh *gatewayHook) RequestHandled(ctx context.Context, svc *runtime.Service, m *runtime.Method,
	resp proto.Message, meta *runtime.ServerMetadata, err error) {
	ms := float64(-1)
	// traceid.
	tid := trace.GetTraceIdOrEmpty(ctx)
	// client.
	clt := ""
	s := "Y"
	tm, ok1 := ctx.Value(RequestReceivedTime).(time.Time)
	code := gr.Code(err)
	if md, ok := metadata.FromOutgoingContext(ctx); ok && ok1 {
		clt = getClientFroMd(md)
		// prometheus metrics.
		ms = addMetrics(ctx, svc, m, code, tm, clt)
	}

	if err != nil {
		s = "N"
		util.Logef(util.ErrorLogger, util.ErrorFormat, tid, fmt.Sprintf("RequestHandled error: %v.", err))
	}
	// record stat logs.
	util.Logf(util.StatLogger, util.StatFormat, tid, svc.Spec.GetServiceName(), m.HttpMethod, m.Path, clt, s, code, ms)
	// record rest logs.
	util.Logf(util.RestLogger, util.ResponseRestFormat, tid, code, gr.ErrorDesc(err))
}

// NewGatewayHook returns a new gatewayHook.
func NewGatewayHook(mux *runtime.ServeMux, host string) runtime.GatewayServiceHook {
	return &gatewayHook{
		mux:  mux,
		host: host,
	}
}

// addMetrics add metrics to prometheus for ease-gateway.
func addMetrics(ctx context.Context, svc *runtime.Service, m *runtime.Method, code codes.Code, startTime time.Time, clt string) float64 {
	rp := &metrics.ReporterParam{StartTime: startTime, ServiceName: svc.Spec.GetServiceName(), Url: m.Path, HttpMethod: m.HttpMethod, Code: strconv.FormatUint(uint64(code), 10), Client: clt}
	return rp.RequestComplete()
}

// getClient returns client value who request ease-gateway from Md.
func getClientFroMd(md metadata.MD) string {
	if s, ok := md[XSource]; ok && len(s) > 0 && s[0] != ResourceClient {
		return s[0]
	}
	if s, ok := md[XClient]; ok && len(s) > 0 {
		return s[0]
	}
	return ""
}

func grpcError(rpcCode codes.Code, pbCode fpb.ErrorCode, params []string) error {
	e := fpb.Error{
		Code:   pbCode,
		Params: params,
	}
	return grpc.ToGrpcError(rpcCode, &e)
}
