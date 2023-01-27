package grpc

//go:generate rm -f gorm.db
//go:generate find pb -type f -name *.go -delete
//go:generate buf generate protos
//go:generate rm -rf pb/gorm
//go:generate go mod tidy
