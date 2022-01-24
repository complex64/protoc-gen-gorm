# About

This is the plugin itself.

Simply speaking, a plugin is any program that reads [CodeGeneratorRequests](https://pkg.go.dev/github.com/golang/protobuf/protoc-gen-go/plugin#CodeGeneratorRequest) on standard input, and returns [CodeGeneratorResponses](https://pkg.go.dev/github.com/golang/protobuf/protoc-gen-go/plugin#CodeGeneratorResponse) on standard output.

The Go authors support writing plugins with their [protogen package](https://pkg.go.dev/google.golang.org/protobuf/compiler/protogen).

## Example

A simple does-nothing plugin:

```go
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	opts := protogen.Options{}
	opts.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			if f.Generate {
			    // Do something with `plugin` and `f`...
			}
		}
		return nil
	})
}
```
