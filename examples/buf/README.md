# protoc

To use protoc-gen-gorm with [buf](https://docs.buf.build/introduction):
- Make sure `buf` and `protoc-gen-gorm` is in your path (see [Install](https://docs.buf.build/installation), as well as [how to install buf](https://docs.buf.build/installation))
- [Specify the plugins](/examples/buf/buf.gen.yaml) you'd like to use
- Invoke `buf generate`

This compiles our [GORM v2 model](/examples/buf/pb/models_gorm.pb.go) (and [Go bindings](/examples/buf/pb/models.pb.go)) that we make use of in [main.go](/examples/buf/main.go):

```go
if err := db.Create(&pb.FileModel{
    Uuid:  uuid.NewString(),
    Name:  path,
    Bytes: info.Size(),
}).Error; err != nil {
    return err
}
```

When run:

```
$ go generate
$ go run main.go
Inserting...
Files in database:
27cdb9a5-6cc3-4b95-9935-56d1efa22958    9       4:46PM  4:46PM  .gitignore
b8871f9f-ca96-4e7b-94cf-6c1f36e191ac    19      4:46PM  4:46PM  protos/gorm
0209bc87-80ab-4a61-91e5-fd2d917ace41    151     4:46PM  4:46PM  buf.gen.yaml
0cf9ef72-2d5c-46a6-93c3-f1a611b00a4b    432     4:46PM  4:46PM  go.mod
105350a3-c827-4d05-96c1-9a08fe53986a    728     4:46PM  4:46PM  protos/models.proto
42fd81b0-6e0d-40ca-a3e9-5be7a9360841    733     4:46PM  4:46PM  README.md
a01e9b6e-d41a-418a-b360-98446c6bc324    1328    4:46PM  4:46PM  pb/models_gorm.pb.go
3047f3b7-bde7-4edf-b1b6-23ea2d990c58    1399    4:46PM  4:46PM  main.go
4a6544fa-601e-4de8-8698-2481488595c0    2142    4:46PM  4:46PM  go.sum
88266425-43cb-432e-889c-0491de9aafcb    6706    4:46PM  4:46PM  pb/models.pb.go
afdb7bef-1a16-4df2-a427-b14b19ddae1c    16384   4:46PM  4:46PM  gorm.db
```
