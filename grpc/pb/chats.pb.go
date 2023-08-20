//
// Chats proto file

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: chats.proto

package pb

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

type Chat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Messages     []*MessageDTO `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
	Participants []*UserDTO    `protobuf:"bytes,3,rep,name=participants,proto3" json:"participants,omitempty"`
	CreatedAt    string        `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string        `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt    string        `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Chat) Reset() {
	*x = Chat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{0}
}

func (x *Chat) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Chat) GetMessages() []*MessageDTO {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *Chat) GetParticipants() []*UserDTO {
	if x != nil {
		return x.Participants
	}
	return nil
}

func (x *Chat) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Chat) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Chat) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type ListChatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListChatsRequest) Reset() {
	*x = ListChatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChatsRequest) ProtoMessage() {}

func (x *ListChatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChatsRequest.ProtoReflect.Descriptor instead.
func (*ListChatsRequest) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{1}
}

type ListChatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chats []*Chat `protobuf:"bytes,1,rep,name=chats,proto3" json:"chats,omitempty"`
}

func (x *ListChatsResponse) Reset() {
	*x = ListChatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListChatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListChatsResponse) ProtoMessage() {}

func (x *ListChatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListChatsResponse.ProtoReflect.Descriptor instead.
func (*ListChatsResponse) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{2}
}

func (x *ListChatsResponse) GetChats() []*Chat {
	if x != nil {
		return x.Chats
	}
	return nil
}

type GetChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetChatRequest) Reset() {
	*x = GetChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatRequest) ProtoMessage() {}

func (x *GetChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatRequest.ProtoReflect.Descriptor instead.
func (*GetChatRequest) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{3}
}

func (x *GetChatRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat *Chat `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *GetChatResponse) Reset() {
	*x = GetChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatResponse) ProtoMessage() {}

func (x *GetChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatResponse.ProtoReflect.Descriptor instead.
func (*GetChatResponse) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{4}
}

func (x *GetChatResponse) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type CreateChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat *Chat `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *CreateChatRequest) Reset() {
	*x = CreateChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatRequest) ProtoMessage() {}

func (x *CreateChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatRequest.ProtoReflect.Descriptor instead.
func (*CreateChatRequest) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{5}
}

func (x *CreateChatRequest) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type CreateChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat *Chat `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *CreateChatResponse) Reset() {
	*x = CreateChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateChatResponse) ProtoMessage() {}

func (x *CreateChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateChatResponse.ProtoReflect.Descriptor instead.
func (*CreateChatResponse) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{6}
}

func (x *CreateChatResponse) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type AddUsersToChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId  uint64   `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	UserIds []uint64 `protobuf:"varint,2,rep,packed,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *AddUsersToChatRequest) Reset() {
	*x = AddUsersToChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUsersToChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUsersToChatRequest) ProtoMessage() {}

func (x *AddUsersToChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUsersToChatRequest.ProtoReflect.Descriptor instead.
func (*AddUsersToChatRequest) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{7}
}

func (x *AddUsersToChatRequest) GetChatId() uint64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *AddUsersToChatRequest) GetUserIds() []uint64 {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type AddUsersToChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chat *Chat `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
}

func (x *AddUsersToChatResponse) Reset() {
	*x = AddUsersToChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUsersToChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUsersToChatResponse) ProtoMessage() {}

func (x *AddUsersToChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUsersToChatResponse.ProtoReflect.Descriptor instead.
func (*AddUsersToChatResponse) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{8}
}

func (x *AddUsersToChatResponse) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type DeleteChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteChatRequest) Reset() {
	*x = DeleteChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChatRequest) ProtoMessage() {}

func (x *DeleteChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChatRequest.ProtoReflect.Descriptor instead.
func (*DeleteChatRequest) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteChatRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteChatResponse) Reset() {
	*x = DeleteChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteChatResponse) ProtoMessage() {}

func (x *DeleteChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteChatResponse.ProtoReflect.Descriptor instead.
func (*DeleteChatResponse) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{10}
}

type GetChatMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetChatMessagesRequest) Reset() {
	*x = GetChatMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatMessagesRequest) ProtoMessage() {}

func (x *GetChatMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetChatMessagesRequest) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{11}
}

func (x *GetChatMessagesRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetChatMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*MessageDTO `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *GetChatMessagesResponse) Reset() {
	*x = GetChatMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatMessagesResponse) ProtoMessage() {}

func (x *GetChatMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetChatMessagesResponse) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{12}
}

func (x *GetChatMessagesResponse) GetMessages() []*MessageDTO {
	if x != nil {
		return x.Messages
	}
	return nil
}

type GetChatUserMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChatId uint64 `protobuf:"varint,1,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetChatUserMessagesRequest) Reset() {
	*x = GetChatUserMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatUserMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatUserMessagesRequest) ProtoMessage() {}

func (x *GetChatUserMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatUserMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetChatUserMessagesRequest) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{13}
}

func (x *GetChatUserMessagesRequest) GetChatId() uint64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

func (x *GetChatUserMessagesRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetChatUserMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*MessageDTO `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *GetChatUserMessagesResponse) Reset() {
	*x = GetChatUserMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chats_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetChatUserMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatUserMessagesResponse) ProtoMessage() {}

func (x *GetChatUserMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chats_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatUserMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetChatUserMessagesResponse) Descriptor() ([]byte, []int) {
	return file_chats_proto_rawDescGZIP(), []int{14}
}

func (x *GetChatUserMessagesResponse) GetMessages() []*MessageDTO {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_chats_proto protoreflect.FileDescriptor

var file_chats_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x68, 0x61, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x01, 0x0a, 0x04, 0x43,
	0x68, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x27, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44,
	0x54, 0x4f, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x0c,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x52, 0x0c, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x05, 0x63, 0x68, 0x61, 0x74, 0x73, 0x22, 0x29, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x63,
	0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x2e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x63,
	0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x2f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04,
	0x63, 0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x4b, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x54, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x73, 0x22, 0x33, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x54, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42,
	0x0a, 0x17, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x54, 0x4f, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x22, 0x4e, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x46, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x54, 0x4f,
	0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0xc0, 0x03, 0x0a, 0x05, 0x43,
	0x68, 0x61, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74,
	0x73, 0x12, 0x11, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x68, 0x61, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x74, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x54,
	0x6f, 0x43, 0x68, 0x61, 0x74, 0x12, 0x16, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x54, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e,
	0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x1b, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x27, 0x5a,
	0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x4c, 0x6f, 0x70,
	0x65, 0x7a, 0x6f, 0x73, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x77, 0x61, 0x72, 0x73, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chats_proto_rawDescOnce sync.Once
	file_chats_proto_rawDescData = file_chats_proto_rawDesc
)

func file_chats_proto_rawDescGZIP() []byte {
	file_chats_proto_rawDescOnce.Do(func() {
		file_chats_proto_rawDescData = protoimpl.X.CompressGZIP(file_chats_proto_rawDescData)
	})
	return file_chats_proto_rawDescData
}

var file_chats_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_chats_proto_goTypes = []interface{}{
	(*Chat)(nil),                        // 0: Chat
	(*ListChatsRequest)(nil),            // 1: ListChatsRequest
	(*ListChatsResponse)(nil),           // 2: ListChatsResponse
	(*GetChatRequest)(nil),              // 3: GetChatRequest
	(*GetChatResponse)(nil),             // 4: GetChatResponse
	(*CreateChatRequest)(nil),           // 5: CreateChatRequest
	(*CreateChatResponse)(nil),          // 6: CreateChatResponse
	(*AddUsersToChatRequest)(nil),       // 7: AddUsersToChatRequest
	(*AddUsersToChatResponse)(nil),      // 8: AddUsersToChatResponse
	(*DeleteChatRequest)(nil),           // 9: DeleteChatRequest
	(*DeleteChatResponse)(nil),          // 10: DeleteChatResponse
	(*GetChatMessagesRequest)(nil),      // 11: GetChatMessagesRequest
	(*GetChatMessagesResponse)(nil),     // 12: GetChatMessagesResponse
	(*GetChatUserMessagesRequest)(nil),  // 13: GetChatUserMessagesRequest
	(*GetChatUserMessagesResponse)(nil), // 14: GetChatUserMessagesResponse
	(*MessageDTO)(nil),                  // 15: MessageDTO
	(*UserDTO)(nil),                     // 16: UserDTO
}
var file_chats_proto_depIdxs = []int32{
	15, // 0: Chat.messages:type_name -> MessageDTO
	16, // 1: Chat.participants:type_name -> UserDTO
	0,  // 2: ListChatsResponse.chats:type_name -> Chat
	0,  // 3: GetChatResponse.chat:type_name -> Chat
	0,  // 4: CreateChatRequest.chat:type_name -> Chat
	0,  // 5: CreateChatResponse.chat:type_name -> Chat
	0,  // 6: AddUsersToChatResponse.chat:type_name -> Chat
	15, // 7: GetChatMessagesResponse.messages:type_name -> MessageDTO
	15, // 8: GetChatUserMessagesResponse.messages:type_name -> MessageDTO
	1,  // 9: Chats.ListChats:input_type -> ListChatsRequest
	3,  // 10: Chats.GetChat:input_type -> GetChatRequest
	5,  // 11: Chats.CreateChat:input_type -> CreateChatRequest
	7,  // 12: Chats.AddUsersToChat:input_type -> AddUsersToChatRequest
	9,  // 13: Chats.DeleteChat:input_type -> DeleteChatRequest
	11, // 14: Chats.GetChatMessages:input_type -> GetChatMessagesRequest
	13, // 15: Chats.GetChatUserMessages:input_type -> GetChatUserMessagesRequest
	2,  // 16: Chats.ListChats:output_type -> ListChatsResponse
	4,  // 17: Chats.GetChat:output_type -> GetChatResponse
	6,  // 18: Chats.CreateChat:output_type -> CreateChatResponse
	8,  // 19: Chats.AddUsersToChat:output_type -> AddUsersToChatResponse
	10, // 20: Chats.DeleteChat:output_type -> DeleteChatResponse
	12, // 21: Chats.GetChatMessages:output_type -> GetChatMessagesResponse
	14, // 22: Chats.GetChatUserMessages:output_type -> GetChatUserMessagesResponse
	16, // [16:23] is the sub-list for method output_type
	9,  // [9:16] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_chats_proto_init() }
func file_chats_proto_init() {
	if File_chats_proto != nil {
		return
	}
	file_messages_proto_init()
	file_users_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chats_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chat); i {
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
		file_chats_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChatsRequest); i {
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
		file_chats_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListChatsResponse); i {
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
		file_chats_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatRequest); i {
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
		file_chats_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatResponse); i {
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
		file_chats_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatRequest); i {
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
		file_chats_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateChatResponse); i {
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
		file_chats_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUsersToChatRequest); i {
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
		file_chats_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUsersToChatResponse); i {
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
		file_chats_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChatRequest); i {
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
		file_chats_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteChatResponse); i {
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
		file_chats_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatMessagesRequest); i {
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
		file_chats_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatMessagesResponse); i {
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
		file_chats_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatUserMessagesRequest); i {
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
		file_chats_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetChatUserMessagesResponse); i {
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
			RawDescriptor: file_chats_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chats_proto_goTypes,
		DependencyIndexes: file_chats_proto_depIdxs,
		MessageInfos:      file_chats_proto_msgTypes,
	}.Build()
	File_chats_proto = out.File
	file_chats_proto_rawDesc = nil
	file_chats_proto_goTypes = nil
	file_chats_proto_depIdxs = nil
}
