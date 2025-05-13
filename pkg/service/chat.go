package service

import (
	"WhyAi/models"
	"WhyAi/pkg/repository"
)

type ChatService struct {
	repo repository.Repository
}

func NewChatService(repo repository.Repository) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) Chat(taskId, userId int) ([]models.Message, error) {
	return s.repo.GetChat(taskId, userId)
}
func (s *ChatService) AddMessage(taskId, userId int, message models.Message) error {
	return s.repo.AddMessage(taskId, userId, message)
}
func (s *ChatService) ChatExist(taskId, userId int) (bool, error) {
	return s.repo.ChatExist(taskId, userId)
}
func (s *ChatService) ClearContext(taskId, userId int) error {
	return s.repo.ClearContext(taskId, userId)
}
