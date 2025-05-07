package service

import (
	"WhyAi/models"
	"WhyAi/pkg/repository"
	"fmt"
	"github.com/sirupsen/logrus"
)

type ChatService struct {
	repo repository.Chat
}

func NewChatService(repo repository.Chat) *ChatService {
	return &ChatService{repo: repo}
}

func (s *ChatService) Chat(taskId, userId int) ([]models.Message, error) {
	exist, err := s.repo.ChatExist(taskId, userId)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	if exist { //Если чат существует - выводим
		fmt.Println("чат есть...")
		return s.repo.GetChat(taskId, userId)

	} else {
		fmt.Println("чата нету")
		_, err := s.repo.CreateChat(taskId, userId)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
}
func (s *ChatService) AddMessage(taskId, userId int, message models.Message) error {
	return s.repo.AddMessage(taskId, userId, message)
}
