// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: community.proto

package community

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

type DEFAULTS int32

const (
	DEFAULTS_STRING_DEFAULT DEFAULTS = 0
	DEFAULTS_INT32_DEFAULT  DEFAULTS = 1
)

// Enum value maps for DEFAULTS.
var (
	DEFAULTS_name = map[int32]string{
		0: "STRING_DEFAULT",
		1: "INT32_DEFAULT",
	}
	DEFAULTS_value = map[string]int32{
		"STRING_DEFAULT": 0,
		"INT32_DEFAULT":  1,
	}
)

func (x DEFAULTS) Enum() *DEFAULTS {
	p := new(DEFAULTS)
	*p = x
	return p
}

func (x DEFAULTS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DEFAULTS) Descriptor() protoreflect.EnumDescriptor {
	return file_community_proto_enumTypes[0].Descriptor()
}

func (DEFAULTS) Type() protoreflect.EnumType {
	return &file_community_proto_enumTypes[0]
}

func (x DEFAULTS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DEFAULTS.Descriptor instead.
func (DEFAULTS) EnumDescriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{0}
}

type NUMBER_NUM int32

const (
	NUMBER_NUM_one   NUMBER_NUM = 0
	NUMBER_NUM_two   NUMBER_NUM = 1
	NUMBER_NUM_three NUMBER_NUM = 2
)

// Enum value maps for NUMBER_NUM.
var (
	NUMBER_NUM_name = map[int32]string{
		0: "one",
		1: "two",
		2: "three",
	}
	NUMBER_NUM_value = map[string]int32{
		"one":   0,
		"two":   1,
		"three": 2,
	}
)

func (x NUMBER_NUM) Enum() *NUMBER_NUM {
	p := new(NUMBER_NUM)
	*p = x
	return p
}

func (x NUMBER_NUM) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NUMBER_NUM) Descriptor() protoreflect.EnumDescriptor {
	return file_community_proto_enumTypes[1].Descriptor()
}

func (NUMBER_NUM) Type() protoreflect.EnumType {
	return &file_community_proto_enumTypes[1]
}

func (x NUMBER_NUM) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NUMBER_NUM.Descriptor instead.
func (NUMBER_NUM) EnumDescriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{1}
}

type STATUS int32

const (
	STATUS_ERROR        STATUS = 0
	STATUS_NOT_FOUND    STATUS = 1
	STATUS_WRONG_FORMAT STATUS = 2
	STATUS_OK           STATUS = 3
)

// Enum value maps for STATUS.
var (
	STATUS_name = map[int32]string{
		0: "ERROR",
		1: "NOT_FOUND",
		2: "WRONG_FORMAT",
		3: "OK",
	}
	STATUS_value = map[string]int32{
		"ERROR":        0,
		"NOT_FOUND":    1,
		"WRONG_FORMAT": 2,
		"OK":           3,
	}
)

func (x STATUS) Enum() *STATUS {
	p := new(STATUS)
	*p = x
	return p
}

func (x STATUS) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (STATUS) Descriptor() protoreflect.EnumDescriptor {
	return file_community_proto_enumTypes[2].Descriptor()
}

func (STATUS) Type() protoreflect.EnumType {
	return &file_community_proto_enumTypes[2]
}

func (x STATUS) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use STATUS.Descriptor instead.
func (STATUS) EnumDescriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{2}
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status STATUS `protobuf:"varint,1,opt,name=status,proto3,enum=community.STATUS" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetStatus() STATUS {
	if x != nil {
		return x.Status
	}
	return STATUS_ERROR
}

// temp
type SpacesCompact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SpacesCompact) Reset() {
	*x = SpacesCompact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpacesCompact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpacesCompact) ProtoMessage() {}

func (x *SpacesCompact) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpacesCompact.ProtoReflect.Descriptor instead.
func (*SpacesCompact) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{1}
}

func (x *SpacesCompact) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// request community data : Request
type CommunityGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestType    NUMBER_NUM `protobuf:"varint,1,opt,name=requestType,proto3,enum=community.NUMBER_NUM" json:"requestType,omitempty"`           // either 1 or 0
	Id             *string    `protobuf:"bytes,2,opt,name=id,proto3,oneof" json:"id,omitempty"`                                                  // search by id : requestType must be 0
	Query          *string    `protobuf:"bytes,3,opt,name=query,proto3,oneof" json:"query,omitempty"`                                            // search by query : requestType must be 1
	PageNumber     *string    `protobuf:"bytes,4,opt,name=page_number,json=pageNumber,proto3,oneof" json:"page_number,omitempty"`                // parameters for query
	RequestPerPage *int32     `protobuf:"varint,5,opt,name=request_per_page,json=requestPerPage,proto3,oneof" json:"request_per_page,omitempty"` // parameters for query
}

func (x *CommunityGetRequest) Reset() {
	*x = CommunityGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunityGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunityGetRequest) ProtoMessage() {}

func (x *CommunityGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommunityGetRequest.ProtoReflect.Descriptor instead.
func (*CommunityGetRequest) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{2}
}

func (x *CommunityGetRequest) GetRequestType() NUMBER_NUM {
	if x != nil {
		return x.RequestType
	}
	return NUMBER_NUM_one
}

func (x *CommunityGetRequest) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *CommunityGetRequest) GetQuery() string {
	if x != nil && x.Query != nil {
		return *x.Query
	}
	return ""
}

func (x *CommunityGetRequest) GetPageNumber() string {
	if x != nil && x.PageNumber != nil {
		return *x.PageNumber
	}
	return ""
}

func (x *CommunityGetRequest) GetRequestPerPage() int32 {
	if x != nil && x.RequestPerPage != nil {
		return *x.RequestPerPage
	}
	return 0
}

// get community data : Response
type CommunityMetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      *STATUS          `protobuf:"varint,1,opt,name=status,proto3,enum=community.STATUS,oneof" json:"status,omitempty"`
	Id          string           `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`                         // unique id of community
	Name        string           `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                     // name of community
	Description *string          `protobuf:"bytes,4,opt,name=Description,proto3,oneof" json:"Description,omitempty"` // description of community
	Banner      *string          `protobuf:"bytes,5,opt,name=banner,proto3,oneof" json:"banner,omitempty"`           // banner source in server
	Spaces      []*SpacesCompact `protobuf:"bytes,6,rep,name=spaces,proto3" json:"spaces,omitempty"`                 // spaces id
}

func (x *CommunityMetaData) Reset() {
	*x = CommunityMetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_community_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommunityMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommunityMetaData) ProtoMessage() {}

func (x *CommunityMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_community_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommunityMetaData.ProtoReflect.Descriptor instead.
func (*CommunityMetaData) Descriptor() ([]byte, []int) {
	return file_community_proto_rawDescGZIP(), []int{3}
}

func (x *CommunityMetaData) GetStatus() STATUS {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return STATUS_ERROR
}

func (x *CommunityMetaData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CommunityMetaData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CommunityMetaData) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *CommunityMetaData) GetBanner() string {
	if x != nil && x.Banner != nil {
		return *x.Banner
	}
	return ""
}

func (x *CommunityMetaData) GetSpaces() []*SpacesCompact {
	if x != nil {
		return x.Spaces
	}
	return nil
}

var File_community_proto protoreflect.FileDescriptor

var file_community_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x22, 0x33, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x2e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x1f, 0x0a, 0x0d, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x89, 0x02, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x0b, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x4e, 0x55, 0x4d, 0x42,
	0x45, 0x52, 0x5f, 0x4e, 0x55, 0x4d, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x10, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x65,
	0x72, 0x50, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x22, 0x83,
	0x02, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x2e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x88, 0x01, 0x01, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x02, 0x52, 0x06, 0x62, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x06,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x70, 0x61, 0x63, 0x65, 0x73, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x63, 0x74, 0x52, 0x06, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x62, 0x61,
	0x6e, 0x6e, 0x65, 0x72, 0x2a, 0x31, 0x0a, 0x08, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x53,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x4e, 0x54, 0x33, 0x32, 0x5f, 0x44, 0x45,
	0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x01, 0x2a, 0x29, 0x0a, 0x0a, 0x4e, 0x55, 0x4d, 0x42, 0x45,
	0x52, 0x5f, 0x4e, 0x55, 0x4d, 0x12, 0x07, 0x0a, 0x03, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x74, 0x77, 0x6f, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x74, 0x68, 0x72, 0x65, 0x65,
	0x10, 0x02, 0x2a, 0x3c, 0x0a, 0x06, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x5f,
	0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x10, 0x02, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x03,
	0x32, 0xb9, 0x02, 0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74,
	0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69,
	0x74, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x4d, 0x65,
	0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x0c, 0x4e, 0x65,
	0x77, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x08, 0x41, 0x64, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x1c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x4d,
	0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e,
	0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x42, 0x37, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x62, 0x68, 0x69, 0x72,
	0x61, 0x6a, 0x72, 0x61, 0x6e, 0x6a, 0x61, 0x6e, 0x2f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_community_proto_rawDescOnce sync.Once
	file_community_proto_rawDescData = file_community_proto_rawDesc
)

func file_community_proto_rawDescGZIP() []byte {
	file_community_proto_rawDescOnce.Do(func() {
		file_community_proto_rawDescData = protoimpl.X.CompressGZIP(file_community_proto_rawDescData)
	})
	return file_community_proto_rawDescData
}

var file_community_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_community_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_community_proto_goTypes = []interface{}{
	(DEFAULTS)(0),               // 0: community.DEFAULTS
	(NUMBER_NUM)(0),             // 1: community.NUMBER_NUM
	(STATUS)(0),                 // 2: community.STATUS
	(*Status)(nil),              // 3: community.Status
	(*SpacesCompact)(nil),       // 4: community.SpacesCompact
	(*CommunityGetRequest)(nil), // 5: community.CommunityGetRequest
	(*CommunityMetaData)(nil),   // 6: community.CommunityMetaData
}
var file_community_proto_depIdxs = []int32{
	2, // 0: community.Status.status:type_name -> community.STATUS
	1, // 1: community.CommunityGetRequest.requestType:type_name -> community.NUMBER_NUM
	2, // 2: community.CommunityMetaData.status:type_name -> community.STATUS
	4, // 3: community.CommunityMetaData.spaces:type_name -> community.SpacesCompact
	5, // 4: community.CommunityService.GetCommunity:input_type -> community.CommunityGetRequest
	5, // 5: community.CommunityService.CommunitySearch:input_type -> community.CommunityGetRequest
	6, // 6: community.CommunityService.NewCommunity:input_type -> community.CommunityMetaData
	6, // 7: community.CommunityService.AddSpace:input_type -> community.CommunityMetaData
	6, // 8: community.CommunityService.GetCommunity:output_type -> community.CommunityMetaData
	6, // 9: community.CommunityService.CommunitySearch:output_type -> community.CommunityMetaData
	3, // 10: community.CommunityService.NewCommunity:output_type -> community.Status
	3, // 11: community.CommunityService.AddSpace:output_type -> community.Status
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_community_proto_init() }
func file_community_proto_init() {
	if File_community_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_community_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_community_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpacesCompact); i {
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
		file_community_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunityGetRequest); i {
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
		file_community_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommunityMetaData); i {
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
	file_community_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_community_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_community_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_community_proto_goTypes,
		DependencyIndexes: file_community_proto_depIdxs,
		EnumInfos:         file_community_proto_enumTypes,
		MessageInfos:      file_community_proto_msgTypes,
	}.Build()
	File_community_proto = out.File
	file_community_proto_rawDesc = nil
	file_community_proto_goTypes = nil
	file_community_proto_depIdxs = nil
}
