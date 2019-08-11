package(default_visibility = ["//visibility:public"])

load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["reporter.go"],
    importpath = "github.com/binchencoder/ease-gateway/integrate/metrics",
    deps = [
        "@binchencoder_letsgo//time:go_default_library",
        "@binchencoder_third_party_go//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
    ],
)
