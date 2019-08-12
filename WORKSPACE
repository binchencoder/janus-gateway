workspace(name = "binchencoder_ease_gateway")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# ----------从github下载扩展 io_bazel_rules_go ----------
http_archive(
    name = "io_bazel_rules_go",
    urls = [
        "https://github.com/bazelbuild/rules_go/releases/download/0.17.8/rules_go-0.17.8.tar.gz",
    ],
    sha256 = "38113392bac83252d2e6450b0056e41f35b2469903e319688883598ce38f0377",
)
# 从下载的扩展里载入 go_rules_dependencies go_register_toolchains 函数
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

# ---------- io_bazel_rules_docker ----------
# Download the rules_docker repository at release v0.9.0
# http_archive(
#     name = "io_bazel_rules_docker",
#     sha256 = "e513c0ac6534810eb7a14bf025a0f159726753f97f74ab7863c650d26e01d677",
#     strip_prefix = "rules_docker-0.9.0",
#     urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.9.0.tar.gz"],
# )
# load(
#     "@io_bazel_rules_docker//repositories:repositories.bzl",
#     container_repositories = "repositories",
# )
# container_repositories()

# ---------- bazel_gazelle ----------
# 一般来说都会使用gazelle工具来自动生成 BUILD 文件, 而不是手写.
http_archive(
    name = "bazel_gazelle",
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
)
# 从gazelle中加载gazelle_dependencies
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

# ---------- com_github_bazelbuild_buildtools ----------
http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "86592d703ecbe0c5cbb5139333a63268cf58d7efd2c459c8be8e69e77d135e29",
    strip_prefix = "buildtools-0.26.0",
    urls = ["https://github.com/bazelbuild/buildtools/archive/0.26.0.tar.gz"],
)
load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")
buildifier_dependencies()

# Overriding dependencies go_rules_dependencies
http_archive(
    name = "com_github_golang_protobuf",
    urls = [
        "https://github.com/golang/protobuf/archive/v1.2.0.tar.gz"
    ],
    strip_prefix = "protobuf-1.2.0",
    patches = [
        "@io_bazel_rules_go//third_party:com_github_golang_protobuf-gazelle.patch",
        "@io_bazel_rules_go//third_party:com_github_golang_protobuf-extras.patch",
    ],
    patch_args = ["-p1"],
    # gazelle args: -go_prefix github.com/golang/protobuf -proto disable_global
)
http_archive(
    name = "org_golang_google_grpc",
    urls = [
        "https://codeload.github.com/grpc/grpc-go/tar.gz/ee87494b1f58190a421bb41cce5ccbe8e833c04b",
    ],
    strip_prefix = "grpc-go-ee87494b1f58190a421bb41cce5ccbe8e833c04b",
    type = "tar.gz",
    patches = [
        "@io_bazel_rules_go//third_party:org_golang_google_grpc-gazelle.patch",
        "@io_bazel_rules_go//third_party:org_golang_google_grpc-crosscompile.patch",
    ],
    patch_args = ["-p1"],
    # gazelle args: -go_prefix google.golang.org/grpc -proto disable
)
http_archive(
    name = "org_golang_x_sys",
    urls = [
        "https://codeload.github.com/golang/sys/tar.gz/2be51725563103c17124a318f1745b66f2347acb",
    ],
    strip_prefix = "sys-2be51725563103c17124a318f1745b66f2347acb",
    type = "tar.gz",
    patches = ["@io_bazel_rules_go//third_party:org_golang_x_sys-gazelle.patch"],
    patch_args = ["-p1"],
    # gazelle args: -go_prefix golang.org/x/sys
)
http_archive(
    name = "org_golang_x_text",
    urls = [
        "https://codeload.github.com/golang/text/tar.gz/f21a4dfb5e38f5895301dc265a8def02365cc3d0",
    ],
    strip_prefix = "text-f21a4dfb5e38f5895301dc265a8def02365cc3d0",
    type = "tar.gz",
    patches = ["@io_bazel_rules_go//third_party:org_golang_x_text-gazelle.patch"],
    patch_args = ["-p1"],
    # gazelle args: -go_prefix golang.org/x/text
)

# 注册一堆常用依赖 如github.com/google/protobuf golang.org/x/net
go_rules_dependencies()
# 下载golang工具链
go_register_toolchains()
# 加载gazelle依赖
gazelle_dependencies()

# Use gazelle to declare Go dependencies in Bazel.
# gazelle:repository_macro repositories.bzl%go_repositories
load("//:repositories.bzl", "go_repositories")
go_repositories()