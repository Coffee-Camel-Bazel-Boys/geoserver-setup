# GoLang rewrite
load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "mapControllerGo",
    srcs = [
      "MapController.go",
      "//:goSecrets",
      "Repo.go"  
    ],
    importpath = "mapController",
    deps = [
      "@com_github_pq//:go_default_library"
    ],
    visibility = ["//visibility:public"],
)