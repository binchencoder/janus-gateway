load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "binchencoder_third_party_go",
        importpath = "github.com/binchencoder/third-party-go",
        urls = [
            "https://codeload.github.com/binchencoder/third-party-go/tar.gz/4e9c6ce6b9edd7289966dda9be983f12a063584c",
        ],
        strip_prefix = "third-party-go-4e9c6ce6b9edd7289966dda9be983f12a063584c",
        type = "tar.gz",
    )
    go_repository(
        name = "binchencoder_letsgo",
        importpath = "github.com/binchencoder/letsgo",
        urls = [
            "https://codeload.github.com/binchencoder/letsgo/tar.gz/d43bf202de7e0bd45f50810bad8aa83a5813c941",
        ],
        strip_prefix = "letsgo-d43bf202de7e0bd45f50810bad8aa83a5813c941",
        type = "tar.gz",
    )
    go_repository(
        name = "binchencoder_skylb_api",
        importpath = "github.com/binchencoder/skylb-api",
        urls = [
            "https://codeload.github.com/binchencoder/skylb-api/tar.gz/f6b037bb0a48844b2624b561feb88ee5ff223e17",
        ],
        strip_prefix = "skylb-api-f6b037bb0a48844b2624b561feb88ee5ff223e17",
        type = "tar.gz",
    )

    go_repository(
        name = "grpc_ecosystem_grpc_gateway",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
        urls = [
            "https://codeload.github.com/grpc-ecosystem/grpc-gateway/tar.gz/fdf063599d922ec89a70819e2d5b7b4b5c642b92",
        ],
        strip_prefix = "grpc-gateway-fdf063599d922ec89a70819e2d5b7b4b5c642b92",
        type = "tar.gz",
    )