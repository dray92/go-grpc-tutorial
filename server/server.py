
from concurrent import futures
import time

import grpc


from pb.service_pb2 import Message
from pb.service_pb2_grpc import (
    HelloServiceServicer,
    add_HelloServiceServicer_to_server
)

class HelloServer(HelloServiceServicer):
    def Hello(self, request, context):
        # being lazy and returning same message
        return request


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    add_HelloServiceServicer_to_server(HelloServer(), server)
    server.add_insecure_port('[::]:50000')
    server.start()
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()