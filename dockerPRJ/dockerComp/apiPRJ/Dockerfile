FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/apiPRJ

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o apiPRJ .

##测试consul
#FROM consul:1.5.2
#
#WORKDIR /root/consul
#
#EXPOSE 8500

FROM alpine:3.9 as prod

WORKDIR /root/apiPRJ

COPY --from=0 /go/src/apiPRJ .

CMD ["./apiPRJ"]




