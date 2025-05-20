package service

import (
	"WhyAi/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"unicode"
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
		Model:       "deepseek-chat",
		Temperature: 0,
		Messages:    messages,
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
	//return &res.Choices[0].Message, nil
	ans := &res.Choices[0].Message
	fmt.Println(ans)
	ans.Content, _ = extractJSONFromMarkdown(ans.Content)
	return ans, nil
}

func cleanResponseWithMarkdown(raw string) string {
	// Удаляем бинарные/непечатные символы, сохраняя Markdown
	cleaned := strings.Map(func(r rune) rune {
		// Сохраняем символы Markdown: # * ` [ ] ! и др.
		if (r >= ' ' && r <= '~') || unicode.IsSpace(r) ||
			r == '#' || r == '*' || r == '`' || r == '[' ||
			r == ']' || r == '!' || r == '_' || r == '-' {
			return r
		}
		return -1
	}, raw)

	// Дополнительная обработка частых проблем
	cleaned = strings.ReplaceAll(cleaned, "“", `"`) // Замена "умных" кавычек
	cleaned = strings.ReplaceAll(cleaned, "”", `"`)
	return cleaned
}
func extractJSONFromMarkdown(raw string) (string, error) {
	// Регулярное выражение для поиска JSON-блока
	re := regexp.MustCompile(`(?s)\x60{0,3}json\s*(\{.*?\})\x60{0,3}`)
	matches := re.FindStringSubmatch(raw)

	if len(matches) > 1 {
		return matches[1], nil
	}

	// Если JSON не найден, возвращаем очищенный текст
	return cleanResponseWithMarkdown(raw), nil
}
