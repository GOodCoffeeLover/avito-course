FROM golang:1.20 as build

COPY ./go.mod ./go.sum /

RUN go mod download

COPY ./main.go /

ENV CGO_ENABLED=0

RUN ["go", "build",   "-o", "/sender", "/main.go"]

FROM scratch

COPY  --from=build /sender /

CMD ["/sender"]