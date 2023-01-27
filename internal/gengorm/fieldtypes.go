package gengorm

import (
	"fmt"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	KnownTypesTimestampPkg = "google.golang.org/protobuf/types/known/timestamppb"
	KnownTypesWrappersPkg  = "google.golang.org/protobuf/types/known/wrapperspb"
)

func NewFieldType(field *Field) (*FieldType, error) {
	types := &FieldType{
		field: field,
		Go:    field.proto.GoIdent,
		JSON:  field.opts.Json,
	}
	if err := types.init(); err != nil {
		return nil, err
	}
	return types, nil
}

type FieldType struct {
	field *Field
	Go    protogen.GoIdent
	JSON  bool

	Gorm     protogen.GoIdent
	Pointer  bool
	Enum     bool
	Custom   bool
	External bool
}

func (t *FieldType) init() error {
	if t.JSON {
		t.Gorm.GoName = "[]byte"
		t.Pointer = false
		return nil
	}

	switch t.field.proto.Desc.Kind() {
	case protoreflect.BoolKind:
		t.Gorm.GoName = "bool"
	case protoreflect.EnumKind:
		t.Gorm.GoName = "int32"
		t.Enum = true
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
			t.Gorm.GoName = "NullTime"
			t.Gorm.GoImportPath = "database/sql"
			return nil
		}
		if t.isDoubleValueWrapper() {
			t.Gorm.GoName = "NullFloat64"
			t.Gorm.GoImportPath = "database/sql"
			return nil
		}
		if t.isFloatValueWrapper() {
			t.Gorm.GoName = "NullFloat64"
			t.Gorm.GoImportPath = "database/sql"
			return nil
		}
		if t.isInt64ValueWrapper() {
			t.Gorm.GoName = "NullInt64"
			t.Gorm.GoImportPath = "database/sql"
			return nil
		}
		if t.isUInt64ValueWrapper() {
			t.Gorm.GoName = "NullInt64"
			t.Gorm.GoImportPath = "database/sql"
			return nil
		}
		if t.isInt32ValueWrapper() {
			t.Gorm.GoName = "NullInt32"
			t.Gorm.GoImportPath = "database/sql"
			return nil
		}
		if t.isUInt32ValueWrapper() {
			t.Gorm.GoName = "NullInt64"
			t.Gorm.GoImportPath = "database/sql"
			return nil
		}
		if t.isBoolValueWrapper() {
			t.Gorm.GoName = "NullBool"
			t.Gorm.GoImportPath = "database/sql"
			return nil
		}
		if t.isStringValueWrapper() {
			t.Gorm.GoName = "NullString"
			t.Gorm.GoImportPath = "database/sql"
			return nil
		}
		if t.isBytesValueWrapper() {
			t.Gorm.GoName = "[]byte"
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
		panic(fmt.Sprintf("TODO: External custom types: %+v", t.Go))
		// t.Gorm.GoName = t.alias()

	case t.Custom && !t.External:
		panic(fmt.Sprintf("TODO: Internal custom types: %+v", t.Go))
		// t.Gorm.GoName = t.Go.GoName
	}
	return nil
}

// TODO: Refactor
func (t *FieldType) String() string {
	if t.Gorm.GoImportPath != "" {
		id := t.field.msg.file.out.QualifiedGoIdent(t.Gorm)
		if t.Pointer {
			id = "*" + id
		}
		return id
	}
	if t.Pointer {
		return "*" + t.Gorm.GoName
	}
	return t.Gorm.GoName
}

func (t *FieldType) isTimestamp() bool {
	if t.field.proto.Message == nil {
		return false
	}
	var (
		path = t.field.proto.Message.GoIdent.GoImportPath
		name = t.field.proto.Message.GoIdent.GoName
	)
	return path == KnownTypesTimestampPkg &&
		name == "Timestamp"
}

func (t *FieldType) isDoubleValueWrapper() bool {
	if t.field.proto.Message == nil {
		return false
	}
	var (
		path = t.field.proto.Message.GoIdent.GoImportPath
		name = t.field.proto.Message.GoIdent.GoName
	)
	return path == KnownTypesWrappersPkg && name == "DoubleValue"
}

func (t *FieldType) isFloatValueWrapper() bool {
	if t.field.proto.Message == nil {
		return false
	}
	var (
		path = t.field.proto.Message.GoIdent.GoImportPath
		name = t.field.proto.Message.GoIdent.GoName
	)
	return path == KnownTypesWrappersPkg && name == "FloatValue"
}

func (t *FieldType) isInt64ValueWrapper() bool {
	if t.field.proto.Message == nil {
		return false
	}
	var (
		path = t.field.proto.Message.GoIdent.GoImportPath
		name = t.field.proto.Message.GoIdent.GoName
	)
	return path == KnownTypesWrappersPkg && name == "Int64Value"
}

func (t *FieldType) isUInt64ValueWrapper() bool {
	if t.field.proto.Message == nil {
		return false
	}
	var (
		path = t.field.proto.Message.GoIdent.GoImportPath
		name = t.field.proto.Message.GoIdent.GoName
	)
	return path == KnownTypesWrappersPkg && name == "UInt64Value"
}

func (t *FieldType) isInt32ValueWrapper() bool {
	if t.field.proto.Message == nil {
		return false
	}
	var (
		path = t.field.proto.Message.GoIdent.GoImportPath
		name = t.field.proto.Message.GoIdent.GoName
	)
	return path == KnownTypesWrappersPkg && name == "Int32Value"
}

func (t *FieldType) isUInt32ValueWrapper() bool {
	if t.field.proto.Message == nil {
		return false
	}
	var (
		path = t.field.proto.Message.GoIdent.GoImportPath
		name = t.field.proto.Message.GoIdent.GoName
	)
	return path == KnownTypesWrappersPkg && name == "UInt32Value"
}

func (t *FieldType) isBoolValueWrapper() bool {
	if t.field.proto.Message == nil {
		return false
	}
	var (
		path = t.field.proto.Message.GoIdent.GoImportPath
		name = t.field.proto.Message.GoIdent.GoName
	)
	return path == KnownTypesWrappersPkg && name == "BoolValue"
}

func (t *FieldType) isStringValueWrapper() bool {
	if t.field.proto.Message == nil {
		return false
	}
	var (
		path = t.field.proto.Message.GoIdent.GoImportPath
		name = t.field.proto.Message.GoIdent.GoName
	)
	return path == KnownTypesWrappersPkg && name == "StringValue"
}

func (t *FieldType) isBytesValueWrapper() bool {
	if t.field.proto.Message == nil {
		return false
	}
	var (
		path = t.field.proto.Message.GoIdent.GoImportPath
		name = t.field.proto.Message.GoIdent.GoName
	)
	return path == KnownTypesWrappersPkg && name == "BytesValue"
}

func (t *FieldType) Gen() {
	// TODO: Call
	// TODO: Generate fully qualified imports
	// TODO: Generate custom types
}
