# Строим Go проект в первом контейнере
FROM golang:1.23.1-alpine AS builder

WORKDIR /app

# Копируем исходники
COPY . .

# Загружаем зависимости
RUN go mod download

# Собираем бинарный файл
RUN go build -o library_service cmd/main.go

# Создаем финальный минимальный образ
FROM alpine:latest

WORKDIR /root/

# Копируем бинарный файл из предыдущего контейнера
COPY --from=builder /app/library_service .

# Копируем скрипт миграции для использования
COPY wait-for-it.sh .

# Делаем скрипт исполняемым
RUN ["chmod", "+x", "/node/execute.sh"]

# Указываем точку входа
ENTRYPOINT ["./library_service"]
