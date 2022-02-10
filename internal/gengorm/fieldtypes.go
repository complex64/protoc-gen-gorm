package gengorm

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func NewFieldType(field *Field) (*FieldType, error) {
	types := &FieldType{field: field}
	if err := types.init(); err != nil {
		return nil, err
	}
	return types, nil
}

type FieldType struct {
	field *Field

	golang  protogen.GoIdent
	gorm    protogen.GoIdent
	custom  bool
	pointer bool
	json    bool
}

func (ft *FieldType) init() error {

	var (
		field = ft.field
		proto = field.proto
		desc  = proto.Desc
	)
	ft.golang = proto.GoIdent
	ft.pointer = desc.HasPresence()
	ft.json = field.opts.Json

	switch desc.Kind() {
	case protoreflect.BoolKind:
		ft.gorm.GoName = "bool"
	case protoreflect.EnumKind:
		ft.gorm.GoName = "int32"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		ft.gorm.GoName = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		ft.gorm.GoName = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		ft.gorm.GoName = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		ft.gorm.GoName = "uint64"
	case protoreflect.FloatKind:
		ft.gorm.GoName = "float32"
	case protoreflect.DoubleKind:
		ft.gorm.GoName = "float64"
	case protoreflect.StringKind:
		ft.gorm.GoName = "string"
	case protoreflect.BytesKind:
		ft.gorm.GoName = "[]byte"
		ft.pointer = false

	case protoreflect.MessageKind, protoreflect.GroupKind:
		if ft.isTimestamp() {
			ft.gorm.GoName = "Time"
			ft.gorm.GoImportPath = "time"
			ft.pointer = false
		}

	default:
		panic(fmt.Sprintf("TODO: %+v", desc))
	}

	foundMapping := ft.gorm.GoName != ""
	if !foundMapping {
		ft.gorm.GoName = ft.alias()
		ft.custom = true
		ft.pointer = false
	}

	return nil
}

func (ft *FieldType) alias() string {
	if ft.json {
		return "string"
	}

	if ft.field.msg.file.proto.GoImportPath == ft.golang.GoImportPath {
		return ft.field.proto.Message.GoIdent.GoName
	}
	return fmt.Sprintf("%s%s", ft.field.msg.ModelName(), ft.field.Name())
}

func (ft *FieldType) isTimestamp() bool {
	var (
		path = ft.field.proto.Message.GoIdent.GoImportPath
		name = ft.field.proto.Message.GoIdent.GoName
	)
	return path == "google.golang.org/protobuf/types/known/timestamppb" &&
		name == "Timestamp"
}

func (ft *FieldType) Gen() {
	// TODO: Call
	// TODO: Generate fully qualified imports
	// TODO: Generate custom types
}

func (ft *FieldType) Gorm() string {
	if ft.gorm.GoImportPath != "" {
		return string(ft.gorm.GoImportPath) + "." + ft.gorm.GoName
	}
	return ft.gorm.GoName
}
