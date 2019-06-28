# Building the Default Authorization Plugin

## Prerequisites

Building the Default Authorization Plugin Docker image requires:
- GNU `make`
- Docker CLI

If you want to make modifications to the source code and/or build the server
for your local host system (not in a Docker image), you will most likely also need:
- [dep](https://github.com/golang/dep)
- A recent version of [Go](https://golang.org) (tested with 1.12, but anything >= 1.9 should work just fine)
- [golint](https://github.com/golang/lint)

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
