# protoc-gen-gorm

[![Tests](https://github.com/complex64/protoc-gen-gorm/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/complex64/protoc-gen-gorm/actions/workflows/tests.yml) [![Linters](https://github.com/complex64/protoc-gen-gorm/actions/workflows/linters.yml/badge.svg?branch=main)](https://github.com/complex64/protoc-gen-gorm/actions/workflows/linters.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/complex64/protoc-gen-gorm)](https://goreportcard.com/report/github.com/complex64/protoc-gen-gorm) [![Maintainability](https://api.codeclimate.com/v1/badges/69739915a43041e34892/maintainability)](https://codeclimate.com/github/complex64/protoc-gen-gorm/maintainability) [![Go Reference](https://pkg.go.dev/badge/github.com/complex64/protoc-gen-gorm.svg)](https://pkg.go.dev/github.com/complex64/protoc-gen-gorm)

Generate [GORM v2 Models](https://gorm.io/docs/models.html) and APIs from your .proto files.

`protoc-gen-gorm` is a plugin for [protoc](https://grpc.io/docs/protoc-installation/) (or [buf](https://docs.buf.build/introduction)), a [Protocol Buffer](https://developers.google.com/protocol-buffers) ("proto") compiler.

**Under active development as of November 2022.**

Used in production at [GameAnalytics](https://github.com/GameAnalytics) in 2022.

## [Documentation](https://complex64.github.io/protoc-gen-gorm/)

Please find the documentation [here](https://complex64.github.io/protoc-gen-gorm/).

Additionally, there are several [code examples](/examples), including a full [gRPC/REST/CRUD service](/examples/grpc).

## Install

```
go install github.com/complex64/protoc-gen-gorm@latest
```

## Features

- Targets Gorm v2 (`gorm.io/gorm`)
- Generates GORM-compatible [model struct types](https://gorm.io/docs/models.html) for your message types: `UserModel` for your `User` message
- Generates methods to convert from proto message to model: `ToModel()` on the proto message, and `ToProto()` on the model
- Generates [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) helper methods for rapid and convenient implementation of APIs: `Create()`, `Get()`, `List()`, `Update()`, `Patch()`, and `Delete()`
- Can encode fields as JSON strings when instructed to

Additionally:

- All features are covered by tests
- Minimal external dependencies

**Planned Features:**

- [ ] Support for record lifecycle hooks (on demand)
- [ ] Support for associations
- [ ] Support for custom types
- [ ] Support for database-specific types
- [ ] Support for embedded structs
- [ ] Code examples
- [ ] Helpful code comments
- [ ] Comments on generated code

## Contributing

TODO

## Notes

TODO
