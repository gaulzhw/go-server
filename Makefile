# Ensure Make is run with bash shell as some syntax below is bash-specific
SHELL:=/usr/bin/env bash

.DEFAULT_GOAL := help

#
# Go.
#
GO_VERSION ?= 1.18-alpine
GO_CONTAINER_IMAGE ?= golang:$(GO_VERSION)

#
# Directories.
#
# Full directory of where the Makefile resides
ROOT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
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
REGISTRY ?= k8scloudplatform

# syncer
IMAGE_NAME_MANAGER ?= go-server
CONTROLLER_IMG_MANAGER ?= $(REGISTRY)/$(IMAGE_NAME_MANAGER)

# release
RELEASE_TAG ?= $(shell git describe --tags)

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
	docker build --build-arg builder_image=$(GO_CONTAINER_IMAGE) --build-arg package=cmd/manager/main.go . -t $(CONTROLLER_IMG_MANAGER):$(RELEASE_TAG)

.PHONY: docker-push
docker-push: ## Push image.
	docker push $(CONTROLLER_IMG_MANAGER):$(RELEASE_TAG)

## --------------------------------------
## Hack / Tools
## --------------------------------------

##@ hack/tools:

code-gen: $(CODE_GEN) ## Build a local copy of code-gen.

$(CODE_GEN): $(TOOLS_DIR)/go.mod # Build code-gen from tools folder.
	cd $(TOOLS_DIR); go build -tags=tools -o $(BIN_DIR)/code-gen github.com/gaulzhw/go-server/hack/tools/cmd/codegen
