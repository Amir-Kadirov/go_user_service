// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: branch.proto

package user_service

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

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Longitude float64 `protobuf:"fixed64,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude  float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

type BranchPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BranchPrimaryKey) Reset() {
	*x = BranchPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchPrimaryKey) ProtoMessage() {}

func (x *BranchPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchPrimaryKey.ProtoReflect.Descriptor instead.
func (*BranchPrimaryKey) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{1}
}

func (x *BranchPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateBranch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Location *Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"` // Note: Changed to lowercase "location"
}

func (x *CreateBranch) Reset() {
	*x = CreateBranch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBranch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBranch) ProtoMessage() {}

func (x *CreateBranch) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBranch.ProtoReflect.Descriptor instead.
func (*CreateBranch) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBranch) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *CreateBranch) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type Branch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Location *Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"` // Note: Changed to lowercase "location"
	Id       string    `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Branch) Reset() {
	*x = Branch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Branch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Branch) ProtoMessage() {}

func (x *Branch) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Branch.ProtoReflect.Descriptor instead.
func (*Branch) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{3}
}

func (x *Branch) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Branch) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

func (x *Branch) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateBranchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address  string    `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Location *Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *UpdateBranchRequest) Reset() {
	*x = UpdateBranchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBranchRequest) ProtoMessage() {}

func (x *UpdateBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBranchRequest.ProtoReflect.Descriptor instead.
func (*UpdateBranchRequest) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBranchRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *UpdateBranchRequest) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type BMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BMessage) Reset() {
	*x = BMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BMessage) ProtoMessage() {}

func (x *BMessage) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BMessage.ProtoReflect.Descriptor instead.
func (*BMessage) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{5}
}

func (x *BMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetListBranchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListBranchRequest) Reset() {
	*x = GetListBranchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListBranchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListBranchRequest) ProtoMessage() {}

func (x *GetListBranchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListBranchRequest.ProtoReflect.Descriptor instead.
func (*GetListBranchRequest) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{6}
}

func (x *GetListBranchRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListBranchRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListBranchRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListBranchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count    int64     `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`      // Note: Changed to lowercase "count"
	Branches []*Branch `protobuf:"bytes,2,rep,name=branches,proto3" json:"branches,omitempty"` // Note: Changed to "branches"
}

func (x *GetListBranchResponse) Reset() {
	*x = GetListBranchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListBranchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListBranchResponse) ProtoMessage() {}

func (x *GetListBranchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_branch_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListBranchResponse.ProtoReflect.Descriptor instead.
func (*GetListBranchResponse) Descriptor() ([]byte, []int) {
	return file_branch_proto_rawDescGZIP(), []int{7}
}

func (x *GetListBranchResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListBranchResponse) GetBranches() []*Branch {
	if x != nil {
		return x.Branches
	}
	return nil
}

var File_branch_proto protoreflect.FileDescriptor

var file_branch_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x44, 0x0a, 0x08,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e,
	0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5c, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x32, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x66, 0x0a, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x63, 0x0a, 0x13,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x24, 0x0a, 0x08, 0x42, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x5f, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x08, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x32, 0xfb, 0x02, 0x0a, 0x0d, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x1a, 0x1e,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x14, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_branch_proto_rawDescOnce sync.Once
	file_branch_proto_rawDescData = file_branch_proto_rawDesc
)

func file_branch_proto_rawDescGZIP() []byte {
	file_branch_proto_rawDescOnce.Do(func() {
		file_branch_proto_rawDescData = protoimpl.X.CompressGZIP(file_branch_proto_rawDescData)
	})
	return file_branch_proto_rawDescData
}

var file_branch_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_branch_proto_goTypes = []interface{}{
	(*Location)(nil),              // 0: user_service.Location
	(*BranchPrimaryKey)(nil),      // 1: user_service.BranchPrimaryKey
	(*CreateBranch)(nil),          // 2: user_service.CreateBranch
	(*Branch)(nil),                // 3: user_service.Branch
	(*UpdateBranchRequest)(nil),   // 4: user_service.UpdateBranchRequest
	(*BMessage)(nil),              // 5: user_service.BMessage
	(*GetListBranchRequest)(nil),  // 6: user_service.GetListBranchRequest
	(*GetListBranchResponse)(nil), // 7: user_service.GetListBranchResponse
}
var file_branch_proto_depIdxs = []int32{
	0, // 0: user_service.CreateBranch.location:type_name -> user_service.Location
	0, // 1: user_service.Branch.location:type_name -> user_service.Location
	0, // 2: user_service.UpdateBranchRequest.location:type_name -> user_service.Location
	3, // 3: user_service.GetListBranchResponse.branches:type_name -> user_service.Branch
	2, // 4: user_service.BranchService.Create:input_type -> user_service.CreateBranch
	1, // 5: user_service.BranchService.GetByID:input_type -> user_service.BranchPrimaryKey
	6, // 6: user_service.BranchService.GetList:input_type -> user_service.GetListBranchRequest
	4, // 7: user_service.BranchService.Update:input_type -> user_service.UpdateBranchRequest
	1, // 8: user_service.BranchService.Delete:input_type -> user_service.BranchPrimaryKey
	1, // 9: user_service.BranchService.Create:output_type -> user_service.BranchPrimaryKey
	3, // 10: user_service.BranchService.GetByID:output_type -> user_service.Branch
	7, // 11: user_service.BranchService.GetList:output_type -> user_service.GetListBranchResponse
	5, // 12: user_service.BranchService.Update:output_type -> user_service.BMessage
	5, // 13: user_service.BranchService.Delete:output_type -> user_service.BMessage
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_branch_proto_init() }
func file_branch_proto_init() {
	if File_branch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_branch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_branch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchPrimaryKey); i {
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
		file_branch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBranch); i {
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
		file_branch_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Branch); i {
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
		file_branch_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBranchRequest); i {
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
		file_branch_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BMessage); i {
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
		file_branch_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListBranchRequest); i {
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
		file_branch_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListBranchResponse); i {
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
			RawDescriptor: file_branch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_branch_proto_goTypes,
		DependencyIndexes: file_branch_proto_depIdxs,
		MessageInfos:      file_branch_proto_msgTypes,
	}.Build()
	File_branch_proto = out.File
	file_branch_proto_rawDesc = nil
	file_branch_proto_goTypes = nil
	file_branch_proto_depIdxs = nil
}
