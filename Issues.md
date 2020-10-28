# Issues

## 依赖版本问题

1. org_golang_google_grpc
	```go
   go_repository(
       name = "org_golang_google_grpc",
       importpath = "google.golang.org/grpc",
       sum = "h1:zvIju4sqAGvwKspUQOhwnpcqSbzi7/H6QomNNjTL4sk=",
       version = "v1.27.1",
   )
   ```

   ```go
   go_repository(
       name = "org_golang_google_grpc",
       importpath = "google.golang.org/grpc",
       sum = "h1:DGeFlSan2f+WEtCERJ4J9GJWk15TxUi8QGagfI87Xyc=",
       version = "v1.33.1",
   )
   ```

   org_golang_google_grpc 最新版本已经升级到v1.33.1，编译会出现如下error:

   ```verilog
   chenbin@chenbin-ThinkPad:~/.../github-workspace/ease-gateway$ bazel build gateway/...
   ERROR: /home/chenbin/.cache/bazel/_bazel_chenbin/95d98bab223e52f58e53a4599e22df3c/external/com_github_binchencoder_skylb_api/balancer/BUILD:5:1: no such package '@org_golang_google_grpc//naming': BUILD file not found in directory 'naming' of external repository @org_golang_google_grpc. Add a BUILD file to a directory to mark it as a package. and referenced by '@com_github_binchencoder_skylb_api//balancer:go_default_library'
   ERROR: Analysis of target '//gateway/runtime:go_default_test' failed; build aborted: no such package '@org_golang_google_grpc//naming': BUILD file not found in directory 'naming' of external repository @org_golang_google_grpc. Add a BUILD file to a directory to mark it as a package.
   INFO: Elapsed time: 3.781s
   INFO: 0 processes.
   FAILED: Build did NOT complete successfully (43 packages loaded, 588 targets configured)
   ```

2. org_golang_google_grpc_cmd_protoc_gen_go_grpc

   ```go
   go_repository(
       name = "org_golang_google_grpc_cmd_protoc_gen_go_grpc",
       importpath = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
       sum = "h1:KNluVV5ay+orsSPJ6XTpwJQ8qBhrBkOTmtBFGeDlBcY=",
       version = "v0.0.0-20200527211525-6c9e30c09db2",
   )
   ```

   ```go
   go_repository(
       name = "org_golang_google_grpc_cmd_protoc_gen_go_grpc",
       importpath = "google.golang.org/grpc/cmd/protoc-gen-go-grpc",
       sum = "h1:lQ+dE99pFsb8osbJB3oRfE5eW4Hx6a/lZQr8Jh+eoT4=",
       version = "v1.0.0",
   )
   ```

   github.com/grpc-ecosystem/grpc-gateway 使用 org_golang_google_grpc_cmd_protoc_gen_go_grpc的版本是v1.0.0，ease-gateway 升级会出现如下error:

   ```verilog
   chenbin@chenbin-ThinkPad:~/.../github-workspace/ease-gateway$ bazel build gateway/...
   INFO: Analyzed 32 targets (103 packages loaded, 1391 targets configured).
   INFO: Found 32 targets...
   ERROR: /home/chenbin/.cache/bazel/_bazel_chenbin/95d98bab223e52f58e53a4599e22df3c/external/com_github_grpc_ecosystem_grpc_gateway/runtime/internal/examplepb/BUILD.bazel:39:1: GoCompilePkg external/com_github_grpc_ecosystem_grpc_gateway/runtime/internal/examplepb/go_default_library.a failed (Exit 1) builder failed: error executing command bazel-out/host/bin/external/go_sdk/builder compilepkg -sdk external/go_sdk -installsuffix linux_amd64 -src ... (remaining 67 argument(s) skipped)
   
   Use --sandbox_debug to see verbose messages from the sandbox
   /home/chenbin/.cache/bazel/_bazel_chenbin/95d98bab223e52f58e53a4599e22df3c/sandbox/linux-sandbox/845/execroot/com_github_binchencoder_ease_gateway/bazel-out/k8-fastbuild/bin/external/com_github_grpc_ecosystem_grpc_gateway/runtime/internal/examplepb/examplepb_go_proto_/github.com/grpc-ecosystem/grpc-gateway/v2/runtime/internal/examplepb/non_standard_names_grpc.pb.go:14:11: undefined: grpc.SupportPackageIsVersion7
   compilepkg: error running subcommand external/go_sdk/pkg/tool/linux_amd64/compile: exit status 2
   INFO: Elapsed time: 11.320s, Critical Path: 1.12s
   INFO: 5 processes: 5 linux-sandbox.
   FAILED: Build did NOT complete successfully
   ```

   