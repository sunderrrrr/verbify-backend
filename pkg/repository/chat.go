package repository

import (
	"WhyAi/models"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ChatPostgres struct {
	db *sqlx.DB
}

func NewChatPostgres(db *sqlx.DB) *ChatPostgres {
	return &ChatPostgres{db: db}
}
func (p *ChatPostgres) ChatExist(taskId, userId int) bool {
	var exists bool
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE user_id = $1 AND task_id = $2)`, chatDb)
	err := p.db.Get(&exists, query, userId, taskId)
	if err != nil {
		panic(err)
	}
	return exists
}

func (p *ChatPostgres) CreateChat(userId, taskId int) (int, error) {
	exists := p.ChatExist(taskId, userId)
	if exists {
		return -1, errors.New("chat already exists")
	}
	query := fmt.Sprintf("INSERT INTO %s (task_id, user_id) VALUES ($1, $2) RETURNING id", chatDb)
	row := p.db.QueryRow(query, taskId, userId)
	if row.Err() != nil {
		return -1, row.Err()
	}
	var chatId int
	err := row.Scan(&chatId)
	if err != nil {
		return -1, err
	}
	return chatId, nil
}

func (p *ChatPostgres) AddMessage(taskId, userId int, message models.Message) error {
	
}
