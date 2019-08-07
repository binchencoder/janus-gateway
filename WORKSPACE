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
http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "e513c0ac6534810eb7a14bf025a0f159726753f97f74ab7863c650d26e01d677",
    strip_prefix = "rules_docker-0.9.0",
    urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.9.0.tar.gz"],
)
load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)
container_repositories()

# ---------- bazel_gazelle ----------
http_archive(
    name = "bazel_gazelle",
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
)
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
gazelle_dependencies()

# ---------- com_github_bazelbuild_buildtools ----------
go_repository(
    name = "com_github_bazelbuild_buildtools",
    importpath = "github.com/bazelbuild/buildtools",
    commit = "36bd730dfa67bff4998fe897ee4bbb529cc9fbee",
)
load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")
buildifier_dependencies()

# ---------- com_google_protobuf ----------
# TODO(chenbin) 2019/08/07
git_repository(
    name = "com_google_protobuf",
    commit = "09745575a923640154bcf307fba8aedff47f240a",
    remote = "https://github.com/protocolbuffers/protobuf",
    shallow_since = "1558721209 -0700",
)
#load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")
#protobuf_deps()

go_repository(
    name = "binchencoder_third_party_java",
    commit = "dcac035f578caefefc6cd12a799cbb400a09f004",
    importpath = "github.com/binchencoder/third-party-java",
)

go_repository(
    name = "grpc_ecosystem_grpc_gateway",
    commit = "ad529a448ba494a88058f9e5be0988713174ac86",
    importpath = "github.com/grpc-ecosystem/grpc-gateway",
)

go_repository(
    name = "com_github_fatih_color",
    commit = "3f9d52f7176a6927daacff70a3e8d1dc2025c53e",
    importpath = "github.com/fatih/color",
)

go_repository(
    name = "com_github_klauspost_compress",
    commit = "ae52aff18558bd92cbe681549bfe9e8cbffd5903",
    importpath = "github.com/klauspost/compress",
)

go_repository(
    name = "com_github_klauspost_cpuid",
    commit = "05a8198c0f5a27739aec358908d7e12c64ce6eb7",
    importpath = "github.com/klauspost/cpuid",
)

go_repository(
    name = "com_github_golang_net",
    commit = "4829fb13d2c62012c17688fa7f629f371014946d",
    importpath = "github.com/golang/net",
)

# Also define in Gopkg.toml
go_repository(
    name = "org_golang_google_genproto",
    commit = "383e8b2c3b9e36c4076b235b32537292176bae20",
    importpath = "google.golang.org/genproto",
)

# Also define in Gopkg.toml
go_repository(
    name = "com_github_rogpeppe_fastuuid",
    commit = "6724a57986aff9bff1a1770e9347036def7c89f6",
    importpath = "github.com/rogpeppe/fastuuid",
)

# Also define in Gopkg.toml
go_repository(
    name = "in_gopkg_resty_v1",
    commit = "fa5875c0caa5c260ab78acec5a244215a730247f",
    importpath = "gopkg.in/resty.v1",
)

# Also define in Gopkg.toml
go_repository(
    name = "com_github_ghodss_yaml",
    commit = "0ca9ea5df5451ffdf184b4428c902747c2c11cd7",
    importpath = "github.com/ghodss/yaml",
)

# Also define in Gopkg.toml
go_repository(
    name = "in_gopkg_yaml_v2",
    commit = "eb3733d160e74a9c7e442f435eb3bea458e1d19f",
    importpath = "gopkg.in/yaml.v2",
)
