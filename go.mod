module github.com/binchencoder/ease-gateway

require (
	github.com/VividCortex/gohistogram v1.0.0 // indirect
	github.com/binchencoder/letsgo v0.0.0-20190813050654-d221d6b03c21
	github.com/binchencoder/skylb-api v0.0.0-20190816070449-75a71296dbf2
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/fatih/color v1.7.0
	github.com/ghodss/yaml v1.0.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.0.0
	github.com/grpc-ecosystem/grpc-gateway v1.9.4
	github.com/grpc-ecosystem/grpc-opentracing v0.0.0-20180507213350-8e809c8a8645 // indirect
	github.com/klauspost/compress v1.5.0
	github.com/klauspost/cpuid v1.2.1 // indirect
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/pborman/uuid v1.2.0
	github.com/prometheus/client_golang v1.1.0
	github.com/uber/jaeger-client-go v2.16.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.0.0+incompatible // indirect
	golang.org/x/net v0.0.0-20190613194153-d28f0bde5980
	google.golang.org/genproto v0.0.0-20190508193815-b515fa19cec8
	google.golang.org/grpc v1.20.1
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.34.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190427214059-a29dc8fdc73485
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190427214059-956cc1757749645f24
	golang.org/x/image => github.com/golang/image v0.0.0-20190427214059-59b11bec70c7cc648c
	golang.org/x/lint => github.com/golang/lint v0.0.0-20181026193005-c67002cb31c3
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190427214059-3e0bab5405d63a8
	golang.org/x/net => github.com/golang/net v0.0.0-20190425155659-4829fb13d2c6
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20180821212333-d2e6202438be
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20181116152217-5ac8a444bdc5
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190425155659-ad9eeb80039afa
	google.golang.org/appengine => github.com/golang/appengine v1.6.1
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190425155659-357c62f0e4bb
	google.golang.org/grpc => github.com/grpc/grpc-go v1.17.0
	gopkg.in/yaml.v2 => github.com/go-yaml/yaml v0.0.0-20181115110504-51d6538a90f8
	honnef.co/go/tools => github.com/dominikh/go-tools v0.0.0-20190425155659-e561f6794a2a09dd
)
