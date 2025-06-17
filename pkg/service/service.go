package service

import (
	"WhyAi/models"
	"WhyAi/pkg/repository"
	"os"
)

type Service struct {
	Theory
	Auth
	LLM
	Chat
	Facts
	Essay
	User
}
type Theory interface {
	SendTheory(n string, forBot bool) (string, error)
}
type Auth interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(user models.AuthUser) (string, error)
	ParseToken(token string) (models.User, error)
}

type LLM interface {
	AskLLM(messages []models.Message, isEssay bool) (*models.Message, error)
}

type Chat interface {
	ChatExist(taskId, userId int) (bool, error)
	Chat(taskId, userId int) ([]models.Message, error)
	AddMessage(taskId int, userId int, message models.Message) error
	ClearContext(taskId, userId int) error
}

type Facts interface {
	GetFacts() ([]models.Fact, error)
}
type Essay interface {
	GetEssayThemes() ([]models.EssayTheme, error)
	GenerateUserPrompt(request models.EssayRequest) (string, error)
}
type User interface {
	ResetPassword(resetModel models.UserReset) error
	ResetPasswordRequest(email models.ResetRequest) error
	GeneratePasswordResetToken(email, signingKey string) (string, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth:   NewAuthService(repo),
		Theory: NewTheoryService(*repo),
		LLM:    NewLLMService(os.Getenv("DEEPSEEK_URL"), os.Getenv("DEEPSEEK_TOKEN")),
		Chat:   NewChatService(*repo),
		Facts:  NewFactService(),
		Essay:  NewEssayService(),
		User:   NewUserService(repo.User),
	}
}
