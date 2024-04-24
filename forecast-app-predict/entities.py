import predict_pb2


class MakePredictRequestEntity(object):
    def __init__(self, request: predict_pb2.MakePredictRequest):
        self.name = request.name
        self.username = request.username
        self.unit = request.unit
        self.period = request.period
        self.items = list(request.items)
