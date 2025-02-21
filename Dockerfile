# Этап сборки
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Копируем файлы с зависимостями
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app

# Финальный этап
FROM alpine:latest

WORKDIR /app

# Копируем бинарный файл из этапа сборки
COPY --from=builder /app/main .

# Указываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"] 