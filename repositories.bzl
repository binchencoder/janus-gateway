load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    # go_repository(
    #     name = "com_github_binchencoder_letsgo",
    #     importpath = "binchencoder.com/letsgo",
    #     urls = [
    #         "https://codeload.github.com/binchencoder/letsgo/tar.gz/3a34eef5d1546b1be444e2e053d425e80afe100e",
    #     ],
    #     strip_prefix = "letsgo-3a34eef5d1546b1be444e2e053d425e80afe100e",
    #     type = "tar.gz",
    #     # gazelle args: -go_prefix binchencoder.com/letsgo
    # )
    # go_repository(
    #     name = "com_github_binchencoder_skylb_api",
    #     importpath = "binchencoder.com/skylb-api",
    #     urls = [
    #         "https://codeload.github.com/binchencoder/skylb-api/tar.gz/43a2566186d2411255f6818afce1cb5639cf42c5",
    #     ],
    #     strip_prefix = "skylb-api-43a2566186d2411255f6818afce1cb5639cf42c5",
    #     type = "tar.gz",
    #     # gazelle args: -go_prefix binchencoder.com/skylb-api
    # )
    # go_repository(
    #     name = "com_github_binchencoder_gateway_proto",
    #     importpath = "binchencoder.com/gateway-proto",
    #     urls = [
    #         "https://codeload.github.com/binchencoder/gateway-proto/tar.gz/c099a5a6646c572557bc8326f4d952fba4165a3b",
    #     ],
    #     strip_prefix = "gateway-proto-c099a5a6646c572557bc8326f4d952fba4165a3b",
    #     type = "tar.gz",
    #     # gazelle args: -go_prefix binchencoder.com/gateway-proto
    # )

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
        name = "com_github_grpc_ecosystem_grpc_opentracing",
        importpath = "github.com/grpc-ecosystem/grpc-opentracing",
        urls = [
            "https://codeload.github.com/grpc-ecosystem/grpc-opentracing/tar.gz/8e809c8a86450a29b90dcc9efbf062d0fe6d9746",
        ],
        strip_prefix = "grpc-opentracing-8e809c8a86450a29b90dcc9efbf062d0fe6d9746",
        type = "tar.gz",
    )
    go_repository(
        name = "com_github_grpc_ecosystem_go_grpc_middleware",
        importpath = "github.com/grpc-ecosystem/go-grpc-middleware",
        urls = [
            "https://codeload.github.com/grpc-ecosystem/go-grpc-middleware/tar.gz/e0797f438f94f4d032395b8f71aae0e73d6efa08",
        ],
        strip_prefix = "go-grpc-middleware-e0797f438f94f4d032395b8f71aae0e73d6efa08",
        type = "tar.gz",
    )
    go_repository(
        name = "com_github_beorn7_perks",
        importpath = "github.com/beorn7/perks",
        urls = ["https://codeload.github.com/beorn7/perks/tar.gz/37c8de3658fcb183f997c4e13e8337516ab753e6"],
        strip_prefix = "perks-37c8de3658fcb183f997c4e13e8337516ab753e6",
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
        name = "com_github_fatih_color",
        importpath = "github.com/fatih/color",
        commit = "3f9d52f7176a6927daacff70a3e8d1dc2025c53e",
    )
    go_repository(
        name = "com_github_ghodss_yaml",
        importpath = "github.com/ghodss/yaml",
        sum = "h1:wQHKEahhL6wmXdzwWG11gIVCkOv05bNOh+Rxn0yngAk=",
        version = "v1.0.0",
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
        name = "com_github_klauspost_cpuid",
        importpath = "github.com/klauspost/cpuid",
        commit = "cf2ded4aa83357cbc18366e6e1c64eda3c26eaf2",
    )
    go_repository(
        name = "com_github_pborman_uuid",
        importpath = "github.com/pborman/uuid",
        commit = "8b1b92947f46224e3b97bb1a3a5b0382be00d31e",
        # gazelle args: -go_prefix github.com/pborman/uuid
    )
    go_repository(
        name = "com_github_go_kit_kit",
        importpath = "github.com/go-kit/kit",
        urls = ["https://codeload.github.com/go-kit/kit/tar.gz/dc489b75b9cdbf29c739534c2aa777cabb034954"],
        strip_prefix = "kit-dc489b75b9cdbf29c739534c2aa777cabb034954",
        type = "tar.gz",
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        commit = "27936f6d90f9c8e1145f11ed52ffffbfdb9e0af7",
    )
    go_repository(
        name = "com_github_vividcortex_gohistogram",
        importpath = "github.com/VividCortex/gohistogram",
        commit = "51564d9861991fb0ad0f531c99ef602d0f9866e6",
    )
    go_repository(
        name = "com_github_klauspost_compress",
        importpath = "github.com/klauspost/compress",
        urls = ["https://codeload.github.com/klauspost/compress/tar.gz/f82c96c236f2249d76676da0d91e798e619acb35"],
        strip_prefix = "compress-f82c96c236f2249d76676da0d91e798e619acb35",
        type = "tar.gz",
    )
    go_repository(
        name = "com_github_opentracing_opentracing_go",
        importpath = "github.com/opentracing/opentracing-go",
        urls = [
            "https://codeload.github.com/opentracing/opentracing-go/tar.gz/135aa78c6f95b4a199daf2f0470d231136cbbd0c",
        ],
        strip_prefix = "opentracing-go-135aa78c6f95b4a199daf2f0470d231136cbbd0c",
        type = "tar.gz",
        # gazelle args: -go_prefix github.com/opentracing/opentracing-go
    )
    go_repository(
        name = "com_github_matttproud_golang_protobuf_extension",
        importpath = "github.com/matttproud/golang_protobuf_extensions",
        commit = "c182affec369e30f25d3eb8cd8a478dee585ae7d",
    )
    go_repository(
        name = "com_github_stretchr_testify",
        importpath = "github.com/stretchr/testify",
        commit = "221dbe5ed46703ee255b1da0dec05086f5035f62",
    )
    go_repository(
        name = "com_github_soheilhy_cmux",
        importpath = "github.com/soheilhy/cmux",
        commit = "8a8ea3c53959009183d7914522833c1ed8835020",
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
        name = "com_github_prometheus_client_golang",
        importpath = "github.com/prometheus/client_golang",
        urls = [
            "https://codeload.github.com/prometheus/client_golang/tar.gz/b12dd9c58c3d7ce96f9e1ede31d02f6df3d50c61",
        ],
        strip_prefix = "client_golang-b12dd9c58c3d7ce96f9e1ede31d02f6df3d50c61",
        type = "tar.gz",
        # gazelle args: -go_prefix github.com/prometheus/client_golang
    )
    go_repository(
        name = "com_github_prometheus_client_model",
        importpath = "github.com/prometheus/client_model",
        urls = [
            "https://codeload.github.com/prometheus/client_model/tar.gz/fd36f4220a901265f90734c3183c5f0c91daa0b8",
        ],
        strip_prefix = "client_model-fd36f4220a901265f90734c3183c5f0c91daa0b8",
        type = "tar.gz",
        # gazelle args: -go_prefix github.com/prometheus/client_model
    )
    go_repository(
        name = "com_github_prometheus_common",
        importpath = "github.com/prometheus/common",
        urls = [
            "https://codeload.github.com/prometheus/common/tar.gz/637d7c34db122e2d1a25d061423098663758d2d3",
        ],
        strip_prefix = "common-637d7c34db122e2d1a25d061423098663758d2d3",
        type = "tar.gz",
    )
    go_repository(
        name = "com_github_prometheus_procfs",
        importpath = "github.com/prometheus/procfs",
        urls = [
            "https://codeload.github.com/prometheus/procfs/tar.gz/6df11039f8de6804bb01c0ebd52cde9c26091e1c",
        ],
        strip_prefix = "procfs-6df11039f8de6804bb01c0ebd52cde9c26091e1c",
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
        name = "org_golang_x_oauth2",
        importpath = "golang.org/x/oauth2",
        urls = [
            "https://codeload.github.com/golang/oauth2/tar.gz/0f29369cfe4552d0e4bcddc57cc75f4d7e672a33",
        ],
        strip_prefix = "oauth2-0f29369cfe4552d0e4bcddc57cc75f4d7e672a33",
        type = "tar.gz",
        # gazelle args: -go_prefix golang.org/x/oauth2
    )