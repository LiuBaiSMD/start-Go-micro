FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/apiPRJ

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o apiPRJ .

FROM alpine:3.9 as prod

WORKDIR /root/apiPRJ

COPY --from=0 /go/src/apiPRJ .

#EXPOSE 8500

CMD ["./apiPRJ"]
