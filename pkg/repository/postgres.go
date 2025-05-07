package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	SSLMode  string
}

const (
	userDb = "users"
	chatDb = "chats"
	msgDb  = "messages"
)

func NewDB(db DB) (*sqlx.DB, error) {
	query := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", db.Host, db.Port, db.Username, db.Database, db.Password, db.SSLMode)
	postgres, err := sqlx.Connect("postgres", query)
	if err != nil {
		return nil, err
	}
	err = postgres.Ping()
	if err != nil {
		log.Fatalf("postgres.go: error connecting to database: %s", err.Error())
	}
	return postgres, nil
}
