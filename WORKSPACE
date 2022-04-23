workspace(name = "com_github_binchencoder_janus_gateway")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

# Define before rules_proto, otherwise we receive the version of com_google_protobuf from there
http_archive(
    name = "com_google_protobuf",
    sha256 = "b07772d38ab07e55eca4d50f4b53da2d998bb221575c60a4f81100242d4b4889",
    strip_prefix = "protobuf-3.20.0",
    urls = ["https://github.com/protocolbuffers/protobuf/archive/v3.20.0.tar.gz"],
)

http_archive(
    name = "bazel_skylib",
    sha256 = "f7be3474d42aae265405a592bb7da8e171919d74c16f082a5457840f06054728",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-skylib/releases/download/1.2.1/bazel-skylib-1.2.1.tar.gz",
        "https://github.com/bazelbuild/bazel-skylib/releases/download/1.2.1/bazel-skylib-1.2.1.tar.gz",
    ],
)

load("@bazel_skylib//:workspace.bzl", "bazel_skylib_workspace")

bazel_skylib_workspace()

http_archive(
    name = "rules_proto",
    sha256 = "66bfdf8782796239d3875d37e7de19b1d94301e8972b3cbd2446b332429b4df1",
    strip_prefix = "rules_proto-4.0.0",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0.tar.gz",
        "https://github.com/bazelbuild/rules_proto/archive/refs/tags/4.0.0.tar.gz",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "f2dcd210c7095febe54b804bb1cd3a58fe8435a909db2ec04e31542631cf715c",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.31.0/rules_go-v0.31.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.31.0/rules_go-v0.31.0.zip",
    ],
)

# ---------- bazel_gazelle ----------
# 一般来说都会使用gazelle工具来自动生成 BUILD 文件, 而不是手写.
# http_archive(
#     name = "bazel_gazelle",
#     sha256 = "b85f48fa105c4403326e9525ad2b2cc437babaa6e15a3fc0b1dbab0ab064bc7c",
#     urls = [
#         "https://storage.googleapis.com/bazel-mirror/github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.2/bazel-gazelle-v0.22.2.tar.gz",
#         "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.22.2/bazel-gazelle-v0.22.2.tar.gz",
#     ],
# )

# TODO: Revert https://github.com/grpc-ecosystem/grpc-gateway/pull/2578/commits/fb9b59be7f2408767657c83c5002bf700ac7c460 once
# https://github.com/bazelbuild/bazel-gazelle/pull/1194 is merged
git_repository(
    name = "bazel_gazelle",
    commit = "f377e6eff8e24508feb1a34b1e5e681982482a9f",
    remote = "https://github.com/bazelbuild/bazel-gazelle",
    shallow_since = "1648046534 -0400",
)

# 从下载的扩展里载入 go_rules_dependencies go_register_toolchains 函数
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
# 注册一堆常用依赖 如github.com/google/protobuf golang.org/x/net
go_rules_dependencies()
# 下载golang工具链
go_register_toolchains(version = "1.17.2")

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

# Use gazelle to declare Go dependencies in Bazel.
# gazelle:repository_macro repositories.bzl%go_repositories

load("//:repositories.bzl", "go_repositories")

go_repositories()

# 加载gazelle依赖
# This must be invoked after our explicit dependencies
# See https://github.com/bazelbuild/bazel-gazelle/issues/1115.
gazelle_dependencies()

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

# ---------- com_github_bazelbuild_buildtools ----------
http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "7f43df3cca7bb4ea443b4159edd7a204c8d771890a69a50a190dc9543760ca21",
    strip_prefix = "buildtools-5.0.1",
    urls = ["https://github.com/bazelbuild/buildtools/archive/5.0.1.tar.gz"],
)

load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")

buildifier_dependencies()

# ---------- local repositories
# local_repository(
#     name = "com_github_binchencoder_gateway_proto",
#     path = "../gateway-proto",
# )

# local_repository(
#     name = "com_github_binchencoder_letsgo",
#     path = "../letsgo",
# )

# local_repository(
#     name = "com_github_binchencoder_skylb_api",
#     path = "../skylb-api",
# )
