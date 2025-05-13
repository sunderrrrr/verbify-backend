package service

import (
	"WhyAi/pkg/repository"
	"fmt"
	"io/ioutil"
	"log"
)

type TheoryService struct {
	repo repository.Repository
}

func NewTheoryService(repo repository.Repository) *TheoryService {
	return &TheoryService{repo: repo}
}

func GetTheory(n string, forBot bool) (string, error) {
	if forBot {
		data, err := ioutil.ReadFile(fmt.Sprintf("./static/%sbot.txt", n))
		if err != nil {
			log.Printf("Ошибка чтения файла: %v", err)
			return "", err
		}
		return string(data), nil
	}
	data, err := ioutil.ReadFile(fmt.Sprintf("./static/%s.txt", n))
	if err != nil {
		log.Printf("Ошибка чтения файла: %v", err)
		return "", err
	}
	return string(data), nil
}

func (t *TheoryService) SendTheory(n string, forBot bool) (string, error) {
	theory, err := GetTheory(n, forBot)
	if err != nil {
		return "", err
	}
	return theory, nil
}
