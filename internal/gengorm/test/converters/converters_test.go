package converters_test

import (
	"testing"
	"time"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm/test/converters"
	requirepb "github.com/complex64/protoc-gen-gorm/internal/require"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestScalarsModel_AsProto(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		m := new(converters.ScalarsModel)
		x, err := m.AsProto()
		require.NoError(t, err)
		require.NotNil(t, x)
		require.EqualValues(t, 0, x.DoubleField)
		require.EqualValues(t, 0, x.FloatField)
		require.EqualValues(t, 0, x.Int32Field)
		require.EqualValues(t, 0, x.Int64Field)
		require.EqualValues(t, 0, x.Uint32Field)
		require.EqualValues(t, 0, x.Uint64Field)
		require.EqualValues(t, 0, x.Sint32Field)
		require.EqualValues(t, 0, x.Sint64Field)
		require.EqualValues(t, 0, x.Fixed32Field)
		require.EqualValues(t, 0, x.Fixed64Field)
		require.EqualValues(t, 0, x.Sfixed32Field)
		require.EqualValues(t, 0, x.Sfixed64Field)
		require.False(t, x.BoolField)
		require.EqualValues(t, "", x.StringField)
		require.EqualValues(t, []byte(nil), x.BytesField)
	})

	t.Run("with values set", func(t *testing.T) {
		m := &converters.ScalarsModel{
			DoubleField:   1,
			FloatField:    2,
			Int32Field:    3,
			Int64Field:    4,
			Uint32Field:   5,
			Uint64Field:   6,
			Sint32Field:   7,
			Sint64Field:   8,
			Fixed32Field:  9,
			Fixed64Field:  10,
			Sfixed32Field: 11,
			Sfixed64Field: 12,
			BoolField:     true,
			StringField:   "abc",
			BytesField:    []byte("def"),
		}
		x, err := m.AsProto()
		require.NoError(t, err)
		require.NotNil(t, x)
		require.EqualValues(t, 1, x.DoubleField)
		require.EqualValues(t, 2, x.FloatField)
		require.EqualValues(t, 3, x.Int32Field)
		require.EqualValues(t, 4, x.Int64Field)
		require.EqualValues(t, 5, x.Uint32Field)
		require.EqualValues(t, 6, x.Uint64Field)
		require.EqualValues(t, 7, x.Sint32Field)
		require.EqualValues(t, 8, x.Sint64Field)
		require.EqualValues(t, 9, x.Fixed32Field)
		require.EqualValues(t, 10, x.Fixed64Field)
		require.EqualValues(t, 11, x.Sfixed32Field)
		require.EqualValues(t, 12, x.Sfixed64Field)
		require.True(t, x.BoolField)
		require.EqualValues(t, "abc", x.StringField)
		require.EqualValues(t, []byte("def"), x.BytesField)
	})
}

func TestScalars_AsModel(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		x := new(converters.Scalars)
		m, err := x.AsModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.EqualValues(t, 0, x.DoubleField)
		require.EqualValues(t, 0, x.FloatField)
		require.EqualValues(t, 0, x.Int32Field)
		require.EqualValues(t, 0, x.Int64Field)
		require.EqualValues(t, 0, x.Uint32Field)
		require.EqualValues(t, 0, x.Uint64Field)
		require.EqualValues(t, 0, x.Sint32Field)
		require.EqualValues(t, 0, x.Sint64Field)
		require.EqualValues(t, 0, x.Fixed32Field)
		require.EqualValues(t, 0, x.Fixed64Field)
		require.EqualValues(t, 0, x.Sfixed32Field)
		require.EqualValues(t, 0, x.Sfixed64Field)
		require.False(t, x.BoolField)
		require.Equal(t, "", x.StringField)
		require.Equal(t, []byte(nil), x.BytesField)
	})

	t.Run("with values set", func(t *testing.T) {
		x := &converters.Scalars{
			DoubleField:   1,
			FloatField:    2,
			Int32Field:    3,
			Int64Field:    4,
			Uint32Field:   5,
			Uint64Field:   6,
			Sint32Field:   7,
			Sint64Field:   8,
			Fixed32Field:  9,
			Fixed64Field:  10,
			Sfixed32Field: 11,
			Sfixed64Field: 12,
			BoolField:     true,
			StringField:   "abc",
			BytesField:    []byte("def"),
		}
		m, err := x.AsModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.EqualValues(t, 1, x.DoubleField)
		require.EqualValues(t, 2, x.FloatField)
		require.EqualValues(t, 3, x.Int32Field)
		require.EqualValues(t, 4, x.Int64Field)
		require.EqualValues(t, 5, x.Uint32Field)
		require.EqualValues(t, 6, x.Uint64Field)
		require.EqualValues(t, 7, x.Sint32Field)
		require.EqualValues(t, 8, x.Sint64Field)
		require.EqualValues(t, 9, x.Fixed32Field)
		require.EqualValues(t, 10, x.Fixed64Field)
		require.EqualValues(t, 11, x.Sfixed32Field)
		require.EqualValues(t, 12, x.Sfixed64Field)
		require.True(t, x.BoolField)
		require.Equal(t, "abc", x.StringField)
		require.Equal(t, []byte("def"), x.BytesField)
	})
}

func TestKnownTypes_AsModel(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		x := new(converters.KnownTypes)
		m, err := x.AsModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, time.Time{}, m.TimestampField)
	})

	t.Run("with values set", func(t *testing.T) {
		now := time.Now().UTC()
		nowpb := timestamppb.New(now)
		x := &converters.KnownTypes{
			TimestampField: nowpb,
		}
		m, err := x.AsModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, now, m.TimestampField)
	})
}

func TestKnownTypesModel_AsProto(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		m := new(converters.KnownTypesModel)
		x, err := m.AsProto()
		require.NoError(t, err)
		require.NotNil(t, x)
		require.Nil(t, x.TimestampField)
	})

	t.Run("with values set", func(t *testing.T) {
		now := time.Now().UTC()
		nowpb := timestamppb.New(now)
		m := &converters.KnownTypesModel{
			TimestampField: now,
		}
		x, err := m.AsProto()
		require.NoError(t, err)
		require.NotNil(t, x)
		requirepb.EqualProtos(t, nowpb, x.TimestampField)
	})
}

func TestEnum_AsModel(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		x := new(converters.Enum)
		m, err := x.AsModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.EqualValues(t, 0, m.EnumField)
	})

	t.Run("with values set", func(t *testing.T) {
		x := &converters.Enum{
			EnumField:       converters.AnEnum_AN_ENUM_VALUE,
			NestedEnumField: converters.Enum_A_NESTED_ENUM_VALUE,
		}
		m, err := x.AsModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.EqualValues(t, converters.AnEnum_AN_ENUM_VALUE, m.EnumField)
		require.EqualValues(t, converters.Enum_A_NESTED_ENUM_VALUE, m.NestedEnumField)
	})
}

func TestEnumModel_AsProto(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		m := new(converters.EnumModel)
		x, err := m.AsProto()
		require.NoError(t, err)
		require.NotNil(t, x)
		require.Equal(t, converters.AnEnum_AN_ENUM_UNSPECIFIED, x.EnumField)
		require.Equal(t, converters.Enum_A_NESTED_ENUM_UNSPECIFIED, x.NestedEnumField)
	})

	t.Run("with values set", func(t *testing.T) {
		m := &converters.EnumModel{
			EnumField:       int32(converters.AnEnum_AN_ENUM_VALUE),
			NestedEnumField: int32(converters.Enum_A_NESTED_ENUM_VALUE),
		}
		x, err := m.AsProto()
		require.NoError(t, err)
		require.NotNil(t, x)
		require.EqualValues(t, converters.AnEnum_AN_ENUM_VALUE, x.EnumField)
		require.EqualValues(t, converters.Enum_A_NESTED_ENUM_VALUE, x.NestedEnumField)
	})
}

func TestJson_AsModel(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		x := new(converters.Json)
		m, err := x.AsModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.EqualValues(t, []byte("null"), m.MapField)
	})

	t.Run("with values set", func(t *testing.T) {
		x := &converters.Json{
			MapField: map[string]string{"foo": "bar"},
		}
		m, err := x.AsModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, []byte(`{"foo":"bar"}`), m.MapField)
	})
}

func TestJsonModel_AsProto(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		m := new(converters.JsonModel)
		x, err := m.AsProto()
		require.NoError(t, err)
		require.NotNil(t, x)
		require.Nil(t, x.MapField)
	})

	t.Run("with values set", func(t *testing.T) {
		m := &converters.JsonModel{
			MapField: []byte(`{"foo":"bar"}`),
		}
		x, err := m.AsProto()
		require.NoError(t, err)
		require.NotNil(t, x)
		require.Equal(t, map[string]string{"foo": "bar"}, x.MapField)
	})
}
