# Overrview

ease-gateway/examples 是使用ease-gateway的一个完整示例，包含gateway-server 和 gRPC-server. 还有Java实现gRPC Server的例子 [https://github.com/binchencoder/spring-boot-grpc/tree/master/spring-boot-grpc-examples]

# Build the example

build gateway server
```
bazel build ease-gateway/examples/cmd/example-gateway-server/... 
```

build gRPC server
```
bazel build ease-gateway/examples/cmd/example-grpc-server/...
```

# Run the example

start gateway server
```
ease-gateway/bazel-bin/examples/cmd/example-gateway-server/linux_amd64_stripped/example-gateway-server -skylb-endpoints="127.0.0.1:1900" -debug-svc-endpoint=custom-ease-gateway-test=localhost:9090
```

start custom-gateway server
```
ease-gateway/bazel-bin/cmd/custom-gateway/linux_amd64_stripped/custom-ease-gateway -skylb-endpoints="127.0.0.1:1900" -debug-service=custom-ease-gateway-test -debug-svc-endpoint=custom-ease-gateway-test=localhost:9090
```

start examples gRPC server for test //examples/cmd/example-gateway-server
```
ease-gateway/bazel-bin/examples/cmd/example-grpc-server/linux_amd64_stripped/example-grpc-server
```

start gRPC server for test //cmd/gateway
```
ease-gateway/bazel-bin/examples/cmd/grpc-server/linux_amd64_stripped/grpc-server -skylb-endpoints="127.0.0.1:1900"
```

start //cmd/gateway
```
ease-gateway/bazel-bin/cmd/gateway/linux_amd64_stripped/gateway -skylb-endpoints="127.0.0.1:1900" -v=2 -log_dir=.
```

# Usage

You can use curl or a browser to test:
```
# List all apis
$ curl http://localhost:8080/swagger/echo_service.swagger.json

# Visit the apis
$ curl -XPOST  -H "x-source: web" http://localhost:8080/v1/example/echo/foo
{"id":"foo"}

$ curl -H "x-source: web"  http://localhost:8080/v1/example/echo/foo/123
{"id":"foo","num":"123"}

$ curl -XDELETE -H "x-source: web"  http://localhost:8080/v1/example/echo_delete
```

> NOTE: 请注意当前用户是否是管理员