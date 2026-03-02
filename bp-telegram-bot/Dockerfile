## Этап сборки
#FROM golang:1.25.4-alpine AS builder
#
#WORKDIR /app
#
## Кэшируем зависимости
#COPY go.mod go.sum ./
#RUN go mod download
#
## Копируем исходники
#COPY . .
#
## Собираем приложение
#RUN CGO_ENABLED=0 GOOS=linux go build -o bp-telegram-bot ./cmd/bot/main.go
#
## Финальный образ
#FROM alpine:3.20
#
## Устанавливаем часовой пояс (опционально)
#RUN apk --no-cache add tzdata
#WORKDIR /app
#
#COPY --from=builder /app/bp-telegram-bot .
#COPY .env .
#
## Запускаем приложение
#EXPOSE 8085
#CMD ["./bp-telegram-bot"]

FROM golang:1.25.4-alpine AS builder

WORKDIR /home/src/app/

COPY bp-telegram-bot/go.mod bp-telegram-bot/go.sum ./bp-telegram-bot_run/

WORKDIR /home/src/app/bp-telegram-bot_run

RUN go mod download

COPY bp-telegram-bot/ .

CMD [ "go", "run", "cmd/bot/main.go" ]