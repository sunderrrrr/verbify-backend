package handler

import (
	"WhyAi/models"
	"WhyAi/pkg/utils/logger"
	"WhyAi/pkg/utils/responser"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	initPrompt = "Ты - самый лучший репетитор, который готовит к ЕГЭ по русскому языку 2025. Твоя задача будет состоять в том, чтобы максимально понятно и подробно объяснить конкретное задание. Пользуйся теорией которая тебе будет отправлена ниже. За работу тебе очень хорошо заплатят"
)

type EmptyChat struct{}

func (h *Handler) GetOrCreateChat(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "invalid user id")
		logger.Log.Errorf("Error while getting user id: %v", err)
		return
	}
	taskId, err := strconv.Atoi(c.Param("id"))

	req, err := h.service.Chat.Chat(taskId, userId)

	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "get chat failed")
		logger.Log.Errorf("Error while getting chat: %v", err)
		return
	}
	if req == nil {
		theory, _ := h.service.Theory.SendTheory(c.Param("id"), false)
		msg := models.Message{
			Role:    "system",
			Content: initPrompt + theory,
		}
		req := h.service.Chat.AddMessage(taskId, userId, msg) //Добавляем вопрос пользователя
		if req != nil {
			responser.NewErrorResponse(c, http.StatusInternalServerError, "send message failed")
			logger.Log.Errorf("Error while getting user id: %v", err)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"result": req})
}

func (h *Handler) SendMessage(c *gin.Context) {
	userId, err := getUserId(c)
	var msg models.Message
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "invalid user id")
		logger.Log.Errorf("Error while getting user id: %v", err)
		return
	}
	taskId, err := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&msg); err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Log.Errorf("Error while binding input: %v", err)
		return
	}
	req := h.service.Chat.AddMessage(taskId, userId, msg) //Добавляем вопрос пользователя
	if req != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "send message failed")
		logger.Log.Errorf("Error while add message: %v", err)
		return
	}
	history, err := h.service.Chat.Chat(taskId, userId)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "failed get history")
		logger.Log.Errorf("Error while get history: %v", err)
		return
	}
	ask, err := h.service.AskLLM(history, false)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "ask fail")
		logger.Log.Errorf("Error while LLM asking: %v", err)
		return
	}
	if ask == nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "llm error")
		logger.Log.Errorf("Ask equals nil: %v", err)
		return
	}
	if err = h.service.Chat.AddMessage(taskId, userId, *ask); err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "add msg failed "+err.Error())
		logger.Log.Errorf("Error while add message: %v", err)
		return
	}
	final, err := h.service.Chat.Chat(taskId, userId)
	c.JSON(http.StatusOK, gin.H{"result": final})
}

func (h *Handler) ClearContext(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "invalid user id")
		logger.Log.Errorf("Error while getting user id: %v", err)
		return
	}
	taskId, err := strconv.Atoi(c.Param("id"))
	del := h.service.Chat.ClearContext(taskId, userId)

	theory, _ := h.service.Theory.SendTheory(c.Param("id"), true)
	msg := models.Message{
		Role:    "system",
		Content: initPrompt + theory,
	}
	req := h.service.Chat.AddMessage(taskId, userId, msg) // Заново добавляем системное сообщение
	if req != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "send message failed "+req.Error())
		logger.Log.Errorf("Error while adding initial prompt: %v", req)
		return
	}

	if del != nil {
		c.JSON(http.StatusOK, gin.H{"result": del})
	}
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}
