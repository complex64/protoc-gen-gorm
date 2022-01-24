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

# Generate files used in tests.
gen: generate
generate:
	cd proto && buf generate
	mv gormpb/v2/gorm/v2/*.pb.go gormpb/v2
	rm -r gormpb/v2/gorm
	cd gormpb/v2 && go mod tidy

	cd cmd/protoc-gen-gorm/test && buf generate
	rm proto/gorm/v2/options.pb.go

	cd internal/require && buf generate

# Remove all generated files.
clean:
	find -name '*.pb.go' -delete

p: proto
.PHONY: proto
proto:
	buf lint proto
	cd internal/require && buf lint
