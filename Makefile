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
