package gengorm

func (m *Message) genConverters() {
	m.genModelToProto()
	m.genProtoToModel()
}

func (m *Message) genModelToProto() {
	m.P(Comment(" ToProto converts a %s to its protobuf representation.", m.ModelName()),
		"func (m *", m.ModelName(), ") ToProto() ", m.proto.GoIdent.GoName, "{")
	m.P("panic(true)")
	m.P("}")
	m.P()
}

func (m *Message) genProtoToModel() {
	m.P(Comment(" ToModel converts a %s to its GORM model.", m.proto.GoIdent.GoName),
		"func (x *", m.proto.GoIdent.GoName, ") ToModel() ", m.ModelName(), "{")
	m.P("panic(true)")
	m.P("}")
	m.P()
}
