// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: proto/api_key.proto

package api_key

import (
	context "github.com/liyouxina/buildzone/protogen/context"
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

type ApiKey_Capability int32

const (
	ApiKey_UNKNOWN_CAPABILITY           ApiKey_Capability = 0
	ApiKey_CACHE_WRITE_CAPABILITY       ApiKey_Capability = 1
	ApiKey_REGISTER_EXECUTOR_CAPABILITY ApiKey_Capability = 2
)

// Enum value maps for ApiKey_Capability.
var (
	ApiKey_Capability_name = map[int32]string{
		0: "UNKNOWN_CAPABILITY",
		1: "CACHE_WRITE_CAPABILITY",
		2: "REGISTER_EXECUTOR_CAPABILITY",
	}
	ApiKey_Capability_value = map[string]int32{
		"UNKNOWN_CAPABILITY":           0,
		"CACHE_WRITE_CAPABILITY":       1,
		"REGISTER_EXECUTOR_CAPABILITY": 2,
	}
)

func (x ApiKey_Capability) Enum() *ApiKey_Capability {
	p := new(ApiKey_Capability)
	*p = x
	return p
}

func (x ApiKey_Capability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApiKey_Capability) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_api_key_proto_enumTypes[0].Descriptor()
}

func (ApiKey_Capability) Type() protoreflect.EnumType {
	return &file_proto_api_key_proto_enumTypes[0]
}

func (x ApiKey_Capability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApiKey_Capability.Descriptor instead.
func (ApiKey_Capability) EnumDescriptor() ([]byte, []int) {
	return file_proto_api_key_proto_rawDescGZIP(), []int{0, 0}
}

type ApiKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value               string              `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Label               string              `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Capability          []ApiKey_Capability `protobuf:"varint,4,rep,packed,name=capability,proto3,enum=api_key.ApiKey_Capability" json:"capability,omitempty"`
	VisibleToDevelopers bool                `protobuf:"varint,5,opt,name=visible_to_developers,json=visibleToDevelopers,proto3" json:"visible_to_developers,omitempty"`
}

func (x *ApiKey) Reset() {
	*x = ApiKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApiKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKey) ProtoMessage() {}

func (x *ApiKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKey.ProtoReflect.Descriptor instead.
func (*ApiKey) Descriptor() ([]byte, []int) {
	return file_proto_api_key_proto_rawDescGZIP(), []int{0}
}

func (x *ApiKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApiKey) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ApiKey) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ApiKey) GetCapability() []ApiKey_Capability {
	if x != nil {
		return x.Capability
	}
	return nil
}

func (x *ApiKey) GetVisibleToDevelopers() bool {
	if x != nil {
		return x.VisibleToDevelopers
	}
	return false
}

type CreateApiKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestContext      *context.RequestContext `protobuf:"bytes,1,opt,name=request_context,json=requestContext,proto3" json:"request_context,omitempty"`
	GroupId             string                  `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Label               string                  `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Capability          []ApiKey_Capability     `protobuf:"varint,4,rep,packed,name=capability,proto3,enum=api_key.ApiKey_Capability" json:"capability,omitempty"`
	VisibleToDevelopers bool                    `protobuf:"varint,5,opt,name=visible_to_developers,json=visibleToDevelopers,proto3" json:"visible_to_developers,omitempty"`
}

func (x *CreateApiKeyRequest) Reset() {
	*x = CreateApiKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_key_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApiKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiKeyRequest) ProtoMessage() {}

func (x *CreateApiKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_key_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiKeyRequest.ProtoReflect.Descriptor instead.
func (*CreateApiKeyRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_key_proto_rawDescGZIP(), []int{1}
}

func (x *CreateApiKeyRequest) GetRequestContext() *context.RequestContext {
	if x != nil {
		return x.RequestContext
	}
	return nil
}

func (x *CreateApiKeyRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *CreateApiKeyRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *CreateApiKeyRequest) GetCapability() []ApiKey_Capability {
	if x != nil {
		return x.Capability
	}
	return nil
}

func (x *CreateApiKeyRequest) GetVisibleToDevelopers() bool {
	if x != nil {
		return x.VisibleToDevelopers
	}
	return false
}

type CreateApiKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseContext *context.ResponseContext `protobuf:"bytes,1,opt,name=response_context,json=responseContext,proto3" json:"response_context,omitempty"`
	ApiKey          *ApiKey                  `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
}

func (x *CreateApiKeyResponse) Reset() {
	*x = CreateApiKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_key_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateApiKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiKeyResponse) ProtoMessage() {}

func (x *CreateApiKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_key_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiKeyResponse.ProtoReflect.Descriptor instead.
func (*CreateApiKeyResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_key_proto_rawDescGZIP(), []int{2}
}

func (x *CreateApiKeyResponse) GetResponseContext() *context.ResponseContext {
	if x != nil {
		return x.ResponseContext
	}
	return nil
}

func (x *CreateApiKeyResponse) GetApiKey() *ApiKey {
	if x != nil {
		return x.ApiKey
	}
	return nil
}

type GetApiKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestContext *context.RequestContext `protobuf:"bytes,1,opt,name=request_context,json=requestContext,proto3" json:"request_context,omitempty"`
	GroupId        string                  `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
}

func (x *GetApiKeysRequest) Reset() {
	*x = GetApiKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_key_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApiKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApiKeysRequest) ProtoMessage() {}

func (x *GetApiKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_key_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApiKeysRequest.ProtoReflect.Descriptor instead.
func (*GetApiKeysRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_key_proto_rawDescGZIP(), []int{3}
}

func (x *GetApiKeysRequest) GetRequestContext() *context.RequestContext {
	if x != nil {
		return x.RequestContext
	}
	return nil
}

func (x *GetApiKeysRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

type GetApiKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseContext *context.ResponseContext `protobuf:"bytes,1,opt,name=response_context,json=responseContext,proto3" json:"response_context,omitempty"`
	ApiKey          []*ApiKey                `protobuf:"bytes,2,rep,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
}

func (x *GetApiKeysResponse) Reset() {
	*x = GetApiKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_key_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetApiKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApiKeysResponse) ProtoMessage() {}

func (x *GetApiKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_key_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApiKeysResponse.ProtoReflect.Descriptor instead.
func (*GetApiKeysResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_key_proto_rawDescGZIP(), []int{4}
}

func (x *GetApiKeysResponse) GetResponseContext() *context.ResponseContext {
	if x != nil {
		return x.ResponseContext
	}
	return nil
}

func (x *GetApiKeysResponse) GetApiKey() []*ApiKey {
	if x != nil {
		return x.ApiKey
	}
	return nil
}

type UpdateApiKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestContext      *context.RequestContext `protobuf:"bytes,1,opt,name=request_context,json=requestContext,proto3" json:"request_context,omitempty"`
	Id                  string                  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Label               string                  `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
	Capability          []ApiKey_Capability     `protobuf:"varint,4,rep,packed,name=capability,proto3,enum=api_key.ApiKey_Capability" json:"capability,omitempty"`
	VisibleToDevelopers bool                    `protobuf:"varint,5,opt,name=visible_to_developers,json=visibleToDevelopers,proto3" json:"visible_to_developers,omitempty"`
}

func (x *UpdateApiKeyRequest) Reset() {
	*x = UpdateApiKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_key_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateApiKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateApiKeyRequest) ProtoMessage() {}

func (x *UpdateApiKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_key_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateApiKeyRequest.ProtoReflect.Descriptor instead.
func (*UpdateApiKeyRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_key_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateApiKeyRequest) GetRequestContext() *context.RequestContext {
	if x != nil {
		return x.RequestContext
	}
	return nil
}

func (x *UpdateApiKeyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateApiKeyRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *UpdateApiKeyRequest) GetCapability() []ApiKey_Capability {
	if x != nil {
		return x.Capability
	}
	return nil
}

func (x *UpdateApiKeyRequest) GetVisibleToDevelopers() bool {
	if x != nil {
		return x.VisibleToDevelopers
	}
	return false
}

type UpdateApiKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseContext *context.ResponseContext `protobuf:"bytes,1,opt,name=response_context,json=responseContext,proto3" json:"response_context,omitempty"`
}

func (x *UpdateApiKeyResponse) Reset() {
	*x = UpdateApiKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_key_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateApiKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateApiKeyResponse) ProtoMessage() {}

func (x *UpdateApiKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_key_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateApiKeyResponse.ProtoReflect.Descriptor instead.
func (*UpdateApiKeyResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_key_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateApiKeyResponse) GetResponseContext() *context.ResponseContext {
	if x != nil {
		return x.ResponseContext
	}
	return nil
}

type DeleteApiKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestContext *context.RequestContext `protobuf:"bytes,1,opt,name=request_context,json=requestContext,proto3" json:"request_context,omitempty"`
	Id             string                  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteApiKeyRequest) Reset() {
	*x = DeleteApiKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_key_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApiKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApiKeyRequest) ProtoMessage() {}

func (x *DeleteApiKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_key_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApiKeyRequest.ProtoReflect.Descriptor instead.
func (*DeleteApiKeyRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_key_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteApiKeyRequest) GetRequestContext() *context.RequestContext {
	if x != nil {
		return x.RequestContext
	}
	return nil
}

func (x *DeleteApiKeyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteApiKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseContext *context.ResponseContext `protobuf:"bytes,1,opt,name=response_context,json=responseContext,proto3" json:"response_context,omitempty"`
}

func (x *DeleteApiKeyResponse) Reset() {
	*x = DeleteApiKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_key_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteApiKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApiKeyResponse) ProtoMessage() {}

func (x *DeleteApiKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_key_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApiKeyResponse.ProtoReflect.Descriptor instead.
func (*DeleteApiKeyResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_key_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteApiKeyResponse) GetResponseContext() *context.ResponseContext {
	if x != nil {
		return x.ResponseContext
	}
	return nil
}

var File_proto_api_key_proto protoreflect.FileDescriptor

var file_proto_api_key_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x1a, 0x13,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x98, 0x02, 0x0a, 0x06, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x2e,
	0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c,
	0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x54, 0x6f,
	0x44, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x73, 0x22, 0x62, 0x0a, 0x0a, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x43, 0x41, 0x50, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x00,
	0x12, 0x1a, 0x0a, 0x16, 0x43, 0x41, 0x43, 0x48, 0x45, 0x5f, 0x57, 0x52, 0x49, 0x54, 0x45, 0x5f,
	0x43, 0x41, 0x50, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x01, 0x12, 0x20, 0x0a, 0x1c,
	0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x4f,
	0x52, 0x5f, 0x43, 0x41, 0x50, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x10, 0x02, 0x22, 0xf8,
	0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x2e, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x44,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x73, 0x22, 0x85, 0x01, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x28, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x6b,
	0x65, 0x79, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x22, 0x70, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x10, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x28, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x22, 0xed, 0x01, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x40, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x2e, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x2e, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x32, 0x0a, 0x15, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x5f, 0x74, 0x6f, 0x5f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x76, 0x69, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x44,
	0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x73, 0x22, 0x5b, 0x0a, 0x14, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x43, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x67, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a,
	0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x5b, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x0f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_key_proto_rawDescOnce sync.Once
	file_proto_api_key_proto_rawDescData = file_proto_api_key_proto_rawDesc
)

func file_proto_api_key_proto_rawDescGZIP() []byte {
	file_proto_api_key_proto_rawDescOnce.Do(func() {
		file_proto_api_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_key_proto_rawDescData)
	})
	return file_proto_api_key_proto_rawDescData
}

var file_proto_api_key_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_api_key_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_api_key_proto_goTypes = []interface{}{
	(ApiKey_Capability)(0),          // 0: api_key.ApiKey.Capability
	(*ApiKey)(nil),                  // 1: api_key.ApiKey
	(*CreateApiKeyRequest)(nil),     // 2: api_key.CreateApiKeyRequest
	(*CreateApiKeyResponse)(nil),    // 3: api_key.CreateApiKeyResponse
	(*GetApiKeysRequest)(nil),       // 4: api_key.GetApiKeysRequest
	(*GetApiKeysResponse)(nil),      // 5: api_key.GetApiKeysResponse
	(*UpdateApiKeyRequest)(nil),     // 6: api_key.UpdateApiKeyRequest
	(*UpdateApiKeyResponse)(nil),    // 7: api_key.UpdateApiKeyResponse
	(*DeleteApiKeyRequest)(nil),     // 8: api_key.DeleteApiKeyRequest
	(*DeleteApiKeyResponse)(nil),    // 9: api_key.DeleteApiKeyResponse
	(*context.RequestContext)(nil),  // 10: context.RequestContext
	(*context.ResponseContext)(nil), // 11: context.ResponseContext
}
var file_proto_api_key_proto_depIdxs = []int32{
	0,  // 0: api_key.ApiKey.capability:type_name -> api_key.ApiKey.Capability
	10, // 1: api_key.CreateApiKeyRequest.request_context:type_name -> context.RequestContext
	0,  // 2: api_key.CreateApiKeyRequest.capability:type_name -> api_key.ApiKey.Capability
	11, // 3: api_key.CreateApiKeyResponse.response_context:type_name -> context.ResponseContext
	1,  // 4: api_key.CreateApiKeyResponse.api_key:type_name -> api_key.ApiKey
	10, // 5: api_key.GetApiKeysRequest.request_context:type_name -> context.RequestContext
	11, // 6: api_key.GetApiKeysResponse.response_context:type_name -> context.ResponseContext
	1,  // 7: api_key.GetApiKeysResponse.api_key:type_name -> api_key.ApiKey
	10, // 8: api_key.UpdateApiKeyRequest.request_context:type_name -> context.RequestContext
	0,  // 9: api_key.UpdateApiKeyRequest.capability:type_name -> api_key.ApiKey.Capability
	11, // 10: api_key.UpdateApiKeyResponse.response_context:type_name -> context.ResponseContext
	10, // 11: api_key.DeleteApiKeyRequest.request_context:type_name -> context.RequestContext
	11, // 12: api_key.DeleteApiKeyResponse.response_context:type_name -> context.ResponseContext
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_proto_api_key_proto_init() }
func file_proto_api_key_proto_init() {
	if File_proto_api_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_key_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApiKey); i {
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
		file_proto_api_key_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApiKeyRequest); i {
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
		file_proto_api_key_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateApiKeyResponse); i {
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
		file_proto_api_key_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApiKeysRequest); i {
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
		file_proto_api_key_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetApiKeysResponse); i {
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
		file_proto_api_key_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateApiKeyRequest); i {
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
		file_proto_api_key_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateApiKeyResponse); i {
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
		file_proto_api_key_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApiKeyRequest); i {
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
		file_proto_api_key_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteApiKeyResponse); i {
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
			RawDescriptor: file_proto_api_key_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_api_key_proto_goTypes,
		DependencyIndexes: file_proto_api_key_proto_depIdxs,
		EnumInfos:         file_proto_api_key_proto_enumTypes,
		MessageInfos:      file_proto_api_key_proto_msgTypes,
	}.Build()
	File_proto_api_key_proto = out.File
	file_proto_api_key_proto_rawDesc = nil
	file_proto_api_key_proto_goTypes = nil
	file_proto_api_key_proto_depIdxs = nil
}
