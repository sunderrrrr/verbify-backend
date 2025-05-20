package service

import (
	"WhyAi/models"
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
		return nil, err
	}
	var themes []models.EssayTheme
	err = json.Unmarshal(data, &themes)
	if err != nil {
		return nil, err
	}
	return themes, nil
}

func (s *EssayService) GenerateUserPrompt(request models.EssayRequest) (string, error) {
	prompt, err := ioutil.ReadFile("./static/theory/essay.txt")
	if err != nil {
		return "", err
	}
	essayContext := string(prompt)
	userPrompt := fmt.Sprintf(essayContext, request.Essay, request.Theme, request.Text)
	return userPrompt, nil
}
