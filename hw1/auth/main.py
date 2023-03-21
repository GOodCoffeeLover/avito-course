import auth_pb2
import auth_pb2_grpc
import grpc
import logging, concurrent.futures


LIST_OF_ALLOWED_USERS = ['admin', 'max', 'alex']

class Auther(auth_pb2_grpc.AutherServicer):
    def AuthByName(self, req: auth_pb2.AuthByNameRequest, context: grpc.RpcContext ) -> auth_pb2.AuthByNameResponse :
        res = (req.name.lower() in LIST_OF_ALLOWED_USERS)
        return auth_pb2.AuthByNameResponse(authed=res)

def serve():
    port = '50051'
    server = grpc.server(concurrent.futures.ThreadPoolExecutor(max_workers=10))
    auth_pb2_grpc.add_AutherServicer_to_server(Auther(), server)
    server.add_insecure_port('[::]:' + port)
    server.start()
    print("Server started, listening on " + port)
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
