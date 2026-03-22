FROM golang:1.25.4-alpine AS builder

WORKDIR /home/src/app/

COPY bp-telegram-bot/go.mod bp-telegram-bot/go.sum ./bp-telegram-bot_run/

WORKDIR /home/src/app/bp-telegram-bot_run

RUN go mod download

COPY bp-telegram-bot/ .

CMD [ "go", "run", "cmd/bot/main.go" ]