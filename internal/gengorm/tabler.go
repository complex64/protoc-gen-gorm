package gengorm

// genTabler implements the Tabler interface if a custom table name is set.
// https://gorm.io/docs/conventions.html#TableName
func (m *Message) genTabler() {
	if m.opts.Table == "" {
		return
	}
	m.P("func (m *", m.ModelName(), ") TableName() string {")
	m.P(`return "`, m.opts.Table, `"`)
	m.P("}")
}
