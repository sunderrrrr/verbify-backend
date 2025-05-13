package service

import (
	"WhyAi/models"
	"github.com/goccy/go-json"
	"io/ioutil"
)

type FactService struct {
}

func NewFactService() *FactService {
	return &FactService{}
}

func (s *FactService) GetFacts() ([]models.Fact, error) {
	data, err := ioutil.ReadFile("./static/facts.json")
	if err != nil {
		return nil, err
	}
	var facts []models.Fact
	err = json.Unmarshal(data, &facts)
	if err != nil {
		return nil, err
	}
	return facts, nil
}
