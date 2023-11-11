load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "controller"
)

# gazelle:prefix

alias(
    name = "worker",
    actual = "//apps/worker:app",
)