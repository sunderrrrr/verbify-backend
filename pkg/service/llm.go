package service

import (
	"WhyAi/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type LLMService struct {
	ApiUrl string
	Token  string
}

func NewLLMService(apiUrl string, token string) *LLMService {
	return &LLMService{ApiUrl: apiUrl, Token: token}
}

func (s *LLMService) SendMessage(message models.Message) (*models.Message, error) {
	request := models.LLMRequest{
		Messages: []models.Message{message},
		MaxToken: 1000,
		Stream:   false,
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, errors.New("request marshal fail")
	}
	req, err := http.NewRequest("POST", s.ApiUrl, bytes.NewReader(body))
	if err != nil {
		return nil, errors.New("fail request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.Token)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.New("fail request")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.New("api return status code not 200")

	}
	var res models.LLMResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, errors.New("fail to decode response")
	}
	if len(res.Choices) == 0 {
		return nil, errors.New("no choices in response")
	}
	return &res.Choices[0].Message, nil
}
