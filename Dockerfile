FROM golang:1.20 as build

WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN ["go", "mod", "download"]

COPY . .
ENV CGO_ENABLED=0
RUN ["go", "build", "-o", "server", "./cmd/main.go"]

# -----------------------------------------------

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY --from=build /app/server /

ENV LISTEN_PORT=7001

CMD ["/server"]
