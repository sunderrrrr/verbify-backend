package handler

import (
	"WhyAi/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) GetEssayTasks(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	themes, err := h.service.Essay.GetEssayThemes()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": themes})
}

func (h *Handler) SendEssay(c *gin.Context) {
	_, err := getUserId(c)
	var input models.EssayRequest
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "no user id")
	}
	if err = c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid request body")
	}
	request, err := h.service.Essay.GenerateUserPrompt(input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	llmRequest := []models.Message{
		models.Message{
			Role:    "user",
			Content: request,
		},
	}
	llmAsk, err := h.service.LLM.AskLLM(llmRequest, true)
	resp := strings.Replace(strings.Replace(strings.Replace(llmAsk.Content, "`", "", -1), "json", "", -1), "Ð", "", -1)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	var eResponse models.EssayResponse
	err = json.Unmarshal([]byte(resp), &eResponse)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": eResponse})
}
