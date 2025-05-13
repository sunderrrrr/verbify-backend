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

func (s *LLMService) AskLLM(messages []models.Message) (*models.Message, error) {
	request := models.LLMRequest{
		Model:    "deepseek-chat",
		Messages: messages,
	}
	//fmt.Println("request", request)
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

	var res models.LLMResponse
	err = json.NewDecoder(resp.Body).Decode(&res)

	if err != nil {
		return nil, errors.New("fail to decode response " + err.Error())
	}
	if len(res.Choices) == 0 {
		return nil, errors.New("no choices in response")
	}
	return &res.Choices[0].Message, nil
}
