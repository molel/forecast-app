import argparse
import multiprocessing

import grpc

import handlers
import predict_pb2
import predict_pb2_grpc
from entities import MakePredictRequestEntity
from logging.handlers import QueueListener


class PredictServer(predict_pb2_grpc.PredictServiceServicer):
    def __init__(self, config: argparse.Namespace, queue: multiprocessing.Queue, log_queue: multiprocessing.Queue):
        self.db = handlers.GetHandler(config, log_queue)
        self.queue = queue
        super(PredictServer, self).__init__()

    @staticmethod
    def predict(request: predict_pb2.MakePredictRequest) -> list[predict_pb2.TimeSeriesItem]:
        return [predict_pb2.TimeSeriesItem()]

    async def GetPredict(
            self,
            request: predict_pb2.GetPredictRequest,
            context: grpc.aio.ServicerContext
    ) -> predict_pb2.GetPredictResponse:
        unit, delimiter, period, items = self.db.get_prediction(request.username, request.name)
        items = [predict_pb2.TimeSeriesItem(ts=ts, value=value) for ts, value in items]
        return predict_pb2.GetPredictResponse(
            unit=unit,
            delimiter=delimiter,
            period=period,
            items=items
        )

    async def GetPredicts(
            self,
            request: predict_pb2.GetPredictRequest,
            context: grpc.aio.ServicerContext
    ) -> predict_pb2.GetPredictsResponse:
        names = self.db.get_predictions(request.username)
        return predict_pb2.GetPredictsResponse(names=names)

    async def MakePredict(
            self,
            request: predict_pb2.MakePredictRequest,
            context: grpc.aio.ServicerContext
    ) -> predict_pb2.Empty:
        self.queue.put(MakePredictRequestEntity(request=request))
        return predict_pb2.Empty()


async def serve(config: argparse.Namespace) -> None:
    manager = multiprocessing.Manager()
    queue = manager.Queue()
    log_queue = manager.Queue()

    grpc_server = grpc.aio.server()
    prediction_server = PredictServer(config, queue, log_queue)
    predict_pb2_grpc.add_PredictServiceServicer_to_server(prediction_server, grpc_server)

    listen_addr = f'0.0.0.0:{config.http}'
    grpc_server.add_insecure_port(listen_addr)
    print(f'Starting grpc_server on {listen_addr}')

    for i in range(config.pool):
        multiprocessing.Process(target=handlers.PredictHandler, args=(config, queue, log_queue)).start()

    log_queue_listener = QueueListener(log_queue)
    log_queue_listener.start()

    await grpc_server.start()
    await grpc_server.wait_for_termination()

    log_queue_listener.stop()
