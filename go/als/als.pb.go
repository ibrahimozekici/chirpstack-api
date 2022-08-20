// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: als/als.proto

package als

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type Alarm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DevEui            string  `protobuf:"bytes,2,opt,name=devEui,proto3" json:"devEui,omitempty"`
	MinTreshold       float32 `protobuf:"fixed32,3,opt,name=minTreshold,proto3" json:"minTreshold,omitempty"`
	MaxTreshold       float32 `protobuf:"fixed32,4,opt,name=maxTreshold,proto3" json:"maxTreshold,omitempty"`
	Sms               bool    `protobuf:"varint,5,opt,name=sms,proto3" json:"sms,omitempty"`
	Email             bool    `protobuf:"varint,6,opt,name=email,proto3" json:"email,omitempty"`
	Notification      bool    `protobuf:"varint,7,opt,name=notification,proto3" json:"notification,omitempty"`
	Temperature       bool    `protobuf:"varint,8,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Humadity          bool    `protobuf:"varint,9,opt,name=humadity,proto3" json:"humadity,omitempty"`
	Ec                bool    `protobuf:"varint,10,opt,name=ec,proto3" json:"ec,omitempty"`
	Door              bool    `protobuf:"varint,11,opt,name=door,proto3" json:"door,omitempty"`
	WLeak             bool    `protobuf:"varint,12,opt,name=wLeak,proto3" json:"wLeak,omitempty"`
	UserID            int64   `protobuf:"varint,13,opt,name=userID,proto3" json:"userID,omitempty"`
	IpAddress         string  `protobuf:"bytes,14,opt,name=ipAddress,proto3" json:"ipAddress,omitempty"`
	IsTimeLimitActive bool    `protobuf:"varint,15,opt,name=isTimeLimitActive,proto3" json:"isTimeLimitActive,omitempty"`
	AlarmStartTime    float32 `protobuf:"fixed32,16,opt,name=alarmStartTime,proto3" json:"alarmStartTime,omitempty"`
	AlarmStopTime     float32 `protobuf:"fixed32,17,opt,name=alarmStopTime,proto3" json:"alarmStopTime,omitempty"`
	ZoneCategoryID    int64   `protobuf:"varint,18,opt,name=zoneCategoryID,proto3" json:"zoneCategoryID,omitempty"`
	IsActive          bool    `protobuf:"varint,19,opt,name=isActive,proto3" json:"isActive,omitempty"`
}

func (x *Alarm) Reset() {
	*x = Alarm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_als_als_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alarm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alarm) ProtoMessage() {}

func (x *Alarm) ProtoReflect() protoreflect.Message {
	mi := &file_als_als_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alarm.ProtoReflect.Descriptor instead.
func (*Alarm) Descriptor() ([]byte, []int) {
	return file_als_als_proto_rawDescGZIP(), []int{0}
}

func (x *Alarm) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Alarm) GetDevEui() string {
	if x != nil {
		return x.DevEui
	}
	return ""
}

func (x *Alarm) GetMinTreshold() float32 {
	if x != nil {
		return x.MinTreshold
	}
	return 0
}

func (x *Alarm) GetMaxTreshold() float32 {
	if x != nil {
		return x.MaxTreshold
	}
	return 0
}

func (x *Alarm) GetSms() bool {
	if x != nil {
		return x.Sms
	}
	return false
}

func (x *Alarm) GetEmail() bool {
	if x != nil {
		return x.Email
	}
	return false
}

func (x *Alarm) GetNotification() bool {
	if x != nil {
		return x.Notification
	}
	return false
}

func (x *Alarm) GetTemperature() bool {
	if x != nil {
		return x.Temperature
	}
	return false
}

func (x *Alarm) GetHumadity() bool {
	if x != nil {
		return x.Humadity
	}
	return false
}

func (x *Alarm) GetEc() bool {
	if x != nil {
		return x.Ec
	}
	return false
}

func (x *Alarm) GetDoor() bool {
	if x != nil {
		return x.Door
	}
	return false
}

func (x *Alarm) GetWLeak() bool {
	if x != nil {
		return x.WLeak
	}
	return false
}

func (x *Alarm) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Alarm) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *Alarm) GetIsTimeLimitActive() bool {
	if x != nil {
		return x.IsTimeLimitActive
	}
	return false
}

func (x *Alarm) GetAlarmStartTime() float32 {
	if x != nil {
		return x.AlarmStartTime
	}
	return 0
}

func (x *Alarm) GetAlarmStopTime() float32 {
	if x != nil {
		return x.AlarmStopTime
	}
	return 0
}

func (x *Alarm) GetZoneCategoryID() int64 {
	if x != nil {
		return x.ZoneCategoryID
	}
	return 0
}

func (x *Alarm) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type ColdRoomRestrictions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DevEui      string `protobuf:"bytes,2,opt,name=devEui,proto3" json:"devEui,omitempty"`
	AlarmID     int64  `protobuf:"varint,3,opt,name=alarmID,proto3" json:"alarmID,omitempty"`
	DefrostTime int64  `protobuf:"varint,4,opt,name=defrostTime,proto3" json:"defrostTime,omitempty"`
	DefrostFrq  int64  `protobuf:"varint,5,opt,name=defrostFrq,proto3" json:"defrostFrq,omitempty"`
	AlarmTime   int64  `protobuf:"varint,6,opt,name=alarmTime,proto3" json:"alarmTime,omitempty"`
}

func (x *ColdRoomRestrictions) Reset() {
	*x = ColdRoomRestrictions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_als_als_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ColdRoomRestrictions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColdRoomRestrictions) ProtoMessage() {}

func (x *ColdRoomRestrictions) ProtoReflect() protoreflect.Message {
	mi := &file_als_als_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColdRoomRestrictions.ProtoReflect.Descriptor instead.
func (*ColdRoomRestrictions) Descriptor() ([]byte, []int) {
	return file_als_als_proto_rawDescGZIP(), []int{1}
}

func (x *ColdRoomRestrictions) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ColdRoomRestrictions) GetDevEui() string {
	if x != nil {
		return x.DevEui
	}
	return ""
}

func (x *ColdRoomRestrictions) GetAlarmID() int64 {
	if x != nil {
		return x.AlarmID
	}
	return 0
}

func (x *ColdRoomRestrictions) GetDefrostTime() int64 {
	if x != nil {
		return x.DefrostTime
	}
	return 0
}

func (x *ColdRoomRestrictions) GetDefrostFrq() int64 {
	if x != nil {
		return x.DefrostFrq
	}
	return 0
}

func (x *ColdRoomRestrictions) GetAlarmTime() int64 {
	if x != nil {
		return x.AlarmTime
	}
	return 0
}

type CreateAlarmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alarm *Alarm `protobuf:"bytes,1,opt,name=alarm,proto3" json:"alarm,omitempty"`
}

func (x *CreateAlarmRequest) Reset() {
	*x = CreateAlarmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_als_als_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlarmRequest) ProtoMessage() {}

func (x *CreateAlarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_als_als_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlarmRequest.ProtoReflect.Descriptor instead.
func (*CreateAlarmRequest) Descriptor() ([]byte, []int) {
	return file_als_als_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAlarmRequest) GetAlarm() *Alarm {
	if x != nil {
		return x.Alarm
	}
	return nil
}

type CreateAlarmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alarm *Alarm `protobuf:"bytes,1,opt,name=alarm,proto3" json:"alarm,omitempty"`
}

func (x *CreateAlarmResponse) Reset() {
	*x = CreateAlarmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_als_als_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlarmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlarmResponse) ProtoMessage() {}

func (x *CreateAlarmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_als_als_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlarmResponse.ProtoReflect.Descriptor instead.
func (*CreateAlarmResponse) Descriptor() ([]byte, []int) {
	return file_als_als_proto_rawDescGZIP(), []int{3}
}

func (x *CreateAlarmResponse) GetAlarm() *Alarm {
	if x != nil {
		return x.Alarm
	}
	return nil
}

type GetAlarmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AlarmID int64 `protobuf:"varint,1,opt,name=alarmID,proto3" json:"alarmID,omitempty"`
}

func (x *GetAlarmRequest) Reset() {
	*x = GetAlarmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_als_als_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlarmRequest) ProtoMessage() {}

func (x *GetAlarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_als_als_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlarmRequest.ProtoReflect.Descriptor instead.
func (*GetAlarmRequest) Descriptor() ([]byte, []int) {
	return file_als_als_proto_rawDescGZIP(), []int{4}
}

func (x *GetAlarmRequest) GetAlarmID() int64 {
	if x != nil {
		return x.AlarmID
	}
	return 0
}

type GetAlarmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alarm *Alarm `protobuf:"bytes,1,opt,name=alarm,proto3" json:"alarm,omitempty"`
}

func (x *GetAlarmResponse) Reset() {
	*x = GetAlarmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_als_als_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlarmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlarmResponse) ProtoMessage() {}

func (x *GetAlarmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_als_als_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlarmResponse.ProtoReflect.Descriptor instead.
func (*GetAlarmResponse) Descriptor() ([]byte, []int) {
	return file_als_als_proto_rawDescGZIP(), []int{5}
}

func (x *GetAlarmResponse) GetAlarm() *Alarm {
	if x != nil {
		return x.Alarm
	}
	return nil
}

type CreateColdRoomRestrictionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ColdRes *ColdRoomRestrictions `protobuf:"bytes,1,opt,name=coldRes,proto3" json:"coldRes,omitempty"`
}

func (x *CreateColdRoomRestrictionsRequest) Reset() {
	*x = CreateColdRoomRestrictionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_als_als_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateColdRoomRestrictionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateColdRoomRestrictionsRequest) ProtoMessage() {}

func (x *CreateColdRoomRestrictionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_als_als_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateColdRoomRestrictionsRequest.ProtoReflect.Descriptor instead.
func (*CreateColdRoomRestrictionsRequest) Descriptor() ([]byte, []int) {
	return file_als_als_proto_rawDescGZIP(), []int{6}
}

func (x *CreateColdRoomRestrictionsRequest) GetColdRes() *ColdRoomRestrictions {
	if x != nil {
		return x.ColdRes
	}
	return nil
}

var File_als_als_proto protoreflect.FileDescriptor

var file_als_als_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x6c, 0x73, 0x2f, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x6c, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xad, 0x04, 0x0a, 0x05, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x65, 0x76, 0x45, 0x75, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65,
	0x76, 0x45, 0x75, 0x69, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x54, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x54, 0x72,
	0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x54, 0x72, 0x65,
	0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6d, 0x61, 0x78,
	0x54, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x6d, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x73, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x75, 0x6d, 0x61, 0x64, 0x69,
	0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x68, 0x75, 0x6d, 0x61, 0x64, 0x69,
	0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x63, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02,
	0x65, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x64, 0x6f, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x4c, 0x65, 0x61, 0x6b, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x77, 0x4c, 0x65, 0x61, 0x6b, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x69, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69,
	0x73, 0x54, 0x69, 0x6d, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0e, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x6c, 0x61, 0x72,
	0x6d, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0d, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x7a, 0x6f, 0x6e, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x7a, 0x6f, 0x6e, 0x65, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6c, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x52,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x65, 0x76, 0x45, 0x75, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76,
	0x45, 0x75, 0x69, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x46, 0x72, 0x71, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x65, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x46, 0x72, 0x71, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x36, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x6c, 0x73, 0x2e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x05,
	0x61, 0x6c, 0x61, 0x72, 0x6d, 0x22, 0x37, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05,
	0x61, 0x6c, 0x61, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x6c,
	0x73, 0x2e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x05, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x22, 0x2b,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x49, 0x44, 0x22, 0x34, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x20, 0x0a, 0x05, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x61, 0x6c, 0x73, 0x2e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x05, 0x61, 0x6c, 0x61, 0x72,
	0x6d, 0x22, 0x58, 0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x64, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x64, 0x52, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x6c, 0x73, 0x2e, 0x43, 0x6f,
	0x6c, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x07, 0x63, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x32, 0xf3, 0x01, 0x0a, 0x12,
	0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72,
	0x6d, 0x12, 0x17, 0x2e, 0x61, 0x6c, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c,
	0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x6c, 0x73,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x61,
	0x72, 0x6d, 0x12, 0x14, 0x2e, 0x61, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x61, 0x72,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x6c, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5e, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x64, 0x52,
	0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x26, 0x2e, 0x61, 0x6c, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x64,
	0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x42, 0x45, 0x0a, 0x15, 0x69, 0x6f, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74, 0x61,
	0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x6c, 0x73, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x62, 0x72, 0x61, 0x68, 0x69, 0x6d, 0x6f, 0x7a,
	0x65, 0x6b, 0x69, 0x63, 0x69, 0x2f, 0x63, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_als_als_proto_rawDescOnce sync.Once
	file_als_als_proto_rawDescData = file_als_als_proto_rawDesc
)

func file_als_als_proto_rawDescGZIP() []byte {
	file_als_als_proto_rawDescOnce.Do(func() {
		file_als_als_proto_rawDescData = protoimpl.X.CompressGZIP(file_als_als_proto_rawDescData)
	})
	return file_als_als_proto_rawDescData
}

var file_als_als_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_als_als_proto_goTypes = []interface{}{
	(*Alarm)(nil),                             // 0: als.Alarm
	(*ColdRoomRestrictions)(nil),              // 1: als.ColdRoomRestrictions
	(*CreateAlarmRequest)(nil),                // 2: als.CreateAlarmRequest
	(*CreateAlarmResponse)(nil),               // 3: als.CreateAlarmResponse
	(*GetAlarmRequest)(nil),                   // 4: als.GetAlarmRequest
	(*GetAlarmResponse)(nil),                  // 5: als.GetAlarmResponse
	(*CreateColdRoomRestrictionsRequest)(nil), // 6: als.CreateColdRoomRestrictionsRequest
	(*empty.Empty)(nil),                       // 7: google.protobuf.Empty
}
var file_als_als_proto_depIdxs = []int32{
	0, // 0: als.CreateAlarmRequest.alarm:type_name -> als.Alarm
	0, // 1: als.CreateAlarmResponse.alarm:type_name -> als.Alarm
	0, // 2: als.GetAlarmResponse.alarm:type_name -> als.Alarm
	1, // 3: als.CreateColdRoomRestrictionsRequest.coldRes:type_name -> als.ColdRoomRestrictions
	2, // 4: als.AlarmServerService.CreateAlarm:input_type -> als.CreateAlarmRequest
	4, // 5: als.AlarmServerService.GetAlarm:input_type -> als.GetAlarmRequest
	6, // 6: als.AlarmServerService.CreateColdRoomRestrictions:input_type -> als.CreateColdRoomRestrictionsRequest
	3, // 7: als.AlarmServerService.CreateAlarm:output_type -> als.CreateAlarmResponse
	5, // 8: als.AlarmServerService.GetAlarm:output_type -> als.GetAlarmResponse
	7, // 9: als.AlarmServerService.CreateColdRoomRestrictions:output_type -> google.protobuf.Empty
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_als_als_proto_init() }
func file_als_als_proto_init() {
	if File_als_als_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_als_als_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alarm); i {
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
		file_als_als_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ColdRoomRestrictions); i {
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
		file_als_als_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlarmRequest); i {
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
		file_als_als_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlarmResponse); i {
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
		file_als_als_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlarmRequest); i {
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
		file_als_als_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlarmResponse); i {
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
		file_als_als_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateColdRoomRestrictionsRequest); i {
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
			RawDescriptor: file_als_als_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_als_als_proto_goTypes,
		DependencyIndexes: file_als_als_proto_depIdxs,
		MessageInfos:      file_als_als_proto_msgTypes,
	}.Build()
	File_als_als_proto = out.File
	file_als_als_proto_rawDesc = nil
	file_als_als_proto_goTypes = nil
	file_als_als_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AlarmServerServiceClient is the client API for AlarmServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlarmServerServiceClient interface {
	// CreateDevice creates the given device.
	CreateAlarm(ctx context.Context, in *CreateAlarmRequest, opts ...grpc.CallOption) (*CreateAlarmResponse, error)
	GetAlarm(ctx context.Context, in *GetAlarmRequest, opts ...grpc.CallOption) (*GetAlarmResponse, error)
	// rpc ListAlarm(ListAlarmRequest) returns (ListAlarmResponse){}
	// rpc UpdateAlarm(UpdateAlarmRequest) returns (GetAlarmResponse){}
	// rpc DeleteAlarm(DeleteAlarmRequest) returns (google.protobuf.Empty){}
	// rpc ListAllOrganizationAlarms(ListOrganizationAlarmRequest) returns (ListOrganizationAlarmResponse){}
	// rpc CreateMultipleAlarms(CreateMultipleAlarmRequest) returns (CreateMultipleAlarmResponse){}
	// rpc DeleteMultipleAlarms(DeleteMultipleAlarmRequest) returns (google.protobuf.Empty){}
	// rpc DeleteSensorAlarms(DeleteMultipleDevAlarmRequest) returns (google.protobuf.Empty){}
	// rpc DeleteZoneAlarms(DeleteMultipleZoneAlarmRequest) returns (google.protobuf.Empty){}
	// rpc CreateAlarmDateTime(CreateAlarmDateTimeRequest) returns (CreateAlarmDateResponse){}
	// rpc CreateAlarmDates(CreateAlarmDateRequest) returns (CreateAlarmDateResponse){}
	// rpc GetAlarmLogs(GetAlarmLogRequest) returns (GetAlarmLogResponse){}
	CreateColdRoomRestrictions(ctx context.Context, in *CreateColdRoomRestrictionsRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type alarmServerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlarmServerServiceClient(cc grpc.ClientConnInterface) AlarmServerServiceClient {
	return &alarmServerServiceClient{cc}
}

func (c *alarmServerServiceClient) CreateAlarm(ctx context.Context, in *CreateAlarmRequest, opts ...grpc.CallOption) (*CreateAlarmResponse, error) {
	out := new(CreateAlarmResponse)
	err := c.cc.Invoke(ctx, "/als.AlarmServerService/CreateAlarm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmServerServiceClient) GetAlarm(ctx context.Context, in *GetAlarmRequest, opts ...grpc.CallOption) (*GetAlarmResponse, error) {
	out := new(GetAlarmResponse)
	err := c.cc.Invoke(ctx, "/als.AlarmServerService/GetAlarm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmServerServiceClient) CreateColdRoomRestrictions(ctx context.Context, in *CreateColdRoomRestrictionsRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/als.AlarmServerService/CreateColdRoomRestrictions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlarmServerServiceServer is the server API for AlarmServerService service.
type AlarmServerServiceServer interface {
	// CreateDevice creates the given device.
	CreateAlarm(context.Context, *CreateAlarmRequest) (*CreateAlarmResponse, error)
	GetAlarm(context.Context, *GetAlarmRequest) (*GetAlarmResponse, error)
	// rpc ListAlarm(ListAlarmRequest) returns (ListAlarmResponse){}
	// rpc UpdateAlarm(UpdateAlarmRequest) returns (GetAlarmResponse){}
	// rpc DeleteAlarm(DeleteAlarmRequest) returns (google.protobuf.Empty){}
	// rpc ListAllOrganizationAlarms(ListOrganizationAlarmRequest) returns (ListOrganizationAlarmResponse){}
	// rpc CreateMultipleAlarms(CreateMultipleAlarmRequest) returns (CreateMultipleAlarmResponse){}
	// rpc DeleteMultipleAlarms(DeleteMultipleAlarmRequest) returns (google.protobuf.Empty){}
	// rpc DeleteSensorAlarms(DeleteMultipleDevAlarmRequest) returns (google.protobuf.Empty){}
	// rpc DeleteZoneAlarms(DeleteMultipleZoneAlarmRequest) returns (google.protobuf.Empty){}
	// rpc CreateAlarmDateTime(CreateAlarmDateTimeRequest) returns (CreateAlarmDateResponse){}
	// rpc CreateAlarmDates(CreateAlarmDateRequest) returns (CreateAlarmDateResponse){}
	// rpc GetAlarmLogs(GetAlarmLogRequest) returns (GetAlarmLogResponse){}
	CreateColdRoomRestrictions(context.Context, *CreateColdRoomRestrictionsRequest) (*empty.Empty, error)
}

// UnimplementedAlarmServerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAlarmServerServiceServer struct {
}

func (*UnimplementedAlarmServerServiceServer) CreateAlarm(context.Context, *CreateAlarmRequest) (*CreateAlarmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAlarm not implemented")
}
func (*UnimplementedAlarmServerServiceServer) GetAlarm(context.Context, *GetAlarmRequest) (*GetAlarmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlarm not implemented")
}
func (*UnimplementedAlarmServerServiceServer) CreateColdRoomRestrictions(context.Context, *CreateColdRoomRestrictionsRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateColdRoomRestrictions not implemented")
}

func RegisterAlarmServerServiceServer(s *grpc.Server, srv AlarmServerServiceServer) {
	s.RegisterService(&_AlarmServerService_serviceDesc, srv)
}

func _AlarmServerService_CreateAlarm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlarmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmServerServiceServer).CreateAlarm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/als.AlarmServerService/CreateAlarm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmServerServiceServer).CreateAlarm(ctx, req.(*CreateAlarmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmServerService_GetAlarm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlarmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmServerServiceServer).GetAlarm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/als.AlarmServerService/GetAlarm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmServerServiceServer).GetAlarm(ctx, req.(*GetAlarmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmServerService_CreateColdRoomRestrictions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateColdRoomRestrictionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmServerServiceServer).CreateColdRoomRestrictions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/als.AlarmServerService/CreateColdRoomRestrictions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmServerServiceServer).CreateColdRoomRestrictions(ctx, req.(*CreateColdRoomRestrictionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlarmServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "als.AlarmServerService",
	HandlerType: (*AlarmServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAlarm",
			Handler:    _AlarmServerService_CreateAlarm_Handler,
		},
		{
			MethodName: "GetAlarm",
			Handler:    _AlarmServerService_GetAlarm_Handler,
		},
		{
			MethodName: "CreateColdRoomRestrictions",
			Handler:    _AlarmServerService_CreateColdRoomRestrictions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "als/als.proto",
}
