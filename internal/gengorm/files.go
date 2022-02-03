package gengorm

import (
	"fmt"

	"github.com/complex64/protoc-gen-gorm/gormpb"
	"github.com/complex64/protoc-gen-gorm/internal/version"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// genHeader generates the comment header at the top of the file.
// We warn to not edit the file and show the tool versions used to generate the file.
func genHeader(plugin *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo) {
	g.P("// Code generated by protoc-gen-gorm. DO NOT EDIT.")
	g.P("// versions:")

	protocGenGormVersion := version.String()
	protocVersion := "(unknown)"

	if v := plugin.Request.GetCompilerVersion(); v != nil {
		protocVersion = fmt.Sprintf(
			"v%v.%v.%v",
			v.GetMajor(),
			v.GetMinor(),
			v.GetPatch(),
		)
		if s := v.GetSuffix(); s != "" {
			protocVersion += "-" + s
		}
	}

	g.P("// \tprotoc-gen-gorm ", protocGenGormVersion)
	g.P("// \tprotoc          ", protocVersion)

	if f.Proto.GetOptions().GetDeprecated() {
		g.P("// ", f.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", f.Desc.Path())
	}

	g.P()
}

// genImport adds necessary import statements to the generated file.
func genImport(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo, imp protoreflect.FileImport) {
	if imp.IsPublic {
		panic("TODO: Implement support for public imports")
	}

	impFile, ok := gen.FilesByPath[imp.Path()]
	if !ok {
		return // .proto file unavailable.
	}

	if impFile.GoImportPath == f.GoImportPath {
		return // Same Go package.
	}

	g.Import(impFile.GoImportPath)
	g.P()
}

// fileInfo wraps the input .proto file and keeps information to generate code.
type fileInfo struct {
	*protogen.File

	allMessages []*messageInfo

	genModel    bool
	genValidate bool
	genCRUD     bool
}

func newFileInfo(file *protogen.File) (*fileInfo, error) {
	f := &fileInfo{File: file}

	if opts := fileOptions(file); opts != nil {
		// Generate a model when using features that need the model.
		implyModel := opts.Validate || opts.Crud
		f.genModel = opts.Model || implyModel
	}

	if err := f.initFileInfo(); err != nil {
		return nil, err
	}

	return f, nil
}

func (f *fileInfo) initFileInfo() error {
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
