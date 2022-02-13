// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: proto/icc/icc.proto

package icc

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

type RichImageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageURL  string   `protobuf:"bytes,1,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	ImageID   string   `protobuf:"bytes,2,opt,name=ImageID,proto3" json:"ImageID,omitempty"`
	Tags      []string `protobuf:"bytes,3,rep,name=Tags,proto3" json:"Tags,omitempty"`
	Timestamp int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *RichImageInfo) Reset() {
	*x = RichImageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_icc_icc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RichImageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RichImageInfo) ProtoMessage() {}

func (x *RichImageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_icc_icc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RichImageInfo.ProtoReflect.Descriptor instead.
func (*RichImageInfo) Descriptor() ([]byte, []int) {
	return file_proto_icc_icc_proto_rawDescGZIP(), []int{0}
}

func (x *RichImageInfo) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *RichImageInfo) GetImageID() string {
	if x != nil {
		return x.ImageID
	}
	return ""
}

func (x *RichImageInfo) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *RichImageInfo) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type GetImagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Query images before.
	// Using Linux timestamp format
	Before int64 `protobuf:"varint,1,opt,name=Before,proto3" json:"Before,omitempty"`
	// The tags that queried images should contains
	Tags []string `protobuf:"bytes,2,rep,name=Tags,proto3" json:"Tags,omitempty"`
	// Result number limit
	Limit uint32 `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetImagesRequest) Reset() {
	*x = GetImagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_icc_icc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImagesRequest) ProtoMessage() {}

func (x *GetImagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_icc_icc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImagesRequest.ProtoReflect.Descriptor instead.
func (*GetImagesRequest) Descriptor() ([]byte, []int) {
	return file_proto_icc_icc_proto_rawDescGZIP(), []int{1}
}

func (x *GetImagesRequest) GetBefore() int64 {
	if x != nil {
		return x.Before
	}
	return 0
}

func (x *GetImagesRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *GetImagesRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetImagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*RichImageInfo `protobuf:"bytes,2,rep,name=Images,proto3" json:"Images,omitempty"`
}

func (x *GetImagesResponse) Reset() {
	*x = GetImagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_icc_icc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetImagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetImagesResponse) ProtoMessage() {}

func (x *GetImagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_icc_icc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetImagesResponse.ProtoReflect.Descriptor instead.
func (*GetImagesResponse) Descriptor() ([]byte, []int) {
	return file_proto_icc_icc_proto_rawDescGZIP(), []int{2}
}

func (x *GetImagesResponse) GetImages() []*RichImageInfo {
	if x != nil {
		return x.Images
	}
	return nil
}

type GetRandomImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags  []string `protobuf:"bytes,1,rep,name=Tags,proto3" json:"Tags,omitempty"`
	Limit int32    `protobuf:"varint,2,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetRandomImageRequest) Reset() {
	*x = GetRandomImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_icc_icc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRandomImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRandomImageRequest) ProtoMessage() {}

func (x *GetRandomImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_icc_icc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRandomImageRequest.ProtoReflect.Descriptor instead.
func (*GetRandomImageRequest) Descriptor() ([]byte, []int) {
	return file_proto_icc_icc_proto_rawDescGZIP(), []int{3}
}

func (x *GetRandomImageRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *GetRandomImageRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetRandomImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Images []*RichImageInfo `protobuf:"bytes,2,rep,name=Images,proto3" json:"Images,omitempty"`
}

func (x *GetRandomImageResponse) Reset() {
	*x = GetRandomImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_icc_icc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRandomImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRandomImageResponse) ProtoMessage() {}

func (x *GetRandomImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_icc_icc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRandomImageResponse.ProtoReflect.Descriptor instead.
func (*GetRandomImageResponse) Descriptor() ([]byte, []int) {
	return file_proto_icc_icc_proto_rawDescGZIP(), []int{4}
}

func (x *GetRandomImageResponse) GetImages() []*RichImageInfo {
	if x != nil {
		return x.Images
	}
	return nil
}

type PreSignObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageType string `protobuf:"bytes,5,opt,name=ImageType,proto3" json:"ImageType,omitempty"`
	MD5Sum    string `protobuf:"bytes,6,opt,name=MD5Sum,proto3" json:"MD5Sum,omitempty"`
}

func (x *PreSignObjectRequest) Reset() {
	*x = PreSignObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_icc_icc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreSignObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreSignObjectRequest) ProtoMessage() {}

func (x *PreSignObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_icc_icc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreSignObjectRequest.ProtoReflect.Descriptor instead.
func (*PreSignObjectRequest) Descriptor() ([]byte, []int) {
	return file_proto_icc_icc_proto_rawDescGZIP(), []int{5}
}

func (x *PreSignObjectRequest) GetImageType() string {
	if x != nil {
		return x.ImageType
	}
	return ""
}

func (x *PreSignObjectRequest) GetMD5Sum() string {
	if x != nil {
		return x.MD5Sum
	}
	return ""
}

type PreSignObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreSignedURI string `protobuf:"bytes,5,opt,name=PreSignedURI,proto3" json:"PreSignedURI,omitempty"`
	ImageID      string `protobuf:"bytes,6,opt,name=ImageID,proto3" json:"ImageID,omitempty"`
}

func (x *PreSignObjectResponse) Reset() {
	*x = PreSignObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_icc_icc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreSignObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreSignObjectResponse) ProtoMessage() {}

func (x *PreSignObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_icc_icc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreSignObjectResponse.ProtoReflect.Descriptor instead.
func (*PreSignObjectResponse) Descriptor() ([]byte, []int) {
	return file_proto_icc_icc_proto_rawDescGZIP(), []int{6}
}

func (x *PreSignObjectResponse) GetPreSignedURI() string {
	if x != nil {
		return x.PreSignedURI
	}
	return ""
}

func (x *PreSignObjectResponse) GetImageID() string {
	if x != nil {
		return x.ImageID
	}
	return ""
}

type CompleteUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageID    string   `protobuf:"bytes,2,opt,name=ImageID,proto3" json:"ImageID,omitempty"`
	ExternalID string   `protobuf:"bytes,3,opt,name=ExternalID,proto3" json:"ExternalID,omitempty"`
	Tags       []string `protobuf:"bytes,4,rep,name=Tags,proto3" json:"Tags,omitempty"`
}

func (x *CompleteUploadRequest) Reset() {
	*x = CompleteUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_icc_icc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteUploadRequest) ProtoMessage() {}

func (x *CompleteUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_icc_icc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteUploadRequest.ProtoReflect.Descriptor instead.
func (*CompleteUploadRequest) Descriptor() ([]byte, []int) {
	return file_proto_icc_icc_proto_rawDescGZIP(), []int{7}
}

func (x *CompleteUploadRequest) GetImageID() string {
	if x != nil {
		return x.ImageID
	}
	return ""
}

func (x *CompleteUploadRequest) GetExternalID() string {
	if x != nil {
		return x.ExternalID
	}
	return ""
}

func (x *CompleteUploadRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type CompleteUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageID  string   `protobuf:"bytes,2,opt,name=ImageID,proto3" json:"ImageID,omitempty"`
	ImageURL string   `protobuf:"bytes,3,opt,name=ImageURL,proto3" json:"ImageURL,omitempty"`
	Tags     []string `protobuf:"bytes,4,rep,name=Tags,proto3" json:"Tags,omitempty"`
}

func (x *CompleteUploadResponse) Reset() {
	*x = CompleteUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_icc_icc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteUploadResponse) ProtoMessage() {}

func (x *CompleteUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_icc_icc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteUploadResponse.ProtoReflect.Descriptor instead.
func (*CompleteUploadResponse) Descriptor() ([]byte, []int) {
	return file_proto_icc_icc_proto_rawDescGZIP(), []int{8}
}

func (x *CompleteUploadResponse) GetImageID() string {
	if x != nil {
		return x.ImageID
	}
	return ""
}

func (x *CompleteUploadResponse) GetImageURL() string {
	if x != nil {
		return x.ImageURL
	}
	return ""
}

func (x *CompleteUploadResponse) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type AddTagToImageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageID string   `protobuf:"bytes,1,opt,name=ImageID,proto3" json:"ImageID,omitempty"`
	Tags    []string `protobuf:"bytes,2,rep,name=Tags,proto3" json:"Tags,omitempty"`
}

func (x *AddTagToImageRequest) Reset() {
	*x = AddTagToImageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_icc_icc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTagToImageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTagToImageRequest) ProtoMessage() {}

func (x *AddTagToImageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_icc_icc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTagToImageRequest.ProtoReflect.Descriptor instead.
func (*AddTagToImageRequest) Descriptor() ([]byte, []int) {
	return file_proto_icc_icc_proto_rawDescGZIP(), []int{9}
}

func (x *AddTagToImageRequest) GetImageID() string {
	if x != nil {
		return x.ImageID
	}
	return ""
}

func (x *AddTagToImageRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type AddTagToImageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImageID string   `protobuf:"bytes,2,opt,name=ImageID,proto3" json:"ImageID,omitempty"`
	Tags    []string `protobuf:"bytes,3,rep,name=Tags,proto3" json:"Tags,omitempty"`
}

func (x *AddTagToImageResponse) Reset() {
	*x = AddTagToImageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_icc_icc_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTagToImageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTagToImageResponse) ProtoMessage() {}

func (x *AddTagToImageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_icc_icc_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTagToImageResponse.ProtoReflect.Descriptor instead.
func (*AddTagToImageResponse) Descriptor() ([]byte, []int) {
	return file_proto_icc_icc_proto_rawDescGZIP(), []int{10}
}

func (x *AddTagToImageResponse) GetImageID() string {
	if x != nil {
		return x.ImageID
	}
	return ""
}

func (x *AddTagToImageResponse) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_proto_icc_icc_proto protoreflect.FileDescriptor

var file_proto_icc_icc_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x63, 0x63, 0x2f, 0x69, 0x63, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x69, 0x63, 0x63, 0x22, 0x77, 0x0a, 0x0d, 0x52, 0x69,
	0x63, 0x68, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49,
	0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x54, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x65, 0x66, 0x6f, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x69, 0x63, 0x63, 0x2e, 0x52, 0x69, 0x63, 0x68, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x44, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x69, 0x63, 0x63, 0x2e, 0x52, 0x69,
	0x63, 0x68, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x22, 0x4c, 0x0a, 0x14, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x44, 0x35,
	0x53, 0x75, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x44, 0x35, 0x53, 0x75,
	0x6d, 0x22, 0x55, 0x0a, 0x15, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72,
	0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x52, 0x49, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x55, 0x52, 0x49, 0x12, 0x18,
	0x0a, 0x07, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x44, 0x22, 0x65, 0x0a, 0x15, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54,
	0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73, 0x22,
	0x62, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x52, 0x4c, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x61, 0x67, 0x73, 0x22, 0x44, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x54, 0x61, 0x67, 0x54, 0x6f, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73, 0x22, 0x45, 0x0a, 0x15, 0x41, 0x64, 0x64,
	0x54, 0x61, 0x67, 0x54, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04,
	0x54, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73,
	0x32, 0xec, 0x02, 0x0a, 0x03, 0x49, 0x43, 0x43, 0x12, 0x4b, 0x0a, 0x12, 0x49, 0x73, 0x73, 0x75,
	0x65, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x19,
	0x2e, 0x69, 0x63, 0x63, 0x2e, 0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x69, 0x63, 0x63, 0x2e,
	0x50, 0x72, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x2e, 0x69, 0x63, 0x63, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x69, 0x63, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x39, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x69,
	0x63, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x69, 0x63, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x2e,
	0x69, 0x63, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x69, 0x63, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x67,
	0x73, 0x54, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x19, 0x2e, 0x69, 0x63, 0x63, 0x2e, 0x41,
	0x64, 0x64, 0x54, 0x61, 0x67, 0x54, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x69, 0x63, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x67,
	0x54, 0x6f, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x79,
	0x6c, 0x6f, 0x6e, 0x78, 0x2f, 0x69, 0x63, 0x63, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x69, 0x63, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_icc_icc_proto_rawDescOnce sync.Once
	file_proto_icc_icc_proto_rawDescData = file_proto_icc_icc_proto_rawDesc
)

func file_proto_icc_icc_proto_rawDescGZIP() []byte {
	file_proto_icc_icc_proto_rawDescOnce.Do(func() {
		file_proto_icc_icc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_icc_icc_proto_rawDescData)
	})
	return file_proto_icc_icc_proto_rawDescData
}

var file_proto_icc_icc_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_icc_icc_proto_goTypes = []interface{}{
	(*RichImageInfo)(nil),          // 0: icc.RichImageInfo
	(*GetImagesRequest)(nil),       // 1: icc.GetImagesRequest
	(*GetImagesResponse)(nil),      // 2: icc.GetImagesResponse
	(*GetRandomImageRequest)(nil),  // 3: icc.GetRandomImageRequest
	(*GetRandomImageResponse)(nil), // 4: icc.GetRandomImageResponse
	(*PreSignObjectRequest)(nil),   // 5: icc.PreSignObjectRequest
	(*PreSignObjectResponse)(nil),  // 6: icc.PreSignObjectResponse
	(*CompleteUploadRequest)(nil),  // 7: icc.CompleteUploadRequest
	(*CompleteUploadResponse)(nil), // 8: icc.CompleteUploadResponse
	(*AddTagToImageRequest)(nil),   // 9: icc.AddTagToImageRequest
	(*AddTagToImageResponse)(nil),  // 10: icc.AddTagToImageResponse
}
var file_proto_icc_icc_proto_depIdxs = []int32{
	0,  // 0: icc.GetImagesResponse.Images:type_name -> icc.RichImageInfo
	0,  // 1: icc.GetRandomImageResponse.Images:type_name -> icc.RichImageInfo
	5,  // 2: icc.ICC.IssuePreSignUpload:input_type -> icc.PreSignObjectRequest
	7,  // 3: icc.ICC.CompleteUpload:input_type -> icc.CompleteUploadRequest
	1,  // 4: icc.ICC.GetImage:input_type -> icc.GetImagesRequest
	3,  // 5: icc.ICC.GetRandomImage:input_type -> icc.GetRandomImageRequest
	9,  // 6: icc.ICC.AddTagsToImage:input_type -> icc.AddTagToImageRequest
	6,  // 7: icc.ICC.IssuePreSignUpload:output_type -> icc.PreSignObjectResponse
	8,  // 8: icc.ICC.CompleteUpload:output_type -> icc.CompleteUploadResponse
	2,  // 9: icc.ICC.GetImage:output_type -> icc.GetImagesResponse
	4,  // 10: icc.ICC.GetRandomImage:output_type -> icc.GetRandomImageResponse
	10, // 11: icc.ICC.AddTagsToImage:output_type -> icc.AddTagToImageResponse
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_proto_icc_icc_proto_init() }
func file_proto_icc_icc_proto_init() {
	if File_proto_icc_icc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_icc_icc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RichImageInfo); i {
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
		file_proto_icc_icc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImagesRequest); i {
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
		file_proto_icc_icc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetImagesResponse); i {
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
		file_proto_icc_icc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRandomImageRequest); i {
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
		file_proto_icc_icc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRandomImageResponse); i {
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
		file_proto_icc_icc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreSignObjectRequest); i {
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
		file_proto_icc_icc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreSignObjectResponse); i {
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
		file_proto_icc_icc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteUploadRequest); i {
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
		file_proto_icc_icc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompleteUploadResponse); i {
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
		file_proto_icc_icc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTagToImageRequest); i {
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
		file_proto_icc_icc_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTagToImageResponse); i {
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
			RawDescriptor: file_proto_icc_icc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_icc_icc_proto_goTypes,
		DependencyIndexes: file_proto_icc_icc_proto_depIdxs,
		MessageInfos:      file_proto_icc_icc_proto_msgTypes,
	}.Build()
	File_proto_icc_icc_proto = out.File
	file_proto_icc_icc_proto_rawDesc = nil
	file_proto_icc_icc_proto_goTypes = nil
	file_proto_icc_icc_proto_depIdxs = nil
}
