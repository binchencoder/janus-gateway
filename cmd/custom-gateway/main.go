package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/golang/glog"

	"binchencoder.com/ease-gateway/gateway/runtime"
	"binchencoder.com/ease-gateway/integrate"
	"binchencoder.com/gateway-proto/data"
	"binchencoder.com/letsgo"
)

var (
	host = flag.String("host", "", "The gateway service host ")
	port = flag.Int("port", 8080, "The gateway service port")
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
	// checkFlags()

	// debugMode := flag.Lookup("debug-mode")
	// debugMode.Value.Set("true")

	runtime.CallerServiceId = data.ServiceId_EASE_GATEWAY
	// integrate.SetAllowCredentials(true)
	// integrate.SetAllowHostsRegexp([]string{"*"})

	glog.Info("***** Ease gateway init. *****")

	hostPort := fmt.Sprintf("%s:%d", *host, *port)
	mux := runtime.NewServeMux()
	runtime.SetGatewayServiceHook(integrate.NewGatewayHook(mux, hostPort))

	glog.Infof("***** Starting custom ease-gateway at %s. *****", hostPort)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	startHTTPGateway(mux, hostPort)
}
