package service

import (
	"WhyAi/models"
	"encoding/json"
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
