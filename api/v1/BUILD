load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "endpoints.go",
        "errors.go",
        "sdk.go",
        "setters.go",
    ],
    importpath = "go-stream/api/v1",
    visibility = ["//visibility:public"],
    deps = ["//api/types:go_default_library"],
)
