package gengorm

import (
	"google.golang.org/protobuf/compiler/protogen"
)

// genConverters adds methods to the message type and model type to convert between the two.
func genConverters(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	genProtoAsModel(g, f, m)
	genModelAsProto(g, f, m)
}

// genProtoAsModel generates a method on the proto message type to convert to the model type.
func genProtoAsModel(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	// TODO: Implement me!
}

// genModelAsProto generates a method on the model type to convert to the proto message type.
func genModelAsProto(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo) {
	// TODO: Implement me!
}
