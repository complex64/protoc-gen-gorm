# Stage 1: Build static executable.
FROM golang:1.19.2 as buildenv
WORKDIR /go/src/app
ADD . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /bin/protoc-gen-gorm ./

# Stage 2: Deliver in minimal image.
FROM gcr.io/distroless/static
COPY --from=buildenv /bin/protoc-gen-gorm /bin
ENTRYPOINT ["protoc-gen-gorm"]
