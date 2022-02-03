package gengorm

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	// Use extension .pb.go, like `protoc-gen-go` does.
	extension = "_gorm.pb.go"
)

func GenerateFile(fs flag.FlagSet, plugin *protogen.Plugin, file *protogen.File) (*protogen.GeneratedFile, error) {
	filename := file.GeneratedFilenamePrefix + extension

	g := plugin.NewGeneratedFile(filename, file.GoImportPath)
	f, err := newFileInfo(file)
	if err != nil {
		return nil, err
	}

	genHeader(plugin, g, f)
	g.P("package ", file.GoPackageName)
	g.P()

	for i, imps := 0, f.Desc.Imports(); i < imps.Len(); i++ {
		genImport(plugin, g, f, imps.Get(i))
	}

	for _, message := range f.allMessages {
		genModels(g, f, message)
	}

	return g, nil
}

func appendDeprecationNotice(prefix protogen.Comments, deprecated bool) protogen.Comments {
	if !deprecated {
		return prefix
	}
	if prefix != "" {
		prefix += "\n"
	}
	return prefix + " Deprecated: Do not use.\n"
}
