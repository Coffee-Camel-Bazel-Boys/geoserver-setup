# Config Will replace with docker secret magic
config_setting(
    name = "prod",
    values = {"define": "prod=true"},
)


filegroup(
    name = "goSecretFile",
    srcs = 
      select({
        ":prod": ["Secrets.prod.go"],
        "//conditions:default": ["Secrets.dev.go"],
      })
)

genrule(
    name = "goSecrets",
    srcs = [
      ":goSecretFile"
    ],
    outs = ["Secrets.go"],
    cmd = "cp $(location :goSecretFile) $@",
    visibility = ["//visibility:public"]
)

# ACTUAL GOLANG

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_binary(
    name = "server",
    srcs = ["Server.go"],
    deps = [
        "//src:mapControllerGo",
        "@com_github_cors//:go_default_library"
    ],  
    visibility = ["//visibility:public"],
)

# DOCKER!!!!!

load("@rules_pkg//:pkg.bzl", "pkg_tar")

pkg_tar(
    name = "tar",
    srcs = [":server"],
    mode = "0755",
)

load("@rules_oci//oci:defs.bzl", "oci_image", "oci_tarball")

oci_image(
    name = "image",
    architecture = select({
        "@platforms//cpu:arm64": "arm64",
        "@platforms//cpu:x86_64": "amd64",
    }),
    base = "@distroless_base",
    tars = [":tar.tar"],
    entrypoint = ["/server"],
    os = "linux",
)

oci_tarball(
    name = "geoserver",
    image = ":image",
    repotags = ["geoserver:latest"],
    visibility = ["//visibility:public"],
)