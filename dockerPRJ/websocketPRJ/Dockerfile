FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/websocketPRJ

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o websocketPRJ .

FROM alpine:3.9 as prod

WORKDIR /root/websocketPRJ

COPY --from=0 /go/src/websocketPRJ .

EXPOSE 8080


CMD ["./websocketPRJ"]
