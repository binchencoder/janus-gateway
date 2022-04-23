/*
 * Command grpc-server is an example gRPC server
 * to be called by //cmd/gateway.
 */
package main

import (
	"flag"
	"fmt"

	examplepb "github.com/binchencoder/janus-gateway/proto/examplepb"
	"github.com/binchencoder/gateway-proto/data"
	skylb "github.com/binchencoder/skylb-api/server"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

var (
	network = flag.String("network", "tcp", "a valid network type which is consistent to -addr")

	port = flag.Int("port", 9090, "The custom gRPC port of the server")
)

func main() {
	flag.Parse()
	defer glog.Flush()

	// Regist to skylbserver
	skylb.Register(data.ServiceId_CUSTOM_JANUS_GATEWAY_TEST, "grpc", *port)
	skylb.EnableHistogram()
	skylb.Start(fmt.Sprintf(":%d", *port), func(s *grpc.Server) error {
		examplepb.RegisterEchoServiceServer(s, NewEchoServer())
		return nil
	})
}
