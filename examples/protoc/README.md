# protoc

To use protoc-gen-gorm with [protoc](https://grpc.io/docs/protoc-installation/), make sure `protoc-gen-gorm` is in your path (see [Install](https://complex64.github.io/protoc-gen-gorm/#install)) and invoke `protoc` with `--gorm_out=dir`, where _dir_ is your desired output directory for the files that protoc-gen-gorm generates.

For this example we invoke (`go generate`):

```protoc -I. --go_out=./ --gorm_out=./ models.proto```

This compiles our [GORM v2 model](/examples/protoc/pb/models_gorm.pb.go) (and [Go bindings](/examples/protoc/pb/models.pb.go)) we make use of in [main.go](/examples/protoc/main.go):

```go
alice := &pb.UserModel{Name: "Alice"}
db.Create(alice)
```

When run:

```shell
$ go run main.go
...
Created first user: Alice

$ go run main.go
The first user is: Alice
```
