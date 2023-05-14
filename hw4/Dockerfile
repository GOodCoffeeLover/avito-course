FROM golang:1.20 as build

WORKDIR /app

COPY ./main.go ./

ENV CGO_ENABLED=0

RUN ["go", "build", "-o", "/app/httpServer", "./main.go"]

#=====================================

FROM scratch

COPY --from=build /app/httpServer /httpServer

CMD ["/httpServer"]