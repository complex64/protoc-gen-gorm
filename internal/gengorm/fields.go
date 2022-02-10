package gengorm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/complex64/protoc-gen-gorm/gormpb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func NewField(msg *Message, proto *protogen.Field) (*Field, error) {
	field := &Field{
		msg:   msg,
		proto: proto,
	}
	if err := field.init(); err != nil {
		return nil, err
	}
	return field, nil
}

type Field struct {
	msg   *Message
	proto *protogen.Field

	opts  *gormpb.FieldOptions
	types *FieldType
}

func (f *Field) init() error {
	f.initOpts()
	if err := f.initTypes(); err != nil {
		return err
	}
	return nil
}

func (f *Field) initOpts() {
	descOpts := f.proto.Desc.Options()
	opts, ok := proto.GetExtension(descOpts, gormpb.E_Field).(*gormpb.FieldOptions)
	if ok && opts != nil {
		f.opts = opts
	} else {
		f.opts = &gormpb.FieldOptions{}
	}
}

func (f *Field) initTypes() error {
	types, err := NewFieldType(f)
	if err != nil {
		return err
	}
	f.types = types
	return nil
}

func (f *Field) Gen() {
	name := f.proto.GoName

	f.Annotate(f.msg.ModelName()+"."+name, f.proto.Location)
	f.P(name, " ", f.types.String(), f.tags())
}

func (f *Field) tags() string {
	if f.opts == nil {
		return ""
	}
	vals := f.tagVals()
	if len(vals) == 0 {
		return ""
	}

	joined := strings.Join(vals, ";")
	quoted := strconv.Quote(joined)
	escaped := strings.Replace(quoted, "`", `\x60`, -1)
	return fmt.Sprintf("`gorm:%s`", escaped)
}

func (f *Field) tagVals() (values []string) {
	if f.opts == nil {
		return
	}

	if f.opts.GetIgnore() {
		values = append(values, "-")
	}
	if col := f.opts.Column; col != "" {
		values = append(values, "column:"+col)
	}
	if f.opts.NotNull {
		values = append(values, "not null")
	}
	if v := f.opts.Default; v != "" {
		values = append(values, "default:"+v)
	}
	if f.opts.Unique {
		values = append(values, "unique")
	}
	if f.opts.PrimaryKey {
		values = append(values, "primaryKey")
	}
	for _, idx := range f.opts.Index {
		value := "index"
		if idx.Name != "" {
			value += ":" + idx.Name
		}
		values = append(values, value)
	}
	for _, idx := range f.opts.UniqueIndex {
		value := "uniqueIndex"
		if idx.Name != "" {
			value += ":" + idx.Name
		}
		values = append(values, value)
	}
	if t := f.opts.AutoCreateTime; t {
		values = append(values, "autoCreateTime")
	}
	if t := f.opts.AutoUpdateTime; t {
		values = append(values, "autoUpdateTime")
	}
	if deny := f.opts.GetDeny(); deny != nil {
		if tag := permTags(deny.Read, deny.Create, deny.Update); tag != "" {
			values = append(values, tag)
		}
	}

	return
}

func (f *Field) Name() string { return f.proto.GoName }

func (f *Field) Annotate(symbol string, loc protogen.Location) { f.msg.Annotate(symbol, loc) }
func (f *Field) P(v ...interface{})                            { f.msg.P(v...) }

// fieldInfo wraps a message's fields from the input .proto file
// and keeps information to generate code.
type fieldInfo struct {
	*protogen.Field

	opts *gormpb.FieldOptions

	gormType   string
	customType bool
	pointer    bool
	json       bool
}

func customTypeAlias(g *protogen.GeneratedFile, field *fieldInfo) string {
	return "*ALIAS_" + g.QualifiedGoIdent(field.Message.GoIdent)
}

func permTags(denyRead, denyCreate, denyUpdate bool) (tag string) {
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
