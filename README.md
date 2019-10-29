# ease-gateway
Gateway service based on grpc-ecosystem/grpc-gateway

## Prepared

**ease-gateway** `在编译时依赖gateway-proto[https://github.com/binchencoder/gateway-proto], 为了编译方便, 在`bazel build`时会依赖本地的reposiroty. 因此在编译ease-gateway 之前需要将**gateway-proto**下载到本地, 并放在与ease-gateway的同级目录
```
git clone https://github.com/binchencoder/gateway-proto.git
```
