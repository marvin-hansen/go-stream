load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "enums.go",
        "functions.go",
        "interface.go",
        "messages.go",
    ],
    importpath = "go-stream/api/types",
    visibility = ["//visibility:public"],
    deps = ["@com_github_shopspring_decimal//:go_default_library"],
)
