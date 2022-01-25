package internal_gengorm

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

const (
	extension = ".gorm.pb.go"
)

func GenerateFile(fs flag.FlagSet, plugin *protogen.Plugin, file *protogen.File) (*protogen.GeneratedFile, error) {
	out := file.GeneratedFilenamePrefix + extension
	f := plugin.NewGeneratedFile(out, file.GoImportPath)
	f.P("package ", file.GoPackageName)
	f.P(`const A = "A"`)
	return f, nil
}
