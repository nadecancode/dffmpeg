## Windows
Create a `.bazelrc.windows.local` file with following content:
```
startup --output_user_root=D:/tmp/bazel
```
Or whatever path but keep it short as workaround for windows path limitation. If you receive error 1104 when compiling Rust then it's likely due to path being too long