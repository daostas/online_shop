// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: setting_service.proto

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

type EmptySettReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptySettReq) Reset() {
	*x = EmptySettReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptySettReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptySettReq) ProtoMessage() {}

func (x *EmptySettReq) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptySettReq.ProtoReflect.Descriptor instead.
func (*EmptySettReq) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{0}
}

type SettRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err string `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *SettRes) Reset() {
	*x = SettRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettRes) ProtoMessage() {}

func (x *SettRes) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettRes.ProtoReflect.Descriptor instead.
func (*SettRes) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{1}
}

func (x *SettRes) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type Language struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LangId    int32  `protobuf:"varint,1,opt,name=lang_id,json=langId,proto3" json:"lang_id,omitempty"`
	Code      string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Image     string `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Locale    string `protobuf:"bytes,4,opt,name=locale,proto3" json:"locale,omitempty"`
	LangName  string `protobuf:"bytes,5,opt,name=lang_name,json=langName,proto3" json:"lang_name,omitempty"`
	SortOrder int32  `protobuf:"varint,6,opt,name=sort_order,json=sortOrder,proto3" json:"sort_order,omitempty"`
}

func (x *Language) Reset() {
	*x = Language{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Language) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Language) ProtoMessage() {}

func (x *Language) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Language.ProtoReflect.Descriptor instead.
func (*Language) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{2}
}

func (x *Language) GetLangId() int32 {
	if x != nil {
		return x.LangId
	}
	return 0
}

func (x *Language) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Language) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Language) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Language) GetLangName() string {
	if x != nil {
		return x.LangName
	}
	return ""
}

func (x *Language) GetSortOrder() int32 {
	if x != nil {
		return x.SortOrder
	}
	return 0
}

type SetDefaultLanguageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LangId int32 `protobuf:"varint,1,opt,name=lang_id,json=langId,proto3" json:"lang_id,omitempty"`
}

func (x *SetDefaultLanguageReq) Reset() {
	*x = SetDefaultLanguageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDefaultLanguageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDefaultLanguageReq) ProtoMessage() {}

func (x *SetDefaultLanguageReq) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDefaultLanguageReq.ProtoReflect.Descriptor instead.
func (*SetDefaultLanguageReq) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{3}
}

func (x *SetDefaultLanguageReq) GetLangId() int32 {
	if x != nil {
		return x.LangId
	}
	return 0
}

type NewLangReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language *Language `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *NewLangReq) Reset() {
	*x = NewLangReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewLangReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewLangReq) ProtoMessage() {}

func (x *NewLangReq) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewLangReq.ProtoReflect.Descriptor instead.
func (*NewLangReq) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{4}
}

func (x *NewLangReq) GetLanguage() *Language {
	if x != nil {
		return x.Language
	}
	return nil
}

type GetListOfLanguagesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Languages []*Language `protobuf:"bytes,1,rep,name=languages,proto3" json:"languages,omitempty"`
	Err       string      `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *GetListOfLanguagesRes) Reset() {
	*x = GetListOfLanguagesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListOfLanguagesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListOfLanguagesRes) ProtoMessage() {}

func (x *GetListOfLanguagesRes) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListOfLanguagesRes.ProtoReflect.Descriptor instead.
func (*GetListOfLanguagesRes) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetListOfLanguagesRes) GetLanguages() []*Language {
	if x != nil {
		return x.Languages
	}
	return nil
}

func (x *GetListOfLanguagesRes) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type ChangeLanguageStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LangId int32 `protobuf:"varint,1,opt,name=lang_id,json=langId,proto3" json:"lang_id,omitempty"`
}

func (x *ChangeLanguageStatusReq) Reset() {
	*x = ChangeLanguageStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setting_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeLanguageStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeLanguageStatusReq) ProtoMessage() {}

func (x *ChangeLanguageStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file_setting_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeLanguageStatusReq.ProtoReflect.Descriptor instead.
func (*ChangeLanguageStatusReq) Descriptor() ([]byte, []int) {
	return file_setting_service_proto_rawDescGZIP(), []int{6}
}

func (x *ChangeLanguageStatusReq) GetLangId() int32 {
	if x != nil {
		return x.LangId
	}
	return 0
}

var File_setting_service_proto protoreflect.FileDescriptor

var file_setting_service_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0e,
	0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x52, 0x65, 0x71, 0x22, 0x1b,
	0x0a, 0x07, 0x53, 0x65, 0x74, 0x74, 0x52, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0xa1, 0x01, 0x0a, 0x08,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x6e, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x61, 0x6e, 0x67, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63,
	0x61, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22,
	0x30, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x6e, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c, 0x61, 0x6e, 0x67, 0x49,
	0x64, 0x22, 0x39, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x2b, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x58, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x32, 0x0a, 0x17, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6c, 0x61, 0x6e, 0x67, 0x49, 0x64, 0x32, 0xd8, 0x02, 0x0a, 0x0e, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a,
	0x12, 0x53, 0x65, 0x74, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x10, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x65, 0x77, 0x4c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4e, 0x65, 0x77, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0b,
	0x4e, 0x65, 0x77, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x4c, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x49, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x14, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x74,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_setting_service_proto_rawDescOnce sync.Once
	file_setting_service_proto_rawDescData = file_setting_service_proto_rawDesc
)

func file_setting_service_proto_rawDescGZIP() []byte {
	file_setting_service_proto_rawDescOnce.Do(func() {
		file_setting_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_setting_service_proto_rawDescData)
	})
	return file_setting_service_proto_rawDescData
}

var file_setting_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_setting_service_proto_goTypes = []interface{}{
	(*EmptySettReq)(nil),            // 0: proto.EmptySettReq
	(*SettRes)(nil),                 // 1: proto.SettRes
	(*Language)(nil),                // 2: proto.Language
	(*SetDefaultLanguageReq)(nil),   // 3: proto.SetDefaultLanguageReq
	(*NewLangReq)(nil),              // 4: proto.NewLangReq
	(*GetListOfLanguagesRes)(nil),   // 5: proto.GetListOfLanguagesRes
	(*ChangeLanguageStatusReq)(nil), // 6: proto.ChangeLanguageStatusReq
}
var file_setting_service_proto_depIdxs = []int32{
	2, // 0: proto.NewLangReq.language:type_name -> proto.Language
	2, // 1: proto.GetListOfLanguagesRes.languages:type_name -> proto.Language
	3, // 2: proto.SettingService.SetDefaultLanguage:input_type -> proto.SetDefaultLanguageReq
	4, // 3: proto.SettingService.FirstNewLanguage:input_type -> proto.NewLangReq
	4, // 4: proto.SettingService.NewLanguage:input_type -> proto.NewLangReq
	0, // 5: proto.SettingService.GetListOfLanguages:input_type -> proto.EmptySettReq
	6, // 6: proto.SettingService.ChangeLanguageStatus:input_type -> proto.ChangeLanguageStatusReq
	1, // 7: proto.SettingService.SetDefaultLanguage:output_type -> proto.SettRes
	1, // 8: proto.SettingService.FirstNewLanguage:output_type -> proto.SettRes
	1, // 9: proto.SettingService.NewLanguage:output_type -> proto.SettRes
	5, // 10: proto.SettingService.GetListOfLanguages:output_type -> proto.GetListOfLanguagesRes
	1, // 11: proto.SettingService.ChangeLanguageStatus:output_type -> proto.SettRes
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_setting_service_proto_init() }
func file_setting_service_proto_init() {
	if File_setting_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_setting_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptySettReq); i {
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
		file_setting_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettRes); i {
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
		file_setting_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Language); i {
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
		file_setting_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDefaultLanguageReq); i {
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
		file_setting_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewLangReq); i {
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
		file_setting_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListOfLanguagesRes); i {
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
		file_setting_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeLanguageStatusReq); i {
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
			RawDescriptor: file_setting_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_setting_service_proto_goTypes,
		DependencyIndexes: file_setting_service_proto_depIdxs,
		MessageInfos:      file_setting_service_proto_msgTypes,
	}.Build()
	File_setting_service_proto = out.File
	file_setting_service_proto_rawDesc = nil
	file_setting_service_proto_goTypes = nil
	file_setting_service_proto_depIdxs = nil
}
