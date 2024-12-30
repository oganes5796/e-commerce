FROM golang:1.23
# Устанавливаем рабочую директорию
WORKDIR /app
# Копируем go.mod и go.sum
COPY go.mod go.sum ./
# Устанавливаем зависимости
RUN go mod download
# Копируем всё приложение
COPY . .
# Собираем бинарник
RUN go build -o main ./cmd/main.go
# Указываем порт, на котором работает приложение
EXPOSE 8080
# Запускаем приложение
CMD ["./main"]