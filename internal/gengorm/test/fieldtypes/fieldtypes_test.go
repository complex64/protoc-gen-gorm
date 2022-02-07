package fieldtypes_test

import (
	"testing"
	"time"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm/test/fieldtypes"
	"github.com/complex64/protoc-gen-gorm/internal/require"
)

// Assert correct (GORM-compatible) types are generated.

func TestWithScalarValues(t *testing.T) {
	msg := &fieldtypes.WithScalarValues{}
	require.FieldType(t, msg, "DoubleField", float64(0))
	require.FieldType(t, msg, "FloatField", float32(0))
	require.FieldType(t, msg, "Int32Field", int32(0))
	require.FieldType(t, msg, "Int64Field", int64(0))
	require.FieldType(t, msg, "Uint32Field", uint32(0))
	require.FieldType(t, msg, "Uint64Field", uint64(0))
	require.FieldType(t, msg, "Sint32Field", int32(0))
	require.FieldType(t, msg, "Sint64Field", int64(0))
	require.FieldType(t, msg, "Fixed32Field", uint32(0))
	require.FieldType(t, msg, "Fixed64Field", uint64(0))
	require.FieldType(t, msg, "Sfixed32Field", int32(0))
	require.FieldType(t, msg, "Sfixed64Field", int64(0))
	require.FieldType(t, msg, "BoolField", true)
	require.FieldType(t, msg, "StringField", "")
	require.FieldType(t, msg, "BytesField", []byte{})
}

func TestWithKnownTypes(t *testing.T) {
	msg := &fieldtypes.WithKnownTypesModel{}
	require.FieldType(t, msg, "TimestampField", time.Time{})
}

func TestWithEnum(t *testing.T) {
	msg := &fieldtypes.WithEnumModel{}
	require.FieldType(t, msg, "EnumField", int32(0))
}
