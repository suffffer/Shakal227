# Стадия сборки
FROM golang:1.23.3 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь код проекта в контейнер
COPY . .

# Компилируем приложение
RUN go build -o server ./cmd/server/main.go

# Стадия запуска
FROM ubuntu:22.04

# Устанавливаем необходимые зависимости
RUN apt-get update && apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем скомпилированное приложение из стадии сборки
COPY --from=builder /app/server .

# Указываем команду для запуска приложения
CMD ["./server"]