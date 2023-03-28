FROM python:latest

WORKDIR /app

RUN python -m pip install grpcio grpcio-tools 

COPY auth_pb2_grpc.py  auth_pb2.py  auth_pb2.pyi  client.py  main.py ./
ENV GRPC_VERBOSITY=debug

CMD python main.py