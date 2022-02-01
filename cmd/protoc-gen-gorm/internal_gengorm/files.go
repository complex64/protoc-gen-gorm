package internal_gengorm

import (
	"github.com/complex64/protoc-gen-gorm/gormpb"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

// fileInfo wraps the input .proto file and keeps information to generate code.
type fileInfo struct {
	*protogen.File

	allMessages []*messageInfo

	genModel    bool
	genHooks    bool
	genValidate bool
	genCRUD     bool
}

func newFileInfo(file *protogen.File) (*fileInfo, error) {
	f := &fileInfo{File: file}

	if opts := fileOptions(file); opts != nil {
		// Generate a model when using features that need the model.
		implyModel := opts.Hooks || opts.Validate || opts.Crud
		f.genModel = opts.Model || implyModel
	}

	if err := f.collectFileInfo(); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *fileInfo) collectFileInfo() error {
	return walkMessages(f.Messages, []func(*protogen.Message) error{
		f.initMessageInfo,
	})
}

func (f *fileInfo) initMessageInfo(message *protogen.Message) error {
	m, err := newMessageInfo(f, message)
	if err != nil {
		return err
	}
	f.allMessages = append(f.allMessages, m)
	return nil
}

// walkMessages invokes all walkFuncs for all messages recursively.
func walkMessages(messages []*protogen.Message, walkFuncs []func(*protogen.Message) error) error {
	for _, message := range messages {
		for _, f := range walkFuncs {
			if err := f(message); err != nil {
				return err
			}
		}
		if err := walkMessages(message.Messages, walkFuncs); err != nil {
			return err
		}
	}
	return nil
}

// fileOptions returns the protoc-gen-gorm options for file.
// Example: option (gorm.file).model = true;
func fileOptions(file *protogen.File) *gormpb.FileOptions {
	opts := file.Desc.Options()
	o, ok := proto.GetExtension(opts, gormpb.E_File).(*gormpb.FileOptions)
	if !ok || o == nil {
		return nil
	}
	return o
}
