// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/api/gatewayProfile.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	duration "google.golang.org/protobuf/ptypes/duration"
	empty "google.golang.org/protobuf/ptypes/empty"
	timestamp "google.golang.org/protobuf/ptypes/timestamp"
	common "github.com/ibrahimozekici/chirpstack-api/go/v5/common"
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

type GatewayProfile struct {
	// Gateway-profile ID (UUID string).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the gateway-profile.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Network-server ID of the gateway-profile.
	NetworkServerId int64 `protobuf:"varint,3,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// Default channels (channels specified by the LoRaWAN Regional Parameters
	// specification) enabled for this configuration.
	Channels []uint32 `protobuf:"varint,4,rep,packed,name=channels,proto3" json:"channels,omitempty"`
	// Extra channels added to the channel-configuration (in case the LoRaWAN
	// region supports adding custom channels).
	ExtraChannels []*GatewayProfileExtraChannel `protobuf:"bytes,5,rep,name=extra_channels,json=extraChannels,proto3" json:"extra_channels,omitempty"`
	// Stats interval.
	// This defines the (expected) stats interval which the gateways using this
	// gateway-profile are using.
	StatsInterval        *duration.Duration `protobuf:"bytes,6,opt,name=stats_interval,json=statsInterval,proto3" json:"stats_interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GatewayProfile) Reset()         { *m = GatewayProfile{} }
func (m *GatewayProfile) String() string { return proto.CompactTextString(m) }
func (*GatewayProfile) ProtoMessage()    {}
func (*GatewayProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{0}
}

func (m *GatewayProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayProfile.Unmarshal(m, b)
}
func (m *GatewayProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayProfile.Marshal(b, m, deterministic)
}
func (m *GatewayProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayProfile.Merge(m, src)
}
func (m *GatewayProfile) XXX_Size() int {
	return xxx_messageInfo_GatewayProfile.Size(m)
}
func (m *GatewayProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayProfile.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayProfile proto.InternalMessageInfo

func (m *GatewayProfile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GatewayProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GatewayProfile) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

func (m *GatewayProfile) GetChannels() []uint32 {
	if m != nil {
		return m.Channels
	}
	return nil
}

func (m *GatewayProfile) GetExtraChannels() []*GatewayProfileExtraChannel {
	if m != nil {
		return m.ExtraChannels
	}
	return nil
}

func (m *GatewayProfile) GetStatsInterval() *duration.Duration {
	if m != nil {
		return m.StatsInterval
	}
	return nil
}

type GatewayProfileListItem struct {
	// Gateway-profile ID (UUID string).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Gateway-profile name,
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Network-server ID on which the gateway-profile is provisioned.
	NetworkServerId int64 `protobuf:"varint,3,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// Network-server name.
	NetworkServerName string `protobuf:"bytes,7,opt,name=network_server_name,json=networkServerName,proto3" json:"network_server_name,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GatewayProfileListItem) Reset()         { *m = GatewayProfileListItem{} }
func (m *GatewayProfileListItem) String() string { return proto.CompactTextString(m) }
func (*GatewayProfileListItem) ProtoMessage()    {}
func (*GatewayProfileListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{1}
}

func (m *GatewayProfileListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayProfileListItem.Unmarshal(m, b)
}
func (m *GatewayProfileListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayProfileListItem.Marshal(b, m, deterministic)
}
func (m *GatewayProfileListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayProfileListItem.Merge(m, src)
}
func (m *GatewayProfileListItem) XXX_Size() int {
	return xxx_messageInfo_GatewayProfileListItem.Size(m)
}
func (m *GatewayProfileListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayProfileListItem.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayProfileListItem proto.InternalMessageInfo

func (m *GatewayProfileListItem) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GatewayProfileListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GatewayProfileListItem) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

func (m *GatewayProfileListItem) GetNetworkServerName() string {
	if m != nil {
		return m.NetworkServerName
	}
	return ""
}

func (m *GatewayProfileListItem) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GatewayProfileListItem) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type GatewayProfileExtraChannel struct {
	// Modulation.
	Modulation common.Modulation `protobuf:"varint,1,opt,name=modulation,proto3,enum=common.Modulation" json:"modulation,omitempty"`
	// Frequency.
	Frequency uint32 `protobuf:"varint,2,opt,name=frequency,proto3" json:"frequency,omitempty"`
	// Bandwidth.
	Bandwidth uint32 `protobuf:"varint,3,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`
	// Bitrate (in case of FSK modulation).
	Bitrate uint32 `protobuf:"varint,4,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	// Spreading factors (in case of LoRa modulation).
	SpreadingFactors     []uint32 `protobuf:"varint,5,rep,packed,name=spreading_factors,json=spreadingFactors,proto3" json:"spreading_factors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayProfileExtraChannel) Reset()         { *m = GatewayProfileExtraChannel{} }
func (m *GatewayProfileExtraChannel) String() string { return proto.CompactTextString(m) }
func (*GatewayProfileExtraChannel) ProtoMessage()    {}
func (*GatewayProfileExtraChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{2}
}

func (m *GatewayProfileExtraChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayProfileExtraChannel.Unmarshal(m, b)
}
func (m *GatewayProfileExtraChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayProfileExtraChannel.Marshal(b, m, deterministic)
}
func (m *GatewayProfileExtraChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayProfileExtraChannel.Merge(m, src)
}
func (m *GatewayProfileExtraChannel) XXX_Size() int {
	return xxx_messageInfo_GatewayProfileExtraChannel.Size(m)
}
func (m *GatewayProfileExtraChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayProfileExtraChannel.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayProfileExtraChannel proto.InternalMessageInfo

func (m *GatewayProfileExtraChannel) GetModulation() common.Modulation {
	if m != nil {
		return m.Modulation
	}
	return common.Modulation_LORA
}

func (m *GatewayProfileExtraChannel) GetFrequency() uint32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

func (m *GatewayProfileExtraChannel) GetBandwidth() uint32 {
	if m != nil {
		return m.Bandwidth
	}
	return 0
}

func (m *GatewayProfileExtraChannel) GetBitrate() uint32 {
	if m != nil {
		return m.Bitrate
	}
	return 0
}

func (m *GatewayProfileExtraChannel) GetSpreadingFactors() []uint32 {
	if m != nil {
		return m.SpreadingFactors
	}
	return nil
}

type CreateGatewayProfileRequest struct {
	// Gateway-profile object to create.
	GatewayProfile       *GatewayProfile `protobuf:"bytes,1,opt,name=gateway_profile,json=gatewayProfile,proto3" json:"gateway_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateGatewayProfileRequest) Reset()         { *m = CreateGatewayProfileRequest{} }
func (m *CreateGatewayProfileRequest) String() string { return proto.CompactTextString(m) }
func (*CreateGatewayProfileRequest) ProtoMessage()    {}
func (*CreateGatewayProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{3}
}

func (m *CreateGatewayProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGatewayProfileRequest.Unmarshal(m, b)
}
func (m *CreateGatewayProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGatewayProfileRequest.Marshal(b, m, deterministic)
}
func (m *CreateGatewayProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGatewayProfileRequest.Merge(m, src)
}
func (m *CreateGatewayProfileRequest) XXX_Size() int {
	return xxx_messageInfo_CreateGatewayProfileRequest.Size(m)
}
func (m *CreateGatewayProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGatewayProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGatewayProfileRequest proto.InternalMessageInfo

func (m *CreateGatewayProfileRequest) GetGatewayProfile() *GatewayProfile {
	if m != nil {
		return m.GatewayProfile
	}
	return nil
}

type CreateGatewayProfileResponse struct {
	// Gateway-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateGatewayProfileResponse) Reset()         { *m = CreateGatewayProfileResponse{} }
func (m *CreateGatewayProfileResponse) String() string { return proto.CompactTextString(m) }
func (*CreateGatewayProfileResponse) ProtoMessage()    {}
func (*CreateGatewayProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{4}
}

func (m *CreateGatewayProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateGatewayProfileResponse.Unmarshal(m, b)
}
func (m *CreateGatewayProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateGatewayProfileResponse.Marshal(b, m, deterministic)
}
func (m *CreateGatewayProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateGatewayProfileResponse.Merge(m, src)
}
func (m *CreateGatewayProfileResponse) XXX_Size() int {
	return xxx_messageInfo_CreateGatewayProfileResponse.Size(m)
}
func (m *CreateGatewayProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateGatewayProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateGatewayProfileResponse proto.InternalMessageInfo

func (m *CreateGatewayProfileResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetGatewayProfileRequest struct {
	// Gateway-profile ID (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetGatewayProfileRequest) Reset()         { *m = GetGatewayProfileRequest{} }
func (m *GetGatewayProfileRequest) String() string { return proto.CompactTextString(m) }
func (*GetGatewayProfileRequest) ProtoMessage()    {}
func (*GetGatewayProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{5}
}

func (m *GetGatewayProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayProfileRequest.Unmarshal(m, b)
}
func (m *GetGatewayProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayProfileRequest.Marshal(b, m, deterministic)
}
func (m *GetGatewayProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayProfileRequest.Merge(m, src)
}
func (m *GetGatewayProfileRequest) XXX_Size() int {
	return xxx_messageInfo_GetGatewayProfileRequest.Size(m)
}
func (m *GetGatewayProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayProfileRequest proto.InternalMessageInfo

func (m *GetGatewayProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetGatewayProfileResponse struct {
	// Gateway-profile object.
	GatewayProfile *GatewayProfile `protobuf:"bytes,1,opt,name=gateway_profile,json=gatewayProfile,proto3" json:"gateway_profile,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetGatewayProfileResponse) Reset()         { *m = GetGatewayProfileResponse{} }
func (m *GetGatewayProfileResponse) String() string { return proto.CompactTextString(m) }
func (*GetGatewayProfileResponse) ProtoMessage()    {}
func (*GetGatewayProfileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{6}
}

func (m *GetGatewayProfileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetGatewayProfileResponse.Unmarshal(m, b)
}
func (m *GetGatewayProfileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetGatewayProfileResponse.Marshal(b, m, deterministic)
}
func (m *GetGatewayProfileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetGatewayProfileResponse.Merge(m, src)
}
func (m *GetGatewayProfileResponse) XXX_Size() int {
	return xxx_messageInfo_GetGatewayProfileResponse.Size(m)
}
func (m *GetGatewayProfileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetGatewayProfileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetGatewayProfileResponse proto.InternalMessageInfo

func (m *GetGatewayProfileResponse) GetGatewayProfile() *GatewayProfile {
	if m != nil {
		return m.GatewayProfile
	}
	return nil
}

func (m *GetGatewayProfileResponse) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *GetGatewayProfileResponse) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type UpdateGatewayProfileRequest struct {
	// Gateway-profile object to update.
	GatewayProfile       *GatewayProfile `protobuf:"bytes,1,opt,name=gateway_profile,json=gatewayProfile,proto3" json:"gateway_profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateGatewayProfileRequest) Reset()         { *m = UpdateGatewayProfileRequest{} }
func (m *UpdateGatewayProfileRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateGatewayProfileRequest) ProtoMessage()    {}
func (*UpdateGatewayProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{7}
}

func (m *UpdateGatewayProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateGatewayProfileRequest.Unmarshal(m, b)
}
func (m *UpdateGatewayProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateGatewayProfileRequest.Marshal(b, m, deterministic)
}
func (m *UpdateGatewayProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateGatewayProfileRequest.Merge(m, src)
}
func (m *UpdateGatewayProfileRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateGatewayProfileRequest.Size(m)
}
func (m *UpdateGatewayProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateGatewayProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateGatewayProfileRequest proto.InternalMessageInfo

func (m *UpdateGatewayProfileRequest) GetGatewayProfile() *GatewayProfile {
	if m != nil {
		return m.GatewayProfile
	}
	return nil
}

type DeleteGatewayProfileRequest struct {
	// Gateway-profile id (UUID string).
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteGatewayProfileRequest) Reset()         { *m = DeleteGatewayProfileRequest{} }
func (m *DeleteGatewayProfileRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteGatewayProfileRequest) ProtoMessage()    {}
func (*DeleteGatewayProfileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{8}
}

func (m *DeleteGatewayProfileRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteGatewayProfileRequest.Unmarshal(m, b)
}
func (m *DeleteGatewayProfileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteGatewayProfileRequest.Marshal(b, m, deterministic)
}
func (m *DeleteGatewayProfileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteGatewayProfileRequest.Merge(m, src)
}
func (m *DeleteGatewayProfileRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteGatewayProfileRequest.Size(m)
}
func (m *DeleteGatewayProfileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteGatewayProfileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteGatewayProfileRequest proto.InternalMessageInfo

func (m *DeleteGatewayProfileRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListGatewayProfilesRequest struct {
	// Max number of items to return.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Network-server ID to filter on (optional).
	NetworkServerId      int64    `protobuf:"varint,3,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListGatewayProfilesRequest) Reset()         { *m = ListGatewayProfilesRequest{} }
func (m *ListGatewayProfilesRequest) String() string { return proto.CompactTextString(m) }
func (*ListGatewayProfilesRequest) ProtoMessage()    {}
func (*ListGatewayProfilesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{9}
}

func (m *ListGatewayProfilesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGatewayProfilesRequest.Unmarshal(m, b)
}
func (m *ListGatewayProfilesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGatewayProfilesRequest.Marshal(b, m, deterministic)
}
func (m *ListGatewayProfilesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGatewayProfilesRequest.Merge(m, src)
}
func (m *ListGatewayProfilesRequest) XXX_Size() int {
	return xxx_messageInfo_ListGatewayProfilesRequest.Size(m)
}
func (m *ListGatewayProfilesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGatewayProfilesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListGatewayProfilesRequest proto.InternalMessageInfo

func (m *ListGatewayProfilesRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListGatewayProfilesRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListGatewayProfilesRequest) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

type ListGatewayProfilesResponse struct {
	// Total number of gateway-profiles.
	TotalCount           int64                     `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	Result               []*GatewayProfileListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ListGatewayProfilesResponse) Reset()         { *m = ListGatewayProfilesResponse{} }
func (m *ListGatewayProfilesResponse) String() string { return proto.CompactTextString(m) }
func (*ListGatewayProfilesResponse) ProtoMessage()    {}
func (*ListGatewayProfilesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a425cd68ff0462c4, []int{10}
}

func (m *ListGatewayProfilesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListGatewayProfilesResponse.Unmarshal(m, b)
}
func (m *ListGatewayProfilesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListGatewayProfilesResponse.Marshal(b, m, deterministic)
}
func (m *ListGatewayProfilesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListGatewayProfilesResponse.Merge(m, src)
}
func (m *ListGatewayProfilesResponse) XXX_Size() int {
	return xxx_messageInfo_ListGatewayProfilesResponse.Size(m)
}
func (m *ListGatewayProfilesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListGatewayProfilesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListGatewayProfilesResponse proto.InternalMessageInfo

func (m *ListGatewayProfilesResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListGatewayProfilesResponse) GetResult() []*GatewayProfileListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*GatewayProfile)(nil), "api.GatewayProfile")
	proto.RegisterType((*GatewayProfileListItem)(nil), "api.GatewayProfileListItem")
	proto.RegisterType((*GatewayProfileExtraChannel)(nil), "api.GatewayProfileExtraChannel")
	proto.RegisterType((*CreateGatewayProfileRequest)(nil), "api.CreateGatewayProfileRequest")
	proto.RegisterType((*CreateGatewayProfileResponse)(nil), "api.CreateGatewayProfileResponse")
	proto.RegisterType((*GetGatewayProfileRequest)(nil), "api.GetGatewayProfileRequest")
	proto.RegisterType((*GetGatewayProfileResponse)(nil), "api.GetGatewayProfileResponse")
	proto.RegisterType((*UpdateGatewayProfileRequest)(nil), "api.UpdateGatewayProfileRequest")
	proto.RegisterType((*DeleteGatewayProfileRequest)(nil), "api.DeleteGatewayProfileRequest")
	proto.RegisterType((*ListGatewayProfilesRequest)(nil), "api.ListGatewayProfilesRequest")
	proto.RegisterType((*ListGatewayProfilesResponse)(nil), "api.ListGatewayProfilesResponse")
}

func init() {
	proto.RegisterFile("as/external/api/gatewayProfile.proto", fileDescriptor_a425cd68ff0462c4)
}

var fileDescriptor_a425cd68ff0462c4 = []byte{
	// 881 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x85, 0x44, 0x59, 0xa9, 0xaf, 0x21, 0xa5, 0x1e, 0x37, 0x2e, 0x43, 0x39, 0xb1, 0x42, 0x74,
	0x21, 0xa8, 0x30, 0x09, 0xc8, 0xe8, 0xa2, 0x45, 0x51, 0x34, 0xb1, 0x13, 0xc3, 0x40, 0x5b, 0x18,
	0x6c, 0xbb, 0x69, 0x16, 0xc2, 0x88, 0x1c, 0x49, 0x03, 0x93, 0x1c, 0x76, 0x66, 0x64, 0xc7, 0x29,
	0xb2, 0xe9, 0xa2, 0xcb, 0x6e, 0xfa, 0x41, 0xfd, 0x84, 0x2e, 0xf2, 0x0b, 0xfd, 0x90, 0x62, 0x1e,
	0xb2, 0x4d, 0x99, 0x54, 0x91, 0x3e, 0x56, 0xd2, 0xbd, 0xf7, 0xcc, 0x7d, 0x9c, 0x39, 0x33, 0x43,
	0xf8, 0x08, 0x8b, 0x90, 0xbc, 0x92, 0x84, 0xe7, 0x38, 0x0d, 0x71, 0x41, 0xc3, 0x19, 0x96, 0xe4,
	0x12, 0x5f, 0x9d, 0x71, 0x36, 0xa5, 0x29, 0x09, 0x0a, 0xce, 0x24, 0x43, 0x0e, 0x2e, 0xa8, 0xb7,
	0x37, 0x63, 0x6c, 0x96, 0x12, 0x8d, 0xc2, 0x79, 0xce, 0x24, 0x96, 0x94, 0xe5, 0xc2, 0x40, 0xbc,
	0x7d, 0x1b, 0xd5, 0xd6, 0x64, 0x31, 0x0d, 0x25, 0xcd, 0x88, 0x90, 0x38, 0x2b, 0x2c, 0xe0, 0xf1,
	0x2a, 0x20, 0x59, 0x70, 0x9d, 0xc1, 0xc6, 0x7b, 0xab, 0x71, 0x92, 0x15, 0xf2, 0xca, 0x06, 0x77,
	0x62, 0x96, 0x65, 0x2c, 0x0f, 0xcd, 0x8f, 0x71, 0xfa, 0xbf, 0x34, 0xa1, 0x7b, 0x52, 0x6a, 0x17,
	0x75, 0xa1, 0x49, 0x13, 0xb7, 0xd1, 0x6f, 0x0c, 0x36, 0xa3, 0x26, 0x4d, 0x10, 0x82, 0x56, 0x8e,
	0x33, 0xe2, 0x36, 0xb5, 0x47, 0xff, 0x47, 0x43, 0xd8, 0xce, 0x89, 0xbc, 0x64, 0xfc, 0x7c, 0x2c,
	0x08, 0xbf, 0x20, 0x7c, 0x4c, 0x13, 0xd7, 0xe9, 0x37, 0x06, 0x4e, 0x74, 0xdf, 0x06, 0xbe, 0xd5,
	0xfe, 0xd3, 0x63, 0xe4, 0xc1, 0x7b, 0xf1, 0x1c, 0xe7, 0x39, 0x49, 0x85, 0xdb, 0xea, 0x3b, 0x83,
	0x4e, 0x74, 0x6d, 0xa3, 0x17, 0xd0, 0x25, 0xaf, 0x24, 0xc7, 0xe3, 0x6b, 0xc4, 0x46, 0xdf, 0x19,
	0x6c, 0x8d, 0xf6, 0x03, 0x5c, 0xd0, 0xa0, 0xdc, 0xd8, 0x73, 0x05, 0x3c, 0x32, 0xb8, 0xa8, 0x43,
	0x6e, 0x59, 0x02, 0x7d, 0x09, 0x5d, 0x21, 0xb1, 0x14, 0x63, 0x9a, 0x4b, 0xc2, 0x2f, 0x70, 0xea,
	0xb6, 0xfb, 0x8d, 0xc1, 0xd6, 0xe8, 0x61, 0x60, 0x18, 0x09, 0x96, 0x8c, 0x04, 0xc7, 0x96, 0xb1,
	0xa8, 0xa3, 0x17, 0x9c, 0x5a, 0xbc, 0xff, 0x6b, 0x13, 0x76, 0xcb, 0xf5, 0xbe, 0xa2, 0x42, 0x9e,
	0x4a, 0x92, 0xfd, 0xe7, 0x84, 0x04, 0xb0, 0xb3, 0x82, 0xd5, 0xe9, 0xee, 0xe9, 0x74, 0xdb, 0x25,
	0xf4, 0x37, 0x2a, 0xf7, 0xa7, 0x00, 0x31, 0x27, 0x58, 0x92, 0x64, 0x8c, 0xa5, 0xbb, 0xa1, 0x07,
	0xf3, 0xee, 0x0c, 0xf6, 0xdd, 0x52, 0x2b, 0xd1, 0xa6, 0x45, 0x3f, 0x95, 0x6a, 0xe9, 0xa2, 0x48,
	0x96, 0x4b, 0xdb, 0x7f, 0xbf, 0xd4, 0xa2, 0x9f, 0x4a, 0xff, 0x6d, 0x03, 0xbc, 0xfa, 0x0d, 0x40,
	0x23, 0x80, 0x8c, 0x25, 0x8b, 0x54, 0x93, 0xa9, 0xc9, 0xe9, 0x8e, 0x50, 0x60, 0xb5, 0xf5, 0xf5,
	0x75, 0x24, 0xba, 0x85, 0x42, 0x7b, 0xb0, 0x39, 0xe5, 0xe4, 0xc7, 0x05, 0xc9, 0xe3, 0x2b, 0xcd,
	0x5e, 0x27, 0xba, 0x71, 0xa8, 0xe8, 0x04, 0xe7, 0xc9, 0x25, 0x4d, 0xe4, 0x5c, 0x53, 0xd7, 0x89,
	0x6e, 0x1c, 0xc8, 0x85, 0x7b, 0x13, 0x2a, 0x39, 0x96, 0xc4, 0x6d, 0xe9, 0xd8, 0xd2, 0x44, 0x1f,
	0xc3, 0xb6, 0x28, 0x38, 0xc1, 0x09, 0xcd, 0x67, 0xe3, 0x29, 0x8e, 0x25, 0xe3, 0x46, 0x46, 0x9d,
	0xe8, 0xfd, 0xeb, 0xc0, 0x0b, 0xe3, 0xf7, 0x5f, 0x42, 0xef, 0x48, 0xb3, 0x53, 0x1e, 0x2d, 0x52,
	0x4d, 0x08, 0x89, 0x3e, 0x87, 0xfb, 0xf6, 0xf0, 0x8e, 0x0b, 0x13, 0xd1, 0xa3, 0x6d, 0x8d, 0x76,
	0x2a, 0x04, 0x19, 0x75, 0xcb, 0x07, 0xdd, 0x0f, 0x60, 0xaf, 0x3a, 0xb9, 0x28, 0x58, 0x2e, 0xee,
	0x9c, 0x2c, 0x7f, 0x08, 0xee, 0x09, 0x91, 0xd5, 0x9d, 0xac, 0x62, 0xff, 0x68, 0xc0, 0xc3, 0x0a,
	0xb0, 0xcd, 0xfc, 0xaf, 0xfa, 0x5e, 0x11, 0x58, 0xf3, 0x9f, 0x0b, 0xcc, 0x79, 0x17, 0x81, 0xbd,
	0x84, 0xde, 0xf7, 0xda, 0xf8, 0x3f, 0xb6, 0xe2, 0x00, 0x7a, 0xc7, 0x24, 0x25, 0x75, 0xc9, 0x57,
	0xd9, 0xbd, 0x00, 0x4f, 0x1d, 0xf7, 0x32, 0x58, 0x2c, 0xd1, 0x1f, 0xc0, 0x46, 0x4a, 0x33, 0x2a,
	0xf5, 0x02, 0x27, 0x32, 0x06, 0xda, 0x85, 0x36, 0x9b, 0x4e, 0x05, 0x31, 0x8c, 0x39, 0x91, 0xb5,
	0xde, 0xe5, 0x2a, 0xf0, 0x05, 0xf4, 0x2a, 0xeb, 0xda, 0x6d, 0xdd, 0x87, 0x2d, 0xc9, 0x24, 0x4e,
	0xc7, 0x31, 0x5b, 0xe4, 0xcb, 0xf2, 0xa0, 0x5d, 0x47, 0xca, 0x83, 0x0e, 0xa1, 0xcd, 0x89, 0x58,
	0xa4, 0xaa, 0x07, 0x75, 0x6f, 0xf6, 0x2a, 0xb8, 0x59, 0xde, 0x63, 0x91, 0x85, 0x8e, 0x7e, 0x6f,
	0xc1, 0x83, 0x32, 0x44, 0xf5, 0x43, 0x63, 0x82, 0x18, 0xb4, 0x8d, 0x80, 0x51, 0x5f, 0x27, 0x5a,
	0x73, 0x54, 0xbc, 0x27, 0x6b, 0x10, 0xa6, 0x7d, 0xbf, 0xff, 0xf3, 0xdb, 0x3f, 0x7f, 0x6b, 0x7a,
	0xfe, 0x83, 0xdb, 0xaf, 0xe2, 0x81, 0xdd, 0x4d, 0xf1, 0x59, 0x63, 0x88, 0xe6, 0xe0, 0x9c, 0x10,
	0x89, 0x1e, 0x99, 0xb6, 0x6b, 0xce, 0x82, 0xf7, 0xb8, 0x2e, 0x6c, 0xeb, 0xf8, 0xba, 0xce, 0x1e,
	0xf2, 0x2a, 0xeb, 0x84, 0x3f, 0xd1, 0xe4, 0x0d, 0xba, 0x82, 0xb6, 0x51, 0x9b, 0x1d, 0x6d, 0x8d,
	0xf4, 0xbc, 0xdd, 0x3b, 0x02, 0x7e, 0xae, 0xde, 0x51, 0xff, 0x13, 0x5d, 0x27, 0xf4, 0x86, 0x35,
	0x75, 0x56, 0xf4, 0x1a, 0xd0, 0xe4, 0x8d, 0x1a, 0x72, 0x0a, 0x6d, 0xa3, 0x45, 0x5b, 0x7a, 0x8d,
	0x30, 0x6b, 0x4b, 0xdb, 0x11, 0x87, 0xeb, 0x46, 0x9c, 0x43, 0x4b, 0xed, 0x35, 0x32, 0x8f, 0x67,
	0xbd, 0x9e, 0xbd, 0x7e, 0x3d, 0xc0, 0x32, 0xfa, 0x48, 0x97, 0xfb, 0x10, 0x55, 0xef, 0xdc, 0xb3,
	0xd7, 0xf0, 0x84, 0xb2, 0x20, 0x9e, 0x53, 0x5e, 0x08, 0x89, 0xe3, 0x73, 0x9d, 0x0f, 0x8b, 0x60,
	0xf9, 0x15, 0xa4, 0xec, 0x67, 0x3b, 0xe5, 0xe4, 0x67, 0x6a, 0xa0, 0xb3, 0xc6, 0x0f, 0x5f, 0xcc,
	0xa8, 0x9c, 0x2f, 0x26, 0xea, 0xa9, 0x08, 0xe9, 0x84, 0xe3, 0x39, 0xcd, 0xd8, 0x6b, 0x72, 0x4e,
	0x63, 0x1a, 0xde, 0x24, 0x3c, 0xd0, 0x65, 0x59, 0x78, 0x71, 0x18, 0xae, 0x7c, 0x5c, 0x4d, 0xda,
	0x9a, 0x99, 0xc3, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x65, 0xde, 0xab, 0x76, 0x09, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GatewayProfileServiceClient is the client API for GatewayProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayProfileServiceClient interface {
	// Create creates the given gateway-profile.
	Create(ctx context.Context, in *CreateGatewayProfileRequest, opts ...grpc.CallOption) (*CreateGatewayProfileResponse, error)
	// Get returns the gateway-profile matching the given id.
	Get(ctx context.Context, in *GetGatewayProfileRequest, opts ...grpc.CallOption) (*GetGatewayProfileResponse, error)
	// Update updates the given gateway-profile.
	Update(ctx context.Context, in *UpdateGatewayProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Delete deletes the gateway-profile matching the given id.
	Delete(ctx context.Context, in *DeleteGatewayProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// List returns the existing gateway-profiles.
	List(ctx context.Context, in *ListGatewayProfilesRequest, opts ...grpc.CallOption) (*ListGatewayProfilesResponse, error)
}

type gatewayProfileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayProfileServiceClient(cc grpc.ClientConnInterface) GatewayProfileServiceClient {
	return &gatewayProfileServiceClient{cc}
}

func (c *gatewayProfileServiceClient) Create(ctx context.Context, in *CreateGatewayProfileRequest, opts ...grpc.CallOption) (*CreateGatewayProfileResponse, error) {
	out := new(CreateGatewayProfileResponse)
	err := c.cc.Invoke(ctx, "/api.GatewayProfileService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayProfileServiceClient) Get(ctx context.Context, in *GetGatewayProfileRequest, opts ...grpc.CallOption) (*GetGatewayProfileResponse, error) {
	out := new(GetGatewayProfileResponse)
	err := c.cc.Invoke(ctx, "/api.GatewayProfileService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayProfileServiceClient) Update(ctx context.Context, in *UpdateGatewayProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.GatewayProfileService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayProfileServiceClient) Delete(ctx context.Context, in *DeleteGatewayProfileRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.GatewayProfileService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayProfileServiceClient) List(ctx context.Context, in *ListGatewayProfilesRequest, opts ...grpc.CallOption) (*ListGatewayProfilesResponse, error) {
	out := new(ListGatewayProfilesResponse)
	err := c.cc.Invoke(ctx, "/api.GatewayProfileService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayProfileServiceServer is the server API for GatewayProfileService service.
type GatewayProfileServiceServer interface {
	// Create creates the given gateway-profile.
	Create(context.Context, *CreateGatewayProfileRequest) (*CreateGatewayProfileResponse, error)
	// Get returns the gateway-profile matching the given id.
	Get(context.Context, *GetGatewayProfileRequest) (*GetGatewayProfileResponse, error)
	// Update updates the given gateway-profile.
	Update(context.Context, *UpdateGatewayProfileRequest) (*empty.Empty, error)
	// Delete deletes the gateway-profile matching the given id.
	Delete(context.Context, *DeleteGatewayProfileRequest) (*empty.Empty, error)
	// List returns the existing gateway-profiles.
	List(context.Context, *ListGatewayProfilesRequest) (*ListGatewayProfilesResponse, error)
}

// UnimplementedGatewayProfileServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGatewayProfileServiceServer struct {
}

func (*UnimplementedGatewayProfileServiceServer) Create(ctx context.Context, req *CreateGatewayProfileRequest) (*CreateGatewayProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedGatewayProfileServiceServer) Get(ctx context.Context, req *GetGatewayProfileRequest) (*GetGatewayProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedGatewayProfileServiceServer) Update(ctx context.Context, req *UpdateGatewayProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedGatewayProfileServiceServer) Delete(ctx context.Context, req *DeleteGatewayProfileRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedGatewayProfileServiceServer) List(ctx context.Context, req *ListGatewayProfilesRequest) (*ListGatewayProfilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterGatewayProfileServiceServer(s *grpc.Server, srv GatewayProfileServiceServer) {
	s.RegisterService(&_GatewayProfileService_serviceDesc, srv)
}

func _GatewayProfileService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGatewayProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayProfileServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayProfileService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayProfileServiceServer).Create(ctx, req.(*CreateGatewayProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayProfileService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayProfileServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayProfileService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayProfileServiceServer).Get(ctx, req.(*GetGatewayProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayProfileService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGatewayProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayProfileServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayProfileService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayProfileServiceServer).Update(ctx, req.(*UpdateGatewayProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayProfileService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGatewayProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayProfileServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayProfileService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayProfileServiceServer).Delete(ctx, req.(*DeleteGatewayProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayProfileService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGatewayProfilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayProfileServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GatewayProfileService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayProfileServiceServer).List(ctx, req.(*ListGatewayProfilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GatewayProfileService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.GatewayProfileService",
	HandlerType: (*GatewayProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _GatewayProfileService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _GatewayProfileService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GatewayProfileService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GatewayProfileService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _GatewayProfileService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as/external/api/gatewayProfile.proto",
}
