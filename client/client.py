from __future__ import absolute_import

import grpc

from pb.service_pb2 import Message
from pb.service_pb2_grpc import HelloServiceStub

def run():
    channel = grpc.insecure_channel('localhost:50000')
    stub = HelloServiceStub(channel)
    message = Message(id='1', msg='world')

    response = stub.Hello(message)
    print("Hello {0}!".format(response.msg))


if __name__ == '__main__':
    run()