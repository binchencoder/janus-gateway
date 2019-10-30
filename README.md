# ease-gateway
Gateway service based on grpc-ecosystem/grpc-gateway

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

## Bazel build gateway

```
cd ease-gateway

bazel build cmd/gateway/...
```