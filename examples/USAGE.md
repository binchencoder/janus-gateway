# Overrview

ease-gateway/examples 是使用ease-gateway的一个完整示例，包含gateway-server 和 gRPC-server

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
ease-gateway/bazel-bin/examples/cmd/example-gateway-server/linux_amd64_stripped/example-gateway-server
```

```
bazel-bin/cmd/custom-gateway/linux_amd64_stripped/ease-gateway-demo -skylb-endpoints="127.0.0.1:1900" -debug-service=custom-ease-gateway-test -debug-svc-endpoint=custom-ease-gateway-test=localhost:9090
```

start gRPC server
```
ease-gateway/bazel-bin/examples/cmd/example-grpc-server/linux_amd64_stripped/example-server -skylb-endpoints="127.0.0.1:1900"
```

# Usage

You can use curl or a browser to test:
```
# List all apis
$ curl http://localhost:8080/swagger/echo_service.swagger.json

# Visit the apis
$ curl -XPOST http://localhost:8080/v1/example/echo/foo
{"id":"foo"}

$ curl  http://localhost:8080/v1/example/echo/foo/123
{"id":"foo","num":"123"}
```

> NOTE: 请注意当前用户是否是管理员