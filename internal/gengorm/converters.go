package gengorm

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func (m *Message) genConverters() {
	m.genModelAsProto()
	m.genAsModel()
}

func (m *Message) genModelAsProto() {
	m.P(Comment(" AsProto converts a %s to its protobuf representation.", m.ModelName()),
		"func (m *", m.ModelName(), ") AsProto() (*", m.proto.GoIdent.GoName, ", error) {")
	m.P("x := new(", m.proto.GoIdent.GoName, ")")
	m.genModelAsProtoFields()
	m.P("return x, nil")
	m.P("}") // func
	m.P()
}

func (m *Message) genModelAsProtoFields() {
	for _, field := range m.fields {
		field.genConvertAsProto()
	}
}

func (f *Field) genConvertAsProto() {
	switch {
	case f.types.JSON:
		f.genConvertJsonAsProto()
	case f.types.Enum:
		f.genEnumAsProto()
	case f.types.isTimestamp():
		f.genConvertTimeAsProto()
	case f.types.isDobuleValueWrapper():
		f.genConvertDoubleValueWrapperAsProto()
	case f.types.isFloatValueWrapper():
		f.genConvertFloatValueWrapperAsProto()
	case f.types.isInt64ValueWrapper():
		f.genConvertInt64ValueWrapperAsProto()
	case f.types.isUInt64ValueWrapper():
		f.genConvertUInt64ValueWrapperAsProto()
	case f.types.isInt32ValueWrapper():
		f.genConvertInt32ValueWrapperAsProto()
	case f.types.isUInt32ValueWrapper():
		f.genConvertUInt32ValueWrapperAsProto()
	case f.types.isBoolValueWrapper():
		f.genConvertBoolValueWrapperAsProto()
	case f.types.isStringValueWrapper():
		f.genConvertStringValueWrapperAsProto()
	case f.types.isBytesValueWrapper():
		f.genConvertBytesValueWrapperAsProto()
	case f.types.Pointer:
		f.P("x.", f.Name(), " = *m.", f.Name())
	default:
		f.P("x.", f.Name(), " = m.", f.Name())
	}
}

func (f *Field) genConvertJsonAsProto() {
	unmarshal := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Unmarshal",
		GoImportPath: "encoding/json",
	})
	f.P("if len(m.", f.Name(), ") > 0 {")
	f.P("if err := ", unmarshal, "(m.", f.Name(), ", &x.", f.Name(), "); err != nil {")
	f.P("return nil, err")
	f.P("}")
	f.P("}")
}

func (f *Field) genConvertTimeAsProto() {
	newTimestamp := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "New",
		GoImportPath: "google.golang.org/protobuf/types/known/timestamppb",
	})
	f.P("if m.", f.Name(), " != (time.Time{}) {")
	f.P("x.", f.Name(), " = ", newTimestamp, "(m.", f.Name(), ")")
	f.P("}")
}

func (f *Field) genConvertDoubleValueWrapperAsProto() {
	newWrapper := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Double",
		GoImportPath: "google.golang.org/protobuf/types/known/wrapperspb",
	})
	f.P("x.", f.Name(), " = ", newWrapper, "(m.", f.Name(), ")")
}

func (f *Field) genConvertFloatValueWrapperAsProto() {
	newWrapper := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Float",
		GoImportPath: "google.golang.org/protobuf/types/known/wrapperspb",
	})
	f.P("x.", f.Name(), " = ", newWrapper, "(m.", f.Name(), ")")
}

func (f *Field) genConvertInt64ValueWrapperAsProto() {
	newWrapper := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Int64",
		GoImportPath: "google.golang.org/protobuf/types/known/wrapperspb",
	})
	f.P("x.", f.Name(), " = ", newWrapper, "(m.", f.Name(), ")")
}

func (f *Field) genConvertUInt64ValueWrapperAsProto() {
	newWrapper := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "UInt64",
		GoImportPath: "google.golang.org/protobuf/types/known/wrapperspb",
	})
	f.P("x.", f.Name(), " = ", newWrapper, "(m.", f.Name(), ")")
}

func (f *Field) genConvertInt32ValueWrapperAsProto() {
	newWrapper := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Int32",
		GoImportPath: "google.golang.org/protobuf/types/known/wrapperspb",
	})
	f.P("x.", f.Name(), " = ", newWrapper, "(m.", f.Name(), ")")
}

func (f *Field) genConvertUInt32ValueWrapperAsProto() {
	newWrapper := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "UInt32",
		GoImportPath: "google.golang.org/protobuf/types/known/wrapperspb",
	})
	f.P("x.", f.Name(), " = ", newWrapper, "(m.", f.Name(), ")")
}

func (f *Field) genConvertBoolValueWrapperAsProto() {
	newWrapper := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Bool",
		GoImportPath: "google.golang.org/protobuf/types/known/wrapperspb",
	})
	f.P("x.", f.Name(), " = ", newWrapper, "(m.", f.Name(), ")")
}

func (f *Field) genConvertStringValueWrapperAsProto() {
	newWrapper := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "String",
		GoImportPath: "google.golang.org/protobuf/types/known/wrapperspb",
	})
	f.P("x.", f.Name(), " = ", newWrapper, "(m.", f.Name(), ")")
}

func (f *Field) genConvertBytesValueWrapperAsProto() {
	newWrapper := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Bytes",
		GoImportPath: "google.golang.org/protobuf/types/known/wrapperspb",
	})
	f.P("x.", f.Name(), " = ", newWrapper, "(m.", f.Name(), ")")
}

func (f *Field) genEnumAsProto() {
	typename := f.proto.Desc.Enum().Name()
	parent := f.proto.Desc.Enum().Parent()
	if x, ok := parent.(protoreflect.MessageDescriptor); ok {
		typename = x.Name() + "_" + typename
	}
	f.P("x.", f.Name(), " = ", typename, "(m.", f.Name(), ")")
}

func (m *Message) genAsModel() {
	m.P(Comment(" AsModel converts a %s to its GORM model.", m.proto.GoIdent.GoName),
		"func (x *", m.proto.GoIdent.GoName, ") AsModel() (*", m.ModelName(), ", error) {")
	m.P("m := new(", m.ModelName(), ")")
	m.genConvertAsModelFields()
	m.P("return m, nil")
	m.P("}")
	m.P()
}

func (m *Message) genConvertAsModelFields() {
	for _, field := range m.fields {
		field.genConvertAsModel()
	}
}

func (f *Field) genConvertAsModel() {
	switch {
	case f.types.JSON:
		f.genConvertJsonAsModel()
	case f.types.Enum:
		f.genEnumAsModel()
	case f.types.isTimestamp():
		f.genConvertTimestampAsModel()
	case f.types.isDobuleValueWrapper(), f.types.isFloatValueWrapper(),
		f.types.isInt64ValueWrapper(), f.types.isUInt64ValueWrapper(),
		f.types.isInt32ValueWrapper(), f.types.isUInt32ValueWrapper(),
		f.types.isBoolValueWrapper(), f.types.isStringValueWrapper(),
		f.types.isBytesValueWrapper():
		f.P("m.", f.Name(), " = x.", f.Name(), ".GetValue()")
	case f.types.Pointer:
		f.P("m.", f.Name(), " = *x.", f.Name())
	default:
		f.P("m.", f.Name(), " = x.", f.Name())
	}
}

func (f *Field) genConvertJsonAsModel() {
	marshal := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Marshal",
		GoImportPath: "encoding/json",
	})
	f.P("if bs, err := ", marshal, "(&x.", f.Name(), "); err != nil {")
	f.P("return nil, err")
	f.P("} else {")
	f.P("m.", f.Name(), " = bs")
	f.P("}") // if
}

func (f *Field) genEnumAsModel() {
	f.P("m.", f.Name(), " = ", "int32(x.", f.Name(), ")")
}

func (f *Field) genConvertTimestampAsModel() {
	f.P("if t := x.", f.Name(), "; t != nil {")
	f.P("m.", f.Name(), " = t.AsTime()")
	f.P("}")
}
