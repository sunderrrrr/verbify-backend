# Этап сборки
FROM golang:1.24 AS builder

WORKDIR /app


COPY go.mod go.sum ./
ENV GOPROXY=direct
RUN go mod download

COPY . .
RUN go build -o server cmd/main.go


FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Копируем бинарник и нужные статические файлы
COPY --from=builder /app/server .
COPY --from=builder /app/static ./static

ENV GIN_MODE=release

EXPOSE 8090

CMD ["./server"]