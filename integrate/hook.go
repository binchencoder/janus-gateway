package integrate

import (
	"flag"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
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
// jingoal.com/janus/gateway/runtime.
type gatewayHook struct {
	mux  *runtime.ServeMux
	host string
}

// NewGatewayHook returns a new gatewayHook.
// func NewGatewayHook(mux *runtime.ServeMux, host string) runtime.GatewayServiceHook {
// 	return &gatewayHook{
// 		mux:  mux,
// 		host: host,
// 	}
// }
