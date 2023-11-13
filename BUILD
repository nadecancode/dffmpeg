load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix controller
gazelle(
    name = "gazelle-controller",
    args = [
        "-proto=disable",
    ],
)

gazelle(name = "gazelle")

alias(
    name = "controller",
    actual = "//apps/controller:app",
)

alias(
    name = "worker",
    actual = "//apps/worker:app",
)
