load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
        "messages.go",
        "wsClient.go",
    ],
    importpath = "go-stream/test",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_gorilla_websocket//:go_default_library",
        "@com_github_shopspring_decimal//:go_default_library",
    ],
)

go_binary(
    name = "test",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)
