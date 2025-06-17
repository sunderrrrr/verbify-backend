package service

import (
	"WhyAi/models"
	"WhyAi/pkg/utils/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type EssayService struct {
}

func NewEssayService() *EssayService {
	return &EssayService{}
}

func (s *EssayService) GetEssayThemes() ([]models.EssayTheme, error) {
	data, err := ioutil.ReadFile("./static/essays.json")
	if err != nil {
		logger.Log.Error("Error while reading essays.json: %v", err)
		return nil, err
	}
	var themes []models.EssayTheme
	err = json.Unmarshal(data, &themes)
	if err != nil {
		logger.Log.Error("Error while parsing essays.json: %v", err)
		return nil, err
	}
	return themes, nil
}

func (s *EssayService) GenerateUserPrompt(request models.EssayRequest) (string, error) {
	prompt, err := ioutil.ReadFile("./static/theory/essay.txt")
	if err != nil {
		logger.Log.Error("Error while reading essay prompt file: %v", err)
		return "", err
	}
	essayContext := string(prompt)
	userPrompt := fmt.Sprintf(essayContext, request.Essay, request.Theme, request.Text)
	return userPrompt, nil
}
