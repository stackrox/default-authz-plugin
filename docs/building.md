# Building the Default Authorization Plugin

## Prerequisites

Building the Default Authorization Plugin Docker image requires:
- GNU `make`
- Docker CLI

If you want to make modifications to the source code and/or build the server
for your local host system (not in a Docker image), you will most likely also need:
- [dep](https://github.com/golang/dep)
- A recent version of [Go](https://golang.org) (minimum 1.11, tested with 1.12, also see the note below)
- [golint](https://github.com/golang/lint)
- [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports)

**Note on required Go version:** The Default Authorization Plugin supports building with Go modules.
This requires a minimum Go version of 1.11. However, if you are using an older version of Go, or do
not want to use Go modules, you can do so by placing the source root directory inside your
`GOPATH` (check the output of `go env GOPATH` for the location on your system). For doing so,
make sure the source root is located at `$GOPATH/src/github.com/stackrox/default-authz-plugin`.

The above only applies to builds for your local host system (via `make`); builds of Docker images
(`make image`) are not affected.

## Build Steps

### Docker Image

The Default Authorization Plugin can be built by running `make image` in the root
source directory. This will result in an image with tag `stackrox/default-authz-plugin:latest`
being created on your local system, which you can then re-tag as you please.

The created image contains nothing but the server binary, which also serves as the entrypoint
for the image. Hence, arguments passed to `docker` (or specified in a Kubernetes deployment)
will be used as arguments for the binary, e.g.,
```bash
$ docker run \
  -v $PWD/examples/config/server-config-plain-anon.json:/server-config.json \
  stackrox/default-authz-plugin:latest \
  -server-config /server-config.json
```

Please note that running `make image` requires all dependencies to be up-to-date. If you have
made changes to the source code or to the dependencies, you might need to run `make deps` first,
which requires the `dep` tool to be installed on your machine.

### Stand-Alone Binary

You can also build a standalone, statically linked binary for your host system. To do so, simply run
`make`. The resulting binary will be named `default-authz-plugin`, located in the `bin/` directory.
