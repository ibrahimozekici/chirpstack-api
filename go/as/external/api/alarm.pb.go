// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: as/external/api/alarm.proto

package api

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Alarm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Alarm id.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Device EUI (HEX encoded).
	DevEui string `protobuf:"bytes,2,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	// Condition of alarm
	MinTreshold float32 `protobuf:"fixed32,3,opt,name=min_treshold,json=minTreshold,proto3" json:"min_treshold,omitempty"`
	// Value of alarm
	MaxTreshold float32 `protobuf:"fixed32,4,opt,name=max_treshold,json=maxTreshold,proto3" json:"max_treshold,omitempty"`
	// Action of alarm
	Action string `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`
	// Target of alarm
	Target string `protobuf:"bytes,6,opt,name=target,proto3" json:"target,omitempty"`
	// Target of alarm
	Section        string               `protobuf:"bytes,7,opt,name=section,proto3" json:"section,omitempty"`
	SubmissionDate *timestamp.Timestamp `protobuf:"bytes,10,opt,name=submission_date,proto3" json:"submission_date,omitempty"`
}

func (x *Alarm) Reset() {
	*x = Alarm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_alarm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alarm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alarm) ProtoMessage() {}

func (x *Alarm) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_alarm_proto_msgTypes[0]
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
	return file_as_external_api_alarm_proto_rawDescGZIP(), []int{0}
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

func (x *Alarm) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Alarm) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Alarm) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *Alarm) GetSubmissionDate() *timestamp.Timestamp {
	if x != nil {
		return x.SubmissionDate
	}
	return nil
}

type CreateAlarmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alarm []*Alarm `protobuf:"bytes,1,rep,name=alarm,proto3" json:"alarm,omitempty"`
}

func (x *CreateAlarmResponse) Reset() {
	*x = CreateAlarmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_alarm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlarmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlarmResponse) ProtoMessage() {}

func (x *CreateAlarmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_alarm_proto_msgTypes[1]
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
	return file_as_external_api_alarm_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAlarmResponse) GetAlarm() []*Alarm {
	if x != nil {
		return x.Alarm
	}
	return nil
}

type CreateAlarmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateAlarm []*Alarm `protobuf:"bytes,1,rep,name=create_alarm,json=createAlarm,proto3" json:"create_alarm,omitempty"`
}

func (x *CreateAlarmRequest) Reset() {
	*x = CreateAlarmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_alarm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlarmRequest) ProtoMessage() {}

func (x *CreateAlarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_alarm_proto_msgTypes[2]
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
	return file_as_external_api_alarm_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAlarmRequest) GetCreateAlarm() []*Alarm {
	if x != nil {
		return x.CreateAlarm
	}
	return nil
}

type GetAlarmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Device EUI (HEX encoded).
	AlarmId string `protobuf:"bytes,1,opt,name=alarm_id,json=alarmID,proto3" json:"alarm_id,omitempty"`
}

func (x *GetAlarmRequest) Reset() {
	*x = GetAlarmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_alarm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlarmRequest) ProtoMessage() {}

func (x *GetAlarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_alarm_proto_msgTypes[3]
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
	return file_as_external_api_alarm_proto_rawDescGZIP(), []int{3}
}

func (x *GetAlarmRequest) GetAlarmId() string {
	if x != nil {
		return x.AlarmId
	}
	return ""
}

type GetAlarmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Alarm object
	Alarm *Alarm `protobuf:"bytes,1,opt,name=alarm,proto3" json:"alarm,omitempty"`
}

func (x *GetAlarmResponse) Reset() {
	*x = GetAlarmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_alarm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlarmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlarmResponse) ProtoMessage() {}

func (x *GetAlarmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_alarm_proto_msgTypes[4]
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
	return file_as_external_api_alarm_proto_rawDescGZIP(), []int{4}
}

func (x *GetAlarmResponse) GetAlarm() *Alarm {
	if x != nil {
		return x.Alarm
	}
	return nil
}

type DeleteAlarmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Alarm Id (HEX encoded).
	AlarmId string `protobuf:"bytes,1,opt,name=alarm_id,json=alarmID,proto3" json:"alarm_id,omitempty"`
}

func (x *DeleteAlarmRequest) Reset() {
	*x = DeleteAlarmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_alarm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAlarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAlarmRequest) ProtoMessage() {}

func (x *DeleteAlarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_alarm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAlarmRequest.ProtoReflect.Descriptor instead.
func (*DeleteAlarmRequest) Descriptor() ([]byte, []int) {
	return file_as_external_api_alarm_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteAlarmRequest) GetAlarmId() string {
	if x != nil {
		return x.AlarmId
	}
	return ""
}

type UpdateAlarmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Alarm object to update.
	Alarm   *Alarm `protobuf:"bytes,1,opt,name=alarm,proto3" json:"alarm,omitempty"`
	AlarmId string `protobuf:"bytes,2,opt,name=alarm_id,json=alarmID,proto3" json:"alarm_id,omitempty"`
}

func (x *UpdateAlarmRequest) Reset() {
	*x = UpdateAlarmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_alarm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAlarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAlarmRequest) ProtoMessage() {}

func (x *UpdateAlarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_alarm_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAlarmRequest.ProtoReflect.Descriptor instead.
func (*UpdateAlarmRequest) Descriptor() ([]byte, []int) {
	return file_as_external_api_alarm_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateAlarmRequest) GetAlarm() *Alarm {
	if x != nil {
		return x.Alarm
	}
	return nil
}

func (x *UpdateAlarmRequest) GetAlarmId() string {
	if x != nil {
		return x.AlarmId
	}
	return ""
}

type ListAlarmRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Max number of devices to return in the result-set.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Device EUI (HEX encoded).
	DevEui string `protobuf:"bytes,3,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
}

func (x *ListAlarmRequest) Reset() {
	*x = ListAlarmRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_alarm_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlarmRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlarmRequest) ProtoMessage() {}

func (x *ListAlarmRequest) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_alarm_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlarmRequest.ProtoReflect.Descriptor instead.
func (*ListAlarmRequest) Descriptor() ([]byte, []int) {
	return file_as_external_api_alarm_proto_rawDescGZIP(), []int{7}
}

func (x *ListAlarmRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListAlarmRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListAlarmRequest) GetDevEui() string {
	if x != nil {
		return x.DevEui
	}
	return ""
}

type ListAlarmResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total number of devices available within the result-set.
	TotalCount int64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// Devices within this result-set.
	Result []*Alarm `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *ListAlarmResponse) Reset() {
	*x = ListAlarmResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_as_external_api_alarm_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAlarmResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAlarmResponse) ProtoMessage() {}

func (x *ListAlarmResponse) ProtoReflect() protoreflect.Message {
	mi := &file_as_external_api_alarm_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAlarmResponse.ProtoReflect.Descriptor instead.
func (*ListAlarmResponse) Descriptor() ([]byte, []int) {
	return file_as_external_api_alarm_proto_rawDescGZIP(), []int{8}
}

func (x *ListAlarmResponse) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ListAlarmResponse) GetResult() []*Alarm {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_as_external_api_alarm_proto protoreflect.FileDescriptor

var file_as_external_api_alarm_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61,
	0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86,
	0x02, 0x0a, 0x05, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x5f,
	0x65, 0x75, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x45, 0x55,
	0x49, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x69, 0x6e, 0x5f, 0x74, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6d, 0x69, 0x6e, 0x54, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x72, 0x65, 0x73,
	0x68, 0x6f, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x54,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x44, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x22, 0x37, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x05, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x05, 0x61, 0x6c, 0x61, 0x72, 0x6d,
	0x22, 0x43, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x6c, 0x61, 0x72, 0x6d, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x61, 0x72,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x61, 0x72,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x61, 0x72,
	0x6d, 0x49, 0x44, 0x22, 0x34, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x61, 0x6c, 0x61, 0x72, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c, 0x61,
	0x72, 0x6d, 0x52, 0x05, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x49, 0x44, 0x22, 0x51, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x05, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x05, 0x61, 0x6c, 0x61,
	0x72, 0x6d, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x49, 0x44, 0x22, 0x59, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x5f, 0x65, 0x75, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x64, 0x65, 0x76, 0x45, 0x55, 0x49, 0x22, 0x58, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0xb6, 0x03, 0x0a, 0x0c, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x52, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6c, 0x61, 0x72, 0x6d, 0x3a, 0x01, 0x2a, 0x12, 0x51, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x14,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x2f,
	0x7b, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x49, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x61,
	0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6c, 0x61, 0x72, 0x6d, 0x12, 0x5a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x1a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c,
	0x61, 0x72, 0x6d, 0x2f, 0x7b, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x01,
	0x2a, 0x12, 0x58, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x61, 0x72, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6c, 0x61, 0x72, 0x6d,
	0x2f, 0x7b, 0x61, 0x6c, 0x61, 0x72, 0x6d, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x63, 0x0a, 0x21, 0x69,
	0x6f, 0x2e, 0x63, 0x68, 0x69, 0x72, 0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x73, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x62, 0x72,
	0x61, 0x68, 0x69, 0x6d, 0x6f, 0x7a, 0x65, 0x6b, 0x69, 0x63, 0x69, 0x2f, 0x63, 0x68, 0x69, 0x72,
	0x70, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x34,
	0x2f, 0x61, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_as_external_api_alarm_proto_rawDescOnce sync.Once
	file_as_external_api_alarm_proto_rawDescData = file_as_external_api_alarm_proto_rawDesc
)

func file_as_external_api_alarm_proto_rawDescGZIP() []byte {
	file_as_external_api_alarm_proto_rawDescOnce.Do(func() {
		file_as_external_api_alarm_proto_rawDescData = protoimpl.X.CompressGZIP(file_as_external_api_alarm_proto_rawDescData)
	})
	return file_as_external_api_alarm_proto_rawDescData
}

var file_as_external_api_alarm_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_as_external_api_alarm_proto_goTypes = []interface{}{
	(*Alarm)(nil),               // 0: api.Alarm
	(*CreateAlarmResponse)(nil), // 1: api.CreateAlarmResponse
	(*CreateAlarmRequest)(nil),  // 2: api.CreateAlarmRequest
	(*GetAlarmRequest)(nil),     // 3: api.GetAlarmRequest
	(*GetAlarmResponse)(nil),    // 4: api.GetAlarmResponse
	(*DeleteAlarmRequest)(nil),  // 5: api.DeleteAlarmRequest
	(*UpdateAlarmRequest)(nil),  // 6: api.UpdateAlarmRequest
	(*ListAlarmRequest)(nil),    // 7: api.ListAlarmRequest
	(*ListAlarmResponse)(nil),   // 8: api.ListAlarmResponse
	(*timestamp.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*empty.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_as_external_api_alarm_proto_depIdxs = []int32{
	9,  // 0: api.Alarm.submission_date:type_name -> google.protobuf.Timestamp
	0,  // 1: api.CreateAlarmResponse.alarm:type_name -> api.Alarm
	0,  // 2: api.CreateAlarmRequest.create_alarm:type_name -> api.Alarm
	0,  // 3: api.GetAlarmResponse.alarm:type_name -> api.Alarm
	0,  // 4: api.UpdateAlarmRequest.alarm:type_name -> api.Alarm
	0,  // 5: api.ListAlarmResponse.result:type_name -> api.Alarm
	2,  // 6: api.AlarmService.Create:input_type -> api.CreateAlarmRequest
	3,  // 7: api.AlarmService.Get:input_type -> api.GetAlarmRequest
	7,  // 8: api.AlarmService.List:input_type -> api.ListAlarmRequest
	6,  // 9: api.AlarmService.Update:input_type -> api.UpdateAlarmRequest
	5,  // 10: api.AlarmService.Delete:input_type -> api.DeleteAlarmRequest
	1,  // 11: api.AlarmService.Create:output_type -> api.CreateAlarmResponse
	4,  // 12: api.AlarmService.Get:output_type -> api.GetAlarmResponse
	8,  // 13: api.AlarmService.List:output_type -> api.ListAlarmResponse
	4,  // 14: api.AlarmService.Update:output_type -> api.GetAlarmResponse
	10, // 15: api.AlarmService.Delete:output_type -> google.protobuf.Empty
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_as_external_api_alarm_proto_init() }
func file_as_external_api_alarm_proto_init() {
	if File_as_external_api_alarm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_as_external_api_alarm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_as_external_api_alarm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_as_external_api_alarm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_as_external_api_alarm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_as_external_api_alarm_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_as_external_api_alarm_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAlarmRequest); i {
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
		file_as_external_api_alarm_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAlarmRequest); i {
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
		file_as_external_api_alarm_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlarmRequest); i {
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
		file_as_external_api_alarm_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAlarmResponse); i {
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
			RawDescriptor: file_as_external_api_alarm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_as_external_api_alarm_proto_goTypes,
		DependencyIndexes: file_as_external_api_alarm_proto_depIdxs,
		MessageInfos:      file_as_external_api_alarm_proto_msgTypes,
	}.Build()
	File_as_external_api_alarm_proto = out.File
	file_as_external_api_alarm_proto_rawDesc = nil
	file_as_external_api_alarm_proto_goTypes = nil
	file_as_external_api_alarm_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AlarmServiceClient is the client API for AlarmService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AlarmServiceClient interface {
	// Create creates the alarm
	Create(ctx context.Context, in *CreateAlarmRequest, opts ...grpc.CallOption) (*CreateAlarmResponse, error)
	// Get return the alarm
	Get(ctx context.Context, in *GetAlarmRequest, opts ...grpc.CallOption) (*GetAlarmResponse, error)
	List(ctx context.Context, in *ListAlarmRequest, opts ...grpc.CallOption) (*ListAlarmResponse, error)
	// Update updates the alarm matching the given alarm_id.
	Update(ctx context.Context, in *UpdateAlarmRequest, opts ...grpc.CallOption) (*GetAlarmResponse, error)
	// Delete deletes the alarm matching the given alarm_id.
	Delete(ctx context.Context, in *DeleteAlarmRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type alarmServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlarmServiceClient(cc grpc.ClientConnInterface) AlarmServiceClient {
	return &alarmServiceClient{cc}
}

func (c *alarmServiceClient) Create(ctx context.Context, in *CreateAlarmRequest, opts ...grpc.CallOption) (*CreateAlarmResponse, error) {
	out := new(CreateAlarmResponse)
	err := c.cc.Invoke(ctx, "/api.AlarmService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmServiceClient) Get(ctx context.Context, in *GetAlarmRequest, opts ...grpc.CallOption) (*GetAlarmResponse, error) {
	out := new(GetAlarmResponse)
	err := c.cc.Invoke(ctx, "/api.AlarmService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmServiceClient) List(ctx context.Context, in *ListAlarmRequest, opts ...grpc.CallOption) (*ListAlarmResponse, error) {
	out := new(ListAlarmResponse)
	err := c.cc.Invoke(ctx, "/api.AlarmService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmServiceClient) Update(ctx context.Context, in *UpdateAlarmRequest, opts ...grpc.CallOption) (*GetAlarmResponse, error) {
	out := new(GetAlarmResponse)
	err := c.cc.Invoke(ctx, "/api.AlarmService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alarmServiceClient) Delete(ctx context.Context, in *DeleteAlarmRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.AlarmService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlarmServiceServer is the server API for AlarmService service.
type AlarmServiceServer interface {
	// Create creates the alarm
	Create(context.Context, *CreateAlarmRequest) (*CreateAlarmResponse, error)
	// Get return the alarm
	Get(context.Context, *GetAlarmRequest) (*GetAlarmResponse, error)
	List(context.Context, *ListAlarmRequest) (*ListAlarmResponse, error)
	// Update updates the alarm matching the given alarm_id.
	Update(context.Context, *UpdateAlarmRequest) (*GetAlarmResponse, error)
	// Delete deletes the alarm matching the given alarm_id.
	Delete(context.Context, *DeleteAlarmRequest) (*empty.Empty, error)
}

// UnimplementedAlarmServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAlarmServiceServer struct {
}

func (*UnimplementedAlarmServiceServer) Create(context.Context, *CreateAlarmRequest) (*CreateAlarmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAlarmServiceServer) Get(context.Context, *GetAlarmRequest) (*GetAlarmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAlarmServiceServer) List(context.Context, *ListAlarmRequest) (*ListAlarmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedAlarmServiceServer) Update(context.Context, *UpdateAlarmRequest) (*GetAlarmResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAlarmServiceServer) Delete(context.Context, *DeleteAlarmRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterAlarmServiceServer(s *grpc.Server, srv AlarmServiceServer) {
	s.RegisterService(&_AlarmService_serviceDesc, srv)
}

func _AlarmService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlarmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AlarmService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmServiceServer).Create(ctx, req.(*CreateAlarmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlarmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AlarmService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmServiceServer).Get(ctx, req.(*GetAlarmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAlarmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AlarmService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmServiceServer).List(ctx, req.(*ListAlarmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlarmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AlarmService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmServiceServer).Update(ctx, req.(*UpdateAlarmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlarmService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAlarmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlarmServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AlarmService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlarmServiceServer).Delete(ctx, req.(*DeleteAlarmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlarmService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AlarmService",
	HandlerType: (*AlarmServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AlarmService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AlarmService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AlarmService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AlarmService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AlarmService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as/external/api/alarm.proto",
}
