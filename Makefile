GOFILES = $(shell find . \( -name vendor -prune \) -o \( -name '*.go' -print \) )

USE_GO_MODULES := $(shell [ ! -f go.mod -o "$$PWD" = "$$(go env GOPATH)/src/github.com/stackrox/default-authz-plugin" ] && echo 0 || echo 1)

ifeq ($(USE_GO_MODULES), 1)
	GOFLAGS := -mod vendor
else
	GOFLAGS :=
endif

all: bin/default-authz-plugin

Gopkg.lock: Gopkg.toml
ifeq ($(USE_GO_MODULES), 1)
	go mod vendor
else
	dep ensure
endif

.PHONY: deps
deps: Gopkg.lock

bin/:
	mkdir -p $@

bin/tools/:
	mkdir -p $@

bin/tools/%-linux: $(GOFILES) bin/tools/
	GOOS=linux CGO_ENABLED=0 go build $(GOFLAGS) -o $@ ./cmd/$*

bin/tools/%-darwin: $(GOFILES) bin/tools/
	GOOS=darwin CGO_ENABLED=0 go build $(GOFLAGS) -o $@ ./cmd/$*

bin/tools/%.exe: $(GOFILES) bin/tools/
	GOOS=windows CGO_ENABLED=0 go build $(GOFLAGS) -o $@ ./cmd/$*

bin/default-authz-plugin: $(GOFILES) Gopkg.lock bin/
	go build $(GOFLAGS) -o $@ .

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
	go vet $(GOFLAGS) ./...

.PHONY: style
style:
	@echo "+ $@"
	@$(MAKE) imports
	@$(MAKE) fmt
	@$(MAKE) -k lint vet

.PHONY: tests
	@echo "+ $@"
	@go test ./...

.PHONY: tag
tag:
	@git describe --tags --abbrev=10 --dirty
