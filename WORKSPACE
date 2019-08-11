workspace(name = "binchencoder_ease_gateway")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

# ---------- io_bazel_rules_go ----------
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "6776d68ebb897625dead17ae510eac3d5f6342367327875210df44dbe2aeeb19",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.17.1/rules_go-0.17.1.tar.gz"],
)
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
go_rules_dependencies()
go_register_toolchains()

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
http_archive(
    name = "bazel_gazelle",
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
)
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
gazelle_dependencies()

# Use gazelle to declare Go dependencies in Bazel.
# gazelle:repository_macro repositories.bzl%go_repositories
load("//:repositories.bzl", "go_repositories")
go_repositories()

# ---------- com_google_protobuf ----------
# git_repository(
#     name = "com_google_protobuf",
#     commit = "c132a4aa165d8ce2b65af62d4bde4a7ce08d07c3",
#     remote = "https://gitee.com/binchencoder/protobuf",
#     shallow_since = "1558721209 -0700",
# )
# load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
# protobuf_deps()

# ---------- com_github_bazelbuild_buildtools ----------
# go_repository(
#     name = "com_github_bazelbuild_buildtools",
#     importpath = "github.com/bazelbuild/buildtools",
#     commit = "36bd730dfa67bff4998fe897ee4bbb529cc9fbee",
# )
git_repository(
    name = "com_github_bazelbuild_buildtools",
    commit = "680ef5165d2bf75d2e2fab17b5a87ce19767aaa6",
    remote = "https://gitee.com/binchencoder/buildtools",
    shallow_since = "1558721209 -0700",
)
load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")
buildifier_dependencies()
