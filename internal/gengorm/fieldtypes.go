package gengorm

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewFieldType(field *Field) (*FieldType, error) {
	types := &FieldType{
		field:   field,
		Go:      field.proto.GoIdent,
		JSON:    field.opts.Json,
		Pointer: field.proto.Desc.HasPresence(),
	}
	if err := types.init(); err != nil {
		return nil, err
	}
	return types, nil
}

type FieldType struct {
	field   *Field
	Go      protogen.GoIdent
	JSON    bool
	Pointer bool

	Gorm     protogen.GoIdent
	Custom   bool
	External bool
}

func (t *FieldType) init() error {
	if t.JSON {
		t.Gorm.GoName = "string"
		return nil
	}

	switch t.field.proto.Desc.Kind() {
	case protoreflect.BoolKind:
		t.Gorm.GoName = "bool"
	case protoreflect.EnumKind:
		t.Gorm.GoName = "int32"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		t.Gorm.GoName = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		t.Gorm.GoName = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		t.Gorm.GoName = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		t.Gorm.GoName = "uint64"
	case protoreflect.FloatKind:
		t.Gorm.GoName = "float32"
	case protoreflect.DoubleKind:
		t.Gorm.GoName = "float64"
	case protoreflect.StringKind:
		t.Gorm.GoName = "string"
	case protoreflect.BytesKind:
		t.Gorm.GoName = "[]byte"
		t.Pointer = false

	case protoreflect.MessageKind, protoreflect.GroupKind:
		if t.isTimestamp() {
			t.Gorm.GoName = "Time"
			t.Gorm.GoImportPath = "time"
			t.Pointer = false
			return nil
		}
		nested := t.field.proto.Message
		t.Go = nested.GoIdent

		filePkg := t.field.msg.file.proto.GoImportPath
		fieldPkg := nested.GoIdent.GoImportPath
		t.External = filePkg != fieldPkg
	}

	if unmapped := t.Gorm.GoName == ""; unmapped {
		t.Custom = true
	}

	switch {
	case t.Custom && t.External:
		panic("TODO: External custom types")
		// t.Gorm.GoName = t.alias()

	case t.Custom && !t.External:
		panic("TODO: External custom types")
		// t.Gorm.GoName = t.Go.GoName
	}
	return nil
}

func (t *FieldType) String() string {
	if t.Gorm.GoImportPath != "" {
		return t.field.msg.file.out.QualifiedGoIdent(t.Gorm)
	}
	return t.Gorm.GoName
}

func (t *FieldType) alias() string {
	return fmt.Sprintf("%s%s", t.field.msg.ModelName(), t.Go.GoName)
}

func (t *FieldType) isTimestamp() bool {
	var (
		path = t.field.proto.Message.GoIdent.GoImportPath
		name = t.field.proto.Message.GoIdent.GoName
	)
	return path == "google.golang.org/protobuf/types/known/timestamppb" &&
		name == "Timestamp"
}

func (t *FieldType) Gen() {
	// TODO: Call
	// TODO: Generate fully qualified imports
	// TODO: Generate custom types
}
