package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/golang/glog"

	"github.com/binchencoder/ease-gateway/gateway/runtime"
	"github.com/binchencoder/ease-gateway/integrate"
	"github.com/binchencoder/ease-gateway/util"
	"github.com/binchencoder/gateway-proto/data"
	"github.com/binchencoder/letsgo"
	"github.com/binchencoder/letsgo/service/naming"
)

var (
	host        = flag.String("host", "", "The ease-gateway service host ")
	port        = flag.Int("port", 8080, "The ease-gateway service port")
	enableHTTPS = flag.Bool("enable-https", false, "Whether to enable https.")
	certFile    = flag.String("cert-file", "", "The TLS cert file.")
	keyFile     = flag.String("key-file", "", "The TLS key file.")
)

func usage() {
	fmt.Println(`Ease Gateway - Universal Gateway of xxx Inc.

Usage:
	ease-gateway [options]

Options:`)

	flag.PrintDefaults()
	os.Exit(2)
}

func checkFlags() {
	if *enableHTTPS {
		if *certFile == "" {
			fmt.Println("Flag --cert-file is required to enable HTTPS.")
			os.Exit(2)
		}
		if *keyFile == "" {
			fmt.Println("Flag --key-file is required to enable HTTPS.")
			os.Exit(2)
		}
	}
}

func startHTTPGateway(mux *runtime.ServeMux, hostPort string) {
	if err := http.ListenAndServe(hostPort, integrate.HttpMux(mux)); err != nil {
		glog.Errorf("Start http gateway error: %v", err)
		shutdown()
		panic(err)
	}
}

func startHTTPSGateway(mux *runtime.ServeMux, hostPort string) {
	if err := http.ListenAndServeTLS(hostPort, *certFile, *keyFile, integrate.HttpMux(mux)); err != nil {
		glog.Errorf("Start https gateway error: %v", err)
		shutdown()
		panic(err)
	}
}

func main() {
	defer letsgo.Cleanup()
	letsgo.Init(letsgo.FlagUsage(usage))
	checkFlags()

	hostPort := fmt.Sprintf("%s:%d", *host, *port)
	runtime.CallerServiceId = data.ServiceId_EASE_GATEWAY
	serviceName, err := naming.ServiceIdToName(runtime.CallerServiceId)
	if err != nil {
		glog.Errorf("Invalid service id %d", runtime.CallerServiceId)
		panic(err)
	}
	util.Logf(util.DefaultLogger, "*****%s init.*****", serviceName)

	// 为了开发测试方便,支持flag传参.
	if *port > 0 {
		hostPort = fmt.Sprintf("%s:%d", *host, *port)
	}
	mux := runtime.NewServeMux()
	runtime.SetGatewayServiceHook(integrate.NewGatewayHook(mux, hostPort))

	util.Logf(util.DefaultLogger, "*****Starting %s at %s.*****", serviceName, hostPort)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	if *enableHTTPS {
		go startHTTPSGateway(mux, hostPort)
	} else {
		go startHTTPGateway(mux, hostPort)
	}

	select {
	case <-signals:
		shutdown()
	}
}

func shutdown() {
	util.Flush()
}
