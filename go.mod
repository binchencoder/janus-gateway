module github.com/binchencoder/ease-gateway

require (
	github.com/Bowery/prompt v0.0.0-20190419144237-972d0ceb96f5 // indirect
	github.com/dchest/safefile v0.0.0-20151022103144-855e8d98f185 // indirect
	github.com/fatih/color v1.7.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.1
	github.com/google/shlex v0.0.0-20181106134648-c34317bd91bf // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.8.6
	github.com/kardianos/govendor v1.0.9 // indirect
	github.com/klauspost/compress v1.5.0
	github.com/klauspost/cpuid v1.2.1 // indirect
	github.com/mattn/go-colorable v0.1.1 // indirect
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	golang.org/x/net v0.0.0-20190425155659-4829fb13d2c6
	google.golang.org/genproto v0.0.0-20190508193815-b515fa19cec8
	google.golang.org/grpc v1.20.1
	gopkg.in/yaml.v2 v2.2.2 // indirect
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
