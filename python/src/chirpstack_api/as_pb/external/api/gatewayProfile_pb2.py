# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chirpstack-api/as_pb/external/api/gatewayProfile.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from chirpstack_api.common import common_pb2 as chirpstack__api_dot_common_dot_common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='chirpstack-api/as_pb/external/api/gatewayProfile.proto',
  package='api',
  syntax='proto3',
  serialized_options=b'\n!io.chirpstack.api.as.external.apiB\023GatewayProfileProtoP\001Z7github.com/ibrahimozekici/chirpstack-api/go/v5/as/external/api',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n6chirpstack-api/as_pb/external/api/gatewayProfile.proto\x12\x03\x61pi\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\"chirpstack-api/common/common.proto\"\xd4\x01\n\x0eGatewayProfile\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12*\n\x11network_server_id\x18\x03 \x01(\x03R\x0fnetworkServerID\x12\x10\n\x08\x63hannels\x18\x04 \x03(\r\x12\x37\n\x0e\x65xtra_channels\x18\x05 \x03(\x0b\x32\x1f.api.GatewayProfileExtraChannel\x12\x31\n\x0estats_interval\x18\x06 \x01(\x0b\x32\x19.google.protobuf.Duration\"\xdb\x01\n\x16GatewayProfileListItem\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12*\n\x11network_server_id\x18\x03 \x01(\x03R\x0fnetworkServerID\x12\x1b\n\x13network_server_name\x18\x07 \x01(\t\x12.\n\ncreated_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\x96\x01\n\x1aGatewayProfileExtraChannel\x12&\n\nmodulation\x18\x01 \x01(\x0e\x32\x12.common.Modulation\x12\x11\n\tfrequency\x18\x02 \x01(\r\x12\x11\n\tbandwidth\x18\x03 \x01(\r\x12\x0f\n\x07\x62itrate\x18\x04 \x01(\r\x12\x19\n\x11spreading_factors\x18\x05 \x03(\r\"K\n\x1b\x43reateGatewayProfileRequest\x12,\n\x0fgateway_profile\x18\x01 \x01(\x0b\x32\x13.api.GatewayProfile\"*\n\x1c\x43reateGatewayProfileResponse\x12\n\n\x02id\x18\x01 \x01(\t\"&\n\x18GetGatewayProfileRequest\x12\n\n\x02id\x18\x01 \x01(\t\"\xa9\x01\n\x19GetGatewayProfileResponse\x12,\n\x0fgateway_profile\x18\x01 \x01(\x0b\x32\x13.api.GatewayProfile\x12.\n\ncreated_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"K\n\x1bUpdateGatewayProfileRequest\x12,\n\x0fgateway_profile\x18\x01 \x01(\x0b\x32\x13.api.GatewayProfile\")\n\x1b\x44\x65leteGatewayProfileRequest\x12\n\n\x02id\x18\x01 \x01(\t\"g\n\x1aListGatewayProfilesRequest\x12\r\n\x05limit\x18\x01 \x01(\x03\x12\x0e\n\x06offset\x18\x02 \x01(\x03\x12*\n\x11network_server_id\x18\x03 \x01(\x03R\x0fnetworkServerID\"_\n\x1bListGatewayProfilesResponse\x12\x13\n\x0btotal_count\x18\x01 \x01(\x03\x12+\n\x06result\x18\x02 \x03(\x0b\x32\x1b.api.GatewayProfileListItem2\xbf\x04\n\x15GatewayProfileService\x12o\n\x06\x43reate\x12 .api.CreateGatewayProfileRequest\x1a!.api.CreateGatewayProfileResponse\" \x82\xd3\xe4\x93\x02\x1a\"\x15/api/gateway-profiles:\x01*\x12h\n\x03Get\x12\x1d.api.GetGatewayProfileRequest\x1a\x1e.api.GetGatewayProfileResponse\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/api/gateway-profiles/{id}\x12y\n\x06Update\x12 .api.UpdateGatewayProfileRequest\x1a\x16.google.protobuf.Empty\"5\x82\xd3\xe4\x93\x02/\x1a*/api/gateway-profiles/{gateway_profile.id}:\x01*\x12\x66\n\x06\x44\x65lete\x12 .api.DeleteGatewayProfileRequest\x1a\x16.google.protobuf.Empty\"\"\x82\xd3\xe4\x93\x02\x1c*\x1a/api/gateway-profiles/{id}\x12h\n\x04List\x12\x1f.api.ListGatewayProfilesRequest\x1a .api.ListGatewayProfilesResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\x12\x15/api/gateway-profilesBs\n!io.chirpstack.api.as.external.apiB\x13GatewayProfileProtoP\x01Z7github.com/ibrahimozekici/chirpstack-api/go/v5/as/external/apib\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,chirpstack__api_dot_common_dot_common__pb2.DESCRIPTOR,])




_GATEWAYPROFILE = _descriptor.Descriptor(
  name='GatewayProfile',
  full_name='api.GatewayProfile',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.GatewayProfile.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='api.GatewayProfile.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='network_server_id', full_name='api.GatewayProfile.network_server_id', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='networkServerID', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='channels', full_name='api.GatewayProfile.channels', index=3,
      number=4, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='extra_channels', full_name='api.GatewayProfile.extra_channels', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='stats_interval', full_name='api.GatewayProfile.stats_interval', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=224,
  serialized_end=436,
)


_GATEWAYPROFILELISTITEM = _descriptor.Descriptor(
  name='GatewayProfileListItem',
  full_name='api.GatewayProfileListItem',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.GatewayProfileListItem.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='api.GatewayProfileListItem.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='network_server_id', full_name='api.GatewayProfileListItem.network_server_id', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='networkServerID', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='network_server_name', full_name='api.GatewayProfileListItem.network_server_name', index=3,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='api.GatewayProfileListItem.created_at', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='api.GatewayProfileListItem.updated_at', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=439,
  serialized_end=658,
)


_GATEWAYPROFILEEXTRACHANNEL = _descriptor.Descriptor(
  name='GatewayProfileExtraChannel',
  full_name='api.GatewayProfileExtraChannel',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='modulation', full_name='api.GatewayProfileExtraChannel.modulation', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='frequency', full_name='api.GatewayProfileExtraChannel.frequency', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='bandwidth', full_name='api.GatewayProfileExtraChannel.bandwidth', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='bitrate', full_name='api.GatewayProfileExtraChannel.bitrate', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='spreading_factors', full_name='api.GatewayProfileExtraChannel.spreading_factors', index=4,
      number=5, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=661,
  serialized_end=811,
)


_CREATEGATEWAYPROFILEREQUEST = _descriptor.Descriptor(
  name='CreateGatewayProfileRequest',
  full_name='api.CreateGatewayProfileRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='gateway_profile', full_name='api.CreateGatewayProfileRequest.gateway_profile', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=813,
  serialized_end=888,
)


_CREATEGATEWAYPROFILERESPONSE = _descriptor.Descriptor(
  name='CreateGatewayProfileResponse',
  full_name='api.CreateGatewayProfileResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.CreateGatewayProfileResponse.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=890,
  serialized_end=932,
)


_GETGATEWAYPROFILEREQUEST = _descriptor.Descriptor(
  name='GetGatewayProfileRequest',
  full_name='api.GetGatewayProfileRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.GetGatewayProfileRequest.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=934,
  serialized_end=972,
)


_GETGATEWAYPROFILERESPONSE = _descriptor.Descriptor(
  name='GetGatewayProfileResponse',
  full_name='api.GetGatewayProfileResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='gateway_profile', full_name='api.GetGatewayProfileResponse.gateway_profile', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='api.GetGatewayProfileResponse.created_at', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='api.GetGatewayProfileResponse.updated_at', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=975,
  serialized_end=1144,
)


_UPDATEGATEWAYPROFILEREQUEST = _descriptor.Descriptor(
  name='UpdateGatewayProfileRequest',
  full_name='api.UpdateGatewayProfileRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='gateway_profile', full_name='api.UpdateGatewayProfileRequest.gateway_profile', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1146,
  serialized_end=1221,
)


_DELETEGATEWAYPROFILEREQUEST = _descriptor.Descriptor(
  name='DeleteGatewayProfileRequest',
  full_name='api.DeleteGatewayProfileRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.DeleteGatewayProfileRequest.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1223,
  serialized_end=1264,
)


_LISTGATEWAYPROFILESREQUEST = _descriptor.Descriptor(
  name='ListGatewayProfilesRequest',
  full_name='api.ListGatewayProfilesRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='limit', full_name='api.ListGatewayProfilesRequest.limit', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='offset', full_name='api.ListGatewayProfilesRequest.offset', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='network_server_id', full_name='api.ListGatewayProfilesRequest.network_server_id', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='networkServerID', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1266,
  serialized_end=1369,
)


_LISTGATEWAYPROFILESRESPONSE = _descriptor.Descriptor(
  name='ListGatewayProfilesResponse',
  full_name='api.ListGatewayProfilesResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='total_count', full_name='api.ListGatewayProfilesResponse.total_count', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='result', full_name='api.ListGatewayProfilesResponse.result', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1371,
  serialized_end=1466,
)

_GATEWAYPROFILE.fields_by_name['extra_channels'].message_type = _GATEWAYPROFILEEXTRACHANNEL
_GATEWAYPROFILE.fields_by_name['stats_interval'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_GATEWAYPROFILELISTITEM.fields_by_name['created_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_GATEWAYPROFILELISTITEM.fields_by_name['updated_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_GATEWAYPROFILEEXTRACHANNEL.fields_by_name['modulation'].enum_type = chirpstack__api_dot_common_dot_common__pb2._MODULATION
_CREATEGATEWAYPROFILEREQUEST.fields_by_name['gateway_profile'].message_type = _GATEWAYPROFILE
_GETGATEWAYPROFILERESPONSE.fields_by_name['gateway_profile'].message_type = _GATEWAYPROFILE
_GETGATEWAYPROFILERESPONSE.fields_by_name['created_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_GETGATEWAYPROFILERESPONSE.fields_by_name['updated_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_UPDATEGATEWAYPROFILEREQUEST.fields_by_name['gateway_profile'].message_type = _GATEWAYPROFILE
_LISTGATEWAYPROFILESRESPONSE.fields_by_name['result'].message_type = _GATEWAYPROFILELISTITEM
DESCRIPTOR.message_types_by_name['GatewayProfile'] = _GATEWAYPROFILE
DESCRIPTOR.message_types_by_name['GatewayProfileListItem'] = _GATEWAYPROFILELISTITEM
DESCRIPTOR.message_types_by_name['GatewayProfileExtraChannel'] = _GATEWAYPROFILEEXTRACHANNEL
DESCRIPTOR.message_types_by_name['CreateGatewayProfileRequest'] = _CREATEGATEWAYPROFILEREQUEST
DESCRIPTOR.message_types_by_name['CreateGatewayProfileResponse'] = _CREATEGATEWAYPROFILERESPONSE
DESCRIPTOR.message_types_by_name['GetGatewayProfileRequest'] = _GETGATEWAYPROFILEREQUEST
DESCRIPTOR.message_types_by_name['GetGatewayProfileResponse'] = _GETGATEWAYPROFILERESPONSE
DESCRIPTOR.message_types_by_name['UpdateGatewayProfileRequest'] = _UPDATEGATEWAYPROFILEREQUEST
DESCRIPTOR.message_types_by_name['DeleteGatewayProfileRequest'] = _DELETEGATEWAYPROFILEREQUEST
DESCRIPTOR.message_types_by_name['ListGatewayProfilesRequest'] = _LISTGATEWAYPROFILESREQUEST
DESCRIPTOR.message_types_by_name['ListGatewayProfilesResponse'] = _LISTGATEWAYPROFILESRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GatewayProfile = _reflection.GeneratedProtocolMessageType('GatewayProfile', (_message.Message,), {
  'DESCRIPTOR' : _GATEWAYPROFILE,
  '__module__' : 'chirpstack_api.as_pb.external.api.gatewayProfile_pb2'
  # @@protoc_insertion_point(class_scope:api.GatewayProfile)
  })
_sym_db.RegisterMessage(GatewayProfile)

GatewayProfileListItem = _reflection.GeneratedProtocolMessageType('GatewayProfileListItem', (_message.Message,), {
  'DESCRIPTOR' : _GATEWAYPROFILELISTITEM,
  '__module__' : 'chirpstack_api.as_pb.external.api.gatewayProfile_pb2'
  # @@protoc_insertion_point(class_scope:api.GatewayProfileListItem)
  })
_sym_db.RegisterMessage(GatewayProfileListItem)

GatewayProfileExtraChannel = _reflection.GeneratedProtocolMessageType('GatewayProfileExtraChannel', (_message.Message,), {
  'DESCRIPTOR' : _GATEWAYPROFILEEXTRACHANNEL,
  '__module__' : 'chirpstack_api.as_pb.external.api.gatewayProfile_pb2'
  # @@protoc_insertion_point(class_scope:api.GatewayProfileExtraChannel)
  })
_sym_db.RegisterMessage(GatewayProfileExtraChannel)

CreateGatewayProfileRequest = _reflection.GeneratedProtocolMessageType('CreateGatewayProfileRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEGATEWAYPROFILEREQUEST,
  '__module__' : 'chirpstack_api.as_pb.external.api.gatewayProfile_pb2'
  # @@protoc_insertion_point(class_scope:api.CreateGatewayProfileRequest)
  })
_sym_db.RegisterMessage(CreateGatewayProfileRequest)

CreateGatewayProfileResponse = _reflection.GeneratedProtocolMessageType('CreateGatewayProfileResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEGATEWAYPROFILERESPONSE,
  '__module__' : 'chirpstack_api.as_pb.external.api.gatewayProfile_pb2'
  # @@protoc_insertion_point(class_scope:api.CreateGatewayProfileResponse)
  })
_sym_db.RegisterMessage(CreateGatewayProfileResponse)

GetGatewayProfileRequest = _reflection.GeneratedProtocolMessageType('GetGatewayProfileRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETGATEWAYPROFILEREQUEST,
  '__module__' : 'chirpstack_api.as_pb.external.api.gatewayProfile_pb2'
  # @@protoc_insertion_point(class_scope:api.GetGatewayProfileRequest)
  })
_sym_db.RegisterMessage(GetGatewayProfileRequest)

GetGatewayProfileResponse = _reflection.GeneratedProtocolMessageType('GetGatewayProfileResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETGATEWAYPROFILERESPONSE,
  '__module__' : 'chirpstack_api.as_pb.external.api.gatewayProfile_pb2'
  # @@protoc_insertion_point(class_scope:api.GetGatewayProfileResponse)
  })
_sym_db.RegisterMessage(GetGatewayProfileResponse)

UpdateGatewayProfileRequest = _reflection.GeneratedProtocolMessageType('UpdateGatewayProfileRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEGATEWAYPROFILEREQUEST,
  '__module__' : 'chirpstack_api.as_pb.external.api.gatewayProfile_pb2'
  # @@protoc_insertion_point(class_scope:api.UpdateGatewayProfileRequest)
  })
_sym_db.RegisterMessage(UpdateGatewayProfileRequest)

DeleteGatewayProfileRequest = _reflection.GeneratedProtocolMessageType('DeleteGatewayProfileRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETEGATEWAYPROFILEREQUEST,
  '__module__' : 'chirpstack_api.as_pb.external.api.gatewayProfile_pb2'
  # @@protoc_insertion_point(class_scope:api.DeleteGatewayProfileRequest)
  })
_sym_db.RegisterMessage(DeleteGatewayProfileRequest)

ListGatewayProfilesRequest = _reflection.GeneratedProtocolMessageType('ListGatewayProfilesRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTGATEWAYPROFILESREQUEST,
  '__module__' : 'chirpstack_api.as_pb.external.api.gatewayProfile_pb2'
  # @@protoc_insertion_point(class_scope:api.ListGatewayProfilesRequest)
  })
_sym_db.RegisterMessage(ListGatewayProfilesRequest)

ListGatewayProfilesResponse = _reflection.GeneratedProtocolMessageType('ListGatewayProfilesResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTGATEWAYPROFILESRESPONSE,
  '__module__' : 'chirpstack_api.as_pb.external.api.gatewayProfile_pb2'
  # @@protoc_insertion_point(class_scope:api.ListGatewayProfilesResponse)
  })
_sym_db.RegisterMessage(ListGatewayProfilesResponse)


DESCRIPTOR._options = None

_GATEWAYPROFILESERVICE = _descriptor.ServiceDescriptor(
  name='GatewayProfileService',
  full_name='api.GatewayProfileService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1469,
  serialized_end=2044,
  methods=[
  _descriptor.MethodDescriptor(
    name='Create',
    full_name='api.GatewayProfileService.Create',
    index=0,
    containing_service=None,
    input_type=_CREATEGATEWAYPROFILEREQUEST,
    output_type=_CREATEGATEWAYPROFILERESPONSE,
    serialized_options=b'\202\323\344\223\002\032\"\025/api/gateway-profiles:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Get',
    full_name='api.GatewayProfileService.Get',
    index=1,
    containing_service=None,
    input_type=_GETGATEWAYPROFILEREQUEST,
    output_type=_GETGATEWAYPROFILERESPONSE,
    serialized_options=b'\202\323\344\223\002\034\022\032/api/gateway-profiles/{id}',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Update',
    full_name='api.GatewayProfileService.Update',
    index=2,
    containing_service=None,
    input_type=_UPDATEGATEWAYPROFILEREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=b'\202\323\344\223\002/\032*/api/gateway-profiles/{gateway_profile.id}:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Delete',
    full_name='api.GatewayProfileService.Delete',
    index=3,
    containing_service=None,
    input_type=_DELETEGATEWAYPROFILEREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=b'\202\323\344\223\002\034*\032/api/gateway-profiles/{id}',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='List',
    full_name='api.GatewayProfileService.List',
    index=4,
    containing_service=None,
    input_type=_LISTGATEWAYPROFILESREQUEST,
    output_type=_LISTGATEWAYPROFILESRESPONSE,
    serialized_options=b'\202\323\344\223\002\027\022\025/api/gateway-profiles',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_GATEWAYPROFILESERVICE)

DESCRIPTOR.services_by_name['GatewayProfileService'] = _GATEWAYPROFILESERVICE

# @@protoc_insertion_point(module_scope)
