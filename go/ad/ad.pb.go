// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ad/ad.proto

package ad

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// The request message containing the sensor ID and temperature.
type AnomalyRequest struct {
	DevEui               string   `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	Temperature          float32  `protobuf:"fixed32,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnomalyRequest) Reset()         { *m = AnomalyRequest{} }
func (m *AnomalyRequest) String() string { return proto.CompactTextString(m) }
func (*AnomalyRequest) ProtoMessage()    {}
func (*AnomalyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd87d7d40ec3472e, []int{0}
}

func (m *AnomalyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnomalyRequest.Unmarshal(m, b)
}
func (m *AnomalyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnomalyRequest.Marshal(b, m, deterministic)
}
func (m *AnomalyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnomalyRequest.Merge(m, src)
}
func (m *AnomalyRequest) XXX_Size() int {
	return xxx_messageInfo_AnomalyRequest.Size(m)
}
func (m *AnomalyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnomalyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnomalyRequest proto.InternalMessageInfo

func (m *AnomalyRequest) GetDevEui() string {
	if m != nil {
		return m.DevEui
	}
	return ""
}

func (m *AnomalyRequest) GetTemperature() float32 {
	if m != nil {
		return m.Temperature
	}
	return 0
}

// The response message containing the result of the anomaly detection.
type AnomalyResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnomalyResponse) Reset()         { *m = AnomalyResponse{} }
func (m *AnomalyResponse) String() string { return proto.CompactTextString(m) }
func (*AnomalyResponse) ProtoMessage()    {}
func (*AnomalyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd87d7d40ec3472e, []int{1}
}

func (m *AnomalyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnomalyResponse.Unmarshal(m, b)
}
func (m *AnomalyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnomalyResponse.Marshal(b, m, deterministic)
}
func (m *AnomalyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnomalyResponse.Merge(m, src)
}
func (m *AnomalyResponse) XXX_Size() int {
	return xxx_messageInfo_AnomalyResponse.Size(m)
}
func (m *AnomalyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnomalyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnomalyResponse proto.InternalMessageInfo

func (m *AnomalyResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*AnomalyRequest)(nil), "ad.AnomalyRequest")
	proto.RegisterType((*AnomalyResponse)(nil), "ad.AnomalyResponse")
}

func init() {
	proto.RegisterFile("ad/ad.proto", fileDescriptor_bd87d7d40ec3472e)
}

var fileDescriptor_bd87d7d40ec3472e = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0x4c, 0xd1, 0x4f,
	0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4a, 0x4c, 0x51, 0xf2, 0xe6, 0xe2, 0x73,
	0xcc, 0xcb, 0xcf, 0x4d, 0xcc, 0xa9, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe7,
	0x62, 0x4f, 0x49, 0x2d, 0x8b, 0x4f, 0x2d, 0xcd, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62,
	0x4b, 0x49, 0x2d, 0x73, 0x2d, 0xcd, 0x14, 0x52, 0xe0, 0xe2, 0x2e, 0x49, 0xcd, 0x2d, 0x48, 0x2d,
	0x4a, 0x2c, 0x29, 0x2d, 0x4a, 0x95, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x0a, 0x42, 0x16, 0x52, 0xd2,
	0xe4, 0xe2, 0x87, 0x1b, 0x56, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc6, 0xc5, 0x56, 0x94,
	0x5a, 0x5c, 0x9a, 0x53, 0x02, 0x33, 0x0c, 0xc2, 0x33, 0xf2, 0xe7, 0x12, 0x80, 0x2a, 0x75, 0x49,
	0x2d, 0x49, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0x13, 0xb2, 0xe6, 0x12, 0x08, 0x29, 0xca, 0x4c, 0x4f,
	0x4f, 0x2d, 0x42, 0x88, 0x09, 0xe9, 0x25, 0xa6, 0xe8, 0xa1, 0xba, 0x50, 0x4a, 0x18, 0x45, 0x0c,
	0x62, 0x51, 0x12, 0x1b, 0xd8, 0x4f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x2f, 0xe7,
	0x22, 0xe2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AnomalyDetectionClient is the client API for AnomalyDetection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnomalyDetectionClient interface {
	// This RPC triggers the anomaly detection process.
	TriggerDetection(ctx context.Context, in *AnomalyRequest, opts ...grpc.CallOption) (*AnomalyResponse, error)
}

type anomalyDetectionClient struct {
	cc grpc.ClientConnInterface
}

func NewAnomalyDetectionClient(cc grpc.ClientConnInterface) AnomalyDetectionClient {
	return &anomalyDetectionClient{cc}
}

func (c *anomalyDetectionClient) TriggerDetection(ctx context.Context, in *AnomalyRequest, opts ...grpc.CallOption) (*AnomalyResponse, error) {
	out := new(AnomalyResponse)
	err := c.cc.Invoke(ctx, "/ad.AnomalyDetection/TriggerDetection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnomalyDetectionServer is the server API for AnomalyDetection service.
type AnomalyDetectionServer interface {
	// This RPC triggers the anomaly detection process.
	TriggerDetection(context.Context, *AnomalyRequest) (*AnomalyResponse, error)
}

// UnimplementedAnomalyDetectionServer can be embedded to have forward compatible implementations.
type UnimplementedAnomalyDetectionServer struct {
}

func (*UnimplementedAnomalyDetectionServer) TriggerDetection(ctx context.Context, req *AnomalyRequest) (*AnomalyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerDetection not implemented")
}

func RegisterAnomalyDetectionServer(s *grpc.Server, srv AnomalyDetectionServer) {
	s.RegisterService(&_AnomalyDetection_serviceDesc, srv)
}

func _AnomalyDetection_TriggerDetection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnomalyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyDetectionServer).TriggerDetection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ad.AnomalyDetection/TriggerDetection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyDetectionServer).TriggerDetection(ctx, req.(*AnomalyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AnomalyDetection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ad.AnomalyDetection",
	HandlerType: (*AnomalyDetectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TriggerDetection",
			Handler:    _AnomalyDetection_TriggerDetection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ad/ad.proto",
}
