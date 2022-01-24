# About

This module declares and documents the options you can use in your `.proto` files.

# Usage

Include `options.proto` in your own `.proto` files and apply options to the file, your messages, or fields to control what `proto-gen-gorm` generates.

```proto
import "gorm/v2/options.proto";
```

Your editor and build setup needs to reference this file.

You have multiple options, depending what tools you use:

- Copy the file into your project, e.g. a `vendor` directory
- Use [Buf's Schema Registry (BSR)](https://docs.buf.build/bsr/introduction)

## Companion Go Module

Any `.proto` file that references this file will generate a Go package that imports `github.com/complex64/protoc-gen-gorm/gormpb/v2`. This package is a [minimal standalone Go module](../gormpb/v2).
