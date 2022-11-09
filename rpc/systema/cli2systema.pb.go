// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: systema/cli2systema.proto

package systema

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CommandMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// intercept is the message to give when a user performs an intercept command.
	Intercept string `protobuf:"bytes,1,opt,name=intercept,proto3" json:"intercept,omitempty"`
}

func (x *CommandMessageResponse) Reset() {
	*x = CommandMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_systema_cli2systema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandMessageResponse) ProtoMessage() {}

func (x *CommandMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_systema_cli2systema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandMessageResponse.ProtoReflect.Descriptor instead.
func (*CommandMessageResponse) Descriptor() ([]byte, []int) {
	return file_systema_cli2systema_proto_rawDescGZIP(), []int{0}
}

func (x *CommandMessageResponse) GetIntercept() string {
	if x != nil {
		return x.Intercept
	}
	return ""
}

var File_systema_cli2systema_proto protoreflect.FileDescriptor

var file_systema_cli2systema_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x61, 0x2f, 0x63, 0x6c, 0x69, 0x32, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x61, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36,
	0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x63, 0x65, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x32, 0x77, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x41, 0x43, 0x6c, 0x69, 0x12, 0x69, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x55, 0x6e, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x2c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x69, 0x6f, 0x2f, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x76, 0x32,
	0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_systema_cli2systema_proto_rawDescOnce sync.Once
	file_systema_cli2systema_proto_rawDescData = file_systema_cli2systema_proto_rawDesc
)

func file_systema_cli2systema_proto_rawDescGZIP() []byte {
	file_systema_cli2systema_proto_rawDescOnce.Do(func() {
		file_systema_cli2systema_proto_rawDescData = protoimpl.X.CompressGZIP(file_systema_cli2systema_proto_rawDescData)
	})
	return file_systema_cli2systema_proto_rawDescData
}

var file_systema_cli2systema_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_systema_cli2systema_proto_goTypes = []interface{}{
	(*CommandMessageResponse)(nil), // 0: telepresence.systema.CommandMessageResponse
	(*emptypb.Empty)(nil),          // 1: google.protobuf.Empty
}
var file_systema_cli2systema_proto_depIdxs = []int32{
	1, // 0: telepresence.systema.SystemACli.GetUnauthenticatedCommandMessages:input_type -> google.protobuf.Empty
	0, // 1: telepresence.systema.SystemACli.GetUnauthenticatedCommandMessages:output_type -> telepresence.systema.CommandMessageResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_systema_cli2systema_proto_init() }
func file_systema_cli2systema_proto_init() {
	if File_systema_cli2systema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_systema_cli2systema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandMessageResponse); i {
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
			RawDescriptor: file_systema_cli2systema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_systema_cli2systema_proto_goTypes,
		DependencyIndexes: file_systema_cli2systema_proto_depIdxs,
		MessageInfos:      file_systema_cli2systema_proto_msgTypes,
	}.Build()
	File_systema_cli2systema_proto = out.File
	file_systema_cli2systema_proto_rawDesc = nil
	file_systema_cli2systema_proto_goTypes = nil
	file_systema_cli2systema_proto_depIdxs = nil
}
