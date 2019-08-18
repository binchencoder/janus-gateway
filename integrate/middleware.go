package integrate

import (
	"flag"
	"net/http"

	"binchencoder.com/ease-gateway/gateway/runtime"
)

var (
	enableGzip = flag.Bool("gzip", true, "Whether to enable gzip.")
	enableCors = flag.Bool("enable-cors", false, "Whether to enable HTTP access control.")
)

// HttpMux 返回所需的封装Handler.
func HttpMux(mux *runtime.ServeMux) http.Handler {
	var sm http.Handler = mux
	if *enableGzip {
		sm = &GzipMiddleware{sm}
	}

	// if *enableCors {
	// 	sm = &CorsMiddleware{sm}
	// }
	return sm
}
