package gengorm

import (
	"fmt"

	"github.com/complex64/protoc-gen-gorm/gormpb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

// genModels generates GORM models and supporting APIs.
func genModels(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	if !m.genModel {
		return
	}

	// Message type declaration.
	g.Annotate(m.modelName, m.Location)
	g.P(m.leadingComment(f), "type ", m.modelName, " struct {")
	genMessageFields(g, f, m)
	g.P("}")
	g.P()

	// TODO: Decide what to do about hooks.
	genTabler(g, f, m)
	genConverters(g, f, m)
	genCRUD(g, f, m)
}

// messageInfo wraps a message from the input .proto file and keeps information to generate code.
type messageInfo struct {
	*protogen.Message

	opts   *gormpb.MessageOptions
	fields []*fieldInfo

	modelName   string
	genModel    bool
	genValidate bool
	genCRUD     bool
}

func newMessageInfo(f *fileInfo, message *protogen.Message) (*messageInfo, error) {
	m := &messageInfo{
		Message: message,
		opts:    &gormpb.MessageOptions{},
	}

	m.modelName = fmt.Sprintf("%sModel", m.GoIdent.GoName)

	// File flags override message flags.
	m.genModel = f.genModel
	m.genValidate = f.genValidate
	m.genCRUD = f.genCRUD

	if opts := messageOptions(message); opts != nil {
		m.opts = opts

		// Generate a model when using features that need the model.
		implyModel := opts.Validate || opts.Crud
		m.genModel = m.genModel || opts.Model || implyModel
		m.genValidate = m.genValidate || opts.Validate
		m.genCRUD = m.genCRUD || opts.Crud
	}

	if err := m.initMessageInfo(); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *messageInfo) initMessageInfo() error {
	for _, field := range m.Fields {
		if f, err := newFieldInfo(m, field); err != nil {
			return err
		} else {
			m.fields = append(m.fields, f)
		}
	}
	return nil
}

// messageOptions returns the protoc-gen-gorm options set for message.
// Example: message MyMessage { option (gorm.message).model = true; }
func messageOptions(message *protogen.Message) *gormpb.MessageOptions {
	opts := message.Desc.Options()
	o, ok := proto.GetExtension(opts, gormpb.E_Message).(*gormpb.MessageOptions)
	if !ok || o == nil {
		return nil
	}
	return o
}

func (m *messageInfo) leadingComment(f *fileInfo) protogen.Comments {
	leading := protogen.Comments(
		fmt.Sprintf(
			" %s is the GORM model for %s.%s.",
			m.modelName,
			f.GoPackageName,
			m.GoIdent.GoName,
		),
	)
	return appendDeprecationNotice(
		leading,
		m.Desc.Options().(*descriptorpb.MessageOptions).GetDeprecated(),
	)
}
