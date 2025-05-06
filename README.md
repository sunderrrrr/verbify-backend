docker run --name whyai-db -e POSTGRES_DB=whyai-db -e POSTGRES_USER=qwerty -e POSTGRES_PASSWORD=qwerty1 -p 5432:5432 -d postgres

migrate -path ./schema -database "postgres://qwerty:qwerty1@localhost:5432/whyai-db?sslmode=disable" up