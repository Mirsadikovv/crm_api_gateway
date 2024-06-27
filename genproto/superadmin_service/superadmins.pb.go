// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: superadmins.proto

package superadmin_service

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type SuperadminPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SuperadminPrimaryKey) Reset() {
	*x = SuperadminPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_superadmins_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuperadminPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuperadminPrimaryKey) ProtoMessage() {}

func (x *SuperadminPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_superadmins_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuperadminPrimaryKey.ProtoReflect.Descriptor instead.
func (*SuperadminPrimaryKey) Descriptor() ([]byte, []int) {
	return file_superadmins_proto_rawDescGZIP(), []int{0}
}

func (x *SuperadminPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateSuperadmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Birthday     string `protobuf:"bytes,1,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender       string `protobuf:"bytes,2,opt,name=gender,proto3" json:"gender,omitempty"`
	Fullname     string `protobuf:"bytes,3,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email        string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone        string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	UserPassword string `protobuf:"bytes,6,opt,name=user_password,json=userPassword,proto3" json:"user_password,omitempty"`
}

func (x *CreateSuperadmin) Reset() {
	*x = CreateSuperadmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_superadmins_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSuperadmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSuperadmin) ProtoMessage() {}

func (x *CreateSuperadmin) ProtoReflect() protoreflect.Message {
	mi := &file_superadmins_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSuperadmin.ProtoReflect.Descriptor instead.
func (*CreateSuperadmin) Descriptor() ([]byte, []int) {
	return file_superadmins_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSuperadmin) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *CreateSuperadmin) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *CreateSuperadmin) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *CreateSuperadmin) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateSuperadmin) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateSuperadmin) GetUserPassword() string {
	if x != nil {
		return x.UserPassword
	}
	return ""
}

type GetSuperadmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserLogin string `protobuf:"bytes,1,opt,name=user_login,json=userLogin,proto3" json:"user_login,omitempty"`
	Birthday  string `protobuf:"bytes,2,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender    string `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Fullname  string `protobuf:"bytes,4,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Id        string `protobuf:"bytes,9,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSuperadmin) Reset() {
	*x = GetSuperadmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_superadmins_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSuperadmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSuperadmin) ProtoMessage() {}

func (x *GetSuperadmin) ProtoReflect() protoreflect.Message {
	mi := &file_superadmins_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSuperadmin.ProtoReflect.Descriptor instead.
func (*GetSuperadmin) Descriptor() ([]byte, []int) {
	return file_superadmins_proto_rawDescGZIP(), []int{2}
}

func (x *GetSuperadmin) GetUserLogin() string {
	if x != nil {
		return x.UserLogin
	}
	return ""
}

func (x *GetSuperadmin) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *GetSuperadmin) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *GetSuperadmin) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *GetSuperadmin) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetSuperadmin) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetSuperadmin) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetSuperadmin) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *GetSuperadmin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateSuperadmin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Birthday string `protobuf:"bytes,1,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender   string `protobuf:"bytes,2,opt,name=gender,proto3" json:"gender,omitempty"`
	Fullname string `protobuf:"bytes,3,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone    string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	Id       string `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateSuperadmin) Reset() {
	*x = UpdateSuperadmin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_superadmins_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSuperadmin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSuperadmin) ProtoMessage() {}

func (x *UpdateSuperadmin) ProtoReflect() protoreflect.Message {
	mi := &file_superadmins_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSuperadmin.ProtoReflect.Descriptor instead.
func (*UpdateSuperadmin) Descriptor() ([]byte, []int) {
	return file_superadmins_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSuperadmin) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *UpdateSuperadmin) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UpdateSuperadmin) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *UpdateSuperadmin) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateSuperadmin) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateSuperadmin) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_superadmins_proto protoreflect.FileDescriptor

var file_superadmins_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x75, 0x70, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x75, 0x70, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x14, 0x53, 0x75, 0x70, 0x65, 0x72,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0xb3, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x65, 0x72, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xf8, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70,
	0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64,
	0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64,
	0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75,
	0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x9e, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x65, 0x72,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c,
	0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c,
	0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x32, 0xfa, 0x02, 0x0a, 0x11, 0x53, 0x75, 0x70, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x27, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x70, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x24, 0x2e, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x22, 0x00, 0x12, 0x5e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x2b, 0x2e,
	0x73, 0x75, 0x70, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x53, 0x75, 0x70, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x24, 0x2e, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x75, 0x70, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x22, 0x00, 0x12, 0x59, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x27, 0x2e, 0x73,
	0x75, 0x70, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x75, 0x70, 0x65, 0x72,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x24, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x75, 0x70, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x22, 0x00, 0x12, 0x4f, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e,
	0x53, 0x75, 0x70, 0x65, 0x72, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1d,
	0x5a, 0x1b, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x75, 0x70, 0x65, 0x72,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_superadmins_proto_rawDescOnce sync.Once
	file_superadmins_proto_rawDescData = file_superadmins_proto_rawDesc
)

func file_superadmins_proto_rawDescGZIP() []byte {
	file_superadmins_proto_rawDescOnce.Do(func() {
		file_superadmins_proto_rawDescData = protoimpl.X.CompressGZIP(file_superadmins_proto_rawDescData)
	})
	return file_superadmins_proto_rawDescData
}

var file_superadmins_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_superadmins_proto_goTypes = []interface{}{
	(*SuperadminPrimaryKey)(nil), // 0: superadmin_service_go.SuperadminPrimaryKey
	(*CreateSuperadmin)(nil),     // 1: superadmin_service_go.CreateSuperadmin
	(*GetSuperadmin)(nil),        // 2: superadmin_service_go.GetSuperadmin
	(*UpdateSuperadmin)(nil),     // 3: superadmin_service_go.UpdateSuperadmin
	(*empty.Empty)(nil),          // 4: google.protobuf.Empty
}
var file_superadmins_proto_depIdxs = []int32{
	1, // 0: superadmin_service_go.SuperadminService.Create:input_type -> superadmin_service_go.CreateSuperadmin
	0, // 1: superadmin_service_go.SuperadminService.GetByID:input_type -> superadmin_service_go.SuperadminPrimaryKey
	3, // 2: superadmin_service_go.SuperadminService.Update:input_type -> superadmin_service_go.UpdateSuperadmin
	0, // 3: superadmin_service_go.SuperadminService.Delete:input_type -> superadmin_service_go.SuperadminPrimaryKey
	2, // 4: superadmin_service_go.SuperadminService.Create:output_type -> superadmin_service_go.GetSuperadmin
	2, // 5: superadmin_service_go.SuperadminService.GetByID:output_type -> superadmin_service_go.GetSuperadmin
	2, // 6: superadmin_service_go.SuperadminService.Update:output_type -> superadmin_service_go.GetSuperadmin
	4, // 7: superadmin_service_go.SuperadminService.Delete:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_superadmins_proto_init() }
func file_superadmins_proto_init() {
	if File_superadmins_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_superadmins_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SuperadminPrimaryKey); i {
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
		file_superadmins_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSuperadmin); i {
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
		file_superadmins_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSuperadmin); i {
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
		file_superadmins_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSuperadmin); i {
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
			RawDescriptor: file_superadmins_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_superadmins_proto_goTypes,
		DependencyIndexes: file_superadmins_proto_depIdxs,
		MessageInfos:      file_superadmins_proto_msgTypes,
	}.Build()
	File_superadmins_proto = out.File
	file_superadmins_proto_rawDesc = nil
	file_superadmins_proto_goTypes = nil
	file_superadmins_proto_depIdxs = nil
}