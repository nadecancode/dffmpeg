load("@bazel_gazelle//:def.bzl", "gazelle")

gazelle(
    name = "gazelle-controller",
)

# gazelle:prefix

alias(
    name = "controller",
    actual = "//apps/controller:app"
)

alias(
    name = "worker",
    actual = "//apps/worker:app",
)
