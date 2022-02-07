package gengorm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/complex64/protoc-gen-gorm/gormpb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// genMessageFields generates all fields for the GORM model m.
func genMessageFields(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	for _, field := range m.fields {
		genMessageField(g, f, m, field)
	}
}

func genMessageField(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo, field *fieldInfo) {
	if field.Desc.IsMap() {
		panic("TODO: maps")
	}

	goType, pointer := fieldGormType(g, f, field.Field)
	if pointer {
		goType = "*" + goType
	}

	tags := fieldGormTag(field)

	name := field.GoName
	g.Annotate(m.modelName+"."+name, field.Location)
	g.P(name, " ", goType, tags)
}

// fieldInfo wraps a message's fields from the input .proto file
// and keeps information to generate code.
type fieldInfo struct {
	*protogen.Field

	opts *gormpb.FieldOptions
}

func newFieldInfo(m *messageInfo, field *protogen.Field) (*fieldInfo, error) {
	f := &fieldInfo{
		Field: field,
		opts:  fieldOptions(field),
	}
	return f, nil
}

// fieldOptions returns the protoc-gen-gorm options set for a message's field.
// Example: message MyMessage {string field = 1 [(gorm.v2.field).not_null = true];}
func fieldOptions(field *protogen.Field) *gormpb.FieldOptions {
	opts := field.Desc.Options()
	o, ok := proto.GetExtension(opts, gormpb.E_Field).(*gormpb.FieldOptions)
	if !ok || o == nil {
		return new(gormpb.FieldOptions)
	}
	return o
}

// TODO
// fieldGormType returns the Go type used for a field.
//
// If it returns pointer=true, the struct field is a pointer to the type.
func fieldGormType(g *protogen.GeneratedFile, f *fileInfo, field *protogen.Field) (goType string, pointer bool) {
	pointer = field.Desc.HasPresence()
	switch field.Desc.Kind() {
	case protoreflect.BoolKind:
		goType = "bool"
	case protoreflect.EnumKind:
		goType = "int32"
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		goType = "int32"
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		goType = "uint32"
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		goType = "int64"
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		goType = "uint64"
	case protoreflect.FloatKind:
		goType = "float32"
	case protoreflect.DoubleKind:
		goType = "float64"
	case protoreflect.StringKind:
		goType = "string"
	case protoreflect.BytesKind:
		goType = "[]byte"
		pointer = false // rely on nullability of slices for presence
	case protoreflect.MessageKind, protoreflect.GroupKind:
		// TODO: Refactor!
		if string(field.Message.GoIdent.GoImportPath) == "google.golang.org/protobuf/types/known/timestamppb" &&
			string(field.Message.GoIdent.GoName) == "Timestamp" {
			g.QualifiedGoIdent(protogen.GoIdent{
				GoName:       "Time",
				GoImportPath: "time",
			})
			goType = "time.Time"
			pointer = false
		} else {
			panic("TODO: message/oneof fields: " + field.Message.GoIdent.String())
			// goType = "*" + g.QualifiedGoIdent(field.Message.GoIdent)
			// pointer = false // pointer captured as part of the type
		}
	}
	switch {
	case field.Desc.IsList():
		panic("TODO: slice fields: " + field.GoName)
		// return "[]" + goType, false
	case field.Desc.IsMap():
		panic("TODO: map fields: " + field.GoName)
		// keyType, _ := fieldGormType(g, f, field.Message.Fields[0])
		// valType, _ := fieldGormType(g, f, field.Message.Fields[1])
		// return fmt.Sprintf("map[%v]%v", keyType, valType), false
	}
	return goType, pointer
}

func fieldGormTag(f *fieldInfo) string {
	if f.opts == nil {
		return ""
	}
	return fmt.Sprintf("`gorm:%s`", fieldGormTagValue(f))
}

func fieldGormTagValue(f *fieldInfo) string {
	var (
		opts   = f.opts
		values []string
	)
	if opts == nil {
		return ""
	}
	if opts.GetIgnore() {
		values = append(values, "-")
	}
	if col := opts.Column; col != "" {
		values = append(values, "column:"+col)
	}
	if opts.NotNull {
		values = append(values, "not null")
	}
	if v := opts.Default; v != "" {
		values = append(values, "default:"+v)
	}
	if opts.Unique {
		values = append(values, "unique")
	}
	for _, idx := range opts.Index {
		value := "index"
		if idx.Name != "" {
			value += ":" + idx.Name
		}
		values = append(values, value)
	}
	for _, idx := range opts.UniqueIndex {
		value := "uniqueIndex"
		if idx.Name != "" {
			value += ":" + idx.Name
		}
		values = append(values, value)
	}
	if t := opts.AutoCreateTime; t {
		values = append(values, "autoCreateTime")
	}
	if t := opts.AutoUpdateTime; t {
		values = append(values, "autoUpdateTime")
	}
	if deny := opts.GetDeny(); deny != nil {
		if tag := permissionTag(deny.Read, deny.Create, deny.Update); tag != "" {
			values = append(values, tag)
		}
	}
	all := strings.Join(values, ";")
	quoted := strconv.Quote(all)
	escaped := strings.Replace(quoted, "`", `\x60`, -1)
	return escaped
}

func permissionTag(denyRead, denyCreate, denyUpdate bool) (tag string) {
	type allowed struct{ read, create, update bool }
	perms := allowed{
		// Easier to argue about non-inverted booleans.
		!denyRead, !denyCreate, !denyUpdate,
	}
	return map[allowed]string{
		{read: true, create: true, update: true}:    "",                   // read & write
		{read: true, create: true, update: false}:   "<-:create",          // read & create, no updates
		{read: true, create: false, update: true}:   "<-:update",          // read & update, no creates
		{read: true, create: false, update: false}:  "->",                 // read-only
		{read: false, create: true, update: true}:   "->:false;<-",        // write-only
		{read: false, create: true, update: false}:  "->:false;<-:create", // create-only
		{read: false, create: false, update: true}:  "->:false;<-:update", // update-only
		{read: false, create: false, update: false}: "-",                  // ignore
	}[perms]
}
