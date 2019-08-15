load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "com_github_binchencoder_letsgo",
        importpath = "github.com/binchencoder/letsgo",
        urls = [
            "https://codeload.github.com/binchencoder/letsgo/tar.gz/4aa9d379feec705d2d0d168c3c2266bed4521fb4",
        ],
        strip_prefix = "letsgo-4aa9d379feec705d2d0d168c3c2266bed4521fb4",
        type = "tar.gz",
    )
    go_repository(
        name = "com_github_binchencoder_skylb_api",
        importpath = "github.com/binchencoder/skylb-api",
        urls = [
            "https://codeload.github.com/binchencoder/skylb-api/tar.gz/f8d8abf7d25490abaeecf5ca1a57c96ab6025237",
        ],
        strip_prefix = "skylb-api-f8d8abf7d25490abaeecf5ca1a57c96ab6025237",
        type = "tar.gz",
    )
    go_repository(
        name = "com_github_binchencoder_gateway_proto",
        importpath = "github.com/binchencoder/gateway-proto",
        commit = "1ee4b0a8951fda57f986695253374d7847adbec6",
    )

    go_repository(
        name = "com_github_grpc_ecosystem_grpc_gateway",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
        urls = [
            "https://codeload.github.com/grpc-ecosystem/grpc-gateway/tar.gz/fdf063599d922ec89a70819e2d5b7b4b5c642b92",
        ],
        strip_prefix = "grpc-gateway-fdf063599d922ec89a70819e2d5b7b4b5c642b92",
        type = "tar.gz",
    )
    go_repository(
        name = "com_github_cenkalti_backoff",
        importpath = "github.com/cenkalti/backoff",
        urls = ["https://github.com/cenkalti/backoff/archive/v2.2.1.tar.gz"],
        strip_prefix = "backoff-2.2.1",
        type = "tar.gz",
    )
    go_repository(
        name = "in_gopkg_yaml_v2",
        importpath = "gopkg.in/yaml.v2",
        replace = "github.com/go-yaml/yaml",
        sum = "h1:eZqMvILvSB6AhTa+FGXHupLRXfU8SFxBP4IW1wetpT4=",
        version = "v2.0.0-20170812160011-eb3733d160e7",
        # gazelle args: -go-prefix gopkg.in/yaml.v2
    )
    go_repository(
        name = "org_golang_google_grpc",
        importpath = "google.golang.org/grpc",
        urls = [
            "https://codeload.github.com/grpc/grpc-go/tar.gz/df014850f6dee74ba2fc94874043a9f3f75fbfd8",
        ],
        strip_prefix = "grpc-go-df014850f6dee74ba2fc94874043a9f3f75fbfd8", # v1.17.0, latest as of 2019-01-15
        type = "tar.gz",
        # gazelle args: -go_prefix google.golang.org/grpc -proto disable
    )
    go_repository(
        name = "com_github_golang_glog",
        importpath = "github.com/golang/glog",
        sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
        version = "v0.0.0-20160126235308-23def4e6c14b",
    )
    go_repository(
        name = "com_github_google_uuid",
        importpath = "github.com/google/uuid",
        commit = "c2e93f3ae59f2904160ceaab466009f965df46d6",
        # gazelle args: -go_prefix github.com/google/uuid
    )
    go_repository(
        name = "com_github_pborman_uuid",
        importpath = "github.com/pborman/uuid",
        commit = "8b1b92947f46224e3b97bb1a3a5b0382be00d31e",
        # gazelle args: -go_prefix github.com/pborman/uuid
    )
    go_repository(
        name = "com_github_klauspost_compress",
        importpath = "github.com/klauspost/compress",
        commit = "f82c96c236f2249d76676da0d91e798e619acb35",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        commit = "221dbe5ed46703ee255b1da0dec05086f5035f62",
    )
    go_repository(
        name = "com_github_uber_jaeger_client_go",
        importpath = "github.com/uber/jaeger-client-go",
        urls = [
            "https://codeload.github.com/jaegertracing/jaeger-client-go/tar.gz/d8999ab8c9e71b2d71022f26f21bf39a3c428301",
        ],
        strip_prefix = "jaeger-client-go-d8999ab8c9e71b2d71022f26f21bf39a3c428301",
        type = "tar.gz",
        # gazelle args: -go_prefix github.com/uber/jaeger-client-go
    )
    go_repository(
        name = "com_github_uber_jaeger_lib",
        importpath = "github.com/uber/jaeger-lib",
        urls = [
            "https://codeload.github.com/jaegertracing/jaeger-lib/tar.gz/ec4562394c7d7c18dc238aad0fc921a4325a8b0a",
        ],
        strip_prefix = "jaeger-lib-ec4562394c7d7c18dc238aad0fc921a4325a8b0a",
        type = "tar.gz",
        # gazelle args: -go-prefix github.com/uber/jaeger-lib
    )
    go_repository(
        name = "org_golang_x_sys",
        importpath = "golang.org/x/sys",
        urls = [
            "https://codeload.github.com/golang/sys/tar.gz/fde4db37ae7ad8191b03d30d27f258b5291ae4e3",
        ],
        strip_prefix = "sys-fde4db37ae7ad8191b03d30d27f258b5291ae4e3",
        type = "tar.gz",
        # gazelle args: -go_prefix golang.org/x/sys
    )