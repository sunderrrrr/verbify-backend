package service

import (
	"WhyAi/models"
	"WhyAi/pkg/utils/logger"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type LLMService struct {
	ApiUrl string
	Token  string
}

func NewLLMService(apiUrl string, token string) *LLMService {
	return &LLMService{ApiUrl: apiUrl, Token: token}
}

func (s *LLMService) AskLLM(messages []models.Message, isEssay bool) (*models.Message, error) {
	request := models.LLMRequest{
		Model:       "deepseek-chat",
		Temperature: 0,
		Messages:    messages,
	}
	var res models.LLMResponse
	body, err := json.Marshal(request)
	if err != nil {
		logger.Log.Errorf("Error while marshaling LLM request: %v", err)
		return nil, errors.New("request marshal fail")
	}

	llmRequest, err := http.NewRequest("POST", s.ApiUrl, bytes.NewReader(body))
	llmRequest.Header.Set("Content-Type", "application/json")
	llmRequest.Header.Set("Authorization", "Bearer "+s.Token)
	resp, err := http.DefaultClient.Do(llmRequest)
	if err != nil {
		logger.Log.Errorf("Error while LLM request: %v", err)
		return nil, errors.New("fail request")
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		logger.Log.Errorf("Error while decoding LLM response: %v", err)
		return nil, errors.New("fail to decode response " + err.Error())
	}
	if len(res.Choices) == 0 {
		logger.Log.Error("LLM response has no choices")
		return nil, errors.New("no choices in response")
	}
	ans := &res.Choices[0].Message
	if isEssay {
		ans.Content, _ = getJson(ans.Content)
	}
	return ans, nil
}

func cleanText(text string) string {
	result := ""
	for _, char := range text {
		if char >= ' ' && char <= '~' {
			result += string(char)
		} else if char == '\n' || char == ' ' {
			result += string(char)
		}
	}
	return result
}
func getJson(text string) (string, error) {
	start := strings.Index(text, "{")
	end := strings.Index(text, "}")

	if start != -1 && end != -1 && start < end {
		return text[start : end+1], nil
	}
	return cleanText(text), nil
}
