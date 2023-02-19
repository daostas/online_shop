// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: client_service.proto

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

type ClientRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *ClientRes) Reset() {
	*x = ClientRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRes) ProtoMessage() {}

func (x *ClientRes) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRes.ProtoReflect.Descriptor instead.
func (*ClientRes) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{0}
}

func (x *ClientRes) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type UpdateClientInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Number  string `protobuf:"bytes,3,opt,name=number,proto3" json:"number,omitempty"`
	Email   string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Dob     string `protobuf:"bytes,5,opt,name=dob,proto3" json:"dob,omitempty"`
	Address string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *UpdateClientInfoReq) Reset() {
	*x = UpdateClientInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClientInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClientInfoReq) ProtoMessage() {}

func (x *UpdateClientInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClientInfoReq.ProtoReflect.Descriptor instead.
func (*UpdateClientInfoReq) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateClientInfoReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateClientInfoReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateClientInfoReq) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *UpdateClientInfoReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateClientInfoReq) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *UpdateClientInfoReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type UpdateClientPassReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Pass  string `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`
	Pass1 string `protobuf:"bytes,3,opt,name=pass1,proto3" json:"pass1,omitempty"`
	Pass2 string `protobuf:"bytes,4,opt,name=pass2,proto3" json:"pass2,omitempty"`
}

func (x *UpdateClientPassReq) Reset() {
	*x = UpdateClientPassReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClientPassReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClientPassReq) ProtoMessage() {}

func (x *UpdateClientPassReq) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClientPassReq.ProtoReflect.Descriptor instead.
func (*UpdateClientPassReq) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateClientPassReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateClientPassReq) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *UpdateClientPassReq) GetPass1() string {
	if x != nil {
		return x.Pass1
	}
	return ""
}

func (x *UpdateClientPassReq) GetPass2() string {
	if x != nil {
		return x.Pass2
	}
	return ""
}

type DeleteClientReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
}

func (x *DeleteClientReq) Reset() {
	*x = DeleteClientReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClientReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClientReq) ProtoMessage() {}

func (x *DeleteClientReq) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClientReq.ProtoReflect.Descriptor instead.
func (*DeleteClientReq) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteClientReq) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type AddToBasketReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId int32 `protobuf:"varint,1,opt,name=Client_id,json=ClientId,proto3" json:"Client_id,omitempty"`
	ProdId   int32 `protobuf:"varint,2,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
}

func (x *AddToBasketReq) Reset() {
	*x = AddToBasketReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddToBasketReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddToBasketReq) ProtoMessage() {}

func (x *AddToBasketReq) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddToBasketReq.ProtoReflect.Descriptor instead.
func (*AddToBasketReq) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{4}
}

func (x *AddToBasketReq) GetClientId() int32 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

func (x *AddToBasketReq) GetProdId() int32 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

type GetGroupsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId    int32 `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	LanguageId int32 `protobuf:"varint,2,opt,name=language_id,json=languageId,proto3" json:"language_id,omitempty"`
}

func (x *GetGroupsReq) Reset() {
	*x = GetGroupsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupsReq) ProtoMessage() {}

func (x *GetGroupsReq) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupsReq.ProtoReflect.Descriptor instead.
func (*GetGroupsReq) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetGroupsReq) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GetGroupsReq) GetLanguageId() int32 {
	if x != nil {
		return x.LanguageId
	}
	return 0
}

type GetGroupsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Groups []*GetGroupsRes_Group `protobuf:"bytes,1,rep,name=groups,proto3" json:"groups,omitempty"`
	Status int32                 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Err    string                `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *GetGroupsRes) Reset() {
	*x = GetGroupsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupsRes) ProtoMessage() {}

func (x *GetGroupsRes) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupsRes.ProtoReflect.Descriptor instead.
func (*GetGroupsRes) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetGroupsRes) GetGroups() []*GetGroupsRes_Group {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *GetGroupsRes) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetGroupsRes) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type GetGroupsRes_Group struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId     int32  `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *GetGroupsRes_Group) Reset() {
	*x = GetGroupsRes_Group{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupsRes_Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupsRes_Group) ProtoMessage() {}

func (x *GetGroupsRes_Group) ProtoReflect() protoreflect.Message {
	mi := &file_client_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupsRes_Group.ProtoReflect.Descriptor instead.
func (*GetGroupsRes_Group) Descriptor() ([]byte, []int) {
	return file_client_service_proto_rawDescGZIP(), []int{6, 0}
}

func (x *GetGroupsRes_Group) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GetGroupsRes_Group) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetGroupsRes_Group) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_client_service_proto protoreflect.FileDescriptor

var file_client_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a,
	0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x93, 0x01, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x6f, 0x62, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x65, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x61, 0x73, 0x73, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61,
	0x73, 0x73, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x73, 0x73, 0x32, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x73, 0x73, 0x32, 0x22, 0x27, 0x0a, 0x0f, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x22, 0x46, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x54, 0x6f, 0x42, 0x61, 0x73, 0x6b, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0xc7, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x65, 0x72, 0x72, 0x1a, 0x5a, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x19, 0x0a,
	0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x32, 0xcd, 0x01, 0x0a, 0x07, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x42, 0x0a, 0x10,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x73, 0x73, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x32, 0x47, 0x0a, 0x0c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73,
	0x12, 0x37, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x13, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_service_proto_rawDescOnce sync.Once
	file_client_service_proto_rawDescData = file_client_service_proto_rawDesc
)

func file_client_service_proto_rawDescGZIP() []byte {
	file_client_service_proto_rawDescOnce.Do(func() {
		file_client_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_service_proto_rawDescData)
	})
	return file_client_service_proto_rawDescData
}

var file_client_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_client_service_proto_goTypes = []interface{}{
	(*ClientRes)(nil),           // 0: proto.ClientRes
	(*UpdateClientInfoReq)(nil), // 1: proto.UpdateClientInfoReq
	(*UpdateClientPassReq)(nil), // 2: proto.UpdateClientPassReq
	(*DeleteClientReq)(nil),     // 3: proto.DeleteClientReq
	(*AddToBasketReq)(nil),      // 4: proto.AddToBasketReq
	(*GetGroupsReq)(nil),        // 5: proto.GetGroupsReq
	(*GetGroupsRes)(nil),        // 6: proto.GetGroupsRes
	(*GetGroupsRes_Group)(nil),  // 7: proto.GetGroupsRes.Group
}
var file_client_service_proto_depIdxs = []int32{
	7, // 0: proto.GetGroupsRes.groups:type_name -> proto.GetGroupsRes.Group
	1, // 1: proto.Clients.UpdateClientInfo:input_type -> proto.UpdateClientInfoReq
	2, // 2: proto.Clients.UpdateClientPass:input_type -> proto.UpdateClientPassReq
	3, // 3: proto.Clients.DeleteClient:input_type -> proto.DeleteClientReq
	5, // 4: proto.ClientGroups.GetGroups:input_type -> proto.GetGroupsReq
	0, // 5: proto.Clients.UpdateClientInfo:output_type -> proto.ClientRes
	0, // 6: proto.Clients.UpdateClientPass:output_type -> proto.ClientRes
	0, // 7: proto.Clients.DeleteClient:output_type -> proto.ClientRes
	6, // 8: proto.ClientGroups.GetGroups:output_type -> proto.GetGroupsRes
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_client_service_proto_init() }
func file_client_service_proto_init() {
	if File_client_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRes); i {
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
		file_client_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClientInfoReq); i {
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
		file_client_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClientPassReq); i {
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
		file_client_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClientReq); i {
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
		file_client_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddToBasketReq); i {
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
		file_client_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupsReq); i {
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
		file_client_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupsRes); i {
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
		file_client_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupsRes_Group); i {
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
			RawDescriptor: file_client_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_client_service_proto_goTypes,
		DependencyIndexes: file_client_service_proto_depIdxs,
		MessageInfos:      file_client_service_proto_msgTypes,
	}.Build()
	File_client_service_proto = out.File
	file_client_service_proto_rawDesc = nil
	file_client_service_proto_goTypes = nil
	file_client_service_proto_depIdxs = nil
}
