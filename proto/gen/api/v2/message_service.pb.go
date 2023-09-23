// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: api/v2/message_service.proto

package apiv2

import (
	_ "github.com/Nicknamezz00/mercury/proto/gen/google/api"
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

type Visibility int32

const (
	Visibility_VISIBILITY_UNSPECIFIED Visibility = 0
	Visibility_PRIVATE                Visibility = 1
	Visibility_PROTECTED              Visibility = 2
	Visibility_PUBLIC                 Visibility = 3
)

// Enum value maps for Visibility.
var (
	Visibility_name = map[int32]string{
		0: "VISIBILITY_UNSPECIFIED",
		1: "PRIVATE",
		2: "PROTECTED",
		3: "PUBLIC",
	}
	Visibility_value = map[string]int32{
		"VISIBILITY_UNSPECIFIED": 0,
		"PRIVATE":                1,
		"PROTECTED":              2,
		"PUBLIC":                 3,
	}
)

func (x Visibility) Enum() *Visibility {
	p := new(Visibility)
	*p = x
	return p
}

func (x Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v2_message_service_proto_enumTypes[0].Descriptor()
}

func (Visibility) Type() protoreflect.EnumType {
	return &file_api_v2_message_service_proto_enumTypes[0]
}

func (x Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Visibility.Descriptor instead.
func (Visibility) EnumDescriptor() ([]byte, []int) {
	return file_api_v2_message_service_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RowStatus        RowStatus  `protobuf:"varint,2,opt,name=row_status,json=rowStatus,proto3,enum=mercury.api.v2.RowStatus" json:"row_status,omitempty"`
	CreatorId        int32      `protobuf:"varint,3,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	CreatorTimestamp int64      `protobuf:"varint,4,opt,name=creator_timestamp,json=creatorTimestamp,proto3" json:"creator_timestamp,omitempty"`
	UpdatedTimestamp int64      `protobuf:"varint,5,opt,name=updated_timestamp,json=updatedTimestamp,proto3" json:"updated_timestamp,omitempty"`
	Content          string     `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Visibility       Visibility `protobuf:"varint,7,opt,name=visibility,proto3,enum=mercury.api.v2.Visibility" json:"visibility,omitempty"`
	Pinned           bool       `protobuf:"varint,8,opt,name=pinned,proto3" json:"pinned,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_message_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_message_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_api_v2_message_service_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetRowStatus() RowStatus {
	if x != nil {
		return x.RowStatus
	}
	return RowStatus_ROW_STATUS_UNSPECIFIED
}

func (x *Message) GetCreatorId() int32 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *Message) GetCreatorTimestamp() int64 {
	if x != nil {
		return x.CreatorTimestamp
	}
	return 0
}

func (x *Message) GetUpdatedTimestamp() int64 {
	if x != nil {
		return x.UpdatedTimestamp
	}
	return 0
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetVisibility() Visibility {
	if x != nil {
		return x.Visibility
	}
	return Visibility_VISIBILITY_UNSPECIFIED
}

func (x *Message) GetPinned() bool {
	if x != nil {
		return x.Pinned
	}
	return false
}

type ListMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Filter   string `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListMessagesRequest) Reset() {
	*x = ListMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_message_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessagesRequest) ProtoMessage() {}

func (x *ListMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_message_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessagesRequest.ProtoReflect.Descriptor instead.
func (*ListMessagesRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_message_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListMessagesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListMessagesRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListMessagesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *ListMessagesResponse) Reset() {
	*x = ListMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_message_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessagesResponse) ProtoMessage() {}

func (x *ListMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_message_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessagesResponse.ProtoReflect.Descriptor instead.
func (*ListMessagesResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_message_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListMessagesResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type GetMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMessageRequest) Reset() {
	*x = GetMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_message_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequest) ProtoMessage() {}

func (x *GetMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_message_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequest.ProtoReflect.Descriptor instead.
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_v2_message_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetMessageRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GetMessageResponse) Reset() {
	*x = GetMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v2_message_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageResponse) ProtoMessage() {}

func (x *GetMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v2_message_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageResponse.ProtoReflect.Descriptor instead.
func (*GetMessageResponse) Descriptor() ([]byte, []int) {
	return file_api_v2_message_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetMessageResponse) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

var File_api_v2_message_service_proto protoreflect.FileDescriptor

var file_api_v2_message_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a, 0x13,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x02, 0x0a, 0x07, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x0a, 0x72, 0x6f, 0x77, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x72,
	0x63, 0x75, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x6f, 0x77, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x09, 0x72, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x2b, 0x0a, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2b, 0x0a, 0x11,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x52, 0x0a, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x22, 0x5e, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x33, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x32, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x31, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2a, 0x50, 0x0a, 0x0a, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x1a, 0x0a, 0x16, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x52, 0x4f,
	0x54, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c,
	0x49, 0x43, 0x10, 0x03, 0x32, 0xfe, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x75, 0x72,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6d,
	0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x77, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x2e, 0x6d, 0x65, 0x72,
	0x63, 0x75, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x6d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x22, 0xda, 0x41, 0x02, 0x69, 0x64, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0xbb, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65,
	0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x42, 0x13, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x7a, 0x7a, 0x30, 0x30, 0x2f, 0x6d, 0x65, 0x72,
	0x63, 0x75, 0x72, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x32, 0x3b, 0x61, 0x70, 0x69, 0x76, 0x32, 0xa2, 0x02, 0x03, 0x4d, 0x41,
	0x58, 0xaa, 0x02, 0x0e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x2e, 0x41, 0x70, 0x69, 0x2e,
	0x56, 0x32, 0xca, 0x02, 0x0e, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x5c, 0x41, 0x70, 0x69,
	0x5c, 0x56, 0x32, 0xe2, 0x02, 0x1a, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x5c, 0x41, 0x70,
	0x69, 0x5c, 0x56, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x10, 0x4d, 0x65, 0x72, 0x63, 0x75, 0x72, 0x79, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x56, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v2_message_service_proto_rawDescOnce sync.Once
	file_api_v2_message_service_proto_rawDescData = file_api_v2_message_service_proto_rawDesc
)

func file_api_v2_message_service_proto_rawDescGZIP() []byte {
	file_api_v2_message_service_proto_rawDescOnce.Do(func() {
		file_api_v2_message_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v2_message_service_proto_rawDescData)
	})
	return file_api_v2_message_service_proto_rawDescData
}

var file_api_v2_message_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_v2_message_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_v2_message_service_proto_goTypes = []interface{}{
	(Visibility)(0),              // 0: mercury.api.v2.Visibility
	(*Message)(nil),              // 1: mercury.api.v2.Message
	(*ListMessagesRequest)(nil),  // 2: mercury.api.v2.ListMessagesRequest
	(*ListMessagesResponse)(nil), // 3: mercury.api.v2.ListMessagesResponse
	(*GetMessageRequest)(nil),    // 4: mercury.api.v2.GetMessageRequest
	(*GetMessageResponse)(nil),   // 5: mercury.api.v2.GetMessageResponse
	(RowStatus)(0),               // 6: mercury.api.v2.RowStatus
}
var file_api_v2_message_service_proto_depIdxs = []int32{
	6, // 0: mercury.api.v2.Message.row_status:type_name -> mercury.api.v2.RowStatus
	0, // 1: mercury.api.v2.Message.visibility:type_name -> mercury.api.v2.Visibility
	1, // 2: mercury.api.v2.ListMessagesResponse.messages:type_name -> mercury.api.v2.Message
	1, // 3: mercury.api.v2.GetMessageResponse.message:type_name -> mercury.api.v2.Message
	2, // 4: mercury.api.v2.MercuryService.ListMessages:input_type -> mercury.api.v2.ListMessagesRequest
	4, // 5: mercury.api.v2.MercuryService.GetMessage:input_type -> mercury.api.v2.GetMessageRequest
	3, // 6: mercury.api.v2.MercuryService.ListMessages:output_type -> mercury.api.v2.ListMessagesResponse
	5, // 7: mercury.api.v2.MercuryService.GetMessage:output_type -> mercury.api.v2.GetMessageResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_v2_message_service_proto_init() }
func file_api_v2_message_service_proto_init() {
	if File_api_v2_message_service_proto != nil {
		return
	}
	file_api_v2_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_v2_message_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_api_v2_message_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMessagesRequest); i {
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
		file_api_v2_message_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMessagesResponse); i {
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
		file_api_v2_message_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageRequest); i {
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
		file_api_v2_message_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMessageResponse); i {
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
			RawDescriptor: file_api_v2_message_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v2_message_service_proto_goTypes,
		DependencyIndexes: file_api_v2_message_service_proto_depIdxs,
		EnumInfos:         file_api_v2_message_service_proto_enumTypes,
		MessageInfos:      file_api_v2_message_service_proto_msgTypes,
	}.Build()
	File_api_v2_message_service_proto = out.File
	file_api_v2_message_service_proto_rawDesc = nil
	file_api_v2_message_service_proto_goTypes = nil
	file_api_v2_message_service_proto_depIdxs = nil
}