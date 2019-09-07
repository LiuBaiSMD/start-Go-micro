FROM golang:1.12.4-alpine3.9 as builder

WORKDIR /go/src/myOwnStation

COPY . .

ENV dockerMode on

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myOwnStation .

FROM alpine:3.9 as prod

WORKDIR /root/myOwnStation

COPY --from=0 /go/src/myOwnStation .

CMD ["./userwebPRJ"]