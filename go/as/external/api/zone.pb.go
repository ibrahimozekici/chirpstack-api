// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: as/external/api/zone.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Zone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Zone ID
	ZoneId int64 `protobuf:"varint,1,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// Name of Zone
	ZoneName string `protobuf:"bytes,2,opt,name=zone_name,json=zoneName,proto3" json:"zone_name,omitempty"`
	// Organization ID
	OrgId int64 `protobuf:"varint,3,opt,name=org_id,json=orgID,proto3" json:"org_id,omitempty"`
	// Devices
	Devices []string `protobuf:"bytes,4,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *Zone) Reset() {
	*x = Zone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_zone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Zone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Zone) ProtoMessage() {}

func (x *Zone) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_zone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Zone.ProtoReflect.Descriptor instead.
func (*Zone) Descriptor() ([]byte, []int) {
	return file_as_external_api_zone_proto_rawDescGZIP(), []int{0}
}

func (x *Zone) GetZoneId() int64 {
	if x != nil {
		return x.ZoneId
	}
	return 0
}

func (x *Zone) GetZoneName() string {
	if x != nil {
		return x.ZoneName
	}
	return ""
}

func (x *Zone) GetOrgId() int64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *Zone) GetDevices() []string {
	if x != nil {
		return x.Devices
	}
	return nil
}

type CreateZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Zone object to create
	Zone *Zone `protobuf:"bytes,1,opt,name=zone,proto3" json:"zone,omitempty"`
}

func (x *CreateZoneRequest) Reset() {
	*x = CreateZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_zone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateZoneRequest) ProtoMessage() {}

func (x *CreateZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_zone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateZoneRequest.ProtoReflect.Descriptor instead.
func (*CreateZoneRequest) Descriptor() ([]byte, []int) {
	return file_as_external_api_zone_proto_rawDescGZIP(), []int{1}
}

func (x *CreateZoneRequest) GetZone() *Zone {
	if x != nil {
		return x.Zone
	}
	return nil
}

type GetZoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Zone *Zone `protobuf:"bytes,1,opt,name=zone,proto3" json:"zone,omitempty"`
}

func (x *GetZoneResponse) Reset() {
	*x = GetZoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_zone_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetZoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetZoneResponse) ProtoMessage() {}

func (x *GetZoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_zone_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetZoneResponse.ProtoReflect.Descriptor instead.
func (*GetZoneResponse) Descriptor() ([]byte, []int) {
	return file_as_external_api_zone_proto_rawDescGZIP(), []int{2}
}

func (x *GetZoneResponse) GetZone() *Zone {
	if x != nil {
		return x.Zone
	}
	return nil
}

type GetZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneId int64 `protobuf:"varint,1,opt,name=zone_id,json=zoneID,proto3" json:"zone_id,omitempty"`
}

func (x *GetZoneRequest) Reset() {
	*x = GetZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_zone_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetZoneRequest) ProtoMessage() {}

func (x *GetZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_zone_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetZoneRequest.ProtoReflect.Descriptor instead.
func (*GetZoneRequest) Descriptor() ([]byte, []int) {
	return file_as_external_api_zone_proto_rawDescGZIP(), []int{3}
}

func (x *GetZoneRequest) GetZoneId() int64 {
	if x != nil {
		return x.ZoneId
	}
	return 0
}

type ListZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Max number of devices to return in the result-set.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Organization ID filter
	OrganizationId int64 `protobuf:"varint,3,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
}

func (x *ListZoneRequest) Reset() {
	*x = ListZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_zone_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListZoneRequest) ProtoMessage() {}

func (x *ListZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_zone_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListZoneRequest.ProtoReflect.Descriptor instead.
func (*ListZoneRequest) Descriptor() ([]byte, []int) {
	return file_as_external_api_zone_proto_rawDescGZIP(), []int{4}
}

func (x *ListZoneRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListZoneRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListZoneRequest) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

type ListZoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Zones []*Zone `protobuf:"bytes,1,rep,name=zones,proto3" json:"zones,omitempty"`
}

func (x *ListZoneResponse) Reset() {
	*x = ListZoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_zone_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListZoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListZoneResponse) ProtoMessage() {}

func (x *ListZoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_zone_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListZoneResponse.ProtoReflect.Descriptor instead.
func (*ListZoneResponse) Descriptor() ([]byte, []int) {
	return file_as_external_api_zone_proto_rawDescGZIP(), []int{5}
}

func (x *ListZoneResponse) GetZones() []*Zone {
	if x != nil {
		return x.Zones
	}
	return nil
}

type DeleteZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneId int64 `protobuf:"varint,1,opt,name=zone_id,json=zoneID,proto3" json:"zone_id,omitempty"`
}

func (x *DeleteZoneRequest) Reset() {
	*x = DeleteZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_zone_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteZoneRequest) ProtoMessage() {}

func (x *DeleteZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_zone_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteZoneRequest.ProtoReflect.Descriptor instead.
func (*DeleteZoneRequest) Descriptor() ([]byte, []int) {
	return file_as_external_api_zone_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteZoneRequest) GetZoneId() int64 {
	if x != nil {
		return x.ZoneId
	}
	return 0
}

type UpdateZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Zone   *Zone `protobuf:"bytes,1,opt,name=zone,proto3" json:"zone,omitempty"`
	ZoneId int64 `protobuf:"varint,2,opt,name=zone_id,json=zoneID,proto3" json:"zone_id,omitempty"`
}

func (x *UpdateZoneRequest) Reset() {
	*x = UpdateZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_zone_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateZoneRequest) ProtoMessage() {}

func (x *UpdateZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_zone_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateZoneRequest.ProtoReflect.Descriptor instead.
func (*UpdateZoneRequest) Descriptor() ([]byte, []int) {
	return file_as_external_api_zone_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateZoneRequest) GetZone() *Zone {
	if x != nil {
		return x.Zone
	}
	return nil
}

func (x *UpdateZoneRequest) GetZoneId() int64 {
	if x != nil {
		return x.ZoneId
	}
	return 0
}

type AddUserToZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneId int64 `protobuf:"varint,1,opt,name=zone_id,json=zoneID,proto3" json:"zone_id,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userID,proto3" json:"user_id,omitempty"`
}

func (x *AddUserToZoneRequest) Reset() {
	*x = AddUserToZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_zone_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToZoneRequest) ProtoMessage() {}

func (x *AddUserToZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_zone_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToZoneRequest.ProtoReflect.Descriptor instead.
func (*AddUserToZoneRequest) Descriptor() ([]byte, []int) {
	return file_as_external_api_zone_proto_rawDescGZIP(), []int{8}
}

func (x *AddUserToZoneRequest) GetZoneId() int64 {
	if x != nil {
		return x.ZoneId
	}
	return 0
}

func (x *AddUserToZoneRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type AddUserToZoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserName string  `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	ZoneList []int64 `protobuf:"varint,2,rep,packed,name=zone_list,json=zoneList,proto3" json:"zone_list,omitempty"`
}

func (x *AddUserToZoneResponse) Reset() {
	*x = AddUserToZoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_zone_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToZoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToZoneResponse) ProtoMessage() {}

func (x *AddUserToZoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_zone_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToZoneResponse.ProtoReflect.Descriptor instead.
func (*AddUserToZoneResponse) Descriptor() ([]byte, []int) {
	return file_as_external_api_zone_proto_rawDescGZIP(), []int{9}
}

func (x *AddUserToZoneResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AddUserToZoneResponse) GetZoneList() []int64 {
	if x != nil {
		return x.ZoneList
	}
	return nil
}

var File_as_external_api_zone_proto protoreflect.FileDescriptor

var file_as_external_api_zone_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70,
	0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x04,
	0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x7a, 0x6f, 0x6e, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x32, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1d, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x22,
	0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x7a, 0x6f, 0x6e,
	0x65, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x44, 0x22, 0x68, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x27, 0x0a,
	0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x22, 0x33, 0x0a, 0x10, 0x6c, 0x69, 0x73, 0x74, 0x5a, 0x6f,
	0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x05, 0x7a, 0x6f,
	0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x22, 0x2c, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x44, 0x22, 0x4b, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x44, 0x22, 0x48, 0x0a, 0x14, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x22, 0x51, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x5a, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x7a, 0x6f, 0x6e, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x32, 0x97, 0x04, 0x0a, 0x0b, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x5a,
	0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x3a,
	0x01, 0x2a, 0x12, 0x4e, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x7b, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x47, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12,
	0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x2a, 0x14, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x7b, 0x7a, 0x6f, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x7d, 0x12, 0x57, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x5a,
	0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x1a, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f,
	0x7b, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x0d,
	0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x5a, 0x6f, 0x6e,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x1a, 0x1c, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x73, 0x2f, 0x61, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x2f, 0x7b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_as_external_api_zone_proto_rawDescOnce sync.Once
	file_as_external_api_zone_proto_rawDescData = file_as_external_api_zone_proto_rawDesc
)

func file_as_external_api_zone_proto_rawDescGZIP() []byte {
	file_as_external_api_zone_proto_rawDescOnce.Do(func() {
		file_as_external_api_zone_proto_rawDescData = protoimpl.X.CompressGZIP(file_as_external_api_zone_proto_rawDescData)
	})
	return file_as_external_api_zone_proto_rawDescData
}

var file_as_external_api_zone_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_as_external_api_zone_proto_goTypes = []interface{}{
	(*Zone)(nil),                  // 0: api.Zone
	(*CreateZoneRequest)(nil),     // 1: api.CreateZoneRequest
	(*GetZoneResponse)(nil),       // 2: api.GetZoneResponse
	(*GetZoneRequest)(nil),        // 3: api.GetZoneRequest
	(*ListZoneRequest)(nil),       // 4: api.ListZoneRequest
	(*ListZoneResponse)(nil),      // 5: api.listZoneResponse
	(*DeleteZoneRequest)(nil),     // 6: api.DeleteZoneRequest
	(*UpdateZoneRequest)(nil),     // 7: api.UpdateZoneRequest
	(*AddUserToZoneRequest)(nil),  // 8: api.AddUserToZoneRequest
	(*AddUserToZoneResponse)(nil), // 9: api.AddUserToZoneResponse
	(*empty.Empty)(nil),           // 10: google.protobuf.Empty
}
var file_as_external_api_zone_proto_depIdxs = []int32{
	0,  // 0: api.CreateZoneRequest.zone:type_name -> api.Zone
	0,  // 1: api.GetZoneResponse.zone:type_name -> api.Zone
	0,  // 2: api.listZoneResponse.zones:type_name -> api.Zone
	0,  // 3: api.UpdateZoneRequest.zone:type_name -> api.Zone
	1,  // 4: api.ZoneService.Create:input_type -> api.CreateZoneRequest
	3,  // 5: api.ZoneService.Get:input_type -> api.GetZoneRequest
	4,  // 6: api.ZoneService.List:input_type -> api.ListZoneRequest
	6,  // 7: api.ZoneService.Delete:input_type -> api.DeleteZoneRequest
	7,  // 8: api.ZoneService.Update:input_type -> api.UpdateZoneRequest
	8,  // 9: api.ZoneService.AddUserToZone:input_type -> api.AddUserToZoneRequest
	2,  // 10: api.ZoneService.Create:output_type -> api.GetZoneResponse
	2,  // 11: api.ZoneService.Get:output_type -> api.GetZoneResponse
	5,  // 12: api.ZoneService.List:output_type -> api.listZoneResponse
	10, // 13: api.ZoneService.Delete:output_type -> google.protobuf.Empty
	2,  // 14: api.ZoneService.Update:output_type -> api.GetZoneResponse
	9,  // 15: api.ZoneService.AddUserToZone:output_type -> api.AddUserToZoneResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_as_external_api_zone_proto_init() }
func file_as_external_api_zone_proto_init() {
	if File_as_external_api_zone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_as_external_api_zone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Zone); i {
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
		file_as_external_api_zone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateZoneRequest); i {
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
		file_as_external_api_zone_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetZoneResponse); i {
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
		file_as_external_api_zone_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetZoneRequest); i {
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
		file_as_external_api_zone_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListZoneRequest); i {
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
		file_as_external_api_zone_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListZoneResponse); i {
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
		file_as_external_api_zone_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteZoneRequest); i {
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
		file_as_external_api_zone_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateZoneRequest); i {
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
		file_as_external_api_zone_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToZoneRequest); i {
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
		file_as_external_api_zone_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToZoneResponse); i {
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
			RawDescriptor: file_as_external_api_zone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_as_external_api_zone_proto_goTypes,
		DependencyIndexes: file_as_external_api_zone_proto_depIdxs,
		MessageInfos:      file_as_external_api_zone_proto_msgTypes,
	}.Build()
	File_as_external_api_zone_proto = out.File
	file_as_external_api_zone_proto_rawDesc = nil
	file_as_external_api_zone_proto_goTypes = nil
	file_as_external_api_zone_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ZoneServiceClient is the client API for ZoneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ZoneServiceClient interface {
	// Create creates the given zone
	Create(ctx context.Context, in *CreateZoneRequest, opts ...grpc.CallOption) (*GetZoneResponse, error)
	// Get returns the zone matching the given zone id
	Get(ctx context.Context, in *GetZoneRequest, opts ...grpc.CallOption) (*GetZoneResponse, error)
	// List returns the available zones.
	List(ctx context.Context, in *ListZoneRequest, opts ...grpc.CallOption) (*ListZoneResponse, error)
	// Delete deletes the zones matching the given id.
	Delete(ctx context.Context, in *DeleteZoneRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Update updates the zone matching the given id.
	Update(ctx context.Context, in *UpdateZoneRequest, opts ...grpc.CallOption) (*GetZoneResponse, error)
	AddUserToZone(ctx context.Context, in *AddUserToZoneRequest, opts ...grpc.CallOption) (*AddUserToZoneResponse, error)
}

type zoneServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewZoneServiceClient(cc grpc.ClientConnInterface) ZoneServiceClient {
	return &zoneServiceClient{cc}
}

func (c *zoneServiceClient) Create(ctx context.Context, in *CreateZoneRequest, opts ...grpc.CallOption) (*GetZoneResponse, error) {
	out := new(GetZoneResponse)
	err := c.cc.Invoke(ctx, "/api.ZoneService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneServiceClient) Get(ctx context.Context, in *GetZoneRequest, opts ...grpc.CallOption) (*GetZoneResponse, error) {
	out := new(GetZoneResponse)
	err := c.cc.Invoke(ctx, "/api.ZoneService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneServiceClient) List(ctx context.Context, in *ListZoneRequest, opts ...grpc.CallOption) (*ListZoneResponse, error) {
	out := new(ListZoneResponse)
	err := c.cc.Invoke(ctx, "/api.ZoneService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneServiceClient) Delete(ctx context.Context, in *DeleteZoneRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.ZoneService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneServiceClient) Update(ctx context.Context, in *UpdateZoneRequest, opts ...grpc.CallOption) (*GetZoneResponse, error) {
	out := new(GetZoneResponse)
	err := c.cc.Invoke(ctx, "/api.ZoneService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneServiceClient) AddUserToZone(ctx context.Context, in *AddUserToZoneRequest, opts ...grpc.CallOption) (*AddUserToZoneResponse, error) {
	out := new(AddUserToZoneResponse)
	err := c.cc.Invoke(ctx, "/api.ZoneService/AddUserToZone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZoneServiceServer is the server API for ZoneService service.
type ZoneServiceServer interface {
	// Create creates the given zone
	Create(context.Context, *CreateZoneRequest) (*GetZoneResponse, error)
	// Get returns the zone matching the given zone id
	Get(context.Context, *GetZoneRequest) (*GetZoneResponse, error)
	// List returns the available zones.
	List(context.Context, *ListZoneRequest) (*ListZoneResponse, error)
	// Delete deletes the zones matching the given id.
	Delete(context.Context, *DeleteZoneRequest) (*empty.Empty, error)
	// Update updates the zone matching the given id.
	Update(context.Context, *UpdateZoneRequest) (*GetZoneResponse, error)
	AddUserToZone(context.Context, *AddUserToZoneRequest) (*AddUserToZoneResponse, error)
}

// UnimplementedZoneServiceServer can be embedded to have forward compatible implementations.
type UnimplementedZoneServiceServer struct {
}

func (*UnimplementedZoneServiceServer) Create(context.Context, *CreateZoneRequest) (*GetZoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedZoneServiceServer) Get(context.Context, *GetZoneRequest) (*GetZoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedZoneServiceServer) List(context.Context, *ListZoneRequest) (*ListZoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedZoneServiceServer) Delete(context.Context, *DeleteZoneRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedZoneServiceServer) Update(context.Context, *UpdateZoneRequest) (*GetZoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedZoneServiceServer) AddUserToZone(context.Context, *AddUserToZoneRequest) (*AddUserToZoneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserToZone not implemented")
}

func RegisterZoneServiceServer(s *grpc.Server, srv ZoneServiceServer) {
	s.RegisterService(&_ZoneService_serviceDesc, srv)
}

func _ZoneService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ZoneService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).Create(ctx, req.(*CreateZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ZoneService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).Get(ctx, req.(*GetZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ZoneService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).List(ctx, req.(*ListZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ZoneService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).Delete(ctx, req.(*DeleteZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ZoneService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).Update(ctx, req.(*UpdateZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneService_AddUserToZone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserToZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).AddUserToZone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ZoneService/AddUserToZone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).AddUserToZone(ctx, req.(*AddUserToZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ZoneService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ZoneService",
	HandlerType: (*ZoneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ZoneService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ZoneService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ZoneService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ZoneService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ZoneService_Update_Handler,
		},
		{
			MethodName: "AddUserToZone",
			Handler:    _ZoneService_AddUserToZone_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as/external/api/zone.proto",
}