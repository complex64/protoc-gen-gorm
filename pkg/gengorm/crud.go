package gengorm

import (
	"google.golang.org/protobuf/types/known/fieldmaskpb"
	"gorm.io/gorm"
)

type CreateOption func(tx *gorm.DB)
type GetOption func(tx *gorm.DB)
type ListOption func(tx *gorm.DB)
type UpdateOption func(tx *gorm.DB)
type PatchOption func(tx *gorm.DB)
type DeleteOption func(tx *gorm.DB)

// https://google.aip.dev/132#filtering
// https://google.aip.dev/160

func WithFilter(filter string) ListOption {
	return func(tx *gorm.DB) {

	}
}

// https://google.aip.dev/132#ordering

func WithOrder(orderBy string) ListOption {
	return func(tx *gorm.DB) {

	}
}

// https://google.aip.dev/161

func WithListFieldMask(mask *fieldmaskpb.FieldMask) ListOption {
	return func(tx *gorm.DB) {

	}
}

func WithGetFieldMask(mask *fieldmaskpb.FieldMask) GetOption {
	return func(tx *gorm.DB) {

	}
}
