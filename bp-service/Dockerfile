FROM golang:1.25.4-alpine AS builder

WORKDIR /home/src/app/

COPY bp-service/go.mod bp-service/go.sum ./bp-service_run/

WORKDIR /home/src/app/bp-service_run

RUN go mod download

COPY bp-service/ .

EXPOSE 8085
EXPOSE 5432
CMD [ "go", "run", "cmd/main.go" ]