package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/golang/glog"

	"github.com/binchencoder/janus-gateway/gateway/runtime"
	"github.com/binchencoder/janus-gateway/integrate"
	"github.com/binchencoder/janus-gateway/util"
	"github.com/binchencoder/gateway-proto/data"
	"github.com/binchencoder/letsgo"
)

var (
	host = flag.String("host", "", "The gateway service host ")
	port = flag.Int("port", 8080, "The gateway service port")
)

func usage() {
	fmt.Println(`JanusGateway - Janus Gateway of binchencoder.

Usage:
	janus-gateway [options]

Options:`)

	flag.PrintDefaults()
	os.Exit(2)
}

func startHTTPGateway(mux *runtime.ServeMux, hostPort string) {
	if err := http.ListenAndServe(hostPort, integrate.HttpMux(mux)); err != nil {
		glog.Errorf("Start http gateway error: %v", err)
		shutdown()
		panic(err)
	}
}

func main() {
	letsgo.Init(letsgo.FlagUsage(usage))
	// checkFlags()

	// debugMode := flag.Lookup("debug-mode")
	// debugMode.Value.Set("true")

	runtime.CallerServiceId = data.ServiceId_JANUS_GATEWAY
	// integrate.SetAllowCredentials(true)
	// integrate.SetAllowHostsRegexp([]string{"*"})

	glog.Info("***** Janus gateway init. *****")

	hostPort := fmt.Sprintf("%s:%d", *host, *port)
	mux := runtime.NewServeMux()
	runtime.SetGatewayServiceHook(integrate.NewGatewayHook(mux, hostPort))

	glog.Infof("***** Starting custom janus-gateway at %s. *****", hostPort)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	go startHTTPGateway(mux, hostPort)

	select {
	case <-signals:
		shutdown()
	}
}

func shutdown() {
	util.Flush()
}
