package service

import (
	"WhyAi/models"
	"WhyAi/pkg/repository"
)

type Service struct {
	Theory
	Auth
}
type Theory interface {
	SendTheory(n string) (string, error)
}
type Auth interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(user models.AuthUser) (string, error)
	ParseToken(token string) (models.User, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{Auth: NewAuthService(repo),
		Theory: NewTheoryService()}
}
