workspace(name = "dffmpeg")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "91585017debb61982f7054c9688857a2ad1fd823fc3f9cb05048b0025c47d023",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.42.0/rules_go-v0.42.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.42.0/rules_go-v0.42.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "b7387f72efb59f876e4daae42f1d3912d0d45563eac7cb23d1de0b094ab588cf",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.34.0/bazel-gazelle-v0.34.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.34.0/bazel-gazelle-v0.34.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_rules_dependencies()

go_register_toolchains(version = "1.21.1")

gazelle_dependencies()

http_archive(
    name = "rules_rust",
    sha256 = "6357de5982dd32526e02278221bb8d6aa45717ba9bbacf43686b130aa2c72e1e",
    urls = ["https://github.com/bazelbuild/rules_rust/releases/download/0.30.0/rules_rust-v0.30.0.tar.gz"],
)

load("@rules_rust//rust:repositories.bzl", "rules_rust_dependencies", "rust_register_toolchains")

rules_rust_dependencies()

rust_register_toolchains()

GAZELLE_RUST_COMMIT = "b764243b3f22e8ad52f79466ba77f06797d0c3d4"
GAZELLE_RUST_SHA256 = "0586cdf986fb74301a4cde44be0280d5d7feeba40e8dd4a6eccdc8986e7cf152"

http_archive(
    name = "gazelle_rust",
    sha256 = GAZELLE_RUST_SHA256,
    strip_prefix = "gazelle_rust-{}".format(GAZELLE_RUST_COMMIT),
    url = "https://github.com/Calsign/gazelle_rust/archive/{}.zip".format(GAZELLE_RUST_COMMIT),
)

load("@gazelle_rust//:deps1.bzl", "gazelle_rust_dependencies1")
gazelle_rust_dependencies1()

load("@gazelle_rust//:deps2.bzl", "gazelle_rust_dependencies2")
gazelle_rust_dependencies2()