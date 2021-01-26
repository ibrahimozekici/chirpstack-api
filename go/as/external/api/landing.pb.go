// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: as/external/api/landing.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type WhoAreYouParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WhoAreYouParams) Reset() {
	*x = WhoAreYouParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_landing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoAreYouParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAreYouParams) ProtoMessage() {}

func (x *WhoAreYouParams) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_landing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAreYouParams.ProtoReflect.Descriptor instead.
func (*WhoAreYouParams) Descriptor() ([]byte, []int) {
	return file_as_external_api_landing_proto_rawDescGZIP(), []int{0}
}

type GetLandingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User ID.
	// Will be set automatically on create.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// E-mail of the user.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	//  Organization List
	// repeated LandingOrganization organizations = 3 [json_name = "organizations"];
	OrganizationList *LandingOrganizationList `protobuf:"bytes,3,opt,name=organizationList,proto3" json:"organizationList,omitempty"`
}

func (x *GetLandingResponse) Reset() {
	*x = GetLandingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_landing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLandingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLandingResponse) ProtoMessage() {}

func (x *GetLandingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_landing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLandingResponse.ProtoReflect.Descriptor instead.
func (*GetLandingResponse) Descriptor() ([]byte, []int) {
	return file_as_external_api_landing_proto_rawDescGZIP(), []int{1}
}

func (x *GetLandingResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetLandingResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetLandingResponse) GetOrganizationList() *LandingOrganizationList {
	if x != nil {
		return x.OrganizationList
	}
	return nil
}

type LandingOrganizationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organizations []*LandingOrganization `protobuf:"bytes,3,rep,name=organizations,proto3" json:"organizations,omitempty"`
}

func (x *LandingOrganizationList) Reset() {
	*x = LandingOrganizationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_landing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LandingOrganizationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LandingOrganizationList) ProtoMessage() {}

func (x *LandingOrganizationList) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_landing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LandingOrganizationList.ProtoReflect.Descriptor instead.
func (*LandingOrganizationList) Descriptor() ([]byte, []int) {
	return file_as_external_api_landing_proto_rawDescGZIP(), []int{2}
}

func (x *LandingOrganizationList) GetOrganizations() []*LandingOrganization {
	if x != nil {
		return x.Organizations
	}
	return nil
}

type LandingOrganization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Organization ID.
	OrganizationId int64 `protobuf:"varint,1,opt,name=organization_id,proto3" json:"organization_id,omitempty"`
	// Organization name.
	OrganizationName string `protobuf:"bytes,2,opt,name=organization_name,proto3" json:"organization_name,omitempty"`
	// Organization display name.
	OrganizationDisplayName string `protobuf:"bytes,3,opt,name=organization_display_name,proto3" json:"organization_display_name,omitempty"`
	// Organization Applications
	Applications []*LandingApplication `protobuf:"bytes,4,rep,name=applications,proto3" json:"applications,omitempty"`
}

func (x *LandingOrganization) Reset() {
	*x = LandingOrganization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_landing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LandingOrganization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LandingOrganization) ProtoMessage() {}

func (x *LandingOrganization) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_landing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LandingOrganization.ProtoReflect.Descriptor instead.
func (*LandingOrganization) Descriptor() ([]byte, []int) {
	return file_as_external_api_landing_proto_rawDescGZIP(), []int{3}
}

func (x *LandingOrganization) GetOrganizationId() int64 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *LandingOrganization) GetOrganizationName() string {
	if x != nil {
		return x.OrganizationName
	}
	return ""
}

func (x *LandingOrganization) GetOrganizationDisplayName() string {
	if x != nil {
		return x.OrganizationDisplayName
	}
	return ""
}

func (x *LandingOrganization) GetApplications() []*LandingApplication {
	if x != nil {
		return x.Applications
	}
	return nil
}

type LandingApplication struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Application ID.
	ApplicationId int64 `protobuf:"varint,1,opt,name=application_id,proto3" json:"application_id,omitempty"`
	// Name of the application.
	ApplicationName string `protobuf:"bytes,2,opt,name=application_name,proto3" json:"application_name,omitempty"`
	// Description of the application.
	ApplicationDescription string `protobuf:"bytes,3,opt,name=application_description,proto3" json:"application_description,omitempty"`
	// service_profile_id
	ServiceProfileId string `protobuf:"bytes,4,opt,name=service_profile_id,proto3" json:"service_profile_id,omitempty"`
	// ID of the organization to which the application belongs.
	Devices []*LandingDevice `protobuf:"bytes,5,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *LandingApplication) Reset() {
	*x = LandingApplication{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_landing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LandingApplication) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LandingApplication) ProtoMessage() {}

func (x *LandingApplication) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_landing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LandingApplication.ProtoReflect.Descriptor instead.
func (*LandingApplication) Descriptor() ([]byte, []int) {
	return file_as_external_api_landing_proto_rawDescGZIP(), []int{4}
}

func (x *LandingApplication) GetApplicationId() int64 {
	if x != nil {
		return x.ApplicationId
	}
	return 0
}

func (x *LandingApplication) GetApplicationName() string {
	if x != nil {
		return x.ApplicationName
	}
	return ""
}

func (x *LandingApplication) GetApplicationDescription() string {
	if x != nil {
		return x.ApplicationDescription
	}
	return ""
}

func (x *LandingApplication) GetServiceProfileId() string {
	if x != nil {
		return x.ServiceProfileId
	}
	return ""
}

func (x *LandingApplication) GetDevices() []*LandingDevice {
	if x != nil {
		return x.Devices
	}
	return nil
}

type LandingDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device EUI (HEX encoded).
	DeviceDevEui string `protobuf:"bytes,1,opt,name=device_dev_eui,proto3" json:"device_dev_eui,omitempty"`
	// device_created_at
	DeviceCreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=device_created_at,proto3" json:"device_created_at,omitempty"`
	// device_updated_at
	DeviceUpdatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=device_updated_at,proto3" json:"device_updated_at,omitempty"`
	// Device-profile ID attached to the device.
	DeviceProfileId string `protobuf:"bytes,4,opt,name=device_profile_id,proto3" json:"device_profile_id,omitempty"`
	// Name of the device.
	DeviceName string `protobuf:"bytes,5,opt,name=device_name,proto3" json:"device_name,omitempty"`
	// Description of the device.
	DeviceDescription string `protobuf:"bytes,6,opt,name=device_description,proto3" json:"device_description,omitempty"`
	// device_last_seen_at
	DeviceLastSeenAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=device_last_seen_at,proto3" json:"device_last_seen_at,omitempty"`
	// data time
	DeviceDataTime int64 `protobuf:"varint,8,opt,name=device_data_time,proto3" json:"device_data_time,omitempty"`
	// device lat value
	DeviceLat float64 `protobuf:"fixed64,9,opt,name=device_lat,proto3" json:"device_lat,omitempty"`
	// device lng value
	DeviceLng float64 `protobuf:"fixed64,10,opt,name=device_lng,proto3" json:"device_lng,omitempty"`
}

func (x *LandingDevice) Reset() {
	*x = LandingDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_landing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LandingDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LandingDevice) ProtoMessage() {}

func (x *LandingDevice) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_landing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LandingDevice.ProtoReflect.Descriptor instead.
func (*LandingDevice) Descriptor() ([]byte, []int) {
	return file_as_external_api_landing_proto_rawDescGZIP(), []int{5}
}

func (x *LandingDevice) GetDeviceDevEui() string {
	if x != nil {
		return x.DeviceDevEui
	}
	return ""
}

func (x *LandingDevice) GetDeviceCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeviceCreatedAt
	}
	return nil
}

func (x *LandingDevice) GetDeviceUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeviceUpdatedAt
	}
	return nil
}

func (x *LandingDevice) GetDeviceProfileId() string {
	if x != nil {
		return x.DeviceProfileId
	}
	return ""
}

func (x *LandingDevice) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *LandingDevice) GetDeviceDescription() string {
	if x != nil {
		return x.DeviceDescription
	}
	return ""
}

func (x *LandingDevice) GetDeviceLastSeenAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeviceLastSeenAt
	}
	return nil
}

func (x *LandingDevice) GetDeviceDataTime() int64 {
	if x != nil {
		return x.DeviceDataTime
	}
	return 0
}

func (x *LandingDevice) GetDeviceLat() float64 {
	if x != nil {
		return x.DeviceLat
	}
	return 0
}

func (x *LandingDevice) GetDeviceLng() float64 {
	if x != nil {
		return x.DeviceLng
	}
	return 0
}

var File_as_external_api_landing_proto protoreflect.FileDescriptor

var file_as_external_api_landing_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x11, 0x0a, 0x0f, 0x57, 0x68, 0x6f, 0x41, 0x72, 0x65, 0x59, 0x6f, 0x75, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x48, 0x0a, 0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x10, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x59, 0x0a, 0x17, 0x4c, 0x61,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xe8, 0x01, 0x0a, 0x13, 0x4c, 0x61, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a,
	0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x19, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x80, 0x02, 0x0a, 0x12, 0x4c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x12,
	0x2a, 0x0a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x61, 0x70,
	0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x61, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x22, 0x85, 0x04, 0x0a, 0x0d, 0x4c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x12, 0x48, 0x0a,
	0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x12, 0x48, 0x0a, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x12, 0x2c, 0x0a, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2e, 0x0a, 0x12, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x4c, 0x0a, 0x13, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x65, 0x6e, 0x5f, 0x61, 0x74, 0x12,
	0x2a, 0x0a, 0x10, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x61, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6c, 0x6e, 0x67, 0x32, 0x63, 0x0a, 0x0e, 0x4c,
	0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x57, 0x68, 0x6f, 0x41, 0x72, 0x65, 0x59, 0x6f, 0x75, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x6e, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6c, 0x61, 0x6e, 0x64, 0x69, 0x6e, 0x67,
	0x42, 0x63, 0x0a, 0x21, 0x69, 0x6f, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x73, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x62, 0x72, 0x61, 0x68, 0x69, 0x6d, 0x6f, 0x7a, 0x65, 0x6b, 0x69, 0x63, 0x69,
	0x2f, 0x63, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6f, 0x2f, 0x76, 0x34, 0x2f, 0x61, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_as_external_api_landing_proto_rawDescOnce sync.Once
	file_as_external_api_landing_proto_rawDescData = file_as_external_api_landing_proto_rawDesc
)

func file_as_external_api_landing_proto_rawDescGZIP() []byte {
	file_as_external_api_landing_proto_rawDescOnce.Do(func() {
		file_as_external_api_landing_proto_rawDescData = protoimpl.X.CompressGZIP(file_as_external_api_landing_proto_rawDescData)
	})
	return file_as_external_api_landing_proto_rawDescData
}

var file_as_external_api_landing_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_as_external_api_landing_proto_goTypes = []interface{}{
	(*WhoAreYouParams)(nil),         // 0: api.WhoAreYouParams
	(*GetLandingResponse)(nil),      // 1: api.GetLandingResponse
	(*LandingOrganizationList)(nil), // 2: api.LandingOrganizationList
	(*LandingOrganization)(nil),     // 3: api.LandingOrganization
	(*LandingApplication)(nil),      // 4: api.LandingApplication
	(*LandingDevice)(nil),           // 5: api.LandingDevice
	(*timestamp.Timestamp)(nil),     // 6: google.protobuf.Timestamp
}
var file_as_external_api_landing_proto_depIdxs = []int32{
	2, // 0: api.GetLandingResponse.organizationList:type_name -> api.LandingOrganizationList
	3, // 1: api.LandingOrganizationList.organizations:type_name -> api.LandingOrganization
	4, // 2: api.LandingOrganization.applications:type_name -> api.LandingApplication
	5, // 3: api.LandingApplication.devices:type_name -> api.LandingDevice
	6, // 4: api.LandingDevice.device_created_at:type_name -> google.protobuf.Timestamp
	6, // 5: api.LandingDevice.device_updated_at:type_name -> google.protobuf.Timestamp
	6, // 6: api.LandingDevice.device_last_seen_at:type_name -> google.protobuf.Timestamp
	0, // 7: api.LandingService.GetLanding:input_type -> api.WhoAreYouParams
	1, // 8: api.LandingService.GetLanding:output_type -> api.GetLandingResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_as_external_api_landing_proto_init() }
func file_as_external_api_landing_proto_init() {
	if File_as_external_api_landing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_as_external_api_landing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoAreYouParams); i {
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
		file_as_external_api_landing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLandingResponse); i {
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
		file_as_external_api_landing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LandingOrganizationList); i {
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
		file_as_external_api_landing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LandingOrganization); i {
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
		file_as_external_api_landing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LandingApplication); i {
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
		file_as_external_api_landing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LandingDevice); i {
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
			RawDescriptor: file_as_external_api_landing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_as_external_api_landing_proto_goTypes,
		DependencyIndexes: file_as_external_api_landing_proto_depIdxs,
		MessageInfos:      file_as_external_api_landing_proto_msgTypes,
	}.Build()
	File_as_external_api_landing_proto = out.File
	file_as_external_api_landing_proto_rawDesc = nil
	file_as_external_api_landing_proto_goTypes = nil
	file_as_external_api_landing_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LandingServiceClient is the client API for LandingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LandingServiceClient interface {
	// Get returns the landing data
	GetLanding(ctx context.Context, in *WhoAreYouParams, opts ...grpc.CallOption) (*GetLandingResponse, error)
}

type landingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLandingServiceClient(cc grpc.ClientConnInterface) LandingServiceClient {
	return &landingServiceClient{cc}
}

func (c *landingServiceClient) GetLanding(ctx context.Context, in *WhoAreYouParams, opts ...grpc.CallOption) (*GetLandingResponse, error) {
	out := new(GetLandingResponse)
	err := c.cc.Invoke(ctx, "/api.LandingService/GetLanding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LandingServiceServer is the server API for LandingService service.
type LandingServiceServer interface {
	// Get returns the landing data
	GetLanding(context.Context, *WhoAreYouParams) (*GetLandingResponse, error)
}

// UnimplementedLandingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLandingServiceServer struct {
}

func (*UnimplementedLandingServiceServer) GetLanding(context.Context, *WhoAreYouParams) (*GetLandingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLanding not implemented")
}

func RegisterLandingServiceServer(s *grpc.Server, srv LandingServiceServer) {
	s.RegisterService(&_LandingService_serviceDesc, srv)
}

func _LandingService_GetLanding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoAreYouParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LandingServiceServer).GetLanding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LandingService/GetLanding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LandingServiceServer).GetLanding(ctx, req.(*WhoAreYouParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _LandingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.LandingService",
	HandlerType: (*LandingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLanding",
			Handler:    _LandingService_GetLanding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as/external/api/landing.proto",
}
