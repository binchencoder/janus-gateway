package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/golang/glog"

	"github.com/binchencoder/ease-gateway/gateway/runtime"
	"github.com/binchencoder/ease-gateway/proto/data"
	"github.com/binchencoder/ease-gateway/integrate"
	"github.com/binchencoder/letsgo"
)

var (
	host        = flag.String("host", "", "The janus service host ")
	port        = flag.Int("port", 6666, "The janus service port")
)

func usage() {
	fmt.Println(`EaseGateway - Ease Gateway of binchencoder.

Usage:
	ease-gateway [options]

Options:`)

	flag.PrintDefaults()
	os.Exit(2)
}

func startHTTPGateway(mux *runtime.ServeMux, hostPort string) {
	if err := http.ListenAndServe(hostPort, integrate.HttpMux(mux)); err != nil {
		panic(err)
	}
}

func main() {
	letsgo.Init(letsgo.FlagUsage(usage))
	checkFlags()

	debugMode := flag.Lookup("debug-mode")
	debugMode.Value.Set("true")

	runtime.CallerServiceId = data.ServiceId_CUSTOM_GATEWAY
	// integrate.SetAllowCredentials(true)
	// integrate.SetAllowHostsRegexp([]string{"*"})

	glog.Info("***** Ease gateway init. *****")

	hostPort := fmt.Sprintf("%s:%d", *host, *port)
	mux := runtime.NewServeMux()
	runtime.SetGatewayServiceHook(integrate.NewGatewayHook(mux, hostPort))

	glog.Infof("***** Starting janus at %s. *****", hostPort)

	startHTTPGateway(mux, hostPort)
}
