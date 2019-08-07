package integrate

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var (
	allowHosts         = []string{"*.xxx.com", "localhost", "localhost:8080", "192.168.*"}
	allowMethods       = []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	allowExposeHeaders = []string{XRequestId}

	allowHostsRegexp = regexp.MustCompile(getHostRegStr(allowHosts))

	allowCredentials = false
)

// SetAllowHostsRegexp sets the allowed hosts for CORS.
// Note: this should only be used by custom gateway and ldap gateway.
func SetAllowHostsRegexp(hosts []string) {
	allowHostsRegexp = regexp.MustCompile(getHostRegStr(hosts))
}

// SetAllowCredentials sets to allow CORS credentials.
func SetAllowCredentials(allow bool) {
	allowCredentials = allow
}

// http请求处理中间件
type CorsMiddleware struct {
	Handler http.Handler
}

func (m *CorsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if origin := r.Header.Get("Origin"); origin != "" {
		// 请求域非法.
		if !isOriginOk(origin) {
			// util.Logf(util.DefaultLogger, "The request source is illegal. %s", origin)
			return
		}

		if allowCredentials {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}

		w.Header().Set("Access-Control-Allow-Origin", origin)
		// 预检请求.
		if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
			return
		}

		exHeaders := allowExposeHeaders
		w.Header().Set("Access-Control-Expose-Headers", strings.Join(exHeaders, ","))
	}
	m.Handler.ServeHTTP(w, r)
}

func isOriginOk(origin string) bool {
	// if *debugMode {
	// 	// util.Logf(util.ErrorLogger, "CORS: in debug mode, return true regardless.")
	// 	return true
	// }

	u, err := url.Parse(origin)
	if err != nil {
		return false
	}

	return hostMatch(u.Hostname())
}

// hostMatch 判断host是否与正则表达式匹配.
func hostMatch(host string) bool {
	return allowHostsRegexp.Match([]byte(host))
}

// getHostRegStr 获取host模板列表的正则匹配表达式.
func getHostRegStr(temps []string) string {
	ts := []string{}
	for i := range temps {
		ts = append(ts, "^"+strings.Replace(temps[i], "*", "[\\w\\.]+", -1)+"$")
	}

	return strings.Join(ts, "|")
}
