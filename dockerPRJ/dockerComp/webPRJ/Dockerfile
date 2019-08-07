FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/webPRJ

COPY ../webPRJ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o webPRJ .

FROM alpine:3.9 as prod

WORKDIR /root/webPRJ

COPY --from=0 /go/src/webPRJ .

#EXPOSE 8500

CMD ["./webPRJ"]