# grpc-ecosystem/grpc-gateway

grpc-ecosystem/grpc-gateway: grpc-gateway is a plugin of protoc. It reads
gRPC service definition, and generates a reverse-proxy server which translates
a RESTful JSON API into gRPC. This server is generated according to custom
options in your gRPC definition.

Hosted on https://github.com/grpc-ecosystem/grpc-gateway.

## Update History

### 2019/08/09

Based on commit ID fdf063599d922ec89a70819e2d5b7b4b5c642b92

Commit message
```
commit fdf063599d922ec89a70819e2d5b7b4b5c642b92 (HEAD -> master, origin/master, origin/HEAD)
Author: Johan Brandhorst <johan.brandhorst@gmail.com>
Date:   Mon Aug 5 16:08:44 2019 +0100

    Fix release script
    
    Specify an "id" for each build, as this was otherwise inferred
    and duplicate for our two builds.
    
    Also remove the use of the deprecated "archive" instruction
    in favour of "archives".
    
    Fixes #981
```