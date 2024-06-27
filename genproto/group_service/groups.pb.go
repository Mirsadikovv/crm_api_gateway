// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: groups.proto

package group_service

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

type GroupPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GroupPrimaryKey) Reset() {
	*x = GroupPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupPrimaryKey) ProtoMessage() {}

func (x *GroupPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupPrimaryKey.ProtoReflect.Descriptor instead.
func (*GroupPrimaryKey) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{0}
}

func (x *GroupPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId         string `protobuf:"bytes,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	TeacherId        string `protobuf:"bytes,2,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
	SupportTeacherId string `protobuf:"bytes,3,opt,name=support_teacher_id,json=supportTeacherId,proto3" json:"support_teacher_id,omitempty"`
	GroupName        string `protobuf:"bytes,4,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	GroupLevel       string `protobuf:"bytes,5,opt,name=group_level,json=groupLevel,proto3" json:"group_level,omitempty"`
	StartedAt        string `protobuf:"bytes,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt       string `protobuf:"bytes,7,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
}

func (x *CreateGroup) Reset() {
	*x = CreateGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroup) ProtoMessage() {}

func (x *CreateGroup) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroup.ProtoReflect.Descriptor instead.
func (*CreateGroup) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGroup) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *CreateGroup) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

func (x *CreateGroup) GetSupportTeacherId() string {
	if x != nil {
		return x.SupportTeacherId
	}
	return ""
}

func (x *CreateGroup) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *CreateGroup) GetGroupLevel() string {
	if x != nil {
		return x.GroupLevel
	}
	return ""
}

func (x *CreateGroup) GetStartedAt() string {
	if x != nil {
		return x.StartedAt
	}
	return ""
}

func (x *CreateGroup) GetFinishedAt() string {
	if x != nil {
		return x.FinishedAt
	}
	return ""
}

type GetGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId         string `protobuf:"bytes,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	TeacherId        string `protobuf:"bytes,2,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
	SupportTeacherId string `protobuf:"bytes,3,opt,name=support_teacher_id,json=supportTeacherId,proto3" json:"support_teacher_id,omitempty"`
	GroupName        string `protobuf:"bytes,4,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	GroupLevel       string `protobuf:"bytes,5,opt,name=group_level,json=groupLevel,proto3" json:"group_level,omitempty"`
	StartedAt        string `protobuf:"bytes,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt       string `protobuf:"bytes,7,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	CreatedAt        string `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        string `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Id               string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGroup) Reset() {
	*x = GetGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroup) ProtoMessage() {}

func (x *GetGroup) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroup.ProtoReflect.Descriptor instead.
func (*GetGroup) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{2}
}

func (x *GetGroup) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *GetGroup) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

func (x *GetGroup) GetSupportTeacherId() string {
	if x != nil {
		return x.SupportTeacherId
	}
	return ""
}

func (x *GetGroup) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *GetGroup) GetGroupLevel() string {
	if x != nil {
		return x.GroupLevel
	}
	return ""
}

func (x *GetGroup) GetStartedAt() string {
	if x != nil {
		return x.StartedAt
	}
	return ""
}

func (x *GetGroup) GetFinishedAt() string {
	if x != nil {
		return x.FinishedAt
	}
	return ""
}

func (x *GetGroup) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetGroup) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *GetGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdateGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BranchId         string `protobuf:"bytes,1,opt,name=branch_id,json=branchId,proto3" json:"branch_id,omitempty"`
	TeacherId        string `protobuf:"bytes,2,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
	SupportTeacherId string `protobuf:"bytes,3,opt,name=support_teacher_id,json=supportTeacherId,proto3" json:"support_teacher_id,omitempty"`
	GroupName        string `protobuf:"bytes,4,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	GroupLevel       string `protobuf:"bytes,5,opt,name=group_level,json=groupLevel,proto3" json:"group_level,omitempty"`
	StartedAt        string `protobuf:"bytes,6,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt       string `protobuf:"bytes,7,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	Id               string `protobuf:"bytes,8,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateGroup) Reset() {
	*x = UpdateGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroup) ProtoMessage() {}

func (x *UpdateGroup) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroup.ProtoReflect.Descriptor instead.
func (*UpdateGroup) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateGroup) GetBranchId() string {
	if x != nil {
		return x.BranchId
	}
	return ""
}

func (x *UpdateGroup) GetTeacherId() string {
	if x != nil {
		return x.TeacherId
	}
	return ""
}

func (x *UpdateGroup) GetSupportTeacherId() string {
	if x != nil {
		return x.SupportTeacherId
	}
	return ""
}

func (x *UpdateGroup) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *UpdateGroup) GetGroupLevel() string {
	if x != nil {
		return x.GroupLevel
	}
	return ""
}

func (x *UpdateGroup) GetStartedAt() string {
	if x != nil {
		return x.StartedAt
	}
	return ""
}

func (x *UpdateGroup) GetFinishedAt() string {
	if x != nil {
		return x.FinishedAt
	}
	return ""
}

func (x *UpdateGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetListGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListGroupRequest) Reset() {
	*x = GetListGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListGroupRequest) ProtoMessage() {}

func (x *GetListGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListGroupRequest.ProtoReflect.Descriptor instead.
func (*GetListGroupRequest) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{4}
}

func (x *GetListGroupRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListGroupRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListGroupRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count  int64       `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	Groups []*GetGroup `protobuf:"bytes,2,rep,name=Groups,proto3" json:"Groups,omitempty"`
}

func (x *GetListGroupResponse) Reset() {
	*x = GetListGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_groups_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListGroupResponse) ProtoMessage() {}

func (x *GetListGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_groups_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListGroupResponse.ProtoReflect.Descriptor instead.
func (*GetListGroupResponse) Descriptor() ([]byte, []int) {
	return file_groups_proto_rawDescGZIP(), []int{5}
}

func (x *GetListGroupResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListGroupResponse) GetGroups() []*GetGroup {
	if x != nil {
		return x.Groups
	}
	return nil
}

var File_groups_proto protoreflect.FileDescriptor

var file_groups_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a,
	0x0f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xf7, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc2, 0x02, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74,
	0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x87, 0x02, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x1b, 0x0a, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x12, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5b, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x60, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x52, 0x06, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x32, 0x8b, 0x03, 0x0a, 0x0c, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00,
	0x12, 0x4a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x21, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1a,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x67, 0x6f, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_groups_proto_rawDescOnce sync.Once
	file_groups_proto_rawDescData = file_groups_proto_rawDesc
)

func file_groups_proto_rawDescGZIP() []byte {
	file_groups_proto_rawDescOnce.Do(func() {
		file_groups_proto_rawDescData = protoimpl.X.CompressGZIP(file_groups_proto_rawDescData)
	})
	return file_groups_proto_rawDescData
}

var file_groups_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_groups_proto_goTypes = []interface{}{
	(*GroupPrimaryKey)(nil),      // 0: group_service_go.GroupPrimaryKey
	(*CreateGroup)(nil),          // 1: group_service_go.CreateGroup
	(*GetGroup)(nil),             // 2: group_service_go.GetGroup
	(*UpdateGroup)(nil),          // 3: group_service_go.UpdateGroup
	(*GetListGroupRequest)(nil),  // 4: group_service_go.GetListGroupRequest
	(*GetListGroupResponse)(nil), // 5: group_service_go.GetListGroupResponse
	(*empty.Empty)(nil),          // 6: google.protobuf.Empty
}
var file_groups_proto_depIdxs = []int32{
	2, // 0: group_service_go.GetListGroupResponse.Groups:type_name -> group_service_go.GetGroup
	1, // 1: group_service_go.GroupService.Create:input_type -> group_service_go.CreateGroup
	0, // 2: group_service_go.GroupService.GetByID:input_type -> group_service_go.GroupPrimaryKey
	4, // 3: group_service_go.GroupService.GetList:input_type -> group_service_go.GetListGroupRequest
	3, // 4: group_service_go.GroupService.Update:input_type -> group_service_go.UpdateGroup
	0, // 5: group_service_go.GroupService.Delete:input_type -> group_service_go.GroupPrimaryKey
	2, // 6: group_service_go.GroupService.Create:output_type -> group_service_go.GetGroup
	2, // 7: group_service_go.GroupService.GetByID:output_type -> group_service_go.GetGroup
	5, // 8: group_service_go.GroupService.GetList:output_type -> group_service_go.GetListGroupResponse
	2, // 9: group_service_go.GroupService.Update:output_type -> group_service_go.GetGroup
	6, // 10: group_service_go.GroupService.Delete:output_type -> google.protobuf.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_groups_proto_init() }
func file_groups_proto_init() {
	if File_groups_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_groups_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupPrimaryKey); i {
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
		file_groups_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroup); i {
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
		file_groups_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroup); i {
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
		file_groups_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroup); i {
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
		file_groups_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListGroupRequest); i {
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
		file_groups_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListGroupResponse); i {
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
			RawDescriptor: file_groups_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_groups_proto_goTypes,
		DependencyIndexes: file_groups_proto_depIdxs,
		MessageInfos:      file_groups_proto_msgTypes,
	}.Build()
	File_groups_proto = out.File
	file_groups_proto_rawDesc = nil
	file_groups_proto_goTypes = nil
	file_groups_proto_depIdxs = nil
}
