FROM golang:1.20-alpine

WORKDIR /app

# Копируем ВСЕ файлы бэкенда
COPY . .

# Собираем приложение
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/backend

EXPOSE 8080

CMD ["/app/backend"]