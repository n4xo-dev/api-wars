//*
// Comments proto file

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: comments.proto

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

type CommentDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content   string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	UserId    uint64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostId    uint64 `protobuf:"varint,4,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *CommentDTO) Reset() {
	*x = CommentDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentDTO) ProtoMessage() {}

func (x *CommentDTO) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentDTO.ProtoReflect.Descriptor instead.
func (*CommentDTO) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{0}
}

func (x *CommentDTO) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentDTO) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CommentDTO) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CommentDTO) GetPostId() uint64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *CommentDTO) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CommentDTO) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ListCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCommentsRequest) Reset() {
	*x = ListCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentsRequest) ProtoMessage() {}

func (x *ListCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentsRequest.ProtoReflect.Descriptor instead.
func (*ListCommentsRequest) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{1}
}

type ListCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*CommentDTO `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *ListCommentsResponse) Reset() {
	*x = ListCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentsResponse) ProtoMessage() {}

func (x *ListCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentsResponse.ProtoReflect.Descriptor instead.
func (*ListCommentsResponse) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{2}
}

func (x *ListCommentsResponse) GetComments() []*CommentDTO {
	if x != nil {
		return x.Comments
	}
	return nil
}

type GetCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCommentRequest) Reset() {
	*x = GetCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentRequest) ProtoMessage() {}

func (x *GetCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentRequest.ProtoReflect.Descriptor instead.
func (*GetCommentRequest) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{3}
}

func (x *GetCommentRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *CommentDTO `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *GetCommentResponse) Reset() {
	*x = GetCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentResponse) ProtoMessage() {}

func (x *GetCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentResponse.ProtoReflect.Descriptor instead.
func (*GetCommentResponse) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{4}
}

func (x *GetCommentResponse) GetComment() *CommentDTO {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CreateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	UserId  uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostId  uint64 `protobuf:"varint,3,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *CreateCommentRequest) Reset() {
	*x = CreateCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentRequest) ProtoMessage() {}

func (x *CreateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCommentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateCommentRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateCommentRequest) GetPostId() uint64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type CreateCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *CommentDTO `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *CreateCommentResponse) Reset() {
	*x = CreateCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentResponse) ProtoMessage() {}

func (x *CreateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentResponse.ProtoReflect.Descriptor instead.
func (*CreateCommentResponse) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{6}
}

func (x *CreateCommentResponse) GetComment() *CommentDTO {
	if x != nil {
		return x.Comment
	}
	return nil
}

type UpdateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	UserId  uint64 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostId  uint64 `protobuf:"varint,4,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *UpdateCommentRequest) Reset() {
	*x = UpdateCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentRequest) ProtoMessage() {}

func (x *UpdateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentRequest.ProtoReflect.Descriptor instead.
func (*UpdateCommentRequest) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCommentRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCommentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpdateCommentRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateCommentRequest) GetPostId() uint64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type UpdateCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *CommentDTO `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *UpdateCommentResponse) Reset() {
	*x = UpdateCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentResponse) ProtoMessage() {}

func (x *UpdateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentResponse.ProtoReflect.Descriptor instead.
func (*UpdateCommentResponse) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateCommentResponse) GetComment() *CommentDTO {
	if x != nil {
		return x.Comment
	}
	return nil
}

type DeleteCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCommentRequest) Reset() {
	*x = DeleteCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentRequest) ProtoMessage() {}

func (x *DeleteCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentRequest.ProtoReflect.Descriptor instead.
func (*DeleteCommentRequest) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteCommentRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted bool `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeleteCommentResponse) Reset() {
	*x = DeleteCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentResponse) ProtoMessage() {}

func (x *DeleteCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentResponse.ProtoReflect.Descriptor instead.
func (*DeleteCommentResponse) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteCommentResponse) GetDeleted() bool {
	if x != nil {
		return x.Deleted
	}
	return false
}

var File_comments_proto protoreflect.FileDescriptor

var file_comments_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa6, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x54, 0x4f, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x3f, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x44, 0x54, 0x4f, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x54, 0x4f, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x62, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x54, 0x4f, 0x52, 0x07,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x72, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x15, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x44,
	0x54, 0x4f, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x32, 0xc5, 0x02, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x15, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x15, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27,
	0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x4c, 0x6f,
	0x70, 0x65, 0x7a, 0x6f, 0x73, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2d, 0x77, 0x61, 0x72, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comments_proto_rawDescOnce sync.Once
	file_comments_proto_rawDescData = file_comments_proto_rawDesc
)

func file_comments_proto_rawDescGZIP() []byte {
	file_comments_proto_rawDescOnce.Do(func() {
		file_comments_proto_rawDescData = protoimpl.X.CompressGZIP(file_comments_proto_rawDescData)
	})
	return file_comments_proto_rawDescData
}

var file_comments_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_comments_proto_goTypes = []interface{}{
	(*CommentDTO)(nil),            // 0: CommentDTO
	(*ListCommentsRequest)(nil),   // 1: ListCommentsRequest
	(*ListCommentsResponse)(nil),  // 2: ListCommentsResponse
	(*GetCommentRequest)(nil),     // 3: GetCommentRequest
	(*GetCommentResponse)(nil),    // 4: GetCommentResponse
	(*CreateCommentRequest)(nil),  // 5: CreateCommentRequest
	(*CreateCommentResponse)(nil), // 6: CreateCommentResponse
	(*UpdateCommentRequest)(nil),  // 7: UpdateCommentRequest
	(*UpdateCommentResponse)(nil), // 8: UpdateCommentResponse
	(*DeleteCommentRequest)(nil),  // 9: DeleteCommentRequest
	(*DeleteCommentResponse)(nil), // 10: DeleteCommentResponse
}
var file_comments_proto_depIdxs = []int32{
	0,  // 0: ListCommentsResponse.comments:type_name -> CommentDTO
	0,  // 1: GetCommentResponse.comment:type_name -> CommentDTO
	0,  // 2: CreateCommentResponse.comment:type_name -> CommentDTO
	0,  // 3: UpdateCommentResponse.comment:type_name -> CommentDTO
	1,  // 4: CommentsService.ListComments:input_type -> ListCommentsRequest
	3,  // 5: CommentsService.GetComment:input_type -> GetCommentRequest
	5,  // 6: CommentsService.CreateComment:input_type -> CreateCommentRequest
	7,  // 7: CommentsService.UpdateComment:input_type -> UpdateCommentRequest
	9,  // 8: CommentsService.DeleteComment:input_type -> DeleteCommentRequest
	2,  // 9: CommentsService.ListComments:output_type -> ListCommentsResponse
	4,  // 10: CommentsService.GetComment:output_type -> GetCommentResponse
	6,  // 11: CommentsService.CreateComment:output_type -> CreateCommentResponse
	8,  // 12: CommentsService.UpdateComment:output_type -> UpdateCommentResponse
	10, // 13: CommentsService.DeleteComment:output_type -> DeleteCommentResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_comments_proto_init() }
func file_comments_proto_init() {
	if File_comments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentDTO); i {
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
		file_comments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentsRequest); i {
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
		file_comments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCommentsResponse); i {
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
		file_comments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentRequest); i {
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
		file_comments_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentResponse); i {
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
		file_comments_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentRequest); i {
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
		file_comments_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentResponse); i {
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
		file_comments_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentRequest); i {
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
		file_comments_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentResponse); i {
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
		file_comments_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentRequest); i {
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
		file_comments_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentResponse); i {
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
			RawDescriptor: file_comments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comments_proto_goTypes,
		DependencyIndexes: file_comments_proto_depIdxs,
		MessageInfos:      file_comments_proto_msgTypes,
	}.Build()
	File_comments_proto = out.File
	file_comments_proto_rawDesc = nil
	file_comments_proto_goTypes = nil
	file_comments_proto_depIdxs = nil
}
