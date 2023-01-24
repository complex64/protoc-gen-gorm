# Build static executable.
FROM golang:1.19 as buildenv
WORKDIR /go/src/app
ADD . .
RUN GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=0 \
    go build -ldflags="-w -s" -o /bin/protoc-gen-gorm ./

# Compress binaries:
FROM alpine as upxenv
ADD https://github.com/upx/upx/releases/download/v4.0.1/upx-4.0.1-amd64_linux.tar.xz /tmp
COPY --from=buildenv /bin/protoc-gen-gorm /bin
RUN tar -C / -Jxpf /tmp/upx-4.0.1-amd64_linux.tar.xz \
    && /upx-4.0.1-amd64_linux/upx -q -9 /bin/protoc-gen-gorm

# Deliver binaries in minimal image:
FROM gcr.io/distroless/static
COPY --from=upxenv /bin/protoc-gen-gorm /bin
ENTRYPOINT ["protoc-gen-gorm"]
