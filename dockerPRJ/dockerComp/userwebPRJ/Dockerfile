FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/userwebPRJ

COPY . .

ENV dockerMode on

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o userwebPRJ .

FROM alpine:3.9 as prod

WORKDIR /root/userwebPRJ

COPY --from=0 /go/src/userwebPRJ .

CMD ["./userwebPRJ"]