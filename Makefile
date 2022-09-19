# Note: github actions handles creating the image now
VERSION=v0.0.6
MAIN=main.go
IMAGE_NAME=go-api
IMAGE_REPO=ghcr.io/ericbutera
IMAGE_TAG=${VERSION}

.DEFAULT_GOAL := help

.PHONY: build
build: ## Build the binary
	go build -o bin/app

.PHONY: clean
clean: ## Remove builds
	rm bin/*

.PHONY: docker-build
docker-build: ## Build the docker image
	docker build -t ${IMAGE_NAME} .

.PHONY: docker-run
docker-run: ## Run the docker image
	docker run --rm -p 8080:8080 ${IMAGE_NAME}

# .PHONY: docs-install
# docs-install:
# 	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: docs
docs: ## generate docs
	swag init

.PHONY: help
help: ## Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: image-build
image-build: ## Build docker image
	docker build -t ${IMAGE_REPO}/${IMAGE_NAME}:${IMAGE_TAG} .

.PHONY: image-push
 image-push: ## Push docker image
	docker push ${IMAGE_REPO}/${IMAGE_NAME}:${IMAGE_TAG}

.PHONY: image-run-local
 image-run-local: ## Run docker image locally
	docker run ${IMAGE_NAME}:${IMAGE_TAG}

.PHONY: image-run
 image-run: ## Run docker image
	docker run ${IMAGE_REPO}/${IMAGE_NAME}:${IMAGE_TAG}

.PHONY: run
run: ## Run app
	go run ${MAIN} server

.PHONY: test
test: ## Run tests
	go test ./...

.PHONY: tidy
tidy: ## go mod tidy
	go mod tidy
