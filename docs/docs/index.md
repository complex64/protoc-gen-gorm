# Introduction

`protoc-gen-gorm` is a plugin for [`protoc`](https://grpc.io/docs/protoc-installation/), the [Protocol Buffer](https://developers.google.com/protocol-buffers) ("proto") Compiler.

The plugin generates [GORM v2](https://gorm.io/) [models](https://gorm.io/docs/models.html) and supporting code, depending on the [options](options.md) you use.

## Install

```
go install github.com/complex64/protoc-gen-gorm@latest
```

- TODO: Point to Usage for GHA/containerized

## Features

- Targets Gorm v2 (`gorm.io/gorm`)
- Generates GORM-compatible model struct types for your message types: `UserModel` for your `User` message
- Generates methods to convert from proto message to model: `ToModel()` on the proto message, and `ToProto()` on the model
- Generates [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) helper methods for rapid and convenient implementation of APIs: `Create()`, `Get()`, `List()`, `Update()`, `Patch()`, and `Delete()`
- Can encode fields as JSON strings when instructed to

## Usage

- TODO: protoc
- TODO: buf
- TODO: GHA

## About

- TODO: Repeat notes from README
