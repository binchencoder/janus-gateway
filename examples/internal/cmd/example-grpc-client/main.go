package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	prom "github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	hpb "google.golang.org/grpc/health/grpc_health_v1"

	examplepb "github.com/binchencoder/ease-gateway/examples/internal/proto/examplepb"
	vexpb "github.com/binchencoder/gateway-proto/data"
	"github.com/binchencoder/letsgo"
	skylb "github.com/binchencoder/skylb-api/client"
	skypb "github.com/binchencoder/skylb-api/proto"
)

var (
	nBatchRequest  = flag.Int("n-batch-request", 10000, "The number of batched request")
	requestSleep   = flag.Duration("request-sleep", 100*time.Millisecond, "The sleep time after each request")
	requestTimeout = flag.Duration("request-timeout", 100*time.Millisecond, "The timeout of each request")

	spec = skylb.NewServiceSpec(skylb.DefaultNameSpace, vexpb.ServiceId_CUSTOM_EASE_GATEWAY_TEST, skylb.DefaultPortName)

	grpcFailCount = prom.NewCounter(
		prom.CounterOpts{
			Namespace: "skytest",
			Subsystem: "client",
			Name:      "grpc_call_failure",
			Help:      "The number of failed gRPC calls.",
		},
	)
)

func startSkylb(sid vexpb.ServiceId) (skylb.ServiceCli, examplepb.EchoServiceClient, hpb.HealthClient) {
	skycli := skylb.NewServiceCli(vexpb.ServiceId_SHARED_TEST_CLIENT_SERVICE)
	options := []grpc.DialOption{}
	options = append(options, grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin": {}}]}`))
	skycli.Resolve(skylb.NewServiceSpec(skylb.DefaultNameSpace, sid, skylb.DefaultPortName), options...)
	skycli.EnableHistogram()
	var cli examplepb.EchoServiceClient
	var healthCli hpb.HealthClient
	skycli.Start(func(spec *skypb.ServiceSpec, conn *grpc.ClientConn) {
		cli = examplepb.NewEchoServiceClient(conn)
		healthCli = hpb.NewHealthClient(conn)
	})
	return skycli, cli, healthCli
}

func usage() {
	fmt.Println(`Skytest gRPC client.

Usage:
	skytest-client [options]

Options:`)

	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	letsgo.Init(letsgo.FlagUsage(usage))

	testClient()
}

func testClient() {
	sl, cli, healthCli := startSkylb(vexpb.ServiceId_CUSTOM_EASE_GATEWAY_TEST)
	for {
		for i := 0; i < *nBatchRequest; i++ {
			req := examplepb.SimpleMessage{
				Id: fmt.Sprintf("John Doe %d", time.Now().Second()),
			}
			ctx, cancel := context.WithTimeout(context.Background(), *requestTimeout)
			_, err := cli.Echo(ctx, &req, grpc.FailFast(false))
			if err != nil {
				cancel()
				glog.Errorf("Failed to greet service, %v \n", err)
				grpcFailCount.Inc()
				time.Sleep(*requestTimeout)
				continue
			}

			// glog.Infof("Greeting resp: %v \n", resp)

			healthCli.Check(context.Background(), &hpb.HealthCheckRequest{})
			time.Sleep(*requestSleep)
		}
	}

	sl.Shutdown()
}
