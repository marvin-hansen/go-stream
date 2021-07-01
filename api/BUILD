load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "sdk.go",
    ],
    importpath = "go-stream/api",
    visibility = ["//visibility:public"],
    deps = [
        "//api/types:go_default_library",
        "//api/v1:go_default_library",
    ],
)
