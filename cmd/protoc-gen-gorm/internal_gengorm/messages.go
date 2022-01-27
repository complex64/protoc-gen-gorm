package internal_gengorm

import (
	"github.com/complex64/protoc-gen-gorm/gormpb/v2"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

// messageInfo wraps a message from the input .proto file and keeps information to generate code.
type messageInfo struct {
	*protogen.Message

	genModel    bool
	genHooks    bool
	genValidate bool
	genCRUD     bool
}

func newMessageInfo(f *fileInfo, message *protogen.Message) (*messageInfo, error) {
	m := &messageInfo{Message: message}

	// File flags override message flags.
	m.genModel = f.genModel
	m.genHooks = f.genHooks
	m.genValidate = f.genValidate
	m.genCRUD = f.genCRUD

	if opts := messageOptions(message); opts != nil {
		// Generate a model when using features that need the model.
		implyModel := opts.Hooks || opts.Validate || opts.Crud
		m.genModel = m.genModel || opts.Model || implyModel
		m.genHooks = m.genHooks || opts.Hooks
		m.genValidate = m.genValidate || opts.Validate
		m.genCRUD = m.genCRUD || opts.Crud
	}

	return m, nil
}

// messageOptions returns the protoc-gen-gorm options set for message.
// Example: message MyMessage { option (gorm.v2.message).model = true; }
func messageOptions(message *protogen.Message) *gormpb.MessageOptions {
	opts := message.Desc.Options()
	o, ok := proto.GetExtension(opts, gormpb.E_Message).(*gormpb.MessageOptions)
	if !ok || o == nil {
		return nil
	}
	return o
}
