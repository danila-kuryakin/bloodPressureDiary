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
#RUN CGO_ENABLED=0 GOOS=linux go build -o bp-service ./cmd/main.go
#
## Финальный образ
#FROM alpine:3.20
#
## Устанавливаем часовой пояс (опционально)
#RUN apk --no-cache add tzdata
#WORKDIR /app
#
#COPY --from=builder /app/bp-service .
#COPY --from=builder /app/migrations ./migrations
#COPY --from=builder /app/configs/config.yml ./configs/
##COPY .env .
#
## Запускаем приложение
#EXPOSE 8085
#EXPOSE 5432
#CMD ["./bp-service"]


FROM golang:1.25.4-alpine AS builder

WORKDIR /home/src/app/

COPY bp-service/go.mod bp-service/go.sum ./bp-service_run/

WORKDIR /home/src/app/bp-service_run

RUN go mod download

COPY bp-service/ .

EXPOSE 8085
EXPOSE 5432
CMD [ "go", "run", "cmd/main.go" ]