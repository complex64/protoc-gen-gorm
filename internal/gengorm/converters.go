package gengorm

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (m *Message) genConverters() {
	m.genModelToProto()
	m.genToModel()
}

func (m *Message) genModelToProto() {
	m.P(Comment(" ToProto converts a %s to its protobuf representation.", m.ModelName()),
		"func (m *", m.ModelName(), ") ToProto() (*", m.proto.GoIdent.GoName, ", error) {")
	m.P("p := new(", m.proto.GoIdent.GoName, ")")
	m.genModelToProtoFields()
	m.P("return p, nil")
	m.P("}") // func
	m.P()
}

func (m *Message) genModelToProtoFields() {
	for _, field := range m.fields {
		field.genConvertToProto()
	}
}

func (f *Field) genConvertToProto() {
	switch {
	case f.types.JSON:
		f.genConvertJsonToProto()
	case f.types.Enum:
		f.genEnumToProto()
	case f.types.isTimestamp():
		f.genConvertTimeToProto()
	case f.types.isDoubleValueWrapper():
		f.genConvertDoubleValueWrapperToProto()
	case f.types.isFloatValueWrapper():
		f.genConvertFloatValueWrapperToProto()
	case f.types.isInt64ValueWrapper():
		f.genConvertInt64ValueWrapperToProto()
	case f.types.isUInt64ValueWrapper():
		f.genConvertUInt64ValueWrapperToProto()
	case f.types.isInt32ValueWrapper():
		f.genConvertInt32ValueWrapperToProto()
	case f.types.isUInt32ValueWrapper():
		f.genConvertUInt32ValueWrapperToProto()
	case f.types.isBoolValueWrapper():
		f.genConvertBoolValueWrapperToProto()
	case f.types.isStringValueWrapper():
		f.genConvertStringValueWrapperToProto()
	case f.types.isBytesValueWrapper():
		f.genConvertBytesValueWrapperToProto()
	case f.types.Pointer:
		f.P("p.", f.Name(), " = *m.", f.Name())
	default:
		f.P("p.", f.Name(), " = m.", f.Name())
	}
}

func (f *Field) genConvertJsonToProto() {
	unmarshal := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Unmarshal",
		GoImportPath: "encoding/json",
	})
	f.P("if len(m.", f.Name(), ") > 0 {")
	f.P("if err := ", unmarshal, "(m.", f.Name(), ", &p.", f.Name(), "); err != nil {")
	f.P("return nil, err")
	f.P("}")
	f.P("}")
}

func (f *Field) genConvertTimeToProto() {
	newTimestamp := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "New",
		GoImportPath: KnownTypesTimestampPkg,
	})
	timeType := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Time",
		GoImportPath: "time",
	})
	f.P("if m.", f.Name(), ".Valid && m.", f.Name(), ".Time != (", timeType, "{}) {")
	f.P("p.", f.Name(), " = ", newTimestamp, "(m.", f.Name(), ".Time)")
	f.P("}")
}

func (f *Field) genConvertDoubleValueWrapperToProto() {
	newWrapper := f.wrappedType("Double")
	f.P("if m.", f.Name(), ".Valid {")
	f.P("p.", f.Name(), " = ", newWrapper, "(m.", f.Name(), ".Float64)")
	f.P("}")
}

func (f *Field) genConvertFloatValueWrapperToProto() {
	newWrapper := f.wrappedType("Float")
	f.P("if m.", f.Name(), ".Valid {")
	f.P("p.", f.Name(), " = ", newWrapper, "(float32(m.", f.Name(), ".Float64))")
	f.P("}")
}

func (f *Field) genConvertInt64ValueWrapperToProto() {
	newWrapper := f.wrappedType("Int64")
	f.P("if m.", f.Name(), ".Valid {")
	f.P("p.", f.Name(), " = ", newWrapper, "(m.", f.Name(), ".Int64)")
	f.P("}")
}

func (f *Field) genConvertUInt64ValueWrapperToProto() {
	newWrapper := f.wrappedType("UInt64")
	f.P("if m.", f.Name(), ".Valid {")
	f.P("p.", f.Name(), " = ", newWrapper, "(uint64(m.", f.Name(), ".Int64))")
	f.P("}")
}

func (f *Field) genConvertInt32ValueWrapperToProto() {
	newWrapper := f.wrappedType("Int32")
	f.P("if m.", f.Name(), ".Valid {")
	f.P("p.", f.Name(), " = ", newWrapper, "(int32(m.", f.Name(), ".Int32))")
	f.P("}")
}

func (f *Field) genConvertUInt32ValueWrapperToProto() {
	newWrapper := f.wrappedType("UInt32")
	f.P("if m.", f.Name(), ".Valid {")
	f.P("p.", f.Name(), " = ", newWrapper, "(uint32(m.", f.Name(), ".Int64))")
	f.P("}")
}

func (f *Field) genConvertBoolValueWrapperToProto() {
	newWrapper := f.wrappedType("Bool")
	f.P("if m.", f.Name(), ".Valid {")
	f.P("p.", f.Name(), " = ", newWrapper, "(m.", f.Name(), ".Bool)")
	f.P("}")
}

func (f *Field) genConvertStringValueWrapperToProto() {
	newWrapper := f.wrappedType("String")
	f.P("if m.", f.Name(), ".Valid {")
	f.P("p.", f.Name(), " = ", newWrapper, "(m.", f.Name(), ".String)")
	f.P("}")
}

func (f *Field) genConvertBytesValueWrapperToProto() {
	newWrapper := f.wrappedType("Bytes")
	f.P("if m.", f.Name(), "!= nil {")
	f.P("p.", f.Name(), " = ", newWrapper, "(m.", f.Name(), ")")
	f.P("}")
}

func (f *Field) genEnumToProto() {
	typename := f.proto.Desc.Enum().Name()
	parent := f.proto.Desc.Enum().Parent()
	if p, ok := parent.(protoreflect.MessageDescriptor); ok {
		typename = p.Name() + "_" + typename
	}
	f.P("p.", f.Name(), " = ", typename, "(m.", f.Name(), ")")
}

func (m *Message) genToModel() {
	m.P(Comment(" ToModel converts a %s to its GORM model.", m.proto.GoIdent.GoName),
		"func (p *", m.proto.GoIdent.GoName, ") ToModel() (*", m.ModelName(), ", error) {")
	m.P("m := new(", m.ModelName(), ")")
	m.genConvertToModelFields()
	m.P("return m, nil")
	m.P("}")
	m.P()
}

func (m *Message) genConvertToModelFields() {
	for _, field := range m.fields {
		field.genConvertToModel()
	}
}

func (f *Field) genConvertToModel() {
	switch {
	case f.types.JSON:
		f.genConvertJsonToModel()
	case f.types.Enum:
		f.genEnumToModel()
	case f.types.isTimestamp():
		f.genConvertTimestampToModel()
	case f.types.isDoubleValueWrapper():
		f.genConvertDoubleValueWrapperToModel()
	case f.types.isFloatValueWrapper():
		f.genConvertFloatValueWrapperToModel()
	case f.types.isInt64ValueWrapper():
		f.genConvertInt64ValueWrapperToModel()
	case f.types.isUInt64ValueWrapper():
		f.genConvertUInt64ValueWrapperToModel()
	case f.types.isInt32ValueWrapper():
		f.genConvertInt32ValueWrapperToModel()
	case f.types.isUInt32ValueWrapper():
		f.genConvertUInt32ValueWrapperToModel()
	case f.types.isBoolValueWrapper():
		f.genConvertBoolValueWrapperToModel()
	case f.types.isStringValueWrapper():
		f.genConvertStringValueWrapperToModel()
	case f.types.isBytesValueWrapper():
		f.genConvertBytesValueWrapperToModel()
	case f.types.Pointer:
		f.P("m.", f.Name(), " = *p.", f.Name())
	default:
		f.P("m.", f.Name(), " = p.", f.Name())
	}
}

func (f *Field) genConvertJsonToModel() {
	marshal := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Marshal",
		GoImportPath: "encoding/json",
	})
	f.P("if bs, err := ", marshal, "(&p.", f.Name(), "); err != nil {")
	f.P("return nil, err")
	f.P("} else {")
	f.P("m.", f.Name(), " = bs")
	f.P("}") // if
}

func (f *Field) genEnumToModel() {
	f.P("m.", f.Name(), " = ", "int32(p.", f.Name(), ")")
}

func (f *Field) genConvertTimestampToModel() {
	nullTime := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "NullTime",
		GoImportPath: "database/sql",
	})
	f.P("if t := p.", f.Name(), "; t != nil {")
	f.P("m.", f.Name(), " = ", nullTime, "{")
	f.P("Valid: true,")
	f.P("Time: t.AsTime(),")
	f.P("}")
	f.P("}")
}

func (f *Field) genConvertDoubleValueWrapperToModel() {
	f.P("if p.", f.Name(), " != nil {")
	f.P("m.", f.Name(), " = ", f.sqlNullableType("NullFloat64"), " {")
	f.P("Valid: true,")
	f.P("Float64: p.", f.Name(), ".GetValue(),")
	f.P("}")
	f.P("}")
}

func (f *Field) genConvertFloatValueWrapperToModel() {
	f.P("if p.", f.Name(), " != nil {")
	f.P("m.", f.Name(), " = ", f.sqlNullableType("NullFloat64"), " {")
	f.P("Valid: true,")
	f.P("Float64: float64(p.", f.Name(), ".GetValue()),")
	f.P("}")
	f.P("}")
}

func (f *Field) genConvertInt64ValueWrapperToModel() {
	f.P("if p.", f.Name(), " != nil {")
	f.P("m.", f.Name(), " = ", f.sqlNullableType("NullInt64"), " {")
	f.P("Valid: true,")
	f.P("Int64: p.", f.Name(), ".GetValue(),")
	f.P("}")
	f.P("}")
}

func (f *Field) genConvertUInt64ValueWrapperToModel() {
	f.P("if p.", f.Name(), " != nil {")
	f.P("m.", f.Name(), " = ", f.sqlNullableType("NullInt64"), " {")
	f.P("Valid: true,")
	f.P("Int64: int64(p.", f.Name(), ".GetValue()),")
	f.P("}")
	f.P("}")
}

func (f *Field) genConvertInt32ValueWrapperToModel() {
	f.P("if p.", f.Name(), " != nil {")
	f.P("m.", f.Name(), " = ", f.sqlNullableType("NullInt32"), " {")
	f.P("Valid: true,")
	f.P("Int32: p.", f.Name(), ".GetValue(),")
	f.P("}")
	f.P("}")
}

func (f *Field) genConvertUInt32ValueWrapperToModel() {
	f.P("if p.", f.Name(), " != nil {")
	f.P("m.", f.Name(), " = ", f.sqlNullableType("NullInt64"), " {")
	f.P("Valid: true,")
	f.P("Int64: int64(p.", f.Name(), ".GetValue()),")
	f.P("}")
	f.P("}")
}

func (f *Field) genConvertBoolValueWrapperToModel() {
	f.P("if p.", f.Name(), " != nil {")
	f.P("m.", f.Name(), " = ", f.sqlNullableType("NullBool"), " {")
	f.P("Valid: true,")
	f.P("Bool: p.", f.Name(), ".GetValue(),")
	f.P("}")
	f.P("}")
}

func (f *Field) genConvertStringValueWrapperToModel() {
	f.P("if p.", f.Name(), " != nil {")
	f.P("m.", f.Name(), " = ", f.sqlNullableType("NullString"), " {")
	f.P("Valid: true,")
	f.P("String: p.", f.Name(), ".GetValue(),")
	f.P("}")
	f.P("}")
}

func (f *Field) genConvertBytesValueWrapperToModel() {
	f.P("if p.", f.Name(), " != nil {")
	f.P("m.", f.Name(), " = p.", f.Name(), ".GetValue()")
	f.P("}")
}

func (f *Field) wrappedType(name string) string {
	return f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       name,
		GoImportPath: KnownTypesWrappersPkg,
	})
}

func (f *Field) sqlNullableType(name string) string {
	return f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       name,
		GoImportPath: "database/sql",
	})
}
