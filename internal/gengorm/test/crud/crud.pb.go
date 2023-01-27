// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: crud/crud.proto

package crud

import (
	_ "github.com/complex64/protoc-gen-gorm/gormpb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Crud struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	StringField string `protobuf:"bytes,2,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	Int32Field  int32  `protobuf:"varint,3,opt,name=int32_field,json=int32Field,proto3" json:"int32_field,omitempty"`
	BoolField   bool   `protobuf:"varint,4,opt,name=bool_field,json=boolField,proto3" json:"bool_field,omitempty"`
}

func (x *Crud) Reset() {
	*x = Crud{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_crud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crud) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crud) ProtoMessage() {}

func (x *Crud) ProtoReflect() protoreflect.Message {
	mi := &file_crud_crud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crud.ProtoReflect.Descriptor instead.
func (*Crud) Descriptor() ([]byte, []int) {
	return file_crud_crud_proto_rawDescGZIP(), []int{0}
}

func (x *Crud) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Crud) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *Crud) GetInt32Field() int32 {
	if x != nil {
		return x.Int32Field
	}
	return 0
}

func (x *Crud) GetBoolField() bool {
	if x != nil {
		return x.BoolField
	}
	return false
}

var File_crud_crud_proto protoreflect.FileDescriptor

var file_crud_crud_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x72, 0x75, 0x64, 0x2f, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x63, 0x72, 0x75, 0x64, 0x1a, 0x12, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x04,
	0x43, 0x72, 0x75, 0x64, 0x12, 0x1a, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xb2, 0xbb, 0x18, 0x02, 0x48, 0x01, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x2c, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x0d, 0xb2, 0xbb, 0x18, 0x09, 0x0a, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x3a, 0x06, 0xaa, 0xbb, 0x18, 0x02, 0x20, 0x01, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x36, 0x34, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x72, 0x6d, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x63, 0x72, 0x75, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crud_crud_proto_rawDescOnce sync.Once
	file_crud_crud_proto_rawDescData = file_crud_crud_proto_rawDesc
)

func file_crud_crud_proto_rawDescGZIP() []byte {
	file_crud_crud_proto_rawDescOnce.Do(func() {
		file_crud_crud_proto_rawDescData = protoimpl.X.CompressGZIP(file_crud_crud_proto_rawDescData)
	})
	return file_crud_crud_proto_rawDescData
}

var file_crud_crud_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_crud_crud_proto_goTypes = []interface{}{
	(*Crud)(nil), // 0: crud.Crud
}
var file_crud_crud_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crud_crud_proto_init() }
func file_crud_crud_proto_init() {
	if File_crud_crud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crud_crud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crud); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crud_crud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crud_crud_proto_goTypes,
		DependencyIndexes: file_crud_crud_proto_depIdxs,
		MessageInfos:      file_crud_crud_proto_msgTypes,
	}.Build()
	File_crud_crud_proto = out.File
	file_crud_crud_proto_rawDesc = nil
	file_crud_crud_proto_goTypes = nil
	file_crud_crud_proto_depIdxs = nil
}
