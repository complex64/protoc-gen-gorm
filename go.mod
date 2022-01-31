module github.com/complex64/protoc-gen-gorm

go 1.17

require (
	github.com/complex64/protoc-gen-gorm/gormpb/v2 v2.0.0-00010101000000-000000000000
	github.com/google/go-cmp v0.5.7
	github.com/stretchr/testify v1.7.0
	google.golang.org/protobuf v1.27.1
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

// Run tests against the local version of the generated code for options.proto for convenient feature development.
// Otherwise we'd have to publish features to options.proto, before being able to test them.
replace github.com/complex64/protoc-gen-gorm/gormpb/v2 => ./gormpb/v2
