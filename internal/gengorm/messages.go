package gengorm

import (
	"fmt"

	"github.com/complex64/protoc-gen-gorm/gormpb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func NewMessage(file *File, proto *protogen.Message) (*Message, error) {
	msg := &Message{
		file:  file,
		proto: proto,
	}
	if err := msg.init(); err != nil {
		return nil, err
	}
	return msg, nil
}

type Message struct {
	file  *File
	proto *protogen.Message

	opts   *gormpb.MessageOptions
	fields []*Field
}

func (m *Message) init() error {
	m.initOpts()
	if err := m.initFields(); err != nil {
		return err
	}
	return nil
}

func (m *Message) initFields() error {
	if !m.model() {
		return nil
	}

	for _, f := range m.proto.Fields {
		if err := m.initField(f); err != nil {
			return err
		}
	}
	return nil
}

func (m *Message) initField(proto *protogen.Field) error {
	field, err := NewField(m, proto)
	if err != nil {
		return err
	}
	m.fields = append(m.fields, field)
	return nil
}

// initOpts reads the protoc-gen-gorm options set for this message.
// Example: message MyMessage { option (gorm.message).model = true; }
func (m *Message) initOpts() {
	descOpts := m.proto.Desc.Options()
	opts, ok := proto.GetExtension(descOpts, gormpb.E_Message).(*gormpb.MessageOptions)
	if ok && opts != nil {
		m.opts = opts
	} else {
		m.opts = &gormpb.MessageOptions{}
	}
}

// Gen generates GORM models and supporting APIs.
func (m *Message) Gen() {
	if !m.model() {
		return
	}
	m.genStruct()
	// m.genCustomTypes() // TODO
	m.genConverters()
	m.genTabler()
	m.genCRUD()
}

func (m *Message) genStruct() {
	m.Annotate(m.ModelName(), m.proto.Location) // Message/model type declaration.
	m.P(m.leadingComment(), "type ", m.ModelName(), " struct {")
	m.genFields()
	m.P("}")
	m.P()
}

func (m *Message) leadingComment() protogen.Comments {
	return appendDeprecationNotice(
		Comment(" %s is the GORM model for %s.%s.",
			m.ModelName(),
			m.file.proto.GoPackageName,
			m.proto.GoIdent.GoName),
		m.deprecated(),
	)
}

func (m *Message) deprecated() bool {
	return m.proto.Desc.Options().(*descriptorpb.MessageOptions).GetDeprecated()
}

func (m *Message) genFields() {
	for _, field := range m.fields {
		field.Gen()
	}
}

func (m *Message) ProtoName() string {
	return m.proto.GoIdent.GoName
}

func (m *Message) ModelName() string {
	return fmt.Sprintf("%sModel", m.ProtoName())
}

func (m *Message) crud() bool     { return m.opts.Crud || m.file.CRUD() }
func (m *Message) validate() bool { return m.opts.Validate || m.file.Validate() }
func (m *Message) model() bool    { return m.crud() || m.validate() || m.opts.Model }

func (m *Message) Annotate(symbol string, loc protogen.Location) { m.file.Annotate(symbol, loc) }
func (m *Message) P(v ...interface{})                            { m.file.P(v...) }
