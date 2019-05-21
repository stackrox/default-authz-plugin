GOFILES = $(shell find . \( -name vendor -prune \) -o \( -name '*.go' -print \) )

all: bin/sample-authz-plugin

Gopkg.lock: Gopkg.toml
	dep ensure

bin/:
	mkdir -p $@

bin/sample-authz-plugin: $(GOFILES) Gopkg.lock bin/
	go build -o $@ .

.PHONY: image
image:
	docker build -t stackrox/sample-authz-plugin .
