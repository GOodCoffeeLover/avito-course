import auth_pb2_grpc
import auth_pb2
import grpc

def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = auth_pb2_grpc.AutherStub(channel)
        response = stub.AuthByName(auth_pb2.AuthByNameRequest(name='admin'))
        print("Auth client received: " + str(response.authed))
        response = stub.AuthByName(auth_pb2.AuthByNameRequest(name='not_admin'))
        print("Auth client received: " + str(response.authed))
if __name__ == '__main__':
    run()