FROM golang:1.22-alpine AS builder

WORKDIR /application

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем приложение
RUN go build -o main ./cmd/taskapi.go

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

RUN mkdir -p /app/logs

# Копируем только бинарный файл
COPY --from=builder /application/main .

EXPOSE 8080

# Запускаем готовый бинарный файл
CMD ["./main"]