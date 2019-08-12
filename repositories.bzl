load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_repositories():
    go_repository(
        name = "binchencoder_third_party_go",
        commit = "4e9c6ce6b9edd7289966dda9be983f12a063584c",
        importpath = "gitee.com/binchencoder/third-party-go",
    )
    go_repository(
        name = "binchencoder_letsgo",
        commit = "d43bf202de7e0bd45f50810bad8aa83a5813c941",
        importpath = "gitee.com/binchencoder/letsgo",
    )
    go_repository(
        name = "binchencoder_skylb_api",
        commit = "f6b037bb0a48844b2624b561feb88ee5ff223e17",
        importpath = "gitee.com/binchencoder/skylb-api",
    )

    go_repository(
        name = "grpc_ecosystem_grpc_gateway",
        commit = "ad529a448ba494a88058f9e5be0988713174ac86",
        importpath = "github.com/grpc-ecosystem/grpc-gateway",
    )