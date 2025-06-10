package repository

import (
	"WhyAi/models"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Chat
	Auth
	User
}

type Chat interface {
	ChatExist(taskId, userId int) (bool, error)
	CreateChat(userId, taskId int) (int, error)
	AddMessage(taskId, userId int, message models.Message) error
	ClearContext(taskId, userId int) error
	GetChat(taskId, userId int) ([]models.Message, error)
}

type Auth interface {
	SignUp(user models.User) (int, error)
	GetUser(username, password string, login bool) (models.User, error)
}

type User interface {
	ResetPassword(username string, newPassword string) error
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthPostgres(db),
		Chat: NewChatPostgres(db),
		User: NewUserRepository(db),
	}

}
