# Используем официальный образ Go
FROM golang:1.23.2

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем всё остальное
COPY . .

# Собираем бинарник
RUN go build -o main ./cmd

# Запускаем бинарник
CMD ["./main"]
