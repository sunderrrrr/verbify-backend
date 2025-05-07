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
func (p *ChatPostgres) ChatExist(taskId, userId int) (bool, error) {
	var exists bool
	query := fmt.Sprintf(`SELECT EXISTS(SELECT 1 FROM %s WHERE user_id = $1 AND task_id = $2)`, chatDb)
	err := p.db.Get(&exists, query, userId, taskId)
	if err != nil {
		return exists, err
	}
	return exists, nil
}

func (p *ChatPostgres) CreateChat(userId, taskId int) (int, error) {
	exists, err := p.ChatExist(taskId, userId)
	if err != nil {
		return -1, err
	}
	if exists {
		return -1, errors.New("chat already exists")
	}
	query := fmt.Sprintf("INSERT INTO %s (task_id, user_id) VALUES ($1, $2)", chatDb)
	row := p.db.QueryRow(query, taskId, userId)
	if row.Err() != nil {
		return -1, row.Err()
	}

	return 0, nil
}

func (p *ChatPostgres) AddMessage(taskId, userId int, message models.Message) error {
	query := fmt.Sprintf(`INSERT INTO %s (task_id, user_id, role, content) VALUES ($1, $2,$3,$4)`, msgDb)
	_, err := p.db.Exec(query, taskId, userId, message.Role, message.Content)
	if err != nil {
		return err
	}
	return nil
}

func (p *ChatPostgres) GetChat(taskId, userId int) ([]models.Message, error) {
	var messages []models.Message
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id=$1 AND task_id=$2 ORDER BY created_at ASC;`, msgDb)
	row := p.db.Select(&messages, query, taskId, userId)
	if row != nil {
		return messages, nil
	}

	return messages, nil

}
