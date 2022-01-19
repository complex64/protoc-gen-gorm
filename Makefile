IMAGE=complex64/protoc-gen-gorm

all:

# Build containerized `protoc-gen-gorm`.
docker:
	docker build -t "${IMAGE}:latest" .
