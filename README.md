# **Развертывание бэкенда Verbify** 📕💻

## **Технологический стек** 🛠️
- **Язык**: Go (Golang) 🐹
- **Фреймворк**: Gin 🍸
- **База данных**: PostgreSQL 🐘
- **Миграции**: golang-migrate 📦

## **Быстрый старт** ⚡

1️⃣ **Запустите PostgreSQL в Docker одной командой**:
```bash
docker run --name whyai-db -e POSTGRES_DB=whyai-db -e POSTGRES_USER=qwerty -e POSTGRES_PASSWORD=qwerty1 -p 5432:5432 -d postgres
```

2️⃣ **Примените миграции**:
```bash
migrate -path ./schema -database "postgres://qwerty:qwerty1@localhost:5432/whyai-db?sslmode=disable" up
```

3️⃣ **Запустите сервер**:
```bash
go mod download && go run main.go
```

## **Дополнительные команды** 🔧

| Команда | Описание |
|---------|----------|
| `docker start whyai-db` | Запустить контейнер с БД |
| `docker stop whyai-db` | Остановить контейнер |
| `migrate ... down` | Откатить миграции |

Сервер будет доступен на `http://localhost:8080` 🎉

> **Примечание**: Примеры API-запросов находятся в папке `requests` 📂