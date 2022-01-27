package internal_gengorm

import (
	"github.com/complex64/protoc-gen-gorm/gormpb/v2"
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
		implyModel := opts.Hooks || opts.Validate || opts.Crud
		f.genModel = opts.Model || implyModel
	}

	// Collect all messages recursively.
	var walkMessages func([]*protogen.Message, func(*protogen.Message) error) error
	walkMessages = func(messages []*protogen.Message, f func(*protogen.Message) error) error {
		for _, m := range messages {
			if err := f(m); err != nil {
				return err
			}
			if err := walkMessages(m.Messages, f); err != nil {
				return err
			}
		}
		return nil
	}

	initMessageInfos := func(messages []*protogen.Message) error {
		for _, message := range messages {
			m, err := newMessageInfo(f, message)
			if err != nil {
				return err
			}
			f.allMessages = append(f.allMessages, m)
		}
		return nil
	}

	if err := initMessageInfos(f.Messages); err != nil {
		return nil, err
	}

	err := walkMessages(f.Messages, func(m *protogen.Message) error {
		if err := initMessageInfos(m.Messages); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return f, nil
}

// fileOptions returns the protoc-gen-gorm options for file.
// Example: option (gorm.v2.file).model = true;
func fileOptions(file *protogen.File) *gormpb.FileOptions {
	opts := file.Desc.Options()
	o, ok := proto.GetExtension(opts, gormpb.E_File).(*gormpb.FileOptions)
	if !ok || o == nil {
		return nil
	}
	return o
}
