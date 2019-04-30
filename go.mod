module github.com/binchencoder/ease-gateway

require (
	github.com/fatih/color v1.7.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.1
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	golang.org/x/net v0.0.0-20190425155659-4829fb13d2c6
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.34.0
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190427214059-a29dc8fdc73485
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190427214059-956cc1757749645f24
	golang.org/x/image => github.com/golang/image v0.0.0-20190427214059-59b11bec70c7cc648c
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190427214059-3e0bab5405d63a8
	golang.org/x/net => github.com/golang/net v0.0.0-20190425155659-4829fb13d2c6
	golang.org/x/sys => github.com/golang/sys v0.0.0-20181116152217-5ac8a444bdc5
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190425155659-ad9eeb80039afa
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190425155659-357c62f0e4bb
	google.golang.org/grpc => github.com/grpc/grpc-go v1.17.0
	gopkg.in/yaml.v2 => github.com/go-yaml/yaml v0.0.0-20181115110504-51d6538a90f8
	honnef.co/go/tools => github.com/dominikh/go-tools v0.0.0-20190425155659-e561f6794a2a09dd
)