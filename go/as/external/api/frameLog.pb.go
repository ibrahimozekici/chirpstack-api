// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/api/frameLog.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/ibrahimozekici/chirpstack-api/go/v5/common"
	gw "github.com/ibrahimozekici/chirpstack-api/go/v5/gw"
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

type RXWindow int32

const (
	RXWindow_RX1 RXWindow = 0
	RXWindow_RX2 RXWindow = 1
)

var RXWindow_name = map[int32]string{
	0: "RX1",
	1: "RX2",
}

var RXWindow_value = map[string]int32{
	"RX1": 0,
	"RX2": 1,
}

func (x RXWindow) String() string {
	return proto.EnumName(RXWindow_name, int32(x))
}

func (RXWindow) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d215818867a60801, []int{0}
}

type UplinkFrameLog struct {
	// TX information of the uplink.
	TxInfo *gw.UplinkTXInfo `protobuf:"bytes,1,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// RX information of the uplink.
	RxInfo []*gw.UplinkRXInfo `protobuf:"bytes,2,rep,name=rx_info,json=rxInfo,proto3" json:"rx_info,omitempty"`
	// LoRaWAN PHYPayload.
	PhyPayloadJson string `protobuf:"bytes,3,opt,name=phy_payload_json,json=phyPayloadJSON,proto3" json:"phy_payload_json,omitempty"`
	// Published at timestamp.
	PublishedAt          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UplinkFrameLog) Reset()         { *m = UplinkFrameLog{} }
func (m *UplinkFrameLog) String() string { return proto.CompactTextString(m) }
func (*UplinkFrameLog) ProtoMessage()    {}
func (*UplinkFrameLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_d215818867a60801, []int{0}
}

func (m *UplinkFrameLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UplinkFrameLog.Unmarshal(m, b)
}
func (m *UplinkFrameLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UplinkFrameLog.Marshal(b, m, deterministic)
}
func (m *UplinkFrameLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UplinkFrameLog.Merge(m, src)
}
func (m *UplinkFrameLog) XXX_Size() int {
	return xxx_messageInfo_UplinkFrameLog.Size(m)
}
func (m *UplinkFrameLog) XXX_DiscardUnknown() {
	xxx_messageInfo_UplinkFrameLog.DiscardUnknown(m)
}

var xxx_messageInfo_UplinkFrameLog proto.InternalMessageInfo

func (m *UplinkFrameLog) GetTxInfo() *gw.UplinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *UplinkFrameLog) GetRxInfo() []*gw.UplinkRXInfo {
	if m != nil {
		return m.RxInfo
	}
	return nil
}

func (m *UplinkFrameLog) GetPhyPayloadJson() string {
	if m != nil {
		return m.PhyPayloadJson
	}
	return ""
}

func (m *UplinkFrameLog) GetPublishedAt() *timestamp.Timestamp {
	if m != nil {
		return m.PublishedAt
	}
	return nil
}

type DownlinkFrameLog struct {
	// TX information of the downlink.
	TxInfo *gw.DownlinkTXInfo `protobuf:"bytes,1,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// LoRaWAN PHYPayload.
	PhyPayloadJson string `protobuf:"bytes,2,opt,name=phy_payload_json,json=phyPayloadJSON,proto3" json:"phy_payload_json,omitempty"`
	// Gateway ID.
	GatewayId string `protobuf:"bytes,3,opt,name=gateway_id,json=gatewayID,proto3" json:"gateway_id,omitempty"`
	// Published at timestamp.
	PublishedAt          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=published_at,json=publishedAt,proto3" json:"published_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DownlinkFrameLog) Reset()         { *m = DownlinkFrameLog{} }
func (m *DownlinkFrameLog) String() string { return proto.CompactTextString(m) }
func (*DownlinkFrameLog) ProtoMessage()    {}
func (*DownlinkFrameLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_d215818867a60801, []int{1}
}

func (m *DownlinkFrameLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownlinkFrameLog.Unmarshal(m, b)
}
func (m *DownlinkFrameLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownlinkFrameLog.Marshal(b, m, deterministic)
}
func (m *DownlinkFrameLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownlinkFrameLog.Merge(m, src)
}
func (m *DownlinkFrameLog) XXX_Size() int {
	return xxx_messageInfo_DownlinkFrameLog.Size(m)
}
func (m *DownlinkFrameLog) XXX_DiscardUnknown() {
	xxx_messageInfo_DownlinkFrameLog.DiscardUnknown(m)
}

var xxx_messageInfo_DownlinkFrameLog proto.InternalMessageInfo

func (m *DownlinkFrameLog) GetTxInfo() *gw.DownlinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *DownlinkFrameLog) GetPhyPayloadJson() string {
	if m != nil {
		return m.PhyPayloadJson
	}
	return ""
}

func (m *DownlinkFrameLog) GetGatewayId() string {
	if m != nil {
		return m.GatewayId
	}
	return ""
}

func (m *DownlinkFrameLog) GetPublishedAt() *timestamp.Timestamp {
	if m != nil {
		return m.PublishedAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.RXWindow", RXWindow_name, RXWindow_value)
	proto.RegisterType((*UplinkFrameLog)(nil), "api.UplinkFrameLog")
	proto.RegisterType((*DownlinkFrameLog)(nil), "api.DownlinkFrameLog")
}

func init() {
	proto.RegisterFile("as/external/api/frameLog.proto", fileDescriptor_d215818867a60801)
}

var fileDescriptor_d215818867a60801 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0xc9, 0x8a, 0x36, 0xe6, 0xc2, 0x14, 0x85, 0x4b, 0x55, 0xc1, 0x28, 0x3b, 0x15, 0x10,
	0xb6, 0x18, 0xe2, 0x08, 0x12, 0xd3, 0x84, 0x34, 0x84, 0xa0, 0x0a, 0x43, 0x54, 0x5c, 0xa2, 0x97,
	0xc4, 0x75, 0x1e, 0x4d, 0xfc, 0x2c, 0xc7, 0x25, 0x2d, 0xff, 0x20, 0x47, 0xfe, 0x25, 0x94, 0x3a,
	0x01, 0xad, 0xea, 0x8d, 0xd3, 0xf3, 0xfb, 0xfc, 0x93, 0xfd, 0x7d, 0x4f, 0x8f, 0x9d, 0x42, 0x2d,
	0xe4, 0xda, 0x49, 0xab, 0xa1, 0x14, 0x60, 0x50, 0x2c, 0x2c, 0x54, 0xf2, 0x03, 0x29, 0x6e, 0x2c,
	0x39, 0x8a, 0x06, 0x60, 0x70, 0xfc, 0x48, 0x11, 0xa9, 0x52, 0x8a, 0xad, 0x94, 0xae, 0x16, 0xc2,
	0x61, 0x25, 0x6b, 0x07, 0x95, 0xf1, 0xd4, 0xf8, 0x74, 0x17, 0xc8, 0x57, 0x16, 0x1c, 0x92, 0xee,
	0xee, 0xef, 0x67, 0x54, 0x55, 0xa4, 0x85, 0x2f, 0x9d, 0x38, 0x54, 0x8d, 0x50, 0x8d, 0x6f, 0xce,
	0x7e, 0x07, 0xec, 0xe4, 0x8b, 0x29, 0x51, 0x2f, 0xdf, 0x75, 0x06, 0xa2, 0x27, 0xec, 0xc8, 0xad,
	0x13, 0xd4, 0x0b, 0x1a, 0x05, 0x93, 0x60, 0x3a, 0x3c, 0x0f, 0xb9, 0x6a, 0xb8, 0x87, 0xae, 0xe7,
	0x57, 0x7a, 0x41, 0xf1, 0xa1, 0x5b, 0xb7, 0xb5, 0x45, 0x6d, 0x87, 0x1e, 0x4c, 0x06, 0x37, 0xd1,
	0xb8, 0x43, 0xad, 0x47, 0xa7, 0x2c, 0x34, 0xc5, 0x26, 0x31, 0xb0, 0x29, 0x09, 0xf2, 0xe4, 0x7b,
	0x4d, 0x7a, 0x34, 0x98, 0x04, 0xd3, 0xe3, 0xf8, 0xc4, 0x14, 0x9b, 0x99, 0x97, 0xdf, 0x7f, 0xfe,
	0xf4, 0x31, 0x7a, 0xcd, 0xee, 0x9a, 0x55, 0x5a, 0x62, 0x5d, 0xc8, 0x3c, 0x01, 0x37, 0xba, 0xbd,
	0x35, 0x31, 0xe6, 0x3e, 0x2b, 0xef, 0xb3, 0xf2, 0xeb, 0x7e, 0x18, 0xf1, 0xf0, 0x2f, 0xff, 0xd6,
	0x9d, 0xfd, 0x0a, 0x58, 0x78, 0x49, 0x8d, 0xbe, 0x91, 0xe9, 0xd9, 0x6e, 0xa6, 0xa8, 0x35, 0xda,
	0x63, 0x3b, 0xa9, 0xf6, 0x59, 0x3d, 0xd8, 0x6b, 0xf5, 0x21, 0x63, 0x0a, 0x9c, 0x6c, 0x60, 0x93,
	0x60, 0xde, 0xc5, 0x39, 0xee, 0x94, 0xab, 0xcb, 0xff, 0x4c, 0xf2, 0xf4, 0x01, 0xbb, 0x13, 0xcf,
	0xbf, 0xa2, 0xce, 0xa9, 0x89, 0x8e, 0xd8, 0x20, 0x9e, 0xbf, 0x08, 0x6f, 0xf9, 0xc3, 0x79, 0x18,
	0x5c, 0x38, 0xf6, 0x18, 0x89, 0x67, 0x05, 0x5a, 0x53, 0x3b, 0xc8, 0x96, 0x1c, 0x0c, 0x72, 0xa8,
	0x79, 0xbf, 0x55, 0x6d, 0x7f, 0x71, 0xaf, 0x9f, 0xc0, 0xac, 0xfd, 0x6b, 0x16, 0x7c, 0x7b, 0xa3,
	0xd0, 0x15, 0xab, 0x94, 0x67, 0x54, 0x09, 0x4c, 0x2d, 0x14, 0x58, 0xd1, 0x4f, 0xb9, 0xc4, 0x0c,
	0xc5, 0xbf, 0xa7, 0x9e, 0xb7, 0x0b, 0xa9, 0x48, 0xfc, 0x78, 0x25, 0x76, 0xd6, 0x34, 0x3d, 0xdc,
	0x9a, 0x7e, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x18, 0x5a, 0x50, 0x3d, 0xc0, 0x02, 0x00, 0x00,
}
