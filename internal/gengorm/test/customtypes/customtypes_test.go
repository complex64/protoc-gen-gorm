package customtypes_test

import (
	"testing"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm/test/customtypes"
	"github.com/complex64/protoc-gen-gorm/internal/require"
)

func TestWithInlineJson(t *testing.T) {
	var (
		msg = &customtypes.WithInlineJsonModel{}
	)
	t.Run("fields are byte slices", func(t *testing.T) {
		require.FieldType(t, msg, "MessageField", []byte{})
		require.FieldType(t, msg, "NestedMessageField", []byte{})
		require.FieldType(t, msg, "MapField", []byte{})
		require.FieldType(t, msg, "RepeatedField", []byte{})
		require.FieldType(t, msg, "StringField", []byte{})
		require.FieldType(t, msg, "Int32Field", []byte{})
		require.FieldType(t, msg, "BoolField", []byte{})
	})
}
