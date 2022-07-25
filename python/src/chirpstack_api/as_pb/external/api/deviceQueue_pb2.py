# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chirpstack-api/as_pb/external/api/deviceQueue.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='chirpstack-api/as_pb/external/api/deviceQueue.proto',
  package='api',
  syntax='proto3',
  serialized_options=b'\n!io.chirpstack.api.as.external.apiB\020DeviceQueueProtoP\001Z7github.com/ibrahimozekici/chirpstack-api/go/v5/as/external/api',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n3chirpstack-api/as_pb/external/api/deviceQueue.proto\x12\x03\x61pi\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\"\x7f\n\x0f\x44\x65viceQueueItem\x12\x17\n\x07\x64\x65v_eui\x18\x01 \x01(\tR\x06\x64\x65vEUI\x12\x11\n\tconfirmed\x18\x02 \x01(\x08\x12\r\n\x05\x66_cnt\x18\x06 \x01(\r\x12\x0e\n\x06\x66_port\x18\x03 \x01(\r\x12\x0c\n\x04\x64\x61ta\x18\x04 \x01(\x0c\x12\x13\n\x0bjson_object\x18\x05 \x01(\t\"P\n\x1d\x45nqueueDeviceQueueItemRequest\x12/\n\x11\x64\x65vice_queue_item\x18\x01 \x01(\x0b\x32\x14.api.DeviceQueueItem\"/\n\x1e\x45nqueueDeviceQueueItemResponse\x12\r\n\x05\x66_cnt\x18\x01 \x01(\r\"2\n\x17\x46lushDeviceQueueRequest\x12\x17\n\x07\x64\x65v_eui\x18\x01 \x01(\tR\x06\x64\x65vEUI\"J\n\x1bListDeviceQueueItemsRequest\x12\x17\n\x07\x64\x65v_eui\x18\x01 \x01(\tR\x06\x64\x65vEUI\x12\x12\n\ncount_only\x18\x02 \x01(\x08\"e\n\x1cListDeviceQueueItemsResponse\x12\x30\n\x12\x64\x65vice_queue_items\x18\x01 \x03(\x0b\x32\x14.api.DeviceQueueItem\x12\x13\n\x0btotal_count\x18\x02 \x01(\r2\xfc\x02\n\x12\x44\x65viceQueueService\x12\x8d\x01\n\x07\x45nqueue\x12\".api.EnqueueDeviceQueueItemRequest\x1a#.api.EnqueueDeviceQueueItemResponse\"9\x82\xd3\xe4\x93\x02\x33\"./api/devices/{device_queue_item.dev_eui}/queue:\x01*\x12\x63\n\x05\x46lush\x12\x1c.api.FlushDeviceQueueRequest\x1a\x16.google.protobuf.Empty\"$\x82\xd3\xe4\x93\x02\x1e*\x1c/api/devices/{dev_eui}/queue\x12q\n\x04List\x12 .api.ListDeviceQueueItemsRequest\x1a!.api.ListDeviceQueueItemsResponse\"$\x82\xd3\xe4\x93\x02\x1e\x12\x1c/api/devices/{dev_eui}/queueBp\n!io.chirpstack.api.as.external.apiB\x10\x44\x65viceQueueProtoP\x01Z7github.com/ibrahimozekici/chirpstack-api/go/v5/as/external/apib\x06proto3'
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,])




_DEVICEQUEUEITEM = _descriptor.Descriptor(
  name='DeviceQueueItem',
  full_name='api.DeviceQueueItem',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='dev_eui', full_name='api.DeviceQueueItem.dev_eui', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='devEUI', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='confirmed', full_name='api.DeviceQueueItem.confirmed', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='f_cnt', full_name='api.DeviceQueueItem.f_cnt', index=2,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='f_port', full_name='api.DeviceQueueItem.f_port', index=3,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='data', full_name='api.DeviceQueueItem.data', index=4,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='json_object', full_name='api.DeviceQueueItem.json_object', index=5,
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
  serialized_start=119,
  serialized_end=246,
)


_ENQUEUEDEVICEQUEUEITEMREQUEST = _descriptor.Descriptor(
  name='EnqueueDeviceQueueItemRequest',
  full_name='api.EnqueueDeviceQueueItemRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='device_queue_item', full_name='api.EnqueueDeviceQueueItemRequest.device_queue_item', index=0,
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
  serialized_start=248,
  serialized_end=328,
)


_ENQUEUEDEVICEQUEUEITEMRESPONSE = _descriptor.Descriptor(
  name='EnqueueDeviceQueueItemResponse',
  full_name='api.EnqueueDeviceQueueItemResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='f_cnt', full_name='api.EnqueueDeviceQueueItemResponse.f_cnt', index=0,
      number=1, type=13, cpp_type=3, label=1,
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
  serialized_start=330,
  serialized_end=377,
)


_FLUSHDEVICEQUEUEREQUEST = _descriptor.Descriptor(
  name='FlushDeviceQueueRequest',
  full_name='api.FlushDeviceQueueRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='dev_eui', full_name='api.FlushDeviceQueueRequest.dev_eui', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='devEUI', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=379,
  serialized_end=429,
)


_LISTDEVICEQUEUEITEMSREQUEST = _descriptor.Descriptor(
  name='ListDeviceQueueItemsRequest',
  full_name='api.ListDeviceQueueItemsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='dev_eui', full_name='api.ListDeviceQueueItemsRequest.dev_eui', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='devEUI', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='count_only', full_name='api.ListDeviceQueueItemsRequest.count_only', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=431,
  serialized_end=505,
)


_LISTDEVICEQUEUEITEMSRESPONSE = _descriptor.Descriptor(
  name='ListDeviceQueueItemsResponse',
  full_name='api.ListDeviceQueueItemsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='device_queue_items', full_name='api.ListDeviceQueueItemsResponse.device_queue_items', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='total_count', full_name='api.ListDeviceQueueItemsResponse.total_count', index=1,
      number=2, type=13, cpp_type=3, label=1,
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
  serialized_start=507,
  serialized_end=608,
)

_ENQUEUEDEVICEQUEUEITEMREQUEST.fields_by_name['device_queue_item'].message_type = _DEVICEQUEUEITEM
_LISTDEVICEQUEUEITEMSRESPONSE.fields_by_name['device_queue_items'].message_type = _DEVICEQUEUEITEM
DESCRIPTOR.message_types_by_name['DeviceQueueItem'] = _DEVICEQUEUEITEM
DESCRIPTOR.message_types_by_name['EnqueueDeviceQueueItemRequest'] = _ENQUEUEDEVICEQUEUEITEMREQUEST
DESCRIPTOR.message_types_by_name['EnqueueDeviceQueueItemResponse'] = _ENQUEUEDEVICEQUEUEITEMRESPONSE
DESCRIPTOR.message_types_by_name['FlushDeviceQueueRequest'] = _FLUSHDEVICEQUEUEREQUEST
DESCRIPTOR.message_types_by_name['ListDeviceQueueItemsRequest'] = _LISTDEVICEQUEUEITEMSREQUEST
DESCRIPTOR.message_types_by_name['ListDeviceQueueItemsResponse'] = _LISTDEVICEQUEUEITEMSRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DeviceQueueItem = _reflection.GeneratedProtocolMessageType('DeviceQueueItem', (_message.Message,), {
  'DESCRIPTOR' : _DEVICEQUEUEITEM,
  '__module__' : 'chirpstack_api.as_pb.external.api.deviceQueue_pb2'
  # @@protoc_insertion_point(class_scope:api.DeviceQueueItem)
  })
_sym_db.RegisterMessage(DeviceQueueItem)

EnqueueDeviceQueueItemRequest = _reflection.GeneratedProtocolMessageType('EnqueueDeviceQueueItemRequest', (_message.Message,), {
  'DESCRIPTOR' : _ENQUEUEDEVICEQUEUEITEMREQUEST,
  '__module__' : 'chirpstack_api.as_pb.external.api.deviceQueue_pb2'
  # @@protoc_insertion_point(class_scope:api.EnqueueDeviceQueueItemRequest)
  })
_sym_db.RegisterMessage(EnqueueDeviceQueueItemRequest)

EnqueueDeviceQueueItemResponse = _reflection.GeneratedProtocolMessageType('EnqueueDeviceQueueItemResponse', (_message.Message,), {
  'DESCRIPTOR' : _ENQUEUEDEVICEQUEUEITEMRESPONSE,
  '__module__' : 'chirpstack_api.as_pb.external.api.deviceQueue_pb2'
  # @@protoc_insertion_point(class_scope:api.EnqueueDeviceQueueItemResponse)
  })
_sym_db.RegisterMessage(EnqueueDeviceQueueItemResponse)

FlushDeviceQueueRequest = _reflection.GeneratedProtocolMessageType('FlushDeviceQueueRequest', (_message.Message,), {
  'DESCRIPTOR' : _FLUSHDEVICEQUEUEREQUEST,
  '__module__' : 'chirpstack_api.as_pb.external.api.deviceQueue_pb2'
  # @@protoc_insertion_point(class_scope:api.FlushDeviceQueueRequest)
  })
_sym_db.RegisterMessage(FlushDeviceQueueRequest)

ListDeviceQueueItemsRequest = _reflection.GeneratedProtocolMessageType('ListDeviceQueueItemsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTDEVICEQUEUEITEMSREQUEST,
  '__module__' : 'chirpstack_api.as_pb.external.api.deviceQueue_pb2'
  # @@protoc_insertion_point(class_scope:api.ListDeviceQueueItemsRequest)
  })
_sym_db.RegisterMessage(ListDeviceQueueItemsRequest)

ListDeviceQueueItemsResponse = _reflection.GeneratedProtocolMessageType('ListDeviceQueueItemsResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTDEVICEQUEUEITEMSRESPONSE,
  '__module__' : 'chirpstack_api.as_pb.external.api.deviceQueue_pb2'
  # @@protoc_insertion_point(class_scope:api.ListDeviceQueueItemsResponse)
  })
_sym_db.RegisterMessage(ListDeviceQueueItemsResponse)


DESCRIPTOR._options = None

_DEVICEQUEUESERVICE = _descriptor.ServiceDescriptor(
  name='DeviceQueueService',
  full_name='api.DeviceQueueService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=611,
  serialized_end=991,
  methods=[
  _descriptor.MethodDescriptor(
    name='Enqueue',
    full_name='api.DeviceQueueService.Enqueue',
    index=0,
    containing_service=None,
    input_type=_ENQUEUEDEVICEQUEUEITEMREQUEST,
    output_type=_ENQUEUEDEVICEQUEUEITEMRESPONSE,
    serialized_options=b'\202\323\344\223\0023\"./api/devices/{device_queue_item.dev_eui}/queue:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='Flush',
    full_name='api.DeviceQueueService.Flush',
    index=1,
    containing_service=None,
    input_type=_FLUSHDEVICEQUEUEREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    serialized_options=b'\202\323\344\223\002\036*\034/api/devices/{dev_eui}/queue',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='List',
    full_name='api.DeviceQueueService.List',
    index=2,
    containing_service=None,
    input_type=_LISTDEVICEQUEUEITEMSREQUEST,
    output_type=_LISTDEVICEQUEUEITEMSRESPONSE,
    serialized_options=b'\202\323\344\223\002\036\022\034/api/devices/{dev_eui}/queue',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_DEVICEQUEUESERVICE)

DESCRIPTOR.services_by_name['DeviceQueueService'] = _DEVICEQUEUESERVICE

# @@protoc_insertion_point(module_scope)
