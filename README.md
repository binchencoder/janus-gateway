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

**ease-gateway** 在编译时需要依赖其他的repo, 为了编译时方便, 在WORKSPACE中定义的是依赖本地的repository, 这样在`bazel build` 时就不会再次从远程仓库下载代码. 因此在编译ease-gateway 之前需要将依赖的repositories下载到本地, 并放在与ease-gateway的同级目录

目前依赖本地的repo有：

- gateway-proto [https://github.com/binchencoder/gateway-proto]
- letsgo [https://github.com/binchencoder/letsgo.git]
- skylb-api [https://github.com/binchencoder/skylb-api.git]

```
git clone https://github.com/binchencoder/gateway-proto.git
git clone https://github.com/binchencoder/letsgo.git
git clone https://github.com/binchencoder/skylb-api.git
```

## Clone code

```
git clone https://github.com/binchencoder/ease-gateway.git

git clone https://github.com/binchencoder/gateway-proto.git
git clone https://github.com/binchencoder/letsgo.git
git clone https://github.com/binchencoder/skylb-api.git
```
> 也可以连同submodules一起clone到本地
>
> git clone --recurse-submodules https://github.com/binchencoder/ease-gateway.git

## Sync submodule

```
git submodule init
git submodule update
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
