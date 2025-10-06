FROM golang:1.25.1-alpine

WORKDIR /app

# Устанавливаем зависимости для postgres
RUN apk add --no-cache gcc musl-dev postgresql-dev

# Копируем файлы зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o main .

EXPOSE 8080

CMD ["./main"]