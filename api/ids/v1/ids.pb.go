// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: api/ids/v1/ids.proto

package v1

import (
	v1 "github.com/MealeyAU/schema-go/common/ids/v1"
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

// Possible status of a verify
type VerifyStatus int32

const (
	// Unknown verification  status
	VerifyStatus_Unknown VerifyStatus = 0
	// Invalid. The id has never existed
	VerifyStatus_Invalid VerifyStatus = 1
	// Active. The id exists and the resource associated with it is not deleted
	VerifyStatus_Active VerifyStatus = 2
	// Deleted. The id exists, but the resource associated with it has been deleted
	VerifyStatus_Deleted VerifyStatus = 3
)

// Enum value maps for VerifyStatus.
var (
	VerifyStatus_name = map[int32]string{
		0: "Unknown",
		1: "Invalid",
		2: "Active",
		3: "Deleted",
	}
	VerifyStatus_value = map[string]int32{
		"Unknown": 0,
		"Invalid": 1,
		"Active":  2,
		"Deleted": 3,
	}
)

func (x VerifyStatus) Enum() *VerifyStatus {
	p := new(VerifyStatus)
	*p = x
	return p
}

func (x VerifyStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VerifyStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_ids_v1_ids_proto_enumTypes[0].Descriptor()
}

func (VerifyStatus) Type() protoreflect.EnumType {
	return &file_api_ids_v1_ids_proto_enumTypes[0]
}

func (x VerifyStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VerifyStatus.Descriptor instead.
func (VerifyStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_ids_v1_ids_proto_rawDescGZIP(), []int{0}
}

// Parameters to be supplied to `Generate`
type GenerateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GenerateRequest) Reset() {
	*x = GenerateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ids_v1_ids_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateRequest) ProtoMessage() {}

func (x *GenerateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ids_v1_ids_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateRequest.ProtoReflect.Descriptor instead.
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return file_api_ids_v1_ids_proto_rawDescGZIP(), []int{0}
}

// Result of the execution of `Generate`
type GenerateResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The newly created id
	Id *v1.Id `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GenerateResult) Reset() {
	*x = GenerateResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ids_v1_ids_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateResult) ProtoMessage() {}

func (x *GenerateResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_ids_v1_ids_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateResult.ProtoReflect.Descriptor instead.
func (*GenerateResult) Descriptor() ([]byte, []int) {
	return file_api_ids_v1_ids_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateResult) GetId() *v1.Id {
	if x != nil {
		return x.Id
	}
	return nil
}

// Parameters to be supplied to `Verify`
type VerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id to verify
	Id *v1.Id `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ids_v1_ids_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ids_v1_ids_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_api_ids_v1_ids_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyRequest) GetId() *v1.Id {
	if x != nil {
		return x.Id
	}
	return nil
}

// Result of the execution of `Verify`
type VerifyResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status of the supplied id
	Status VerifyStatus `protobuf:"varint,1,opt,name=status,proto3,enum=mealeyau.api.ids.v1.VerifyStatus" json:"status,omitempty"`
}

func (x *VerifyResult) Reset() {
	*x = VerifyResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ids_v1_ids_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyResult) ProtoMessage() {}

func (x *VerifyResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_ids_v1_ids_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyResult.ProtoReflect.Descriptor instead.
func (*VerifyResult) Descriptor() ([]byte, []int) {
	return file_api_ids_v1_ids_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyResult) GetStatus() VerifyStatus {
	if x != nil {
		return x.Status
	}
	return VerifyStatus_Unknown
}

// Parameters to be supplied to `Delete`
type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The id to be marked as deleted
	Id *v1.Id `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ids_v1_ids_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_ids_v1_ids_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_api_ids_v1_ids_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetId() *v1.Id {
	if x != nil {
		return x.Id
	}
	return nil
}

// Result of the execution of `Delete`
type DeleteResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResult) Reset() {
	*x = DeleteResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ids_v1_ids_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResult) ProtoMessage() {}

func (x *DeleteResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_ids_v1_ids_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResult.ProtoReflect.Descriptor instead.
func (*DeleteResult) Descriptor() ([]byte, []int) {
	return file_api_ids_v1_ids_proto_rawDescGZIP(), []int{5}
}

var File_api_ids_v1_ids_proto protoreflect.FileDescriptor

var file_api_ids_v1_ids_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6d, 0x65, 0x61, 0x6c, 0x65, 0x79, 0x61, 0x75,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x69, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x65, 0x79, 0x61, 0x75, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x69, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a, 0x0d, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x65, 0x79, 0x61, 0x75, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x69, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x49, 0x0a, 0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x21, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x65, 0x79, 0x61, 0x75, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x69, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3b, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x65, 0x61, 0x6c,
	0x65, 0x79, 0x61, 0x75, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x69, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2a, 0x41, 0x0a, 0x0c, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x10, 0x02, 0x12,
	0x0b, 0x0a, 0x07, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x03, 0x32, 0x84, 0x02, 0x0a,
	0x09, 0x49, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x08, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x24, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x65, 0x79, 0x61,
	0x75, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6d,
	0x65, 0x61, 0x6c, 0x65, 0x79, 0x61, 0x75, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x4f, 0x0a, 0x06, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x12, 0x22, 0x2e, 0x6d, 0x65,
	0x61, 0x6c, 0x65, 0x79, 0x61, 0x75, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x21, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x65, 0x79, 0x61, 0x75, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x4f, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x6d,
	0x65, 0x61, 0x6c, 0x65, 0x79, 0x61, 0x75, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x6d, 0x65, 0x61, 0x6c, 0x65, 0x79, 0x61, 0x75, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x69, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4d, 0x65, 0x61, 0x6c, 0x65, 0x79, 0x41, 0x55, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ids_v1_ids_proto_rawDescOnce sync.Once
	file_api_ids_v1_ids_proto_rawDescData = file_api_ids_v1_ids_proto_rawDesc
)

func file_api_ids_v1_ids_proto_rawDescGZIP() []byte {
	file_api_ids_v1_ids_proto_rawDescOnce.Do(func() {
		file_api_ids_v1_ids_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ids_v1_ids_proto_rawDescData)
	})
	return file_api_ids_v1_ids_proto_rawDescData
}

var file_api_ids_v1_ids_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_ids_v1_ids_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_ids_v1_ids_proto_goTypes = []interface{}{
	(VerifyStatus)(0),       // 0: mealeyau.api.ids.v1.VerifyStatus
	(*GenerateRequest)(nil), // 1: mealeyau.api.ids.v1.GenerateRequest
	(*GenerateResult)(nil),  // 2: mealeyau.api.ids.v1.GenerateResult
	(*VerifyRequest)(nil),   // 3: mealeyau.api.ids.v1.VerifyRequest
	(*VerifyResult)(nil),    // 4: mealeyau.api.ids.v1.VerifyResult
	(*DeleteRequest)(nil),   // 5: mealeyau.api.ids.v1.DeleteRequest
	(*DeleteResult)(nil),    // 6: mealeyau.api.ids.v1.DeleteResult
	(*v1.Id)(nil),           // 7: mealeyau.common.ids.v1.Id
}
var file_api_ids_v1_ids_proto_depIdxs = []int32{
	7, // 0: mealeyau.api.ids.v1.GenerateResult.id:type_name -> mealeyau.common.ids.v1.Id
	7, // 1: mealeyau.api.ids.v1.VerifyRequest.id:type_name -> mealeyau.common.ids.v1.Id
	0, // 2: mealeyau.api.ids.v1.VerifyResult.status:type_name -> mealeyau.api.ids.v1.VerifyStatus
	7, // 3: mealeyau.api.ids.v1.DeleteRequest.id:type_name -> mealeyau.common.ids.v1.Id
	1, // 4: mealeyau.api.ids.v1.IdService.Generate:input_type -> mealeyau.api.ids.v1.GenerateRequest
	3, // 5: mealeyau.api.ids.v1.IdService.Verify:input_type -> mealeyau.api.ids.v1.VerifyRequest
	5, // 6: mealeyau.api.ids.v1.IdService.Delete:input_type -> mealeyau.api.ids.v1.DeleteRequest
	2, // 7: mealeyau.api.ids.v1.IdService.Generate:output_type -> mealeyau.api.ids.v1.GenerateResult
	4, // 8: mealeyau.api.ids.v1.IdService.Verify:output_type -> mealeyau.api.ids.v1.VerifyResult
	6, // 9: mealeyau.api.ids.v1.IdService.Delete:output_type -> mealeyau.api.ids.v1.DeleteResult
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_ids_v1_ids_proto_init() }
func file_api_ids_v1_ids_proto_init() {
	if File_api_ids_v1_ids_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ids_v1_ids_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateRequest); i {
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
		file_api_ids_v1_ids_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateResult); i {
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
		file_api_ids_v1_ids_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRequest); i {
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
		file_api_ids_v1_ids_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyResult); i {
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
		file_api_ids_v1_ids_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_api_ids_v1_ids_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResult); i {
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
			RawDescriptor: file_api_ids_v1_ids_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ids_v1_ids_proto_goTypes,
		DependencyIndexes: file_api_ids_v1_ids_proto_depIdxs,
		EnumInfos:         file_api_ids_v1_ids_proto_enumTypes,
		MessageInfos:      file_api_ids_v1_ids_proto_msgTypes,
	}.Build()
	File_api_ids_v1_ids_proto = out.File
	file_api_ids_v1_ids_proto_rawDesc = nil
	file_api_ids_v1_ids_proto_goTypes = nil
	file_api_ids_v1_ids_proto_depIdxs = nil
}
