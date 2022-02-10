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
	t.Run("fields are strings", func(t *testing.T) {
		require.FieldType(t, msg, "MessageField", "")
		require.FieldType(t, msg, "NestedMessageField", "")
		require.FieldType(t, msg, "MapField", "")
		require.FieldType(t, msg, "RepeatedField", "")
		require.FieldType(t, msg, "StringField", "")
		require.FieldType(t, msg, "Int32Field", "")
		require.FieldType(t, msg, "BoolField", "")
	})
}
