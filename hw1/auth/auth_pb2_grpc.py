# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import auth_pb2 as auth__pb2


class AutherStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.AuthByName = channel.unary_unary(
                '/auth.Auther/AuthByName',
                request_serializer=auth__pb2.AuthByNameRequest.SerializeToString,
                response_deserializer=auth__pb2.AuthByNameResponse.FromString,
                )


class AutherServicer(object):
    """Missing associated documentation comment in .proto file."""

    def AuthByName(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AutherServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'AuthByName': grpc.unary_unary_rpc_method_handler(
                    servicer.AuthByName,
                    request_deserializer=auth__pb2.AuthByNameRequest.FromString,
                    response_serializer=auth__pb2.AuthByNameResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'auth.Auther', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Auther(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def AuthByName(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/auth.Auther/AuthByName',
            auth__pb2.AuthByNameRequest.SerializeToString,
            auth__pb2.AuthByNameResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
