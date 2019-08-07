FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/rpcPRJ

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rpcPRJ .

FROM alpine:3.9 as prod

WORKDIR /root/rpcPRJ

COPY --from=0 /go/src/rpcPRJ .

EXPOSE 8080

CMD ["./rpcPRJ"]