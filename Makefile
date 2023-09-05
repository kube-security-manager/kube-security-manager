MAKEFLAGS += --no-print-directory

BIN := bin

GO ?= go
DOCKER ?= docker
KIND ?= kind

export KUBECONFIG ?= ${HOME}/.kube/config

# Active module mode, as we use Go modules to manage dependencies
export GO111MODULE=on
GOPATH=$(shell go env GOPATH)
GOBIN=$(GOPATH)/bin
GINKGO=$(GOBIN)/ginkgo

SOURCES := $(shell find . -name '*.go')

IMAGE_TAG := dev
IMAGE_NAME := docker.io/danielpacak/kube-security-manager:$(IMAGE_TAG)

.PHONY: all
all: controller-manager node-agent

controller-manager: $(SOURCES)
	CGO_ENABLED=0 GOOS=linux $(GO) build -o $@ ./cmd/controller-manager/main.go

node-agent: $(SOURCES)
	CGO_ENABLED=0 GOOS=linux $(GO) build -o $@ ./cmd/node-agent/main.go

.PHONY: get-ginkgo
## Installs Ginkgo CLI
get-ginkgo:
	$(GO) install github.com/onsi/ginkgo/ginkgo

.PHONY: unit-tests
## Runs unit tests with code coverage enabled
unit-tests: $(SOURCES)
	$(GO) test -v -short -race -timeout 30s -coverprofile=coverage.txt ./...

.PHONY: itests-starboard-operator
## Runs integration tests for Starboard Operator with code coverage enabled
itests-starboard-operator: check-kubeconfig get-ginkgo
	@$(GINKGO) \
	-coverprofile=coverage.txt \
	-coverpkg=github.com/aquasecurity/starboard/pkg/operator,\
	github.com/aquasecurity/starboard/pkg/operator/predicate,\
	github.com/aquasecurity/starboard/pkg/operator/controller,\
	github.com/aquasecurity/starboard/pkg/plugin,\
	github.com/aquasecurity/starboard/pkg/plugin/trivy,\
	github.com/aquasecurity/starboard/pkg/plugin/polaris,\
	github.com/aquasecurity/starboard/pkg/plugin/conftest,\
	github.com/aquasecurity/starboard/pkg/configauditreport,\
	github.com/aquasecurity/starboard/pkg/vulnerabilityreport,\
	github.com/aquasecurity/starboard/pkg/kubebench \
	./itest/starboard-operator

.PHONY: check-kubeconfig
check-kubeconfig:
ifndef KUBECONFIG
	$(error Environment variable KUBECONFIG is not set)
else
	@echo "KUBECONFIG=${KUBECONFIG}"
endif

## Removes build artifacts
clean:
	@rm -r ./bin 2> /dev/null || true
	@rm -r ./dist 2> /dev/null || true
	rm -rf controller-manager
	rm -rf node-agent

.PHONY: docker-build
docker-build: controller-manager node-agent
	$(DOCKER) image build --no-cache -t $(IMAGE_NAME) -f Dockerfile .

.PHONY: kind-load-images
kind-load-images: docker-build
	$(KIND) load docker-image $(IMAGE_NAME)

.PHONY: \
	clean