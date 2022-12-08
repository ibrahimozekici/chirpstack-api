# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chirpstack-api/geo/geo.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from chirpstack_api.gw import gw_pb2 as chirpstack__api_dot_gw_dot_gw__pb2
from chirpstack_api.common import common_pb2 as chirpstack__api_dot_common_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1c\x63hirpstack-api/geo/geo.proto\x12\x03geo\x1a\x1a\x63hirpstack-api/gw/gw.proto\x1a\"chirpstack-api/common/common.proto\"3\n\rResolveResult\x12\"\n\x08location\x18\x01 \x01(\x0b\x32\x10.common.Location\"0\n\x0b\x46rameRXInfo\x12!\n\x07rx_info\x18\x01 \x03(\x0b\x32\x10.gw.UplinkRXInfo\"\x86\x01\n\x12ResolveTDOARequest\x12\x17\n\x07\x64\x65v_eui\x18\x01 \x01(\x0cR\x06\x64\x65vEUI\x12\x34\n\rframe_rx_info\x18\x02 \x01(\x0b\x32\x10.geo.FrameRXInfoR\x0b\x66rameRXInfo\x12!\n\x19\x64\x65vice_reference_altitude\x18\x03 \x01(\x01\"\x97\x01\n\x1cResolveMultiFrameTDOARequest\x12\x17\n\x07\x64\x65v_eui\x18\x01 \x01(\x0cR\x06\x64\x65vEUI\x12;\n\x11\x66rame_rx_info_set\x18\x02 \x03(\x0b\x32\x10.geo.FrameRXInfoR\x0e\x66rameRXInfoSet\x12!\n\x19\x64\x65vice_reference_altitude\x18\x03 \x01(\x01\"9\n\x13ResolveTDOAResponse\x12\"\n\x06result\x18\x01 \x01(\x0b\x32\x12.geo.ResolveResult\"C\n\x1dResolveMultiFrameTDOAResponse\x12\"\n\x06result\x18\x01 \x01(\x0b\x32\x12.geo.ResolveResult2\xc0\x01\n\x18GeolocationServerService\x12\x42\n\x0bResolveTDOA\x12\x17.geo.ResolveTDOARequest\x1a\x18.geo.ResolveTDOAResponse\"\x00\x12`\n\x15ResolveMultiFrameTDOA\x12!.geo.ResolveMultiFrameTDOARequest\x1a\".geo.ResolveMultiFrameTDOAResponse\"\x00\x42^\n\x15io.chirpstack.api.geoB\x16GeolocationServerProtoP\x01Z+github.com/ibrahimozekici/chirpstack-api/go/v5/geob\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'chirpstack_api.geo.geo_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\025io.chirpstack.api.geoB\026GeolocationServerProtoP\001Z+github.com/ibrahimozekici/chirpstack-api/go/v5/geo'
  _RESOLVERESULT._serialized_start=101
  _RESOLVERESULT._serialized_end=152
  _FRAMERXINFO._serialized_start=154
  _FRAMERXINFO._serialized_end=202
  _RESOLVETDOAREQUEST._serialized_start=205
  _RESOLVETDOAREQUEST._serialized_end=339
  _RESOLVEMULTIFRAMETDOAREQUEST._serialized_start=342
  _RESOLVEMULTIFRAMETDOAREQUEST._serialized_end=493
  _RESOLVETDOARESPONSE._serialized_start=495
  _RESOLVETDOARESPONSE._serialized_end=552
  _RESOLVEMULTIFRAMETDOARESPONSE._serialized_start=554
  _RESOLVEMULTIFRAMETDOARESPONSE._serialized_end=621
  _GEOLOCATIONSERVERSERVICE._serialized_start=624
  _GEOLOCATIONSERVERSERVICE._serialized_end=816
# @@protoc_insertion_point(module_scope)
