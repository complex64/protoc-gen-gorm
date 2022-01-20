IMAGE=complex64/protoc-gen-gorm

all: build
build:
	go build ./...

# Build containerized `protoc-gen-gorm`.
docker:
	docker build -t "${IMAGE}:latest" .

# Build and run the aggregated linker.
l: lint
lint: build
	golangci-lint run
