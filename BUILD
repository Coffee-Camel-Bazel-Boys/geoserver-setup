filegroup(
    name = "secretFile",
    srcs = 
      select({
        ":prod": ["Secrets.prod.go"],
        "//conditions:default": ["Secrets.example.go"],
      })
)

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
    name = "server",
    srcs = ["Server.go"],
    deps = [
        "//src:mapControllerGo",
        "@com_github_cors//:go_default_library"
    ],  
)