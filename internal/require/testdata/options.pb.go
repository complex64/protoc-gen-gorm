// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: testdata/options.proto

package testdata

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MessageOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptionField_1 string `protobuf:"bytes,1,opt,name=option_field_1,json=optionField1,proto3" json:"option_field_1,omitempty"`
}

func (x *MessageOptions) Reset() {
	*x = MessageOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageOptions) ProtoMessage() {}

func (x *MessageOptions) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageOptions.ProtoReflect.Descriptor instead.
func (*MessageOptions) Descriptor() ([]byte, []int) {
	return file_testdata_options_proto_rawDescGZIP(), []int{0}
}

func (x *MessageOptions) GetOptionField_1() string {
	if x != nil {
		return x.OptionField_1
	}
	return ""
}

type MyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageField_1 string `protobuf:"bytes,1,opt,name=message_field_1,json=messageField1,proto3" json:"message_field_1,omitempty"`
}

func (x *MyMessage) Reset() {
	*x = MyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyMessage) ProtoMessage() {}

func (x *MyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyMessage.ProtoReflect.Descriptor instead.
func (*MyMessage) Descriptor() ([]byte, []int) {
	return file_testdata_options_proto_rawDescGZIP(), []int{1}
}

func (x *MyMessage) GetMessageField_1() string {
	if x != nil {
		return x.MessageField_1
	}
	return ""
}

var file_testdata_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*MessageOptions)(nil),
		Field:         50000,
		Name:          "testdata.message",
		Tag:           "bytes,50000,opt,name=message",
		Filename:      "testdata/options.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional testdata.MessageOptions message = 50000;
	E_Message = &file_testdata_options_proto_extTypes[0]
)

var File_testdata_options_proto protoreflect.FileDescriptor

var file_testdata_options_proto_rawDesc = []byte{
	0x0a, 0x16, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61,
	0x74, 0x61, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x22, 0x49, 0x0a, 0x09,
	0x4d, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x31, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x31, 0x3a, 0x14, 0x82, 0xb5, 0x18, 0x10, 0x0a, 0x0e, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x20,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x20, 0x31, 0x3a, 0x55, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x49,
	0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x78, 0x36, 0x34, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x3b, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_testdata_options_proto_rawDescOnce sync.Once
	file_testdata_options_proto_rawDescData = file_testdata_options_proto_rawDesc
)

func file_testdata_options_proto_rawDescGZIP() []byte {
	file_testdata_options_proto_rawDescOnce.Do(func() {
		file_testdata_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata_options_proto_rawDescData)
	})
	return file_testdata_options_proto_rawDescData
}

var file_testdata_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_testdata_options_proto_goTypes = []interface{}{
	(*MessageOptions)(nil),              // 0: testdata.MessageOptions
	(*MyMessage)(nil),                   // 1: testdata.MyMessage
	(*descriptorpb.MessageOptions)(nil), // 2: google.protobuf.MessageOptions
}
var file_testdata_options_proto_depIdxs = []int32{
	2, // 0: testdata.message:extendee -> google.protobuf.MessageOptions
	0, // 1: testdata.message:type_name -> testdata.MessageOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_testdata_options_proto_init() }
func file_testdata_options_proto_init() {
	if File_testdata_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testdata_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageOptions); i {
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
		file_testdata_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyMessage); i {
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
			RawDescriptor: file_testdata_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_testdata_options_proto_goTypes,
		DependencyIndexes: file_testdata_options_proto_depIdxs,
		MessageInfos:      file_testdata_options_proto_msgTypes,
		ExtensionInfos:    file_testdata_options_proto_extTypes,
	}.Build()
	File_testdata_options_proto = out.File
	file_testdata_options_proto_rawDesc = nil
	file_testdata_options_proto_goTypes = nil
	file_testdata_options_proto_depIdxs = nil
}
