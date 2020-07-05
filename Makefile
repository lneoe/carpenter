PACKAGE := carpenter
COMMIT_ID := $(shell git rev-parse HEAD)
BUILD_TIME := $(shell date -u "+%Y-%m-%dT%H:%M:%S")
_VENDOR_VERSION := github.com/lneoe/go-help-libs/version
BUILD_ARGS := \
    -ldflags "-X $(_VENDOR_VERSION).Version=$(VERSION) -X $(_VENDOR_VERSION).Revision=$(COMMIT_ID) \
    -X $(_VENDOR_VERSION).Branch=$(VCS_BRANCH) -X $(_VENDOR_VERSION).BuildDate=$(BUILD_TIME) \
    -X $(_VENDOR_VERSION).GRPCStubRevision=$(GRPC_STUB_REVISION)" -tags="openssl_static"
EXTRA_BUILD_ARGS =
OUTPUT_FILE := bin/carpenter
IMAGE_NAME := carpenter
IMAGE_TAG := latest
IMAGE_REGISTRY := ""
IMAGE_ORG := lneoe
IMAGE_FULL_NAME := $(IMAGE_ORG)/$(IMAGE_NAME):$(IMAGE_TAG)

VERSION :=
VCS_BRANCH :=
GRPC_STUB_REVISION :=

.PHONY: fotmat lint test build

default: lint test build

build:
	@echo '+ $@'
	go build $(BUILD_ARGS) $(EXTRA_BUILD_ARGS) -o $(OUTPUT_FILE) main.go

clean:
	@echo "+ $@"
	@if [ -d "./bin" ]; then echo 'rm -r ./bin'; rm -r ./bin; fi

format:
	@find $(CURDIR) -mindepth 1 -maxdepth 1 -type d -not -name vendor -not -name .git -print0 | xargs -0 gofmt -s -w
	@find $(CURDIR) -maxdepth 1 -type f -name '*.go' -print0 | xargs -0 gofmt -s -w

test:
	@echo '+ $@'
	@go test ./...

lint:
	@echo '+ $@'
	golangci-lint run ./...


build-image:
	@docker build -t $(IMAGE_FULL_NAME) . \
	--build-arg COMMIT_ID=$(COMMIT_ID) --build-arg BUILD_TIME=$(BUILD_TIME) \
	--build-arg EXTRA_BUILD_ARGS=$(EXTRA_BUILD_ARGS)
