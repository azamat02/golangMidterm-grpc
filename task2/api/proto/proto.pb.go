// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: api/proto/proto.proto

package servicepb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Number struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Number) Reset() {
	*x = Number{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Number) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Number) ProtoMessage() {}

func (x *Number) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Number.ProtoReflect.Descriptor instead.
func (*Number) Descriptor() ([]byte, []int) {
	return file_api_proto_proto_proto_rawDescGZIP(), []int{0}
}

func (x *Number) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type NumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number *Number `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *NumRequest) Reset() {
	*x = NumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumRequest) ProtoMessage() {}

func (x *NumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumRequest.ProtoReflect.Descriptor instead.
func (*NumRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_proto_proto_rawDescGZIP(), []int{1}
}

func (x *NumRequest) GetNumber() *Number {
	if x != nil {
		return x.Number
	}
	return nil
}

type NumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *NumResponse) Reset() {
	*x = NumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumResponse) ProtoMessage() {}

func (x *NumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumResponse.ProtoReflect.Descriptor instead.
func (*NumResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_proto_proto_rawDescGZIP(), []int{2}
}

func (x *NumResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_api_proto_proto_proto protoreflect.FileDescriptor

var file_api_proto_proto_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x22, 0x20,
	0x0a, 0x06, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x33, 0x0a, 0x0a, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x0b, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x46, 0x0a, 0x0c,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x09,
	0x4c, 0x6f, 0x6e, 0x67, 0x47, 0x72, 0x65, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x2e, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x2e, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_proto_proto_proto_rawDescOnce sync.Once
	file_api_proto_proto_proto_rawDescData = file_api_proto_proto_proto_rawDesc
)

func file_api_proto_proto_proto_rawDescGZIP() []byte {
	file_api_proto_proto_proto_rawDescOnce.Do(func() {
		file_api_proto_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_proto_proto_rawDescData)
	})
	return file_api_proto_proto_proto_rawDescData
}

var file_api_proto_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_proto_proto_proto_goTypes = []interface{}{
	(*Number)(nil),      // 0: greet.Number
	(*NumRequest)(nil),  // 1: greet.NumRequest
	(*NumResponse)(nil), // 2: greet.NumResponse
}
var file_api_proto_proto_proto_depIdxs = []int32{
	0, // 0: greet.NumRequest.number:type_name -> greet.Number
	1, // 1: greet.GreetService.LongGreet:input_type -> greet.NumRequest
	2, // 2: greet.GreetService.LongGreet:output_type -> greet.NumResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_proto_proto_init() }
func file_api_proto_proto_proto_init() {
	if File_api_proto_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Number); i {
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
		file_api_proto_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumRequest); i {
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
		file_api_proto_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumResponse); i {
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
			RawDescriptor: file_api_proto_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_proto_proto_goTypes,
		DependencyIndexes: file_api_proto_proto_proto_depIdxs,
		MessageInfos:      file_api_proto_proto_proto_msgTypes,
	}.Build()
	File_api_proto_proto_proto = out.File
	file_api_proto_proto_proto_rawDesc = nil
	file_api_proto_proto_proto_goTypes = nil
	file_api_proto_proto_proto_depIdxs = nil
}
