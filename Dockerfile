FROM golang:1.19.2-alpine3.16 AS builder

COPY . /github.com/kirill0909/pocket-bot/
WORKDIR /github.com/kirill0909/pocket-bot/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/kirill0909/pocket-bot/bin/bot .
COPY --from=0 /github.com/kirill0909/pocket-bot/configs configs/

EXPOSE 80

CMD ["./bot"]
