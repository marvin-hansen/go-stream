load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "utils.go",
        "websocket.go",
    ],
    importpath = "go-stream/api/web_socket",
    visibility = ["//visibility:public"],
    deps = [
        "//api/types:go_default_library",
        "@com_github_gorilla_websocket//:go_default_library",
    ],
)
