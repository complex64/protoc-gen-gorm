package options_test

import (
	// Assert compilation.

	"testing"

	"github.com/complex64/protoc-gen-gorm/cmd/protoc-gen-gorm/test/options"
	"github.com/complex64/protoc-gen-gorm/gormpb"
	"github.com/complex64/protoc-gen-gorm/internal/require"
)

// Test that all options in the .proto file are present as expected.

func TestFileOptions(t *testing.T) {
	msg := &options.MyMessage{}
	defaults := &gormpb.FileOptions{
		Model:    false,
		Hooks:    false,
		Validate: false,
		Crud:     false,
	}
	require.FileOptions(t, defaults, msg)
}

func TestMessageOptions(t *testing.T) {
	var (
		msg = &options.MyMessage{}
	)
	t.Run("defaults", func(t *testing.T) {
		defaults := &gormpb.MessageOptions{
			Model:    false,
			Hooks:    false,
			Validate: false,
			Crud:     false,
			Table:    "",
		}
		require.MessageOption(t, defaults, msg)
	})
}

func TestFieldOptions(t *testing.T) {
	var (
		msg = &options.MyMessage{}
	)
	t.Run("defaults", func(t *testing.T) {
		defaults := &gormpb.FieldOptions{
			Column:             "",
			NotNull:            false,
			Default:            "",
			Unique:             false,
			PrimaryKey:         false,
			Index:              nil,
			UniqueIndex:        nil,
			AutoCreateTime:     false,
			AutoCreateTimeType: gormpb.TimeType_GO_TIME_TIME,
			AutoUpdateTime:     false,
			AutoUpdateTimeType: gormpb.TimeType_GO_TIME_TIME,
			Permissions:        nil,
			Comment:            "",
		}
		require.FieldOption(t, defaults, msg, "field_1")
	})
}
