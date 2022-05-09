// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: proto/context.proto

package context

import (
	user_id "github.com/liyouxina/buildzone/protogen/user_id"
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

type RequestContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId                *user_id.UserId `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	GroupId               string          `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	ImpersonatingGroupId  string          `protobuf:"bytes,4,opt,name=impersonating_group_id,json=impersonatingGroupId,proto3" json:"impersonating_group_id,omitempty"`
	TimezoneOffsetMinutes int32           `protobuf:"varint,3,opt,name=timezone_offset_minutes,json=timezoneOffsetMinutes,proto3" json:"timezone_offset_minutes,omitempty"`
}

func (x *RequestContext) Reset() {
	*x = RequestContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_context_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestContext) ProtoMessage() {}

func (x *RequestContext) ProtoReflect() protoreflect.Message {
	mi := &file_proto_context_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestContext.ProtoReflect.Descriptor instead.
func (*RequestContext) Descriptor() ([]byte, []int) {
	return file_proto_context_proto_rawDescGZIP(), []int{0}
}

func (x *RequestContext) GetUserId() *user_id.UserId {
	if x != nil {
		return x.UserId
	}
	return nil
}

func (x *RequestContext) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *RequestContext) GetImpersonatingGroupId() string {
	if x != nil {
		return x.ImpersonatingGroupId
	}
	return ""
}

func (x *RequestContext) GetTimezoneOffsetMinutes() int32 {
	if x != nil {
		return x.TimezoneOffsetMinutes
	}
	return 0
}

type ResponseContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResponseContext) Reset() {
	*x = ResponseContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_context_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseContext) ProtoMessage() {}

func (x *ResponseContext) ProtoReflect() protoreflect.Message {
	mi := &file_proto_context_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseContext.ProtoReflect.Descriptor instead.
func (*ResponseContext) Descriptor() ([]byte, []int) {
	return file_proto_context_proto_rawDescGZIP(), []int{1}
}

var File_proto_context_proto protoreflect.FileDescriptor

var file_proto_context_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x1a, 0x13,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc3, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x69,
	0x6d, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x69, 0x6d, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x36, 0x0a, 0x17, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x5f, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x15, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x4f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x4d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x73, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_context_proto_rawDescOnce sync.Once
	file_proto_context_proto_rawDescData = file_proto_context_proto_rawDesc
)

func file_proto_context_proto_rawDescGZIP() []byte {
	file_proto_context_proto_rawDescOnce.Do(func() {
		file_proto_context_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_context_proto_rawDescData)
	})
	return file_proto_context_proto_rawDescData
}

var file_proto_context_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_context_proto_goTypes = []interface{}{
	(*RequestContext)(nil),  // 0: context.RequestContext
	(*ResponseContext)(nil), // 1: context.ResponseContext
	(*user_id.UserId)(nil),  // 2: user_id.UserId
}
var file_proto_context_proto_depIdxs = []int32{
	2, // 0: context.RequestContext.user_id:type_name -> user_id.UserId
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_context_proto_init() }
func file_proto_context_proto_init() {
	if File_proto_context_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_context_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestContext); i {
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
		file_proto_context_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseContext); i {
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
			RawDescriptor: file_proto_context_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_context_proto_goTypes,
		DependencyIndexes: file_proto_context_proto_depIdxs,
		MessageInfos:      file_proto_context_proto_msgTypes,
	}.Build()
	File_proto_context_proto = out.File
	file_proto_context_proto_rawDesc = nil
	file_proto_context_proto_goTypes = nil
	file_proto_context_proto_depIdxs = nil
}
