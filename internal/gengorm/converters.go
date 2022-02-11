package gengorm

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func (m *Message) genConverters() {
	m.genModelToProto()
	m.genToModel()
}

func (m *Message) genModelToProto() {
	m.P(Comment(" ToProto converts a %s to its protobuf representation.", m.ModelName()),
		"func (m *", m.ModelName(), ") ToProto() (*", m.proto.GoIdent.GoName, ", error) {")
	m.P("x := new(", m.proto.GoIdent.GoName, ")")
	m.genModelToProtoFields()
	m.P("return x, nil")
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
	case f.types.Pointer:
		f.P("x.", f.Name(), " = *m.", f.Name())
	default:
		f.P("x.", f.Name(), " = m.", f.Name())
	}
}

func (f *Field) genConvertJsonToProto() {
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

func (f *Field) genConvertTimeToProto() {
	newTimestamp := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "New",
		GoImportPath: "google.golang.org/protobuf/types/known/timestamppb",
	})
	f.P("if m.", f.Name(), " != (time.Time{}) {")
	f.P("x.", f.Name(), " = ", newTimestamp, "(m.", f.Name(), ")")
	f.P("}")
}

func (f *Field) genEnumToProto() {
	f.P("x.", f.Name(), " = ", f.proto.Desc.Enum().Name(), "(m.", f.Name(), ")")
}

func (m *Message) genToModel() {
	m.P(Comment(" ToModel converts a %s to its GORM model.", m.proto.GoIdent.GoName),
		"func (x *", m.proto.GoIdent.GoName, ") ToModel() (*", m.ModelName(), ", error) {")
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
	case f.types.Pointer:
		f.P("m.", f.Name(), " = *x.", f.Name())
	default:
		f.P("m.", f.Name(), " = x.", f.Name())
	}
}

func (f *Field) genConvertJsonToModel() {
	marshal := f.msg.file.out.QualifiedGoIdent(protogen.GoIdent{
		GoName:       "Marshal",
		GoImportPath: "encoding/json",
	})
	f.P("{")
	f.P("if bs, err := ", marshal, "(&x.", f.Name(), "); err != nil {")
	f.P("return nil, err")
	f.P("} else {")
	f.P("m.", f.Name(), " = bs")
	f.P("}") // if
	f.P("}")
}

func (f *Field) genEnumToModel() {
	f.P("m.", f.Name(), " = ", "int32(x.", f.Name(), ")")
}

func (f *Field) genConvertTimestampToModel() {
	f.P("if t := x.", f.Name(), "; t != nil {")
	f.P("m.", f.Name(), " = t.AsTime()")
	f.P("}")
}
