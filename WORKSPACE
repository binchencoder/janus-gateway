workspace(name = "com_github_binchencoder_ease_gateway")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# ----------从github下载扩展 io_bazel_rules_go ----------
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "a8d6b1b354d371a646d2f7927319974e0f9e52f73a2452d2b3877118169eb6bb",
    urls = [
        "https://github.com/bazelbuild/rules_go/releases/download/v0.23.3/rules_go-v0.23.3.tar.gz",
    ],
)

# ---------- bazel_gazelle ----------
# 一般来说都会使用gazelle工具来自动生成 BUILD 文件, 而不是手写.
http_archive(
    name = "bazel_gazelle",
    sha256 = "cdb02a887a7187ea4d5a27452311a75ed8637379a1287d8eeb952138ea485f7d",
    urls = [
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.21.1/bazel-gazelle-v0.21.1.tar.gz",
    ],
)

# 从下载的扩展里载入 go_rules_dependencies go_register_toolchains 函数
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

# 注册一堆常用依赖 如github.com/google/protobuf golang.org/x/net
go_rules_dependencies()

# 下载golang工具链
go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

# 加载gazelle依赖
gazelle_dependencies()

# Use gazelle to declare Go dependencies in Bazel.
# gazelle:repository_macro repositories.bzl%go_repositories
load("//:repositories.bzl", "go_repositories")

go_repositories()

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "com_google_protobuf",
    commit = "09745575a923640154bcf307fba8aedff47f240a",
    remote = "https://github.com/protocolbuffers/protobuf",
    shallow_since = "1558721209 -0700",
)

# go_repository(
#     name = "com_google_protobuf",
#     importpath = "github.com/protocolbuffers/protobuf",
#     sum = "h1:pNPOCD+Nm4NY0R6gdOpwOPpRGUjbPo9SO/UlD56lH+0=",
#     version = "v3.8.0+incompatible",
# )

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

# ---------- com_github_bazelbuild_buildtools ----------
http_archive(
    name = "com_github_bazelbuild_buildtools",
    sha256 = "f11fc80da0681a6d64632a850346ed2d4e5cbb0908306d9a2a2915f707048a10",
    strip_prefix = "buildtools-3.3.0",
    urls = ["https://github.com/bazelbuild/buildtools/archive/3.3.0.tar.gz"],
)

load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")

buildifier_dependencies()

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:vGXIOMxbNfDTk/aXCmfdLgkrSV+Z2tcbze+pEc3v5W4=",
    version = "v0.0.0-20200625001655-4c5254603344",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:Ih9Yo4hSPImZOpfGuA4bR/ORKTAbhZo2AbWNRCnevdo=",
    version = "v0.0.0-20200625212154-ddb9806d33ae",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:cokOdA+Jmi5PJGXLlLllQSgYigAEfHXJAERHVMaCc2k=",
    version = "v0.3.3",
)

go_repository(
    name = "org_golang_x_tools",
    importpath = "golang.org/x/tools",
    sum = "h1:/e4fNMHdLn7SQSxTrRZTma2xjQW6ELdxcnpqMhpo9X4=",
    version = "v0.0.0-20200702044944-0cc1aa72b347",
)

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
