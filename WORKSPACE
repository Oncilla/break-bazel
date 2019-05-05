# Generic stuff for dealing with repositories.
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive", "http_file")

# Bazel rules for Golang
git_repository(
    name = "io_bazel_rules_go",
    remote = "https://github.com/bazelbuild/rules_go.git",
    tag = "0.18.1",
)

load("@io_bazel_rules_go//go:deps.bzl", "go_download_sdk", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_download_sdk(
    name = "go_sdk",
    sdks = {
        "linux_amd64": ("go1.11.5.linux-amd64.tar.gz", "ff54aafedff961eb94792487e827515da683d61a5f9482f668008832631e5d25"),
    },
)

# Gazelle
http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.17.0/bazel-gazelle-0.17.0.tar.gz"],
    sha256 = "3c681998538231a2d24d0c07ed5a7658cb72bfb5fd4bf9911157c0e9ac6a2687",
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()


go_repository(
    name = "com_github_golang_mock",
    commit = "73dc87cad333b55a02058f4b3d872dbbafddc2b0",
    importpath = "github.com/golang/mock",  # gomock
)

