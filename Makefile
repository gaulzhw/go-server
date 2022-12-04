# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL:=/usr/bin/env bash

.DEFAULT_GOAL := help

#
# Go.
#
GO_VERSION ?= 1.19-alpine
GO_CONTAINER_IMAGE ?= golang:$(GO_VERSION)

#
# Directories.
#
# Full directory of where the Makefile resides
BIN_DIR := bin
TOOLS_DIR := hack/tools
TOOLS_BIN_DIR := $(TOOLS_DIR)/$(BIN_DIR)

#
# Binaries.
#
# Note: Need to use abspath so we can invoke these from subdirectories
# code gen
CODE_GEN := $(abspath $(TOOLS_BIN_DIR)/code-gen)

#
# Images.
#
# Define Docker related variables. Releases should modify and double check these vars.
IMAGE_NAME ?= k8scloudplatform/go-server

# release
RELEASE_TAG ?= $(shell git describe --tags)

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
VERSION := $(shell git rev-parse HEAD | cut -b -8)

ifeq ($(shell git status --porcelain 2>/dev/null),)
	GIT_TREESTATE = clean
else
	GIT_TREESTATE = dirty
endif

BUILD_DATE := $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
LDFLAGS := "-X github.com/gaulzhw/go-server/pkg/version.gitVersion=${BRANCH} \
			-X github.com/gaulzhw/go-server/pkg/version.gitCommit=${VERSION} \
			-X github.com/gaulzhw/go-server/pkg/version.gitTreeState=$(GIT_TREESTATE) \
			-X github.com/gaulzhw/go-server/pkg/version.buildDate=${BUILD_DATE}"

help:  # Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[0-9A-Za-z_-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## --------------------------------------
## Generate
## --------------------------------------

##@ generate:

.PHONY: generate
generate: $(CODE_GEN) ## Run all generate targets.
	cd pkg/errno; $(CODE_GEN) -type=errorno; $(CODE_GEN) -type=errorno -doc --output=errno_generated.md

## --------------------------------------
## Modules
## --------------------------------------

##@ modules:

.PHONY: modules
modules: ## Run go mod tidy to ensure modules are up to date.
	go mod tidy
	cd $(TOOLS_DIR); go mod tidy

## --------------------------------------
## Docker
## --------------------------------------

##@ docker:

.PHONY: docker-build
docker-build: ## Build image.
	docker build --build-arg builder_image=$(GO_CONTAINER_IMAGE) --build-arg package=cmd/server/main.go --build-arg ldflags=$(LDFLAGS) . -t $(IMAGE_NAME):$(BRANCH)-$(VERSION)

.PHONY: docker-push
docker-push: ## Push image.
	docker push $(IMAGE_NAME):$(BRANCH)-$(VERSION)

## --------------------------------------
## Hack / Tools
## --------------------------------------

##@ hack/tools:

code-gen: $(CODE_GEN) ## Build a local copy of code-gen.

$(CODE_GEN): $(TOOLS_DIR)/go.mod # Build code-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/code-gen github.com/gaulzhw/go-server/hack/tools/cmd/codegen
