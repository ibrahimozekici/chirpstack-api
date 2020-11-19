# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chirpstack-api/as_pb/external/api/networkServer.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='chirpstack-api/as_pb/external/api/networkServer.proto',
  package='api',
  syntax='proto3',
  serialized_options=b'\n!io.chirpstack.api.as.external.apiZ7github.com/ibrahimozekici/chirpstack-api/go/v4/as/external/api',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n5chirpstack-api/as_pb/external/api/networkServer.proto\x12\x03\x61pi\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bgoogle/protobuf/empty.proto\"\xd2\x03\n\rNetworkServer\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0e\n\x06server\x18\x03 \x01(\t\x12\x0f\n\x07\x63\x61_cert\x18\x04 \x01(\t\x12\x10\n\x08tls_cert\x18\x05 \x01(\t\x12\x0f\n\x07tls_key\x18\x06 \x01(\t\x12\x35\n\x17routing_profile_ca_cert\x18\x07 \x01(\tR\x14routingProfileCACert\x12\x37\n\x18routing_profile_tls_cert\x18\x08 \x01(\tR\x15routingProfileTLSCert\x12\x35\n\x17routing_profile_tls_key\x18\t \x01(\tR\x14routingProfileTLSKey\x12!\n\x19gateway_discovery_enabled\x18\n \x01(\x08\x12\"\n\x1agateway_discovery_interval\x18\x0b \x01(\r\x12\x43\n\x1egateway_discovery_tx_frequency\x18\x0c \x01(\rR\x1bgatewayDiscoveryTXFrequency\x12\x30\n\x14gateway_discovery_dr\x18\r \x01(\rR\x12gatewayDiscoveryDR\"\xa1\x01\n\x15NetworkServerListItem\x12\n\n\x02id\x18\x01 \x01(\x03\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0e\n\x06server\x18\x03 \x01(\t\x12.\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"H\n\x1a\x43reateNetworkServerRequest\x12*\n\x0enetwork_server\x18\x01 \x01(\x0b\x32\x12.api.NetworkServer\")\n\x1b\x43reateNetworkServerResponse\x12\n\n\x02id\x18\x01 \x01(\x03\"%\n\x17GetNetworkServerRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"\xc7\x01\n\x18GetNetworkServerResponse\x12*\n\x0enetwork_server\x18\x01 \x01(\x0b\x32\x12.api.NetworkServer\x12.\n\ncreated_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12.\n\nupdated_at\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x0f\n\x07version\x18\x04 \x01(\t\x12\x0e\n\x06region\x18\x05 \x01(\t\"H\n\x1aUpdateNetworkServerRequest\x12*\n\x0enetwork_server\x18\x01 \x01(\x0b\x32\x12.api.NetworkServer\"(\n\x1a\x44\x65leteNetworkServerRequest\x12\n\n\x02id\x18\x01 \x01(\x03\"b\n\x18ListNetworkServerRequest\x12\r\n\x05limit\x18\x01 \x01(\x03\x12\x0e\n\x06offset\x18\x02 \x01(\x03\x12\'\n\x0forganization_id\x18\x03 \x01(\x03R\x0eorganizationID\"\\\n\x19ListNetworkServerResponse\x12\x13\n\x0btotal_count\x18\x01 \x01(\x03\x12*\n\x06result\x18\x02 \x03(\x0b\x32\x1a.api.NetworkServerListItem2\xae\x04\n\x14NetworkServerService\x12l\n\x06\x43reate\x12\x1f.api.CreateNetworkServerRequest\x1a .api.CreateNetworkServerResponse\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/api/network-servers:\x01*\x12\x65\n\x03Get\x12\x1c.api.GetNetworkServerRequest\x1a\x1d.api.GetNetworkServerResponse\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/api/network-servers/{id}\x12v\n\x06Update\x12\x1f.api.UpdateNetworkServerRequest\x1a\x16.google.protobuf.Empty\"3\x82\xd3\xe4\x93\x02-\x1a(/api/network-servers/{network_server.id}:\x01*\x12\x64\n\x06\x44\x65lete\x12\x1f.api.DeleteNetworkServerRequest\x1a\x16.google.protobuf.Empty\"!\x82\xd3\xe4\x93\x02\x1b*\x19/api/network-servers/{id}\x12\x63\n\x04List\x12\x1d.api.ListNetworkServerRequest\x1a\x1e.api.ListNetworkServerResponse\"\x1c\x82\xd3\xe4\x93\x02\x16\x12\x14/api/network-serversB\\\n!io.chirpstack.api.as.external.apiZ7github.com/ibrahimozekici/chirpstack-api/go/v4/as/external/apib\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,])




_NETWORKSERVER = _descriptor.Descriptor(
  name='NetworkServer',
  full_name='api.NetworkServer',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.NetworkServer.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='api.NetworkServer.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='server', full_name='api.NetworkServer.server', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ca_cert', full_name='api.NetworkServer.ca_cert', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='tls_cert', full_name='api.NetworkServer.tls_cert', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='tls_key', full_name='api.NetworkServer.tls_key', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='routing_profile_ca_cert', full_name='api.NetworkServer.routing_profile_ca_cert', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='routingProfileCACert', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='routing_profile_tls_cert', full_name='api.NetworkServer.routing_profile_tls_cert', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='routingProfileTLSCert', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='routing_profile_tls_key', full_name='api.NetworkServer.routing_profile_tls_key', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='routingProfileTLSKey', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='gateway_discovery_enabled', full_name='api.NetworkServer.gateway_discovery_enabled', index=9,
      number=10, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='gateway_discovery_interval', full_name='api.NetworkServer.gateway_discovery_interval', index=10,
      number=11, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='gateway_discovery_tx_frequency', full_name='api.NetworkServer.gateway_discovery_tx_frequency', index=11,
      number=12, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='gatewayDiscoveryTXFrequency', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='gateway_discovery_dr', full_name='api.NetworkServer.gateway_discovery_dr', index=12,
      number=13, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='gatewayDiscoveryDR', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=155,
  serialized_end=621,
)


_NETWORKSERVERLISTITEM = _descriptor.Descriptor(
  name='NetworkServerListItem',
  full_name='api.NetworkServerListItem',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.NetworkServerListItem.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='api.NetworkServerListItem.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='server', full_name='api.NetworkServerListItem.server', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='api.NetworkServerListItem.created_at', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='api.NetworkServerListItem.updated_at', index=4,
      number=5, type=11, cpp_type=10, label=1,
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
  serialized_start=624,
  serialized_end=785,
)


_CREATENETWORKSERVERREQUEST = _descriptor.Descriptor(
  name='CreateNetworkServerRequest',
  full_name='api.CreateNetworkServerRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='network_server', full_name='api.CreateNetworkServerRequest.network_server', index=0,
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
  serialized_start=787,
  serialized_end=859,
)


_CREATENETWORKSERVERRESPONSE = _descriptor.Descriptor(
  name='CreateNetworkServerResponse',
  full_name='api.CreateNetworkServerResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.CreateNetworkServerResponse.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=861,
  serialized_end=902,
)


_GETNETWORKSERVERREQUEST = _descriptor.Descriptor(
  name='GetNetworkServerRequest',
  full_name='api.GetNetworkServerRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.GetNetworkServerRequest.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=904,
  serialized_end=941,
)


_GETNETWORKSERVERRESPONSE = _descriptor.Descriptor(
  name='GetNetworkServerResponse',
  full_name='api.GetNetworkServerResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='network_server', full_name='api.GetNetworkServerResponse.network_server', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='created_at', full_name='api.GetNetworkServerResponse.created_at', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='updated_at', full_name='api.GetNetworkServerResponse.updated_at', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='version', full_name='api.GetNetworkServerResponse.version', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='region', full_name='api.GetNetworkServerResponse.region', index=4,
      number=5, type=9, cpp_type=9, label=1,
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
  serialized_start=944,
  serialized_end=1143,
)


_UPDATENETWORKSERVERREQUEST = _descriptor.Descriptor(
  name='UpdateNetworkServerRequest',
  full_name='api.UpdateNetworkServerRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='network_server', full_name='api.UpdateNetworkServerRequest.network_server', index=0,
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
  serialized_start=1145,
  serialized_end=1217,
)


_DELETENETWORKSERVERREQUEST = _descriptor.Descriptor(
  name='DeleteNetworkServerRequest',
  full_name='api.DeleteNetworkServerRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='api.DeleteNetworkServerRequest.id', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=1219,
  serialized_end=1259,
)


_LISTNETWORKSERVERREQUEST = _descriptor.Descriptor(
  name='ListNetworkServerRequest',
  full_name='api.ListNetworkServerRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='limit', full_name='api.ListNetworkServerRequest.limit', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='offset', full_name='api.ListNetworkServerRequest.offset', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='organization_id', full_name='api.ListNetworkServerRequest.organization_id', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='organizationID', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=1261,
  serialized_end=1359,
)


_LISTNETWORKSERVERRESPONSE = _descriptor.Descriptor(
  name='ListNetworkServerResponse',
  full_name='api.ListNetworkServerResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='total_count', full_name='api.ListNetworkServerResponse.total_count', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='result', full_name='api.ListNetworkServerResponse.result', index=1,
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
  serialized_start=1361,
  serialized_end=1453,
)

_NETWORKSERVERLISTITEM.fields_by_name['created_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_NETWORKSERVERLISTITEM.fields_by_name['updated_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_CREATENETWORKSERVERREQUEST.fields_by_name['network_server'].message_type = _NETWORKSERVER
_GETNETWORKSERVERRESPONSE.fields_by_name['network_server'].message_type = _NETWORKSERVER
_GETNETWORKSERVERRESPONSE.fields_by_name['created_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_GETNETWORKSERVERRESPONSE.fields_by_name['updated_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_UPDATENETWORKSERVERREQUEST.fields_by_name['network_server'].message_type = _NETWORKSERVER
_LISTNETWORKSERVERRESPONSE.fields_by_name['result'].message_type = _NETWORKSERVERLISTITEM
DESCRIPTOR.message_types_by_name['NetworkServer'] = _NETWORKSERVER
DESCRIPTOR.message_types_by_name['NetworkServerListItem'] = _NETWORKSERVERLISTITEM
DESCRIPTOR.message_types_by_name['CreateNetworkServerRequest'] = _CREATENETWORKSERVERREQUEST
DESCRIPTOR.message_types_by_name['CreateNetworkServerResponse'] = _CREATENETWORKSERVERRESPONSE
DESCRIPTOR.message_types_by_name['GetNetworkServerRequest'] = _GETNETWORKSERVERREQUEST
DESCRIPTOR.message_types_by_name['GetNetworkServerResponse'] = _GETNETWORKSERVERRESPONSE
DESCRIPTOR.message_types_by_name['UpdateNetworkServerRequest'] = _UPDATENETWORKSERVERREQUEST
DESCRIPTOR.message_types_by_name['DeleteNetworkServerRequest'] = _DELETENETWORKSERVERREQUEST
DESCRIPTOR.message_types_by_name['ListNetworkServerRequest'] = _LISTNETWORKSERVERREQUEST
DESCRIPTOR.message_types_by_name['ListNetworkServerResponse'] = _LISTNETWORKSERVERRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

NetworkServer = _reflection.GeneratedProtocolMessageType('NetworkServer', (_message.Message,), {
  'DESCRIPTOR' : _NETWORKSERVER,
  '__module__' : 'chirpstack_api.as_pb.external.api.networkServer_pb2'
  # @@protoc_insertion_point(class_scope:api.NetworkServer)
  })
_sym_db.RegisterMessage(NetworkServer)

NetworkServerListItem = _reflection.GeneratedProtocolMessageType('NetworkServerListItem', (_message.Message,), {
  'DESCRIPTOR' : _NETWORKSERVERLISTITEM,
  '__module__' : 'chirpstack_api.as_pb.external.api.networkServer_pb2'
  # @@protoc_insertion_point(class_scope:api.NetworkServerListItem)
  })
_sym_db.RegisterMessage(NetworkServerListItem)

CreateNetworkServerRequest = _reflection.GeneratedProtocolMessageType('CreateNetworkServerRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATENETWORKSERVERREQUEST,
  '__module__' : 'chirpstack_api.as_pb.external.api.networkServer_pb2'
  # @@protoc_insertion_point(class_scope:api.CreateNetworkServerRequest)
  })
_sym_db.RegisterMessage(CreateNetworkServerRequest)

CreateNetworkServerResponse = _reflection.GeneratedProtocolMessageType('CreateNetworkServerResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATENETWORKSERVERRESPONSE,
  '__module__' : 'chirpstack_api.as_pb.external.api.networkServer_pb2'
  # @@protoc_insertion_point(class_scope:api.CreateNetworkServerResponse)
  })
_sym_db.RegisterMessage(CreateNetworkServerResponse)

GetNetworkServerRequest = _reflection.GeneratedProtocolMessageType('GetNetworkServerRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETNETWORKSERVERREQUEST,
  '__module__' : 'chirpstack_api.as_pb.external.api.networkServer_pb2'
  # @@protoc_insertion_point(class_scope:api.GetNetworkServerRequest)
  })
_sym_db.RegisterMessage(GetNetworkServerRequest)

GetNetworkServerResponse = _reflection.GeneratedProtocolMessageType('GetNetworkServerResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETNETWORKSERVERRESPONSE,
  '__module__' : 'chirpstack_api.as_pb.external.api.networkServer_pb2'
  # @@protoc_insertion_point(class_scope:api.GetNetworkServerResponse)
  })
_sym_db.RegisterMessage(GetNetworkServerResponse)

UpdateNetworkServerRequest = _reflection.GeneratedProtocolMessageType('UpdateNetworkServerRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATENETWORKSERVERREQUEST,
  '__module__' : 'chirpstack_api.as_pb.external.api.networkServer_pb2'
  # @@protoc_insertion_point(class_scope:api.UpdateNetworkServerRequest)
  })
_sym_db.RegisterMessage(UpdateNetworkServerRequest)

DeleteNetworkServerRequest = _reflection.GeneratedProtocolMessageType('DeleteNetworkServerRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETENETWORKSERVERREQUEST,
  '__module__' : 'chirpstack_api.as_pb.external.api.networkServer_pb2'
  # @@protoc_insertion_point(class_scope:api.DeleteNetworkServerRequest)
  })
_sym_db.RegisterMessage(DeleteNetworkServerRequest)

ListNetworkServerRequest = _reflection.GeneratedProtocolMessageType('ListNetworkServerRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTNETWORKSERVERREQUEST,
  '__module__' : 'chirpstack_api.as_pb.external.api.networkServer_pb2'
  # @@protoc_insertion_point(class_scope:api.ListNetworkServerRequest)
  })
_sym_db.RegisterMessage(ListNetworkServerRequest)

ListNetworkServerResponse = _reflection.GeneratedProtocolMessageType('ListNetworkServerResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTNETWORKSERVERRESPONSE,
  '__module__' : 'chirpstack_api.as_pb.external.api.networkServer_pb2'
  # @@protoc_insertion_point(class_scope:api.ListNetworkServerResponse)
  })
_sym_db.RegisterMessage(ListNetworkServerResponse)


DESCRIPTOR._options = None

_NETWORKSERVERSERVICE = _descriptor.ServiceDescriptor(
  name='NetworkServerService',
  full_name='api.NetworkServerService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1456,
  serialized_end=2014,
  methods=[
  _descriptor.MethodDescriptor(
    name='Create',
    full_name='api.NetworkServerService.Create',
    index=0,
    containing_service=None,
    input_type=_CREATENETWORKSERVERREQUEST,
    output_type=_CREATENETWORKSERVERRESPONSE,
    serialized_options=b'\202\323\344\223\002\031\"\024/api/network-servers:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Get',
    full_name='api.NetworkServerService.Get',
    index=1,
    containing_service=None,
    input_type=_GETNETWORKSERVERREQUEST,
    output_type=_GETNETWORKSERVERRESPONSE,
    serialized_options=b'\202\323\344\223\002\033\022\031/api/network-servers/{id}',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Update',
    full_name='api.NetworkServerService.Update',
    index=2,
    containing_service=None,
    input_type=_UPDATENETWORKSERVERREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=b'\202\323\344\223\002-\032(/api/network-servers/{network_server.id}:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Delete',
    full_name='api.NetworkServerService.Delete',
    index=3,
    containing_service=None,
    input_type=_DELETENETWORKSERVERREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=b'\202\323\344\223\002\033*\031/api/network-servers/{id}',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='List',
    full_name='api.NetworkServerService.List',
    index=4,
    containing_service=None,
    input_type=_LISTNETWORKSERVERREQUEST,
    output_type=_LISTNETWORKSERVERRESPONSE,
    serialized_options=b'\202\323\344\223\002\026\022\024/api/network-servers',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_NETWORKSERVERSERVICE)

DESCRIPTOR.services_by_name['NetworkServerService'] = _NETWORKSERVERSERVICE

# @@protoc_insertion_point(module_scope)
