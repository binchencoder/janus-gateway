// Note: this file is for ease-gateway  which are exposed to external users.

package integrate

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/context"
	gr "google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	options "binchencoder.com/ease-gateway/httpoptions"
	"binchencoder.com/ease-gateway/gateway/runtime"
	"binchencoder.com/ease-gateway/util"
	vexpb "binchencoder.com/gateway-proto/data"
	fpb "binchencoder.com/gateway-proto/frontend"
	"binchencoder.com/letsgo/grpc"
	"binchencoder.com/letsgo/trace"
)

// Bootstrap starts the gateway and sets up the housekeeping goroutine.
func (gh *gatewayHook) bootstrap(sgs map[string]*runtime.ServiceGroup) error {
	util.Logf(util.DefaultLogger, "*****debug-mode=%t.All program services size=%v", *debugMode, len(sgs))
	util.Logf(util.ConfigLogger, "----------------REGISTER SERVICE-----------------")
	for _, sg := range sgs {
		spec := sg.Spec
		s := fmt.Sprintf("[serviceName:%s | namespace:%s | portName:%s]",
			spec.ServiceName, spec.Namespace, spec.PortName)
		util.Logf(util.DefaultLogger, s)
		for svcKey, svc := range sg.Services {
			util.Logf(util.DefaultLogger, "[Register rpc] %s", svcKey)
			svc.Register(gh.mux)
			for i := 0; i < len(svc.Methods); {
				m := svc.Methods[i]
				// 判断是否为此网关的注册api, 如果不是则从slice中移除
				if !isGatewayApi(m) {
					svc.Methods = append(svc.Methods[:i], svc.Methods[i+1:]...)
					util.Logf(util.DefaultLogger, "  [Remove]%s,%s,%s,%t,%t,%t.", m.Name, m.Path, m.HttpMethod, m.Enabled, m.LoginRequired, m.ClientSignRequired)
				} else {
					i++
					util.Logf(util.DefaultLogger, "  ======>%s,%s,%s,%t,%t,%t.", m.Name, m.Path, m.HttpMethod, m.Enabled, m.LoginRequired, m.ClientSignRequired)
				}
			}
		}
	}
	if *debugMode { // debug模式,直接启动,不使用etcd配置
		if *debugService == "" {
			panic("The flag debug-service is null.")
		}
		for _, sg := range sgs {
			spec := sg.Spec
			if strings.Contains(*debugService, spec.ServiceName) {
				util.Logf(util.DefaultLogger, "&&&Start service %s.", spec.ServiceName)
				go sg.Enable()
			} else {
				util.Logf(util.DefaultLogger, "&&&Do not service %s.", spec.ServiceName)
			}
		}
	} else {
		// TODO(chenbin): 由于没有后台管理来控制APIs的开启和关闭, 默认全部开启
		for _, sg := range sgs {
			go sg.Enable()
		}

		// initEtcd()

		// initMysql()

		// // 注册服务及接口.
		// if err := cache.RegisterService(sgs); err != nil {
		// 	panic(err)
		// }

		// // 注册etcd通知.
		// notify.RegisterEtcdNotify()

		// // 初始化限流配置.
		// go initRateLimit()

		// // 非debug模式才启动redis连接
		// initRedis()

		// go util.SkylbInit()
	}
	return nil
}

func (gh *gatewayHook) requestAccepted(ctx context.Context, svc *runtime.Service, m *runtime.Method, w http.ResponseWriter,
	r *http.Request) (context.Context, error) {
	// client.
	// clt := getClientFromHeader(r.Header)
	// traceid.
	// tid := trace.GetTraceIdOrEmpty(ctx)

	// 新增debug模式,默认uid和cid
	if *debugMode {
		if r.Header.Get(XUid) == "" {
			r.Header.Set(XUid, *debugUid)
		}
		if r.Header.Get(XCid) == "" {
			r.Header.Set(XCid, *debugCid)
		}
		if r.Header.Get(XAid) == "" {
			r.Header.Set(XAid, *debugAid)
		}
	} else {
		// xt, _ := ctx.Value(RequestReceivedTime).(time.Time)
		// 校验请求的http header.
		if err := verifyHeader(ctx, r.Header, svc, m); err != nil {
			return ctx, err
		}
		// // 获取api存储配置信息.
		// api, ok := cache.GetApi(m.HttpMethod, m.Path)
		// if !ok {
		// 	return apiNil(ctx, svc, m, clt, tid, xt)
		// }
		// // api开启状态.
		// if !api.Enabled {
		// 	return apiDisabled(ctx, svc, m, clt, tid, xt)
		// }

		// // 判断SourceAllow.
		// if !api.IsThirdParty && api.SourceAllow.String != "all" && r.Header.Get(XSource) != api.SourceAllow.String {
		// 	return apiForbidden(ctx, svc, m, clt, tid, xt)
		// }
		// // api验签校验.
		// if api.SignRequired && !api.IsThirdParty {
		// 	if ok, err := checkSign(r); !ok {
		// 		return apiSignErr(ctx, svc, m, clt, tid, xt, err)
		// 	}
		// }

		// // api登录校验.
		// if api.LoginRequired && !api.IsThirdParty {
		// 	authS := NewExternalAuthServer(util.AuthClient, util.AuthNClient, authRedis, mgwRedis)
		// 	if err := authS.Authenticate(r, m); err != nil {
		// 		return apiLoginErr(ctx, svc, m, clt, tid, xt, err)
		// 	}
		// }

		// // api限流.
		// if cof, ok := cache.GetApiLimit(m.HttpMethod, m.Path); ok {
		// 	if err := apiLimit(ctx, r.Header, svc, m, xt, cof); err != nil {
		// 		return ctx, err
		// 	}
		// }
	}

	return ctx, nil
}

func getReqInfo(r *http.Request) string {
	// remote address.
	remoteAddr := r.Header.Get(XForwardedFor)
	if remoteAddr == "" {
		remoteAddr = r.RemoteAddr
	}

	str := fmt.Sprintf("%s,%s,%s,u:%s,c:%s,a:%s,d:%s,%s,%s", r.Header.Get(XSource), r.Header.Get(XClient),
		r.Header.Get(XAppVersion), r.Header.Get(XUid), r.Header.Get(XCid), r.Header.Get(XAid), r.Header.Get(XDid),
		remoteAddr, r.URL)
	return str
}

// verifyHeader verify http header.
func verifyHeader(ctx context.Context, header http.Header, svc *runtime.Service, m *runtime.Method) error {
	s := header.Get(XSource)
	c := header.Get(XClient)
	if !m.IsThirdParty && s == ResourceThird {
		// 非ThirdParty接口不能采用第三方source请求数据.
		return grpcError(codes.InvalidArgument, fpb.ErrorCode_BADPARAM_ERROR, []string{"x-source error."})
	}

	errStr := ""
	// 校验必要header信息.
	if s == "" {
		errStr = fmt.Sprintf("%s is missing.", XSource)
	} else if !IsKnownResource(s) {
		errStr = "Unknown x-source."
	} else if s == ResourceClient {
		if c == "" {
			errStr = fmt.Sprintf("%s is missing.", XClient)
		} else if xa := header.Get(XAppVersion); xa == "" {
			errStr = fmt.Sprintf("%s is missing.", XAppVersion)
		} else if xd := header.Get(XDid); xd == "" {
			errStr = fmt.Sprintf("%s is missing.", XDid)
		}
	}

	ms := float64(-1)
	// client.
	clt := getClientFromHeader(header)
	// traceid.
	tid := trace.GetTraceIdOrEmpty(ctx)

	if errStr != "" {
		// prometheus metrics.
		if tm, ok := ctx.Value(RequestReceivedTime).(time.Time); ok {
			ms = addMetrics(ctx, svc, m, codes.InvalidArgument, tm, clt)
		}

		// record default logs.
		util.Logf(util.DefaultLogger, util.DefaultFormat, tid, errStr)
		// record stat logs.
		util.Logf(util.StatLogger, util.StatFormat, tid, svc.Spec.GetServiceName(), m.HttpMethod, m.Path, clt, "N", codes.InvalidArgument, ms)

		ger := grpcError(codes.InvalidArgument, fpb.ErrorCode_BADPARAM_ERROR, []string{"VerifyHeader failed.", errStr})
		// record rest logs.
		util.Logf(util.RestLogger, util.ResponseRestFormat, tid, codes.InvalidArgument, gr.ErrorDesc(ger))

		return ger
	}

	return nil
}

// getClient returns client value who request ease-gateway from header.
func getClientFromHeader(header http.Header) string {
	xs := header.Get(XSource)
	cl := header.Get(XClient)
	if !IsKnownResource(xs) {
		return ""
	}
	if xs == ResourceClient {
		return cl
	}
	return xs
}

// apiNil处理api不存在情况.
func apiNil(ctx context.Context, svc *runtime.Service, m *runtime.Method, clt, tid string, xt time.Time) (context.Context, error) {
	// prometheus metrics.
	ms := addMetrics(ctx, svc, m, codes.PermissionDenied, xt, clt)

	// record default logs.
	util.Logf(util.DefaultLogger, util.DefaultFormat, tid, "There is no api config.svc="+svc.Spec.GetServiceName()+",path="+m.Path+",method="+m.HttpMethod)
	// record stat logs.
	util.Logf(util.StatLogger, util.StatFormat, tid, svc.Spec.GetServiceName(), m.HttpMethod, m.Path, clt, "N", codes.PermissionDenied, ms)

	ger := grpcError(codes.PermissionDenied, fpb.ErrorCode_RESOURCE_NOT_FOUND, []string{"There is no api config."})
	// record rest logs.
	util.Logf(util.RestLogger, util.ResponseRestFormat, tid, codes.PermissionDenied, gr.ErrorDesc(ger))

	return ctx, ger
}

// apiDisabled处理api关闭情况.
func apiDisabled(ctx context.Context, svc *runtime.Service, m *runtime.Method, clt, tid string, xt time.Time) (context.Context, error) {
	// prometheus metrics.
	ms := addMetrics(ctx, svc, m, codes.PermissionDenied, xt, clt)

	// record default logs.
	util.Logf(util.DefaultLogger, util.DefaultFormat, tid, "The api is disabled.svc="+svc.Spec.GetServiceName()+",path="+m.Path+",method="+m.HttpMethod)
	// record stat logs.
	util.Logf(util.StatLogger, util.StatFormat, tid, svc.Spec.GetServiceName(), m.HttpMethod, m.Path, clt, "N", codes.PermissionDenied, ms)

	ger := grpcError(codes.Unavailable, fpb.ErrorCode_SERVICE_DOWN, []string{"The api is disabled."})
	// record rest logs.
	util.Logf(util.RestLogger, util.ResponseRestFormat, tid, codes.PermissionDenied, gr.ErrorDesc(ger))

	return ctx, ger
}

// apiForbidden 禁止访问.
func apiForbidden(ctx context.Context, svc *runtime.Service, m *runtime.Method, clt, tid string, xt time.Time) (context.Context, error) {
	// prometheus metrics.
	ms := addMetrics(ctx, svc, m, codes.PermissionDenied, xt, clt)

	// record default logs.
	util.Logf(util.DefaultLogger, util.DefaultFormat, tid, "The source is forbidden.svc="+svc.Spec.GetServiceName()+",path="+m.Path+",method="+m.HttpMethod)
	// record stat logs.
	util.Logf(util.StatLogger, util.StatFormat, tid, svc.Spec.GetServiceName(), m.HttpMethod, m.Path, clt, "N", codes.PermissionDenied, ms)

	ger := grpcError(codes.PermissionDenied, fpb.ErrorCode_BAD_REQUEST, []string{"The source is forbidden."})
	// record rest logs.
	util.Logf(util.RestLogger, util.ResponseRestFormat, tid, codes.PermissionDenied, gr.ErrorDesc(ger))

	return ctx, ger
}

// apiSignErr处理api验签失败情况.
func apiSignErr(ctx context.Context, svc *runtime.Service, m *runtime.Method, clt, tid string, xt time.Time, err error) (context.Context, error) {
	// prometheus metrics.
	ms := addMetrics(ctx, svc, m, codes.InvalidArgument, xt, clt)

	// record default logs.
	util.Logf(util.DefaultLogger, util.DefaultFormat, tid, fmt.Sprintf("Check sign fail. The reason is: %v.", err))
	// record stat logs.
	util.Logf(util.StatLogger, util.StatFormat, tid, svc.Spec.GetServiceName(), m.HttpMethod, m.Path, clt, "N", codes.InvalidArgument, ms)

	ger := grpcError(codes.InvalidArgument, fpb.ErrorCode_INVALID_SIGNATURE, []string{"Check sign fail."})
	// record rest logs.
	util.Logf(util.RestLogger, util.ResponseRestFormat, tid, codes.InvalidArgument, gr.ErrorDesc(ger))
	return ctx, ger
}

// apiSignErr处理登录校验失败情况.
func apiLoginErr(ctx context.Context, svc *runtime.Service, m *runtime.Method, clt, tid string, xt time.Time, err error) (context.Context, error) {
	// prometheus metrics.
	ms := addMetrics(ctx, svc, m, codes.Unauthenticated, xt, clt)

	// record default logs.
	util.Logf(util.DefaultLogger, util.DefaultFormat, tid, fmt.Sprintf("Authenticate fail. The reason is: %v.", err))
	// record stat logs.
	util.Logf(util.StatLogger, util.StatFormat, tid, svc.Spec.GetServiceName(), m.HttpMethod, m.Path, clt, "N", codes.Unauthenticated, ms)

	errParam := []string{"Authenticate fail."}
	if _, pbErr := grpc.FromGrpcError(err); pbErr != nil {
		errParam = append(errParam, pbErr.Params...)
	}
	ger := grpcError(codes.Unauthenticated, fpb.ErrorCode_AUTHEN_ERROR, errParam)
	// record rest logs.
	util.Logf(util.RestLogger, util.ResponseRestFormat, tid, codes.Unauthenticated, gr.ErrorDesc(ger))
	return ctx, ger
}

func getHeader(h http.Header, key string) string {
	value := ""
	if v, ok := h[http.CanonicalHeaderKey(key)]; ok && len(v) > 0 {
		value = v[0]
	}
	return value
}

func isGatewayApi(api *runtime.Method) bool {
	if runtime.CallerServiceId == vexpb.ServiceId_EASE_GATEWAY {
		return api.ApiSource == options.ApiSourceType_EASE_GATEWAY
	}
	// else if runtime.CallerServiceId == vexpb.ServiceId_EASE_OPEN_GATEWAY {
	// 	return api.ApiSource == options.ApiSourceType_EASE_OPEN_GATEWAY ||
	// 		api.ApiSource == options.ApiSourceType_OPEN_GATEWAY_PRIVATE
	// }
	return false
}
