# grpc-ecosystem/grpc-gateway

grpc-ecosystem/grpc-gateway: grpc-gateway is a plugin of protoc. It reads
gRPC service definition, and generates a reverse-proxy server which translates
a RESTful JSON API into gRPC. This server is generated according to custom
options in your gRPC definition.

Hosted on https://github.com/grpc-ecosystem/grpc-gateway.

## Update History

### 2020/05/25

Based on commit ID [58a91c2b3020ee25e5d3cf49cf7d342544ab773d](https://github.com/grpc-ecosystem/grpc-gateway/commit/7988867d3206f7e90ca3e8d4ab67e060734a297a)

Commit message

```
commit 7988867d3206f7e90ca3e8d4ab67e060734a297a (HEAD, tag: v2.0.0-beta.2, tag: v1.14.6)
Author: Johan Brandhorst <johan.brandhorst@gmail.com>
Date:   Mon May 25 11:42:27 2020 +0100

    Generate changelog for 1.14.6
    
    Fixed the issue where it would include v2 merges
```

### 2019/11/03

Based on commit ID 58a91c2b3020ee25e5d3cf49cf7d342544ab773d

Commit message

```
commit 58a91c2b3020ee25e5d3cf49cf7d342544ab773d (HEAD -> master, origin/master, origin/HEAD)
Author: Prateek Malhotra <someone1@gmail.com>
Date:   Fri Nov 1 10:32:09 2019 -0400

    annotations: Sort import order.
```

### 2019/08/09

Based on V1.9.5
https://github.com/grpc-ecosystem/grpc-gateway/releases/tag/v1.9.5

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
Pull message
```
来自 https://github.com/grpc-ecosystem/grpc-gateway
   fdf0635..b6e6efb  master     -> origin/master
更新 fdf0635..b6e6efb
Fast-forward
 docs/_docs/customizingyourgateway.md           | 22 ++++++++++++++++++++++
 protoc-gen-swagger/genswagger/BUILD.bazel      |  1 +
 protoc-gen-swagger/genswagger/template.go      | 34 +++++++++++++++++++++++++++++-----
 protoc-gen-swagger/genswagger/template_test.go | 61 +++++++++++++++++++++++++++++++++++++++++++++++++++++++++---
```