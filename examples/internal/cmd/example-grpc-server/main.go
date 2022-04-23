/*
 * Command example-grpc-server is an example grpc server
 * to be called by example-gateway-server.
 */
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	examplepb "github.com/binchencoder/janus-gateway/examples/internal/proto/examplepb"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/binchencoder/gateway-proto/data"
	skylb "github.com/binchencoder/skylb-api/server"

	"github.com/binchencoder/janus-gateway/examples/internal/server"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

var (
	addr       = flag.String("addr", ":9090", "endpoint of the gRPC service")
	network    = flag.String("network", "tcp", "a valid network type which is consistent to -addr")
	scrapeAddr = flag.String("scrape-addr", "0.0.0.0:18001", "The prometheus scrape port")

	port = flag.Int("port", 9090, "The gRPC port of the server")
)

func registerPrometheus() {
	http.Handle("/_/metrics", prometheus.UninstrumentedHandler())
	if err := http.ListenAndServe(*scrapeAddr, nil); err != nil {
		log.Fatal("ListenServerError:", err)
	}
}

func main() {
	flag.Parse()
	defer glog.Flush()

	// Don't regist to skylbserver
	// ctx := context.Background()
	// if err := server.Run(ctx, *network, *addr); err != nil {
	// 	glog.Fatal(err)
	// }

	go registerPrometheus()

	skylb.Register(data.ServiceId_CUSTOM_JANUS_GATEWAY_TEST, "grpc", *port)
	skylb.EnableHistogram()
	skylb.Start(fmt.Sprintf(":%d", *port), func(s *grpc.Server) error {
		examplepb.RegisterEchoServiceServer(s, server.NewEchoServer())
		return nil
	})
}
