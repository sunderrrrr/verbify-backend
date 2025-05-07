package service

import (
	"WhyAi/models"
	"WhyAi/pkg/repository"
)

type Service struct {
	Theory
	Auth
	LLM
	Chat
}
type Theory interface {
	SendTheory(n string) (string, error)
}
type Auth interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(user models.AuthUser) (string, error)
	ParseToken(token string) (models.User, error)
}

type LLM interface {
	SendMessage(message models.Message) (*models.Message, error)
}

type Chat interface {
	Chat(taskId, userId int) ([]models.Message, error)
	AddMessage(taskId, userId int, message models.Message) error
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth:   NewAuthService(repo),
		Theory: NewTheoryService(*repo),
		LLM:    NewLLMService("", ""),
		Chat:   NewChatService(repo),
	}
}
