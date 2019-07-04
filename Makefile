GOFILES = $(shell find . \( -name vendor -prune \) -o \( -name '*.go' -print \) )

all: bin/default-authz-plugin

Gopkg.lock: Gopkg.toml
	dep ensure

.PHONY: deps
deps: Gopkg.lock

bin/:
	mkdir -p $@

bin/default-authz-plugin: $(GOFILES) Gopkg.lock bin/
	go build -o $@ .

.PHONY: image
image:
	docker build -t stackrox/default-authz-plugin .

.PHONY: imports
imports:
	@echo "+ $@"
	goimports -w $(GOFILES)

.PHONY: fmt
fmt:
	@echo "+ $@"
	gofmt -s -w $(GOFILES)

.PHONY: lint
lint:
	@echo "+ $@"
	golint -set_exit_status $(sort $(dir $(GOFILES)))

.PHONY: vet
vet:
	@echo "+ $@"
	go vet ./...

.PHONY: style
style:
	@echo "+ $@"
	@$(MAKE) imports
	@$(MAKE) fmt
	@$(MAKE) -k lint vet

.PHONY: tag
tag:
	@git describe --tags --abbrev=10 --dirty --long --always
