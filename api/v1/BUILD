load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "dbg_utils.go",
        "endpoints.go",
        "errors.go",
        "sdk.go",
        "sdk_utils.go",
        "set_methods.go",
        "websocket.go",
        "ws_connection.go",
        "ws_handler.go",
    ],
    importpath = "go-stream/api/v1",
    visibility = ["//visibility:public"],
    deps = [
        "//api/types:go_default_library",
        "//api/web_socket:go_default_library",
        "@com_github_bitly_go_simplejson//:go_default_library",
    ],
)
