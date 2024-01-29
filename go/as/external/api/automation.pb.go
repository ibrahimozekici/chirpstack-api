// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/api/automation.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Automation struct {
	Id                   int64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SenderSensor         string               `protobuf:"bytes,2,opt,name=sender_sensor,json=senderSensor,proto3" json:"sender_sensor,omitempty"`
	ReceiverSensor       string               `protobuf:"bytes,3,opt,name=receiver_sensor,json=receiverSensor,proto3" json:"receiver_sensor,omitempty"`
	Condition            string               `protobuf:"bytes,4,opt,name=condition,proto3" json:"condition,omitempty"`
	Action               string               `protobuf:"bytes,5,opt,name=action,proto3" json:"action,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	UserId               int64                `protobuf:"varint,8,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrganizationId       int64                `protobuf:"varint,9,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	SenderDeviceType     int64                `protobuf:"varint,10,opt,name=sender_device_type,json=senderDeviceType,proto3" json:"sender_device_type,omitempty"`
	ReceiverDeviceType   int64                `protobuf:"varint,11,opt,name=receiver_device_type,json=receiverDeviceType,proto3" json:"receiver_device_type,omitempty"`
	SenderDeviceName     string               `protobuf:"bytes,12,opt,name=sender_device_name,json=senderDeviceName,proto3" json:"sender_device_name,omitempty"`
	ReceiverDeviceName   string               `protobuf:"bytes,13,opt,name=receiver_device_name,json=receiverDeviceName,proto3" json:"receiver_device_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Automation) Reset()         { *m = Automation{} }
func (m *Automation) String() string { return proto.CompactTextString(m) }
func (*Automation) ProtoMessage()    {}
func (*Automation) Descriptor() ([]byte, []int) {
	return fileDescriptor_5791182e45212fe6, []int{0}
}

func (m *Automation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Automation.Unmarshal(m, b)
}
func (m *Automation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Automation.Marshal(b, m, deterministic)
}
func (m *Automation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Automation.Merge(m, src)
}
func (m *Automation) XXX_Size() int {
	return xxx_messageInfo_Automation.Size(m)
}
func (m *Automation) XXX_DiscardUnknown() {
	xxx_messageInfo_Automation.DiscardUnknown(m)
}

var xxx_messageInfo_Automation proto.InternalMessageInfo

func (m *Automation) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Automation) GetSenderSensor() string {
	if m != nil {
		return m.SenderSensor
	}
	return ""
}

func (m *Automation) GetReceiverSensor() string {
	if m != nil {
		return m.ReceiverSensor
	}
	return ""
}

func (m *Automation) GetCondition() string {
	if m != nil {
		return m.Condition
	}
	return ""
}

func (m *Automation) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Automation) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Automation) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Automation) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Automation) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *Automation) GetSenderDeviceType() int64 {
	if m != nil {
		return m.SenderDeviceType
	}
	return 0
}

func (m *Automation) GetReceiverDeviceType() int64 {
	if m != nil {
		return m.ReceiverDeviceType
	}
	return 0
}

func (m *Automation) GetSenderDeviceName() string {
	if m != nil {
		return m.SenderDeviceName
	}
	return ""
}

func (m *Automation) GetReceiverDeviceName() string {
	if m != nil {
		return m.ReceiverDeviceName
	}
	return ""
}

type CreateAutomationRequest struct {
	Automation           *Automation `protobuf:"bytes,1,opt,name=automation,proto3" json:"automation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateAutomationRequest) Reset()         { *m = CreateAutomationRequest{} }
func (m *CreateAutomationRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAutomationRequest) ProtoMessage()    {}
func (*CreateAutomationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5791182e45212fe6, []int{1}
}

func (m *CreateAutomationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAutomationRequest.Unmarshal(m, b)
}
func (m *CreateAutomationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAutomationRequest.Marshal(b, m, deterministic)
}
func (m *CreateAutomationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAutomationRequest.Merge(m, src)
}
func (m *CreateAutomationRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAutomationRequest.Size(m)
}
func (m *CreateAutomationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAutomationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAutomationRequest proto.InternalMessageInfo

func (m *CreateAutomationRequest) GetAutomation() *Automation {
	if m != nil {
		return m.Automation
	}
	return nil
}

type CreateAutomationResponse struct {
	Automation           *Automation `protobuf:"bytes,1,opt,name=automation,proto3" json:"automation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CreateAutomationResponse) Reset()         { *m = CreateAutomationResponse{} }
func (m *CreateAutomationResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAutomationResponse) ProtoMessage()    {}
func (*CreateAutomationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5791182e45212fe6, []int{2}
}

func (m *CreateAutomationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAutomationResponse.Unmarshal(m, b)
}
func (m *CreateAutomationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAutomationResponse.Marshal(b, m, deterministic)
}
func (m *CreateAutomationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAutomationResponse.Merge(m, src)
}
func (m *CreateAutomationResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAutomationResponse.Size(m)
}
func (m *CreateAutomationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAutomationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAutomationResponse proto.InternalMessageInfo

func (m *CreateAutomationResponse) GetAutomation() *Automation {
	if m != nil {
		return m.Automation
	}
	return nil
}

type GetAutomationRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAutomationRequest) Reset()         { *m = GetAutomationRequest{} }
func (m *GetAutomationRequest) String() string { return proto.CompactTextString(m) }
func (*GetAutomationRequest) ProtoMessage()    {}
func (*GetAutomationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5791182e45212fe6, []int{3}
}

func (m *GetAutomationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAutomationRequest.Unmarshal(m, b)
}
func (m *GetAutomationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAutomationRequest.Marshal(b, m, deterministic)
}
func (m *GetAutomationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAutomationRequest.Merge(m, src)
}
func (m *GetAutomationRequest) XXX_Size() int {
	return xxx_messageInfo_GetAutomationRequest.Size(m)
}
func (m *GetAutomationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAutomationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAutomationRequest proto.InternalMessageInfo

func (m *GetAutomationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetAutomationResponse struct {
	Automation           *Automation `protobuf:"bytes,1,opt,name=automation,proto3" json:"automation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetAutomationResponse) Reset()         { *m = GetAutomationResponse{} }
func (m *GetAutomationResponse) String() string { return proto.CompactTextString(m) }
func (*GetAutomationResponse) ProtoMessage()    {}
func (*GetAutomationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5791182e45212fe6, []int{4}
}

func (m *GetAutomationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAutomationResponse.Unmarshal(m, b)
}
func (m *GetAutomationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAutomationResponse.Marshal(b, m, deterministic)
}
func (m *GetAutomationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAutomationResponse.Merge(m, src)
}
func (m *GetAutomationResponse) XXX_Size() int {
	return xxx_messageInfo_GetAutomationResponse.Size(m)
}
func (m *GetAutomationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAutomationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAutomationResponse proto.InternalMessageInfo

func (m *GetAutomationResponse) GetAutomation() *Automation {
	if m != nil {
		return m.Automation
	}
	return nil
}

type ListAutomationRequest struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DevEui               string   `protobuf:"bytes,2,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListAutomationRequest) Reset()         { *m = ListAutomationRequest{} }
func (m *ListAutomationRequest) String() string { return proto.CompactTextString(m) }
func (*ListAutomationRequest) ProtoMessage()    {}
func (*ListAutomationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5791182e45212fe6, []int{5}
}

func (m *ListAutomationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAutomationRequest.Unmarshal(m, b)
}
func (m *ListAutomationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAutomationRequest.Marshal(b, m, deterministic)
}
func (m *ListAutomationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAutomationRequest.Merge(m, src)
}
func (m *ListAutomationRequest) XXX_Size() int {
	return xxx_messageInfo_ListAutomationRequest.Size(m)
}
func (m *ListAutomationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAutomationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListAutomationRequest proto.InternalMessageInfo

func (m *ListAutomationRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ListAutomationRequest) GetDevEui() string {
	if m != nil {
		return m.DevEui
	}
	return ""
}

type ListAutomationResponse struct {
	Automations          []*Automation `protobuf:"bytes,1,rep,name=automations,proto3" json:"automations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListAutomationResponse) Reset()         { *m = ListAutomationResponse{} }
func (m *ListAutomationResponse) String() string { return proto.CompactTextString(m) }
func (*ListAutomationResponse) ProtoMessage()    {}
func (*ListAutomationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5791182e45212fe6, []int{6}
}

func (m *ListAutomationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListAutomationResponse.Unmarshal(m, b)
}
func (m *ListAutomationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListAutomationResponse.Marshal(b, m, deterministic)
}
func (m *ListAutomationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListAutomationResponse.Merge(m, src)
}
func (m *ListAutomationResponse) XXX_Size() int {
	return xxx_messageInfo_ListAutomationResponse.Size(m)
}
func (m *ListAutomationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListAutomationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListAutomationResponse proto.InternalMessageInfo

func (m *ListAutomationResponse) GetAutomations() []*Automation {
	if m != nil {
		return m.Automations
	}
	return nil
}

type UpdateAutomationRequest struct {
	Automation           *Automation `protobuf:"bytes,1,opt,name=automation,proto3" json:"automation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateAutomationRequest) Reset()         { *m = UpdateAutomationRequest{} }
func (m *UpdateAutomationRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAutomationRequest) ProtoMessage()    {}
func (*UpdateAutomationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5791182e45212fe6, []int{7}
}

func (m *UpdateAutomationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAutomationRequest.Unmarshal(m, b)
}
func (m *UpdateAutomationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAutomationRequest.Marshal(b, m, deterministic)
}
func (m *UpdateAutomationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAutomationRequest.Merge(m, src)
}
func (m *UpdateAutomationRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAutomationRequest.Size(m)
}
func (m *UpdateAutomationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAutomationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAutomationRequest proto.InternalMessageInfo

func (m *UpdateAutomationRequest) GetAutomation() *Automation {
	if m != nil {
		return m.Automation
	}
	return nil
}

type DeleteAutomationRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteAutomationRequest) Reset()         { *m = DeleteAutomationRequest{} }
func (m *DeleteAutomationRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteAutomationRequest) ProtoMessage()    {}
func (*DeleteAutomationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5791182e45212fe6, []int{8}
}

func (m *DeleteAutomationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteAutomationRequest.Unmarshal(m, b)
}
func (m *DeleteAutomationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteAutomationRequest.Marshal(b, m, deterministic)
}
func (m *DeleteAutomationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteAutomationRequest.Merge(m, src)
}
func (m *DeleteAutomationRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteAutomationRequest.Size(m)
}
func (m *DeleteAutomationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteAutomationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteAutomationRequest proto.InternalMessageInfo

func (m *DeleteAutomationRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Automation)(nil), "api.Automation")
	proto.RegisterType((*CreateAutomationRequest)(nil), "api.CreateAutomationRequest")
	proto.RegisterType((*CreateAutomationResponse)(nil), "api.CreateAutomationResponse")
	proto.RegisterType((*GetAutomationRequest)(nil), "api.GetAutomationRequest")
	proto.RegisterType((*GetAutomationResponse)(nil), "api.GetAutomationResponse")
	proto.RegisterType((*ListAutomationRequest)(nil), "api.ListAutomationRequest")
	proto.RegisterType((*ListAutomationResponse)(nil), "api.ListAutomationResponse")
	proto.RegisterType((*UpdateAutomationRequest)(nil), "api.UpdateAutomationRequest")
	proto.RegisterType((*DeleteAutomationRequest)(nil), "api.DeleteAutomationRequest")
}

func init() {
	proto.RegisterFile("as/external/api/automation.proto", fileDescriptor_5791182e45212fe6)
}

var fileDescriptor_5791182e45212fe6 = []byte{
	// 709 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xcf, 0x6e, 0xd3, 0x4e,
	0x10, 0xc7, 0x95, 0xa4, 0x3f, 0xf7, 0x97, 0x49, 0xff, 0xd0, 0x55, 0x9a, 0x18, 0x37, 0xa8, 0x21,
	0x20, 0x28, 0x15, 0xd8, 0x50, 0xc4, 0x01, 0x0e, 0x48, 0x85, 0x56, 0xa5, 0x14, 0x71, 0x48, 0x8b,
	0x84, 0x2a, 0xa4, 0x68, 0xe3, 0x1d, 0xd2, 0x55, 0x6b, 0xaf, 0xf1, 0xae, 0x23, 0x5a, 0xd4, 0x0b,
	0xaf, 0xc0, 0x85, 0xf7, 0xe2, 0x15, 0x78, 0x0a, 0x4e, 0xc8, 0x6b, 0xa7, 0x71, 0x1d, 0x07, 0x04,
	0xe2, 0xb8, 0x3b, 0x9f, 0xfd, 0xce, 0xec, 0x77, 0xc6, 0x5e, 0x68, 0x53, 0xe9, 0xe0, 0x47, 0x85,
	0xa1, 0x4f, 0x4f, 0x1c, 0x1a, 0x70, 0x87, 0x46, 0x4a, 0x78, 0x54, 0x71, 0xe1, 0xdb, 0x41, 0x28,
	0x94, 0x20, 0x15, 0x1a, 0x70, 0xab, 0x35, 0x10, 0x62, 0x70, 0x82, 0x09, 0xe1, 0xfb, 0x42, 0x69,
	0x42, 0x26, 0x88, 0xb5, 0x9a, 0x46, 0xf5, 0xaa, 0x1f, 0xbd, 0x77, 0x14, 0xf7, 0x50, 0x2a, 0xea,
	0x05, 0x29, 0xb0, 0x92, 0x07, 0xd0, 0x0b, 0xd4, 0x69, 0x12, 0xec, 0x7c, 0x9d, 0x01, 0xd8, 0xbc,
	0xc8, 0x4a, 0x16, 0xa0, 0xcc, 0x99, 0x59, 0x6a, 0x97, 0xd6, 0x2a, 0xdd, 0x32, 0x67, 0xe4, 0x06,
	0xcc, 0x4b, 0xf4, 0x19, 0x86, 0x3d, 0x89, 0xbe, 0x14, 0xa1, 0x59, 0x6e, 0x97, 0xd6, 0xaa, 0xdd,
	0xb9, 0x64, 0x73, 0x5f, 0xef, 0x91, 0xdb, 0xb0, 0x18, 0xa2, 0x8b, 0x7c, 0x38, 0xc6, 0x2a, 0x1a,
	0x5b, 0x18, 0x6d, 0xa7, 0x60, 0x0b, 0xaa, 0xae, 0xf0, 0x19, 0x8f, 0x53, 0x99, 0x33, 0x1a, 0x19,
	0x6f, 0x90, 0x06, 0x18, 0xd4, 0xd5, 0xa1, 0xff, 0x74, 0x28, 0x5d, 0x91, 0xc7, 0x00, 0x6e, 0x88,
	0x54, 0x21, 0xeb, 0x51, 0x65, 0x1a, 0xed, 0xd2, 0x5a, 0x6d, 0xc3, 0xb2, 0x93, 0x4b, 0xd9, 0xa3,
	0x4b, 0xd9, 0x07, 0xa3, 0x5b, 0x77, 0xab, 0x29, 0xbd, 0xa9, 0xe2, 0xa3, 0x51, 0xc0, 0x46, 0x47,
	0x67, 0x7f, 0x7f, 0x34, 0xa5, 0x37, 0x15, 0x69, 0xc2, 0x6c, 0x24, 0x31, 0xec, 0x71, 0x66, 0xfe,
	0xaf, 0xed, 0x30, 0xe2, 0xe5, 0x2e, 0x8b, 0x6f, 0x2b, 0xc2, 0x01, 0xf5, 0xf9, 0x99, 0xb6, 0x2c,
	0x06, 0xaa, 0x1a, 0x58, 0xc8, 0x6e, 0xef, 0x32, 0x72, 0x17, 0x48, 0xea, 0x1d, 0xc3, 0x21, 0x77,
	0xb1, 0xa7, 0x4e, 0x03, 0x34, 0x41, 0xb3, 0x57, 0x92, 0xc8, 0x96, 0x0e, 0x1c, 0x9c, 0x06, 0x48,
	0xee, 0x43, 0xfd, 0xc2, 0xc4, 0x2c, 0x5f, 0xd3, 0x3c, 0x19, 0xc5, 0x32, 0x27, 0x26, 0xf4, 0x7d,
	0xea, 0xa1, 0x39, 0xa7, 0xbd, 0xbb, 0xa4, 0xff, 0x9a, 0x7a, 0x85, 0xfa, 0x9a, 0x9f, 0xd7, 0x7c,
	0x4e, 0x3f, 0x3e, 0xd1, 0x79, 0x09, 0xcd, 0xe7, 0xda, 0xc9, 0xf1, 0x7c, 0x74, 0xf1, 0x43, 0x84,
	0x52, 0x11, 0x07, 0x60, 0x3c, 0xaa, 0x7a, 0x5c, 0x6a, 0x1b, 0x8b, 0x36, 0x0d, 0xb8, 0x9d, 0x61,
	0x33, 0x48, 0x67, 0x0f, 0xcc, 0x49, 0x2d, 0x19, 0x08, 0x5f, 0xe2, 0x9f, 0x8b, 0xdd, 0x82, 0xfa,
	0x0e, 0xaa, 0xc9, 0xaa, 0x72, 0xc3, 0xdb, 0x79, 0x01, 0xcb, 0x39, 0xee, 0x6f, 0x33, 0xee, 0xc2,
	0xf2, 0x2b, 0x2e, 0x0b, 0x52, 0x66, 0xa6, 0xa4, 0x74, 0x69, 0x4a, 0x9a, 0x30, 0xcb, 0x70, 0xd8,
	0xc3, 0x88, 0xa7, 0x9f, 0x8c, 0xc1, 0x70, 0xb8, 0x1d, 0xf1, 0xce, 0x1e, 0x34, 0xf2, 0x52, 0x69,
	0x55, 0x0f, 0xa0, 0x36, 0x4e, 0x29, 0xcd, 0x52, 0xbb, 0x52, 0x54, 0x56, 0x96, 0x89, 0x5b, 0xf4,
	0x46, 0x4f, 0xec, 0x3f, 0x68, 0xd1, 0x1d, 0x68, 0x6e, 0xe1, 0x09, 0x16, 0x69, 0xe5, 0x8c, 0xdd,
	0xf8, 0x51, 0x81, 0xa5, 0x31, 0xb5, 0x8f, 0x61, 0x3c, 0x33, 0x84, 0x82, 0x91, 0xf4, 0x98, 0xb4,
	0x74, 0x9e, 0x29, 0xc3, 0x63, 0x5d, 0x9b, 0x12, 0x4d, 0x6c, 0xe8, 0x58, 0x9f, 0xbf, 0x7d, 0xff,
	0x52, 0xae, 0x77, 0x16, 0x73, 0x7f, 0xc4, 0x27, 0xa5, 0x75, 0x72, 0x08, 0x95, 0x1d, 0x54, 0xe4,
	0xaa, 0x56, 0x28, 0x9a, 0x01, 0xcb, 0x2a, 0x0a, 0xa5, 0xca, 0x2d, 0xad, 0xdc, 0x20, 0xf5, 0x9c,
	0xb2, 0xf3, 0x89, 0xb3, 0x73, 0xf2, 0x16, 0x66, 0xe2, 0xc6, 0x90, 0x44, 0xa1, 0xb0, 0xdd, 0xd6,
	0x4a, 0x61, 0x2c, 0x95, 0x6f, 0x6a, 0xf9, 0x25, 0x92, 0x2f, 0x9c, 0xf8, 0x60, 0x24, 0x5d, 0x4a,
	0x8d, 0x99, 0xd2, 0xb2, 0x5f, 0xd6, 0xbe, 0xae, 0xc5, 0x6f, 0x5a, 0xab, 0x13, 0xb5, 0x67, 0xde,
	0x0c, 0xce, 0xce, 0x63, 0x97, 0xde, 0x81, 0x91, 0x74, 0x32, 0xcd, 0x37, 0xa5, 0xad, 0x56, 0x63,
	0xe2, 0x4f, 0xb8, 0x1d, 0xbf, 0x0c, 0x23, 0x9f, 0xd6, 0x0b, 0x7d, 0x7a, 0xe6, 0xc2, 0x75, 0x2e,
	0x6c, 0xf7, 0x88, 0x87, 0x81, 0x54, 0xd4, 0x3d, 0xd6, 0x29, 0xa8, 0xb4, 0x47, 0x0f, 0x59, 0xbc,
	0x3e, 0x7c, 0x3a, 0xe0, 0xea, 0x28, 0xea, 0xdb, 0xae, 0xf0, 0x1c, 0xde, 0x0f, 0xe9, 0x11, 0xf7,
	0xc4, 0x19, 0x1e, 0x73, 0x97, 0x3b, 0xe3, 0x93, 0xf7, 0x62, 0xf9, 0x81, 0x70, 0x86, 0x8f, 0x9c,
	0xdc, 0x43, 0xd8, 0x37, 0x74, 0x49, 0x0f, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xc3, 0x32,
	0x87, 0x22, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AutomationServiceClient is the client API for AutomationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AutomationServiceClient interface {
	Create(ctx context.Context, in *CreateAutomationRequest, opts ...grpc.CallOption) (*CreateAutomationResponse, error)
	// Get return the automation
	Get(ctx context.Context, in *GetAutomationRequest, opts ...grpc.CallOption) (*GetAutomationResponse, error)
	List(ctx context.Context, in *ListAutomationRequest, opts ...grpc.CallOption) (*ListAutomationResponse, error)
	// Update updates the automation matching the given automation_id.
	Update(ctx context.Context, in *UpdateAutomationRequest, opts ...grpc.CallOption) (*GetAutomationResponse, error)
	// Delete deletes the automation matching the given automation_id.
	Delete(ctx context.Context, in *DeleteAutomationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type automationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAutomationServiceClient(cc grpc.ClientConnInterface) AutomationServiceClient {
	return &automationServiceClient{cc}
}

func (c *automationServiceClient) Create(ctx context.Context, in *CreateAutomationRequest, opts ...grpc.CallOption) (*CreateAutomationResponse, error) {
	out := new(CreateAutomationResponse)
	err := c.cc.Invoke(ctx, "/api.AutomationService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationServiceClient) Get(ctx context.Context, in *GetAutomationRequest, opts ...grpc.CallOption) (*GetAutomationResponse, error) {
	out := new(GetAutomationResponse)
	err := c.cc.Invoke(ctx, "/api.AutomationService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationServiceClient) List(ctx context.Context, in *ListAutomationRequest, opts ...grpc.CallOption) (*ListAutomationResponse, error) {
	out := new(ListAutomationResponse)
	err := c.cc.Invoke(ctx, "/api.AutomationService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationServiceClient) Update(ctx context.Context, in *UpdateAutomationRequest, opts ...grpc.CallOption) (*GetAutomationResponse, error) {
	out := new(GetAutomationResponse)
	err := c.cc.Invoke(ctx, "/api.AutomationService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *automationServiceClient) Delete(ctx context.Context, in *DeleteAutomationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.AutomationService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AutomationServiceServer is the server API for AutomationService service.
type AutomationServiceServer interface {
	Create(context.Context, *CreateAutomationRequest) (*CreateAutomationResponse, error)
	// Get return the automation
	Get(context.Context, *GetAutomationRequest) (*GetAutomationResponse, error)
	List(context.Context, *ListAutomationRequest) (*ListAutomationResponse, error)
	// Update updates the automation matching the given automation_id.
	Update(context.Context, *UpdateAutomationRequest) (*GetAutomationResponse, error)
	// Delete deletes the automation matching the given automation_id.
	Delete(context.Context, *DeleteAutomationRequest) (*empty.Empty, error)
}

// UnimplementedAutomationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAutomationServiceServer struct {
}

func (*UnimplementedAutomationServiceServer) Create(ctx context.Context, req *CreateAutomationRequest) (*CreateAutomationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAutomationServiceServer) Get(ctx context.Context, req *GetAutomationRequest) (*GetAutomationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAutomationServiceServer) List(ctx context.Context, req *ListAutomationRequest) (*ListAutomationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedAutomationServiceServer) Update(ctx context.Context, req *UpdateAutomationRequest) (*GetAutomationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedAutomationServiceServer) Delete(ctx context.Context, req *DeleteAutomationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterAutomationServiceServer(s *grpc.Server, srv AutomationServiceServer) {
	s.RegisterService(&_AutomationService_serviceDesc, srv)
}

func _AutomationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAutomationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AutomationService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationServiceServer).Create(ctx, req.(*CreateAutomationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAutomationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AutomationService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationServiceServer).Get(ctx, req.(*GetAutomationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomationService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAutomationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AutomationService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationServiceServer).List(ctx, req.(*ListAutomationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomationService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAutomationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AutomationService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationServiceServer).Update(ctx, req.(*UpdateAutomationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AutomationService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAutomationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AutomationServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.AutomationService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AutomationServiceServer).Delete(ctx, req.(*DeleteAutomationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AutomationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.AutomationService",
	HandlerType: (*AutomationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AutomationService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AutomationService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AutomationService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AutomationService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AutomationService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as/external/api/automation.proto",
}