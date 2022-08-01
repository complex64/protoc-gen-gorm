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

# Lint the .proto files with buf.
b: buf
buf:
	cd proto && buf lint
	cd internal/gengorm/test && buf lint
	cd internal/require && buf lint

# Assumes $GOPATH/bin is in your $PATH!
gen: generate
generate: gormpb install
	# Files used by tests of the plugin implementation.
	cd internal/gengorm/test && buf generate

	# Remove code generated from tests.
	find proto -name '*.go' -delete

	# Files used by tests of the internal packages.
	cd internal/require && buf generate

	# Generated examples
	cd examples/buf && go generate
	cd examples/grpc && go generate
	cd examples/protoc && go generate

.PHONY: gormpb
gormpb:
	# Generate the standalone module and update/lock dependencies.
	cd proto && buf generate
	find proto -type f -name "*.pb.go" -exec mv {} gormpb \;
	cd gormpb && go mod tidy

# Install `protoc-gen-go` into $GOPATH/bin.
i: install
install:
	go install

# Remove all generated files.
clean:
	go clean
	find -name '*.pb.go' -delete
	$(MAKE) gormpb

p: proto
.PHONY: proto
proto:
	buf lint proto
	cd internal/gengorm/test && buf lint
	cd internal/require && buf lint
