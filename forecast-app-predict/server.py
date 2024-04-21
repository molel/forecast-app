import argparse
from datetime import datetime, timedelta
import json

import grpc

import db
import predict_pb2, predict_pb2_grpc


class PredictServer(predict_pb2_grpc.PredictServiceServicer):
    def __init__(self, config: argparse.Namespace):
        self.config = config
        self.db = db.connect_to_db(self.config)
        super(PredictServer, self).__init__()

    @staticmethod
    def predict(request: predict_pb2.MakePredictRequest) -> list[predict_pb2.TimeSeriesItem]:
        return [predict_pb2.TimeSeriesItem()]

    async def MakePredict(
            self,
            request: predict_pb2.MakePredictRequest,
            context: grpc.aio.ServicerContext
    ) -> predict_pb2.MakePredictResponse:
        items = []
        for i in range(100):
            items.append({"ts": (datetime.now() + timedelta(days=i)).timestamp(), "value": i})
        return predict_pb2.MakePredictResponse(
            items=items
        )

    async def GetPredict(
            self,
            request: predict_pb2.GetPredictRequest,
            context: grpc.aio.ServicerContext
    ) -> predict_pb2.GetPredictResponse:
        items = []
        for i in range(100):
            items.append(
                predict_pb2.TimeSeriesItem(ts=int((datetime.now() + timedelta(days=i)).timestamp() * 1e9), value=i))
        return predict_pb2.GetPredictResponse(
            unit='шт',
            delimiter=int(len(items) * 0.8),
            items=items
        )


async def serve(config: argparse.Namespace) -> None:
    server = grpc.aio.server()
    predict_pb2_grpc.add_PredictServiceServicer_to_server(PredictServer(config), server)
    listen_addr = f'localhost:{config.http}'
    server.add_insecure_port(listen_addr)
    print(f'Starting server on {listen_addr}')
    await server.start()
    await server.wait_for_termination()
