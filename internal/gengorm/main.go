package gengorm

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

func GenerateFile(fs flag.FlagSet, plugin *protogen.Plugin, proto *protogen.File) error {
	file, err := NewFile(plugin, proto)
	if err != nil {
		return err
	}
	file.Gen()
	return nil
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
