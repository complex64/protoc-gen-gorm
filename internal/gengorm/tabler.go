package gengorm

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// genTabler implements the Tabler interface if a custom table name is set.
// https://gorm.io/docs/conventions.html#TableName
func genTabler(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	if m.opts.Table == "" {
		return
	}
	g.P("func (m *", m.modelName, ") TableName() string {")
	g.P(`return "`, m.opts.Table, `"`)
	g.P("}")
}
