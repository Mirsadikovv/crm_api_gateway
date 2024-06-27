// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: administrators.proto

package administrator_service

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

type AdministratorPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AdministratorPrimaryKey) Reset() {
	*x = AdministratorPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_administrators_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdministratorPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdministratorPrimaryKey) ProtoMessage() {}

func (x *AdministratorPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_administrators_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdministratorPrimaryKey.ProtoReflect.Descriptor instead.
func (*AdministratorPrimaryKey) Descriptor() ([]byte, []int) {
	return file_administrators_proto_rawDescGZIP(), []int{0}
}

func (x *AdministratorPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateAdministrator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId     string  `protobuf:"bytes,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	Birthday     string  `protobuf:"bytes,2,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender       string  `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Fullname     string  `protobuf:"bytes,4,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email        string  `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone        string  `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	UserPassword string  `protobuf:"bytes,7,opt,name=user_password,json=userPassword,proto3" json:"user_password,omitempty"`
	Salary       float32 `protobuf:"fixed32,8,opt,name=salary,proto3" json:"salary,omitempty"`
	StartWorking string  `protobuf:"bytes,9,opt,name=start_working,json=startWorking,proto3" json:"start_working,omitempty"`
	EndWorking   string  `protobuf:"bytes,10,opt,name=end_working,json=endWorking,proto3" json:"end_working,omitempty"`
}

func (x *CreateAdministrator) Reset() {
	*x = CreateAdministrator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_administrators_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAdministrator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAdministrator) ProtoMessage() {}

func (x *CreateAdministrator) ProtoReflect() protoreflect.Message {
	mi := &file_administrators_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAdministrator.ProtoReflect.Descriptor instead.
func (*CreateAdministrator) Descriptor() ([]byte, []int) {
	return file_administrators_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAdministrator) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *CreateAdministrator) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *CreateAdministrator) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *CreateAdministrator) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *CreateAdministrator) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateAdministrator) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateAdministrator) GetUserPassword() string {
	if x != nil {
		return x.UserPassword
	}
	return ""
}

func (x *CreateAdministrator) GetSalary() float32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *CreateAdministrator) GetStartWorking() string {
	if x != nil {
		return x.StartWorking
	}
	return ""
}

func (x *CreateAdministrator) GetEndWorking() string {
	if x != nil {
		return x.EndWorking
	}
	return ""
}

type GetAdministrator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId     string  `protobuf:"bytes,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	UserLogin    string  `protobuf:"bytes,2,opt,name=user_login,json=userLogin,proto3" json:"user_login,omitempty"`
	Birthday     string  `protobuf:"bytes,3,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender       string  `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	Fullname     string  `protobuf:"bytes,5,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email        string  `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Phone        string  `protobuf:"bytes,7,opt,name=phone,proto3" json:"phone,omitempty"`
	Salary       float32 `protobuf:"fixed32,8,opt,name=salary,proto3" json:"salary,omitempty"`
	StartWorking string  `protobuf:"bytes,9,opt,name=start_working,json=startWorking,proto3" json:"start_working,omitempty"`
	EndWorking   string  `protobuf:"bytes,10,opt,name=end_working,json=endWorking,proto3" json:"end_working,omitempty"`
	CreatedAt    string  `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string  `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Id           string  `protobuf:"bytes,13,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAdministrator) Reset() {
	*x = GetAdministrator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_administrators_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdministrator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdministrator) ProtoMessage() {}

func (x *GetAdministrator) ProtoReflect() protoreflect.Message {
	mi := &file_administrators_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdministrator.ProtoReflect.Descriptor instead.
func (*GetAdministrator) Descriptor() ([]byte, []int) {
	return file_administrators_proto_rawDescGZIP(), []int{2}
}

func (x *GetAdministrator) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *GetAdministrator) GetUserLogin() string {
	if x != nil {
		return x.UserLogin
	}
	return ""
}

func (x *GetAdministrator) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *GetAdministrator) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *GetAdministrator) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *GetAdministrator) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetAdministrator) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *GetAdministrator) GetSalary() float32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *GetAdministrator) GetStartWorking() string {
	if x != nil {
		return x.StartWorking
	}
	return ""
}

func (x *GetAdministrator) GetEndWorking() string {
	if x != nil {
		return x.EndWorking
	}
	return ""
}

func (x *GetAdministrator) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetAdministrator) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *GetAdministrator) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateAdministrator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId     string  `protobuf:"bytes,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	Birthday     string  `protobuf:"bytes,2,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Gender       string  `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Fullname     string  `protobuf:"bytes,4,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Email        string  `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone        string  `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Salary       float32 `protobuf:"fixed32,7,opt,name=salary,proto3" json:"salary,omitempty"`
	StartWorking string  `protobuf:"bytes,8,opt,name=start_working,json=startWorking,proto3" json:"start_working,omitempty"`
	EndWorking   string  `protobuf:"bytes,9,opt,name=end_working,json=endWorking,proto3" json:"end_working,omitempty"`
	Id           string  `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateAdministrator) Reset() {
	*x = UpdateAdministrator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_administrators_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdministrator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdministrator) ProtoMessage() {}

func (x *UpdateAdministrator) ProtoReflect() protoreflect.Message {
	mi := &file_administrators_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdministrator.ProtoReflect.Descriptor instead.
func (*UpdateAdministrator) Descriptor() ([]byte, []int) {
	return file_administrators_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAdministrator) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *UpdateAdministrator) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *UpdateAdministrator) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UpdateAdministrator) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *UpdateAdministrator) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateAdministrator) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateAdministrator) GetSalary() float32 {
	if x != nil {
		return x.Salary
	}
	return 0
}

func (x *UpdateAdministrator) GetStartWorking() string {
	if x != nil {
		return x.StartWorking
	}
	return ""
}

func (x *UpdateAdministrator) GetEndWorking() string {
	if x != nil {
		return x.EndWorking
	}
	return ""
}

func (x *UpdateAdministrator) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetListAdministratorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListAdministratorRequest) Reset() {
	*x = GetListAdministratorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_administrators_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListAdministratorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListAdministratorRequest) ProtoMessage() {}

func (x *GetListAdministratorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_administrators_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListAdministratorRequest.ProtoReflect.Descriptor instead.
func (*GetListAdministratorRequest) Descriptor() ([]byte, []int) {
	return file_administrators_proto_rawDescGZIP(), []int{4}
}

func (x *GetListAdministratorRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListAdministratorRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListAdministratorRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListAdministratorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count          int64               `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	Administrators []*GetAdministrator `protobuf:"bytes,2,rep,name=Administrators,proto3" json:"Administrators,omitempty"`
}

func (x *GetListAdministratorResponse) Reset() {
	*x = GetListAdministratorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_administrators_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListAdministratorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListAdministratorResponse) ProtoMessage() {}

func (x *GetListAdministratorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_administrators_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListAdministratorResponse.ProtoReflect.Descriptor instead.
func (*GetListAdministratorResponse) Descriptor() ([]byte, []int) {
	return file_administrators_proto_rawDescGZIP(), []int{5}
}

func (x *GetListAdministratorResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListAdministratorResponse) GetAdministrators() []*GetAdministrator {
	if x != nil {
		return x.Administrators
	}
	return nil
}

var File_administrators_proto protoreflect.FileDescriptor

var file_administrators_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a,
	0x17, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xb1, 0x02, 0x0a, 0x13, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06,
	0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x65,
	0x6e, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x6e, 0x64, 0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0xf6, 0x02, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c,
	0x61, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72,
	0x79, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x64,
	0x57, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9c, 0x02, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1b, 0x0a,
	0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x73, 0x61, 0x6c, 0x61, 0x72, 0x79, 0x12, 0x23,
	0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x72, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e, 0x64, 0x57, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x63, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x88, 0x01, 0x0a, 0x1c, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x52, 0x0a, 0x0e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x52, 0x0e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x73, 0x32, 0xa3, 0x04, 0x0a, 0x14, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x65, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x67, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x2a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12,
	0x31, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b,
	0x65, 0x79, 0x1a, 0x2a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x00,
	0x12, 0x7a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x36, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67,
	0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x2a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x31, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_administrators_proto_rawDescOnce sync.Once
	file_administrators_proto_rawDescData = file_administrators_proto_rawDesc
)

func file_administrators_proto_rawDescGZIP() []byte {
	file_administrators_proto_rawDescOnce.Do(func() {
		file_administrators_proto_rawDescData = protoimpl.X.CompressGZIP(file_administrators_proto_rawDescData)
	})
	return file_administrators_proto_rawDescData
}

var file_administrators_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_administrators_proto_goTypes = []interface{}{
	(*AdministratorPrimaryKey)(nil),      // 0: administrator_service_go.AdministratorPrimaryKey
	(*CreateAdministrator)(nil),          // 1: administrator_service_go.CreateAdministrator
	(*GetAdministrator)(nil),             // 2: administrator_service_go.GetAdministrator
	(*UpdateAdministrator)(nil),          // 3: administrator_service_go.UpdateAdministrator
	(*GetListAdministratorRequest)(nil),  // 4: administrator_service_go.GetListAdministratorRequest
	(*GetListAdministratorResponse)(nil), // 5: administrator_service_go.GetListAdministratorResponse
	(*empty.Empty)(nil),                  // 6: google.protobuf.Empty
}
var file_administrators_proto_depIdxs = []int32{
	2, // 0: administrator_service_go.GetListAdministratorResponse.Administrators:type_name -> administrator_service_go.GetAdministrator
	1, // 1: administrator_service_go.AdministratorService.Create:input_type -> administrator_service_go.CreateAdministrator
	0, // 2: administrator_service_go.AdministratorService.GetByID:input_type -> administrator_service_go.AdministratorPrimaryKey
	4, // 3: administrator_service_go.AdministratorService.GetList:input_type -> administrator_service_go.GetListAdministratorRequest
	3, // 4: administrator_service_go.AdministratorService.Update:input_type -> administrator_service_go.UpdateAdministrator
	0, // 5: administrator_service_go.AdministratorService.Delete:input_type -> administrator_service_go.AdministratorPrimaryKey
	2, // 6: administrator_service_go.AdministratorService.Create:output_type -> administrator_service_go.GetAdministrator
	2, // 7: administrator_service_go.AdministratorService.GetByID:output_type -> administrator_service_go.GetAdministrator
	5, // 8: administrator_service_go.AdministratorService.GetList:output_type -> administrator_service_go.GetListAdministratorResponse
	2, // 9: administrator_service_go.AdministratorService.Update:output_type -> administrator_service_go.GetAdministrator
	6, // 10: administrator_service_go.AdministratorService.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_administrators_proto_init() }
func file_administrators_proto_init() {
	if File_administrators_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_administrators_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdministratorPrimaryKey); i {
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
		file_administrators_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAdministrator); i {
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
		file_administrators_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAdministrator); i {
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
		file_administrators_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAdministrator); i {
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
		file_administrators_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListAdministratorRequest); i {
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
		file_administrators_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListAdministratorResponse); i {
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
			RawDescriptor: file_administrators_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_administrators_proto_goTypes,
		DependencyIndexes: file_administrators_proto_depIdxs,
		MessageInfos:      file_administrators_proto_msgTypes,
	}.Build()
	File_administrators_proto = out.File
	file_administrators_proto_rawDesc = nil
	file_administrators_proto_goTypes = nil
	file_administrators_proto_depIdxs = nil
}