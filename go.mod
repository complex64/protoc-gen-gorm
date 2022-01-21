module github.com/complex64/protoc-gen-gorm

go 1.17

require (
	github.com/complex64/protoc-gen-gorm/gormpb/v2 v2.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.2
	google.golang.org/protobuf v1.27.1
)

// TODO: Why?
replace github.com/complex64/protoc-gen-gorm/gormpb/v2 => ./gormpb/v2
