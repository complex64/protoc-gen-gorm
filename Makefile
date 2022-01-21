IMAGE=complex64/protoc-gen-gorm

all: lint test
build:
	go build ./...

# Run all tests.
t :test
test:
	go test ./...

# Build containerized `protoc-gen-gorm`.
docker:
	docker build -t "${IMAGE}:latest" .

# Build and run the aggregated linker.
l: lint
lint: build
	golangci-lint run
