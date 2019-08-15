package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/binchencoder/ease-gateway/integrate"
	"github.com/binchencoder/ease-gateway/util"
	"github.com/golang/glog"

	// "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/binchencoder/ease-gateway/gateway/runtime"
)

var (
	host        = flag.String("host", "", "The ease-gateway service host ")
	port        = flag.Int("port", 0, "The ease-gateway service port")
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

func startHttpGateway(mux *runtime.ServeMux, hostPort string) {
	if err := http.ListenAndServe(hostPort, integrate.HttpMux(mux)); err != nil {
		glog.Errorf("Start http gateway error: %v", err)
		shutdown()
		panic(err)
	}
}

func startHttpsGateway(mux *runtime.ServeMux, hostPort string) {
	if err := http.ListenAndServeTLS(hostPort, *certFile, *keyFile, integrate.HttpMux(mux)); err != nil {
		glog.Errorf("Start https gateway error: %v", err)
		shutdown()
		panic(err)
	}
}

func main() {
	glog.Infof("Start http gateway")

	checkFlags()

	hostPort := fmt.Sprintf("%s:%d", *host, *port)

	// 为了开发测试方便,支持flag传参.
	if *port > 0 {
		hostPort = fmt.Sprintf("%s:%d", *host, *port)
	}
	mux := runtime.NewServeMux()
	// runtime.SetGatewayServiceHook(integrate.NewGatewayHook(mux, hostPort))

	// util.Logf(util.DefaultLogger, "*****Starting ease-gateway at %s.*****", hostPort)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)

	if *enableHTTPS {
		go startHttpsGateway(mux, hostPort)
	} else {
		go startHttpGateway(mux, hostPort)
	}

	select {
	case <-signals:
		shutdown()
	}
}

func shutdown() {
	util.Flush()
}
