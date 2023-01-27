// Code generated by protoc-gen-gorm. DO NOT EDIT.
// versions:
// 	protoc-gen-gorm 2.0.0
// 	protoc          (unknown)
// source: fieldtypes/fieldtypes.proto

package fieldtypes

import (
	sql "database/sql"
	_ "github.com/complex64/protoc-gen-gorm/gormpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

// WithScalarValuesModel is the GORM model for fieldtypes.WithScalarValues.
type WithScalarValuesModel struct {
	DoubleField   float64
	FloatField    float32
	Int32Field    int32
	Int64Field    int64
	Uint32Field   uint32
	Uint64Field   uint64
	Sint32Field   int32
	Sint64Field   int64
	Fixed32Field  uint32
	Fixed64Field  uint64
	Sfixed32Field int32
	Sfixed64Field int64
	BoolField     bool
	StringField   string
	BytesField    []byte
}

// ToProto converts a WithScalarValuesModel to its protobuf representation.
func (m *WithScalarValuesModel) ToProto() (*WithScalarValues, error) {
	p := new(WithScalarValues)
	p.DoubleField = m.DoubleField
	p.FloatField = m.FloatField
	p.Int32Field = m.Int32Field
	p.Int64Field = m.Int64Field
	p.Uint32Field = m.Uint32Field
	p.Uint64Field = m.Uint64Field
	p.Sint32Field = m.Sint32Field
	p.Sint64Field = m.Sint64Field
	p.Fixed32Field = m.Fixed32Field
	p.Fixed64Field = m.Fixed64Field
	p.Sfixed32Field = m.Sfixed32Field
	p.Sfixed64Field = m.Sfixed64Field
	p.BoolField = m.BoolField
	p.StringField = m.StringField
	p.BytesField = m.BytesField
	return p, nil
}

// ToModel converts a WithScalarValues to its GORM model.
func (p *WithScalarValues) ToModel() (*WithScalarValuesModel, error) {
	m := new(WithScalarValuesModel)
	m.DoubleField = p.DoubleField
	m.FloatField = p.FloatField
	m.Int32Field = p.Int32Field
	m.Int64Field = p.Int64Field
	m.Uint32Field = p.Uint32Field
	m.Uint64Field = p.Uint64Field
	m.Sint32Field = p.Sint32Field
	m.Sint64Field = p.Sint64Field
	m.Fixed32Field = p.Fixed32Field
	m.Fixed64Field = p.Fixed64Field
	m.Sfixed32Field = p.Sfixed32Field
	m.Sfixed64Field = p.Sfixed64Field
	m.BoolField = p.BoolField
	m.StringField = p.StringField
	m.BytesField = p.BytesField
	return m, nil
}

// WithKnownTypesModel is the GORM model for fieldtypes.WithKnownTypes.
type WithKnownTypesModel struct {
	TimestampField sql.NullTime
}

// ToProto converts a WithKnownTypesModel to its protobuf representation.
func (m *WithKnownTypesModel) ToProto() (*WithKnownTypes, error) {
	p := new(WithKnownTypes)
	if m.TimestampField.Valid && m.TimestampField.Time != (time.Time{}) {
		p.TimestampField = timestamppb.New(m.TimestampField.Time)
	}
	return p, nil
}

// ToModel converts a WithKnownTypes to its GORM model.
func (p *WithKnownTypes) ToModel() (*WithKnownTypesModel, error) {
	m := new(WithKnownTypesModel)
	if t := p.TimestampField; t != nil {
		m.TimestampField = sql.NullTime{
			Valid: true,
			Time:  t.AsTime(),
		}
	}
	return m, nil
}

// WithEnumModel is the GORM model for fieldtypes.WithEnum.
type WithEnumModel struct {
	EnumField int32
}

// ToProto converts a WithEnumModel to its protobuf representation.
func (m *WithEnumModel) ToProto() (*WithEnum, error) {
	p := new(WithEnum)
	p.EnumField = AnEnum(m.EnumField)
	return p, nil
}

// ToModel converts a WithEnum to its GORM model.
func (p *WithEnum) ToModel() (*WithEnumModel, error) {
	m := new(WithEnumModel)
	m.EnumField = int32(p.EnumField)
	return m, nil
}
