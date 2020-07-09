# Overview

grpc-gateway 是一款非常优秀的网关服务器，负责转化和代理转发。让```RESTful API ```和 ```gRPC```可以相互转化，这样可以实现一套```gRPC```接口提供两种接口服务（提供内部的```gRPC```服务和外部```RESTful API```服务），大大提高了开发效率。

但是官方提供的版本还是单机版的，还不支持集群，所以并不能直接运行在生产环境中。```ease-gateway```就是为了解决这些问题应运而生的，在```grpc-gateway```的基础上增加了新的feature

- 支持自定义的loadbalancer
- 支持网关层的parameter validation
- 支持自定义的annotaion

为了支持这些新feature，我不得不将grpc-gateway的源码拷贝到ease-gateway中并进行大量的修改，```//gateway``` 目录下就是从grpc-gateway中拷贝的部分源码

# 修改内容

以下详细说明了我基于grpc-gateway修改了哪些目录下的代码

## httpoptions

这个目录下的内容并不是grpc-gateway的源码，是基于[Google
// APIs](https://github.com/googleapis/googleapis) 修改的，为了支持：

- 网关认证
- loadbalancer
- parameter validation(定义validation rules)

> **NOTE :** 这是新增加的目录

## protoc-gen-grpc-gateway

Link   [grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway/tree/master/protoc-gen-grpc-gateway)

```protoc-gen-grpc-gateway``` 是生成反向代理的工具

```
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  path/to/your_service.proto
```

修改文件:

- BUILD.bazel
- main.go (只修改import path)

### descriptor

修改文件:

- BUILD.bazel
- grpc_api_service.go
- registry.go
- services_test.go
- services.go
- types.go

### generator

修改文件:

- BUILD.bazel
- generator.go  (只修改import path)

### internal/gengateway

**这个目录下改动比较大**, 修改文件:

- BUILD.bazel
- generator_test.go
- generator.go
- template_test.go
- template.go

### httprule

Link [grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/httprule](https://github.com/grpc-ecosystem/grpc-gateway/tree/master/protoc-gen-grpc-gateway/httprule)

该目录下没有修改代码，直接拷贝的。原因是 grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway/httprule  package visibility 只能在本项目内依赖

```
package(default_visibility = ["//:generators"])
```

修改文件:

- BUILD.bazel

## protoc-gen-swagger

Link   [grpc-ecosystem/grpc-gateway/protoc-gen-swagger](https://github.com/grpc-ecosystem/grpc-gateway/tree/master/protoc-gen-swagger)

```protoc-gen-swagger``` 是生成swagger API定义的工具

修改文件:

- BUILD.bazel
- defs.bzl
- main.go

### genswagger

修改文件:

- BUILD.bazel
- generator.go
- template.go
- template_test.go
- types.go

### options

修改文件:

- BUILD.bazel
- annotations.proto
- openapiv2.proto

## runtime

Link   [grpc-ecosystem/grpc-gateway/runtime](https://github.com/grpc-ecosystem/grpc-gateway/tree/master/runtime)

这是```grpc-ecosystem/grpc-gateway/runtime``` 的核心模块

新增文件:

- service.go
- balancer_test.go
- balancer.go
- hook.go
- hook_test.go

删除文件:

- marshal_json_test.go
- marshal_jsonpb_test.go
- marshal_proto_test.go

修改文件:

- BUILD.bazel
- context_test.go
- error_test.go
- handler_test.go
- marshal_httpbodyproto_test.go
- marshaler_registry_test.go
- mux_test.go
- mux.go
- query_test.go

## grpc-ecosystem/grpc-gateway

grpc-ecosystem/grpc-gateway: grpc-gateway is a plugin of protoc. It reads
gRPC service definition, and generates a reverse-proxy server which translates
a RESTful JSON API into gRPC. This server is generated according to custom
options in your gRPC definition.

Hosted on https://github.com/grpc-ecosystem/grpc-gateway.