#
# progress
#

BIN_NAME    := progress
BIN_VERSION := 1.0.0
IMAGE_NAME  := hoanhan101/progress
GIT_COMMIT  ?= $(shell git rev-parse --short HEAD 2> /dev/null || true)
BUILD_DATE  := $(shell date -u +%Y-%m-%dT%T 2> /dev/null)

.PHONY: build
build:  ## Build the executable binary
	go build -o bin/${BIN_NAME} cmd/progress.go

.PHONY: build-linux
build-linux:  ## Build the executable binary for linux/amd64
	GOARCH=amd64 GOOS=linux go build -o bin/${BIN_NAME} cmd/progress.go

.PHONY: clean
clean:  ## Remove temporary files, build artifacts, postgres mounted volumes
	go clean -v ./...
	rm -rf bin
	rm -f coverage.out
	rm -rf deploy/compose/postgres-data

.PHONY: cover
cover: test-unit  ## Run unit tests and open the coverage report
	go tool cover -html=coverage.out

.PHONY: dev
dev: clean docker ## Clean, run docker compose for dev environment
	docker-compose -f deploy/compose/dev.yml up

.PHONY: docker
docker:  ## Build the docker image locally
	docker build -f Dockerfile \
		--label "org.label-schema.build-date=${BUILD_DATE}" \
		--label "org.label-schema.vcs-ref=${GIT_COMMIT}" \
		--label "org.label-schema.version=${BIN_VERSION}" \
        -t ${IMAGE_NAME}:latest .

.PHONY: lint
lint:  ## Lint project source files
	golangci-lint run

.PHONY: test-unit
test-unit:  ## Run unit tests
	go test -race -cover -run Unit -coverprofile=coverage.out -covermode=atomic ./...

.PHONY: test-end
test-end:  ## Run end-to-end tests
	-docker-compose -f deploy/compose/dev.yml rm -fsv
	docker-compose -f deploy/compose/dev.yml up -d
	sleep 6
	go test -race -cover -run E2E ./... || (docker-compose -f deploy/compose/dev.yml stop; exit 1)
	docker-compose -f deploy/compose/dev.yml down

.PHONY: version
version: ## Print the version
	@echo "${BIN_VERSION}"

.PHONY: help
help:  ## Print usage information
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST) | sort

.DEFAULT_GOAL := help
