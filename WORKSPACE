
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

###################################################
# Go lang rewrite
###################################################
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "dd926a88a564a9246713a9c00b35315f54cbd46b31a26d5d8fb264c07045f05d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.38.1/rules_go-v0.38.1.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.38.1/rules_go-v0.38.1.zip",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

go_rules_dependencies()

go_register_toolchains(version = "1.19.5")

http_archive(
    name = "bazel_gazelle",
    sha256 = "ecba0f04f96b4960a5b250c8e8eeec42281035970aa8852dda73098274d14a1d",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.29.0/bazel-gazelle-v0.29.0.tar.gz",
    ],
)

# Load and call Gazelle dependencies
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

go_repository(
    name = "com_github_pq",
    commit = "9eb3fc897d6fd97dd4aad3d0404b54e2f7cc56be",
    importpath = "github.com/lib/pq",
)

go_repository(
    name = "com_github_cors",
    commit = "fdcf4f9773b8d459d3ec085a8c91b34fcd17803d",
    importpath = "github.com/rs/cors",
)


# DOCKER
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "rules_oci",
    sha256 = "4a738bdbeacb0e1df070209dddfa7b55fed9bbc553b905cf3d2dd25115e0b598",
    strip_prefix = "rules_oci-0.3.8",
    url = "https://github.com/bazel-contrib/rules_oci/releases/download/v0.3.8/rules_oci-v0.3.8.tar.gz",
)

load("@rules_oci//oci:dependencies.bzl", "rules_oci_dependencies")

rules_oci_dependencies()

load("@rules_oci//oci:repositories.bzl", "LATEST_CRANE_VERSION", "LATEST_ZOT_VERSION", "oci_register_toolchains")

oci_register_toolchains(
    name = "oci",
    crane_version = LATEST_CRANE_VERSION,
    # If a docker media types compatible registry is desired, just omit `zot_version` and crane registry will be used instead.
    # zot_version = LATEST_ZOT_VERSION,
)

# Optional, for oci_tarball rule
load("@rules_pkg//:deps.bzl", "rules_pkg_dependencies")

rules_pkg_dependencies()


load("@rules_oci//oci:pull.bzl", "oci_pull")

oci_pull(
    name = "distroless_base",
    digest = "sha256:ccaef5ee2f1850270d453fdf700a5392534f8d1a8ca2acda391fbb6a06b81c86",
    image = "gcr.io/distroless/base",
    platforms = ["linux/amd64","linux/arm64"],
)