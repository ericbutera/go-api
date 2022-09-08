VERSION=0.0.3
MAIN=main.go
IMAGE_NAME=api-go
IMAGE_REPO=ericbutera
IMAGE_TAG=${VERSION}

.PHONY: build
build:
	go build -o bin/app

.PHONY: clean
clean:
	rm bin/*

.PHONY: docker-build
docker-build:
	docker build -t ${IMAGE_NAME} .

.PHONY: docker-run
docker-run:
	docker run --rm -p 8080:8080 ${IMAGE_NAME}

.PHONY: docs-install
docs-install:
	go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: docs
docs:
	swag init

.PHONY: help
help:
	echo "Usage: make [target]"

.PHONY: image-build
image-build:
	docker build -t ${IMAGE_REPO}/${IMAGE_NAME}:${IMAGE_TAG} .

.PHONY: image-push
 image-push:
	docker push ${IMAGE_REPO}/${IMAGE_NAME}:${IMAGE_TAG}

.PHONY: image-run-local
 image-run-local:
	docker run ${IMAGE_NAME}:${IMAGE_TAG}

.PHONY: image-run
 image-run:
	docker run ${IMAGE_REPO}/${IMAGE_NAME}:${IMAGE_TAG}


.PHONY: protoc
protoc:
	protoc --go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    protobuf/*.proto

.PHONY: run
run:
	go run ${MAIN}

.PHONY: test
test:
	go test ./...

.PHONY: tidy
tidy:
	go mod tidy