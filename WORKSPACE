workspace(name = "com_github_binchencoder_ease_gateway")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# ----------从github下载扩展 io_bazel_rules_go ----------
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8df59f11fb697743cbb3f26cfb8750395f30471e9eabde0d174c3aebc7a1cd39",
    urls = [
        "https://github.com/bazelbuild/rules_go/releases/download/0.19.1/rules_go-0.19.1.tar.gz",
    ],
)

# 从下载的扩展里载入 go_rules_dependencies go_register_toolchains 函数
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

# ---------- bazel_gazelle ----------
# 一般来说都会使用gazelle工具来自动生成 BUILD 文件, 而不是手写.
http_archive(
    name = "bazel_gazelle",
    sha256 = "be9296bfd64882e3c08e3283c58fcb461fa6dd3c171764fcc4cf322f60615a9b",
    urls = [
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.18.1/bazel-gazelle-0.18.1.tar.gz",
    ],
)

# 从gazelle中加载gazelle_dependencies
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

# Overriding dependencies go_rules_dependencies
http_archive(
    name = "go_googleapis",
    patch_args = [
        "-E",
        "-p1",
    ],
    patches = [
        "@io_bazel_rules_go//third_party:go_googleapis-deletebuild.patch",
        "@io_bazel_rules_go//third_party:go_googleapis-directives.patch",
        "@io_bazel_rules_go//third_party:go_googleapis-gazelle.patch",
        "@io_bazel_rules_go//third_party:go_googleapis-fix.patch",
    ],
    strip_prefix = "googleapis-b4c73face84fefb967ef6c72f0eae64faf67895f",
    type = "tar.gz",
    urls = [
        "https://codeload.github.com/googleapis/googleapis/tar.gz/b4c73face84fefb967ef6c72f0eae64faf67895f",
    ],
    # gazelle args: -go_prefix google.golang.org/genproto/googleapi -proto disable
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    strip_prefix = "net-16b79f2e4e95ea23b2bf9903c9809ff7b013ce85",
    type = "tar.gz",
    urls = [
        "https://codeload.github.com/golang/net/tar.gz/16b79f2e4e95ea23b2bf9903c9809ff7b013ce85",  # master, as of 2019-03-3
    ],
    # gazelle args: -go_prefix golang.org/x/net -proto disable
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    strip_prefix = "sys-fde4db37ae7ad8191b03d30d27f258b5291ae4e3",
    type = "tar.gz",
    urls = [
        "https://codeload.github.com/golang/sys/tar.gz/fde4db37ae7ad8191b03d30d27f258b5291ae4e3",
    ],
    # gazelle args: -go_prefix golang.org/x/sys
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    strip_prefix = "text-f21a4dfb5e38f5895301dc265a8def02365cc3d0",
    type = "tar.gz",
    urls = [
        "https://codeload.github.com/golang/text/tar.gz/f21a4dfb5e38f5895301dc265a8def02365cc3d0",  # v0.3.0, latest as of 2019-03-03
    ],
    # gazelle args: -go_prefix golang.org/x/text -proto disable
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    patch_args = ["-p1"],
    patches = [
        "@io_bazel_rules_go//third_party:org_golang_x_tools-extras.patch",
    ],
    strip_prefix = "tools-c8855242db9c1762032abe33c2dff50de3ec9d05",
    type = "tar.gz",
    urls = [
        "https://codeload.github.com/golang/tools/tar.gz/c8855242db9c1762032abe33c2dff50de3ec9d05",
    ],
    # gazelle args: -go_prefix golang.org/x/tools -proto disable
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

# go_repository(
#     name = "com_google_protobuf",
#     importpath = "github.com/protocolbuffers/protobuf",
#     urls = [
#         "https://codeload.github.com/protocolbuffers/protobuf/tar.gz/09745575a923640154bcf307fba8aedff47f240a",
#     ],
#     strip_prefix = "protobuf-09745575a923640154bcf307fba8aedff47f240a",
#     type = "tar.gz",
# )
local_repository(
    name = "com_google_protobuf",
    path = "third_party/protobuf",
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

# ---------- com_github_bazelbuild_buildtools ----------
http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "86592d703ecbe0c5cbb5139333a63268cf58d7efd2c459c8be8e69e77d135e29",
    strip_prefix = "buildtools-0.26.0",
    urls = ["https://github.com/bazelbuild/buildtools/archive/0.26.0.tar.gz"],
)

load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")

buildifier_dependencies()

# ---------- local repositories
local_repository(
    name = "com_github_binchencoder_gateway_proto",
    path = "../gateway-proto",
)

local_repository(
    name = "com_github_binchencoder_letsgo",
    path = "../letsgo",
)

local_repository(
    name = "com_github_binchencoder_skylb_api",
    path = "../skylb-api",
)
