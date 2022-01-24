package internal_gengorm

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

// GenerateFile does nothing yet.
func GenerateFile(flags flag.FlagSet, plugin *protogen.Plugin, file *protogen.File) (*protogen.GeneratedFile, error) {
	out := file.GeneratedFilenamePrefix + ".gorm.pb.go"
	f := plugin.NewGeneratedFile(out, file.GoImportPath)
	f.P("package ", file.GoPackageName)
	f.P(`const A = "A"`)
	return f, nil
}
