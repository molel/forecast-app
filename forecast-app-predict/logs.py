from logging import getLogger, INFO
from logging.handlers import QueueHandler
import multiprocessing


def register_log_queue(log_queue: multiprocessing.Queue):
    handler = QueueHandler(log_queue)
    root = getLogger()
    root.addHandler(handler)
    root.setLevel(INFO)
