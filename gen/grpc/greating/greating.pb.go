// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: greating.proto

package greating

import (
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

type GreatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GreatingRequest) Reset() {
	*x = GreatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greating_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreatingRequest) ProtoMessage() {}

func (x *GreatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_greating_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreatingRequest.ProtoReflect.Descriptor instead.
func (*GreatingRequest) Descriptor() ([]byte, []int) {
	return file_greating_proto_rawDescGZIP(), []int{0}
}

func (x *GreatingRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GreatingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GreatingResponse) Reset() {
	*x = GreatingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_greating_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreatingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreatingResponse) ProtoMessage() {}

func (x *GreatingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_greating_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreatingResponse.ProtoReflect.Descriptor instead.
func (*GreatingResponse) Descriptor() ([]byte, []int) {
	return file_greating_proto_rawDescGZIP(), []int{1}
}

func (x *GreatingResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_greating_proto protoreflect.FileDescriptor

var file_greating_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x67, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x24, 0x0a, 0x10, 0x47, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x4f, 0x0a, 0x0f, 0x47, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x03, 0x53, 0x61, 0x79, 0x12,
	0x19, 0x2e, 0x67, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x69, 0x6e, 0x64, 0x32, 0x53, 0x63, 0x72, 0x65, 0x65, 0x6e,
	0x2d, 0x44, 0x65, 0x76, 0x2d, 0x54, 0x65, 0x61, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x6b, 0x65,
	0x6c, 0x65, 0x74, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_greating_proto_rawDescOnce sync.Once
	file_greating_proto_rawDescData = file_greating_proto_rawDesc
)

func file_greating_proto_rawDescGZIP() []byte {
	file_greating_proto_rawDescOnce.Do(func() {
		file_greating_proto_rawDescData = protoimpl.X.CompressGZIP(file_greating_proto_rawDescData)
	})
	return file_greating_proto_rawDescData
}

var file_greating_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_greating_proto_goTypes = []any{
	(*GreatingRequest)(nil),  // 0: greating.GreatingRequest
	(*GreatingResponse)(nil), // 1: greating.GreatingResponse
}
var file_greating_proto_depIdxs = []int32{
	0, // 0: greating.GreatingService.Say:input_type -> greating.GreatingRequest
	1, // 1: greating.GreatingService.Say:output_type -> greating.GreatingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_greating_proto_init() }
func file_greating_proto_init() {
	if File_greating_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_greating_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GreatingRequest); i {
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
		file_greating_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GreatingResponse); i {
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
			RawDescriptor: file_greating_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_greating_proto_goTypes,
		DependencyIndexes: file_greating_proto_depIdxs,
		MessageInfos:      file_greating_proto_msgTypes,
	}.Build()
	File_greating_proto = out.File
	file_greating_proto_rawDesc = nil
	file_greating_proto_goTypes = nil
	file_greating_proto_depIdxs = nil
}
