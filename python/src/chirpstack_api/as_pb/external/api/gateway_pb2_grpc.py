# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from chirpstack_api.as_pb.external.api import gateway_pb2 as chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


class GatewayServiceStub(object):
    """GatewayService is the service managing the gateways.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Create = channel.unary_unary(
                '/api.GatewayService/Create',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.CreateGatewayRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.Get = channel.unary_unary(
                '/api.GatewayService/Get',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetGatewayRequest.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetGatewayResponse.FromString,
                )
        self.Update = channel.unary_unary(
                '/api.GatewayService/Update',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.UpdateGatewayRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.Delete = channel.unary_unary(
                '/api.GatewayService/Delete',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.DeleteGatewayRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.List = channel.unary_unary(
                '/api.GatewayService/List',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.ListGatewayRequest.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.ListGatewayResponse.FromString,
                )
        self.GetStats = channel.unary_unary(
                '/api.GatewayService/GetStats',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetGatewayStatsRequest.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetGatewayStatsResponse.FromString,
                )
        self.GetLastPing = channel.unary_unary(
                '/api.GatewayService/GetLastPing',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetLastPingRequest.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetLastPingResponse.FromString,
                )
        self.GenerateGatewayClientCertificate = channel.unary_unary(
                '/api.GatewayService/GenerateGatewayClientCertificate',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GenerateGatewayClientCertificateRequest.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GenerateGatewayClientCertificateResponse.FromString,
                )
        self.StreamFrameLogs = channel.unary_stream(
                '/api.GatewayService/StreamFrameLogs',
                request_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.StreamGatewayFrameLogsRequest.SerializeToString,
                response_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.StreamGatewayFrameLogsResponse.FromString,
                )


class GatewayServiceServicer(object):
    """GatewayService is the service managing the gateways.
    """

    def Create(self, request, context):
        """Create creates the given gateway.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Get(self, request, context):
        """Get returns the gateway for the requested mac address.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Update(self, request, context):
        """Update updates the gateway matching the given mac address.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Delete(self, request, context):
        """Delete deletes the gateway matching the given mac address.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def List(self, request, context):
        """List lists the gateways.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetStats(self, request, context):
        """GetStats lists the gateway stats given the query parameters.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetLastPing(self, request, context):
        """GetLastPing returns the last emitted ping and gateways receiving this ping.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GenerateGatewayClientCertificate(self, request, context):
        """GenerateGatewayClientCertificate returns TLS certificate gateway authentication / authorization.
        This endpoint can ony be used when ChirpStack Network Server is configured with a gateway
        CA certificate and key, which is used for signing the TLS certificate. The returned TLS
        certificate will have the Gateway ID as Common Name.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def StreamFrameLogs(self, request, context):
        """StreamFrameLogs streams the uplink and downlink frame-logs for the given gateway ID.
        Notes:
        * These are the raw LoRaWAN frames and this endpoint is intended for debugging only.
        * This endpoint does not work from a web-browser.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_GatewayServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Create': grpc.unary_unary_rpc_method_handler(
                    servicer.Create,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.CreateGatewayRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'Get': grpc.unary_unary_rpc_method_handler(
                    servicer.Get,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetGatewayRequest.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetGatewayResponse.SerializeToString,
            ),
            'Update': grpc.unary_unary_rpc_method_handler(
                    servicer.Update,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.UpdateGatewayRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'Delete': grpc.unary_unary_rpc_method_handler(
                    servicer.Delete,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.DeleteGatewayRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'List': grpc.unary_unary_rpc_method_handler(
                    servicer.List,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.ListGatewayRequest.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.ListGatewayResponse.SerializeToString,
            ),
            'GetStats': grpc.unary_unary_rpc_method_handler(
                    servicer.GetStats,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetGatewayStatsRequest.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetGatewayStatsResponse.SerializeToString,
            ),
            'GetLastPing': grpc.unary_unary_rpc_method_handler(
                    servicer.GetLastPing,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetLastPingRequest.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetLastPingResponse.SerializeToString,
            ),
            'GenerateGatewayClientCertificate': grpc.unary_unary_rpc_method_handler(
                    servicer.GenerateGatewayClientCertificate,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GenerateGatewayClientCertificateRequest.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GenerateGatewayClientCertificateResponse.SerializeToString,
            ),
            'StreamFrameLogs': grpc.unary_stream_rpc_method_handler(
                    servicer.StreamFrameLogs,
                    request_deserializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.StreamGatewayFrameLogsRequest.FromString,
                    response_serializer=chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.StreamGatewayFrameLogsResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'api.GatewayService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class GatewayService(object):
    """GatewayService is the service managing the gateways.
    """

    @staticmethod
    def Create(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.GatewayService/Create',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.CreateGatewayRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Get(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.GatewayService/Get',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetGatewayRequest.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetGatewayResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Update(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.GatewayService/Update',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.UpdateGatewayRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Delete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.GatewayService/Delete',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.DeleteGatewayRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def List(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.GatewayService/List',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.ListGatewayRequest.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.ListGatewayResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetStats(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.GatewayService/GetStats',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetGatewayStatsRequest.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetGatewayStatsResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetLastPing(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.GatewayService/GetLastPing',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetLastPingRequest.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GetLastPingResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GenerateGatewayClientCertificate(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.GatewayService/GenerateGatewayClientCertificate',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GenerateGatewayClientCertificateRequest.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.GenerateGatewayClientCertificateResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def StreamFrameLogs(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/api.GatewayService/StreamFrameLogs',
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.StreamGatewayFrameLogsRequest.SerializeToString,
            chirpstack__api_dot_as__pb_dot_external_dot_api_dot_gateway__pb2.StreamGatewayFrameLogsResponse.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)