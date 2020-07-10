# ease-gateway

Gateway service based on [grpc-ecosystem/grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway).  This helps you provide your APIs in both gRPC and RESTful style at the same time.

![](./docs/images/grpc-rest-gateway.png)

当我第一次使用gRPC就深深的爱上了她, 她在跨平台、跨语言、面向移动和HTTP/2设计上有着天然的优势, 我更喜欢的是她定义API的方式. 了解更多gRPC信息请查看[英文官方文档](https://www.grpc.io/docs/guides/) 、 [中文官方文档](http://grpc.mydoc.io/)

**grpc-gateway** 的出现更能引起开发者对使用gRPC的兴趣, 她可以帮你在原有gRPC服务的基础上做少量的改动, 便可以将原gRPC服务同时提供RESTful HTTP API, 了解更多RESTful API的例子可以参考[GitHub REST API](https://developer.github.com/v3/) 、[Google REST API](https://developers.google.com/drive/v2/reference/)

**ease-gateway** 是站在巨人的肩膀上实现的, 增加了更符合企业级应用开发的**Features**:

- 支持自定义的LoadBalancer
- 既可以部署单机版模式, 也可注册到注册中心实现集群模式
- 支持网关层的Parameter Validation Rules
- 支持自定义的Annotaion

## Prepared

**ease-gateway** 使用GO MOD来管理Dependencies，clone代码之后直接在本地使用bazel构建

### Build tools

- Bazel 3.1.0+
- Go 1.13.12+

## Clone code

```shell
git clone https://github.com/binchencoder/ease-gateway.git
```

## Bazel build gateway

```
cd ease-gateway

bazel build cmd/gateway/...
```

## Usage

TODO

## Run Examples

See [examples/README.md](https://github.com/binchencoder/ease-gateway/tree/master/examples)
