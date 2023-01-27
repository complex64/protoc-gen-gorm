package converters_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/complex64/protoc-gen-gorm/internal/gengorm/test/converters"
	requirepb "github.com/complex64/protoc-gen-gorm/internal/require"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestScalarsModel_ToProto(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		m := new(converters.ScalarsModel)
		p, err := m.ToProto()
		require.NoError(t, err)
		require.NotNil(t, p)
		require.EqualValues(t, 0, p.DoubleField)
		require.EqualValues(t, 0, p.FloatField)
		require.EqualValues(t, 0, p.Int32Field)
		require.EqualValues(t, 0, p.Int64Field)
		require.EqualValues(t, 0, p.Uint32Field)
		require.EqualValues(t, 0, p.Uint64Field)
		require.EqualValues(t, 0, p.Sint32Field)
		require.EqualValues(t, 0, p.Sint64Field)
		require.EqualValues(t, 0, p.Fixed32Field)
		require.EqualValues(t, 0, p.Fixed64Field)
		require.EqualValues(t, 0, p.Sfixed32Field)
		require.EqualValues(t, 0, p.Sfixed64Field)
		require.False(t, p.BoolField)
		require.EqualValues(t, "", p.StringField)
		require.EqualValues(t, []byte(nil), p.BytesField)
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
		p, err := m.ToProto()
		require.NoError(t, err)
		require.NotNil(t, p)
		require.EqualValues(t, 1, p.DoubleField)
		require.EqualValues(t, 2, p.FloatField)
		require.EqualValues(t, 3, p.Int32Field)
		require.EqualValues(t, 4, p.Int64Field)
		require.EqualValues(t, 5, p.Uint32Field)
		require.EqualValues(t, 6, p.Uint64Field)
		require.EqualValues(t, 7, p.Sint32Field)
		require.EqualValues(t, 8, p.Sint64Field)
		require.EqualValues(t, 9, p.Fixed32Field)
		require.EqualValues(t, 10, p.Fixed64Field)
		require.EqualValues(t, 11, p.Sfixed32Field)
		require.EqualValues(t, 12, p.Sfixed64Field)
		require.True(t, p.BoolField)
		require.EqualValues(t, "abc", p.StringField)
		require.EqualValues(t, []byte("def"), p.BytesField)
	})
}

func TestScalars_ToModel(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		p := new(converters.Scalars)
		m, err := p.ToModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.EqualValues(t, 0, p.DoubleField)
		require.EqualValues(t, 0, p.FloatField)
		require.EqualValues(t, 0, p.Int32Field)
		require.EqualValues(t, 0, p.Int64Field)
		require.EqualValues(t, 0, p.Uint32Field)
		require.EqualValues(t, 0, p.Uint64Field)
		require.EqualValues(t, 0, p.Sint32Field)
		require.EqualValues(t, 0, p.Sint64Field)
		require.EqualValues(t, 0, p.Fixed32Field)
		require.EqualValues(t, 0, p.Fixed64Field)
		require.EqualValues(t, 0, p.Sfixed32Field)
		require.EqualValues(t, 0, p.Sfixed64Field)
		require.False(t, p.BoolField)
		require.Equal(t, "", p.StringField)
		require.Equal(t, []byte(nil), p.BytesField)
	})

	t.Run("with values set", func(t *testing.T) {
		p := &converters.Scalars{
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
		m, err := p.ToModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.EqualValues(t, 1, p.DoubleField)
		require.EqualValues(t, 2, p.FloatField)
		require.EqualValues(t, 3, p.Int32Field)
		require.EqualValues(t, 4, p.Int64Field)
		require.EqualValues(t, 5, p.Uint32Field)
		require.EqualValues(t, 6, p.Uint64Field)
		require.EqualValues(t, 7, p.Sint32Field)
		require.EqualValues(t, 8, p.Sint64Field)
		require.EqualValues(t, 9, p.Fixed32Field)
		require.EqualValues(t, 10, p.Fixed64Field)
		require.EqualValues(t, 11, p.Sfixed32Field)
		require.EqualValues(t, 12, p.Sfixed64Field)
		require.True(t, p.BoolField)
		require.Equal(t, "abc", p.StringField)
		require.Equal(t, []byte("def"), p.BytesField)
	})
}

func TestKnownTypes_ToModel(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		p := new(converters.KnownTypes)
		m, err := p.ToModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, sql.NullTime{}, m.TimestampField)
	})

	t.Run("with values set", func(t *testing.T) {
		now := time.Now().UTC()
		p := &converters.KnownTypes{
			TimestampField: timestamppb.New(now),
		}
		m, err := p.ToModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, sql.NullTime{
			Time:  now,
			Valid: true,
		}, m.TimestampField)
	})
}

func TestKnownTypesModel_ToProto(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		m := new(converters.KnownTypesModel)
		p, err := m.ToProto()
		require.NoError(t, err)
		require.NotNil(t, p)

		require.Nil(t, p.TimestampField)
		require.Equal(t, sql.NullFloat64{Float64: 0, Valid: false}, m.WrappedDouble)
		require.Equal(t, sql.NullFloat64{Float64: 0, Valid: false}, m.WrappedFloat)
		require.Equal(t, sql.NullInt64{Int64: 0, Valid: false}, m.WrappedInt64)
		require.Equal(t, sql.NullInt64{Int64: 0, Valid: false}, m.WrappedUin64)
		require.Equal(t, sql.NullInt32{Int32: 0, Valid: false}, m.WrappedInt32)
		require.Equal(t, sql.NullInt64{Int64: 0, Valid: false}, m.WrappedUint32)
		require.Equal(t, sql.NullBool{Bool: false, Valid: false}, m.WrappedBool)
		require.Equal(t, sql.NullString{String: "", Valid: false}, m.WrappedString)
		require.Nil(t, m.WrappedBytes)
	})

	t.Run("with values set", func(t *testing.T) {
		now := time.Now().UTC()
		nowpb := timestamppb.New(now)
		m := &converters.KnownTypesModel{
			TimestampField: sql.NullTime{
				Time:  now,
				Valid: true,
			},
			WrappedDouble: sql.NullFloat64{
				Float64: 1.0,
				Valid:   true,
			},
			WrappedFloat: sql.NullFloat64{
				Float64: 1.0,
				Valid:   true,
			},
			WrappedInt64: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
			WrappedUin64: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
			WrappedInt32: sql.NullInt32{
				Int32: 1,
				Valid: true,
			},
			WrappedUint32: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
			WrappedBool: sql.NullBool{
				Bool:  true,
				Valid: true,
			},
			WrappedString: sql.NullString{
				String: "",
				Valid:  false,
			},
			WrappedBytes: []byte("s"),
		}
		p, err := m.ToProto()
		require.NoError(t, err)
		require.NotNil(t, p)
		requirepb.EqualProtos(t, nowpb, p.TimestampField)
	})
}

func TestEnum_ToModel(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		p := new(converters.Enum)
		m, err := p.ToModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.EqualValues(t, 0, m.EnumField)
	})

	t.Run("with values set", func(t *testing.T) {
		p := &converters.Enum{
			EnumField:       converters.AnEnum_AN_ENUM_VALUE,
			NestedEnumField: converters.Enum_A_NESTED_ENUM_VALUE,
		}
		m, err := p.ToModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.EqualValues(t, converters.AnEnum_AN_ENUM_VALUE, m.EnumField)
		require.EqualValues(t, converters.Enum_A_NESTED_ENUM_VALUE, m.NestedEnumField)
	})
}

func TestEnumModel_ToProto(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		m := new(converters.EnumModel)
		p, err := m.ToProto()
		require.NoError(t, err)
		require.NotNil(t, p)
		require.Equal(t, converters.AnEnum_AN_ENUM_UNSPECIFIED, p.EnumField)
		require.Equal(t, converters.Enum_A_NESTED_ENUM_UNSPECIFIED, p.NestedEnumField)
	})

	t.Run("with values set", func(t *testing.T) {
		m := &converters.EnumModel{
			EnumField:       int32(converters.AnEnum_AN_ENUM_VALUE),
			NestedEnumField: int32(converters.Enum_A_NESTED_ENUM_VALUE),
		}
		p, err := m.ToProto()
		require.NoError(t, err)
		require.NotNil(t, p)
		require.EqualValues(t, converters.AnEnum_AN_ENUM_VALUE, p.EnumField)
		require.EqualValues(t, converters.Enum_A_NESTED_ENUM_VALUE, p.NestedEnumField)
	})
}

func TestJson_ToModel(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		p := new(converters.Json)
		m, err := p.ToModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.EqualValues(t, []byte("null"), m.MapField)
	})

	t.Run("with values set", func(t *testing.T) {
		p := &converters.Json{
			MapField: map[string]string{"foo": "bar"},
		}
		m, err := p.ToModel()
		require.NoError(t, err)
		require.NotNil(t, m)
		require.Equal(t, []byte(`{"foo":"bar"}`), m.MapField)
	})
}

func TestJsonModel_ToProto(t *testing.T) {
	t.Run("with defaults", func(t *testing.T) {
		m := new(converters.JsonModel)
		p, err := m.ToProto()
		require.NoError(t, err)
		require.NotNil(t, p)
		require.Nil(t, p.MapField)
	})

	t.Run("with values set", func(t *testing.T) {
		m := &converters.JsonModel{
			MapField: []byte(`{"foo":"bar"}`),
		}
		p, err := m.ToProto()
		require.NoError(t, err)
		require.NotNil(t, p)
		require.Equal(t, map[string]string{"foo": "bar"}, p.MapField)
	})
}
