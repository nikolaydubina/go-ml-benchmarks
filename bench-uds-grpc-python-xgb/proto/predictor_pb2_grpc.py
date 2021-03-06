# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto import predictor_pb2 as proto_dot_predictor__pb2


class PredictorStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Predict = channel.unary_unary(
                '/predictor.Predictor/Predict',
                request_serializer=proto_dot_predictor__pb2.PredictRequest.SerializeToString,
                response_deserializer=proto_dot_predictor__pb2.PredictResponse.FromString,
                )
        self.PredictProcessed = channel.unary_unary(
                '/predictor.Predictor/PredictProcessed',
                request_serializer=proto_dot_predictor__pb2.PredictProcessedRequest.SerializeToString,
                response_deserializer=proto_dot_predictor__pb2.PredictResponse.FromString,
                )


class PredictorServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Predict(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PredictProcessed(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_PredictorServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Predict': grpc.unary_unary_rpc_method_handler(
                    servicer.Predict,
                    request_deserializer=proto_dot_predictor__pb2.PredictRequest.FromString,
                    response_serializer=proto_dot_predictor__pb2.PredictResponse.SerializeToString,
            ),
            'PredictProcessed': grpc.unary_unary_rpc_method_handler(
                    servicer.PredictProcessed,
                    request_deserializer=proto_dot_predictor__pb2.PredictProcessedRequest.FromString,
                    response_serializer=proto_dot_predictor__pb2.PredictResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'predictor.Predictor', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Predictor(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Predict(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/predictor.Predictor/Predict',
            proto_dot_predictor__pb2.PredictRequest.SerializeToString,
            proto_dot_predictor__pb2.PredictResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PredictProcessed(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/predictor.Predictor/PredictProcessed',
            proto_dot_predictor__pb2.PredictProcessedRequest.SerializeToString,
            proto_dot_predictor__pb2.PredictResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
