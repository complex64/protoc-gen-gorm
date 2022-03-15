# protoc-gen-gorm

[![Tests](https://github.com/complex64/protoc-gen-gorm/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/complex64/protoc-gen-gorm/actions/workflows/tests.yml) [![Linters](https://github.com/complex64/protoc-gen-gorm/actions/workflows/linters.yml/badge.svg?branch=main)](https://github.com/complex64/protoc-gen-gorm/actions/workflows/linters.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/complex64/protoc-gen-gorm)](https://goreportcard.com/report/github.com/complex64/protoc-gen-gorm) [![Maintainability](https://api.codeclimate.com/v1/badges/69739915a43041e34892/maintainability)](https://codeclimate.com/github/complex64/protoc-gen-gorm/maintainability) [![Go Reference](https://pkg.go.dev/badge/github.com/complex64/protoc-gen-gorm.svg)](https://pkg.go.dev/github.com/complex64/protoc-gen-gorm)

Generate GORM v2 Models and APIs from your .proto files.

```
go install github.com/complex64/protoc-gen-gorm@latest
```

**Under active development as of March 2022.**

---

**Features:**

Existing or planned:

- [x] Uses Gorm v2 (`gorm.io/gorm`)
- [x] Generates Gorm Models
- [x] Generates methods to convert between proto and Gorm
- [x] Generates CRUD methods (on demand)
- [ ] Support for inline-JSON encoding (on demand)
- [ ] Support for record lifecycle hooks (on demand)
- [ ] Support for associations
- [ ] Support for custom types
- [ ] Support for database-specific types
- [ ] Support for embedded structs
- [x] All features covered by tests
- [x] Minimal dependencies

    - 0 baseline
    - 1 when using hooks (Gorm)
    - 3 when using CRUD (aip-go, Gorm, fieldmaskpb)
    - 1 implicit for validation (expects methods from protoc-gen-validate)

- [ ] Documented proto options with examples
- [ ] Documented code
- [ ] Extensive documentation with examples
