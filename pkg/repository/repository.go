package repository

import (
	"WhyAi/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Chat
	Auth
}

type Chat interface {
	ChatExist(taskId, userId int) (bool, error)
	CreateChat(userId, taskId int) (int, error)
	AddMessage(taskId, userId int, message models.Message) error
	GetChat(taskId, userId int) ([]models.Message, error)
}

type Auth interface {
	SignUp(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
		Chat: NewChatPostgres(db),
	}

}
