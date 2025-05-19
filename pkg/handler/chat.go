package handler

import (
	"WhyAi/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		NewErrorResponse(c, http.StatusInternalServerError, "invalid user id")
		return
	}
	taskId, err := strconv.Atoi(c.Param("id"))

	req, err := h.service.Chat.Chat(taskId, userId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	//fmt.Println(req)
	if req == nil {
		theory, _ := h.service.Theory.SendTheory(c.Param("id"), false)
		//fmt.Println(theory)
		msg := models.Message{
			Role:    "system",
			Content: initPrompt + theory,
		}
		//fmt.Println(msg)
		req := h.service.Chat.AddMessage(taskId, userId, msg) //Добавляем вопрос пользователя
		if req != nil {
			NewErrorResponse(c, http.StatusInternalServerError, "send message failed "+req.Error())
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"result": req})
}

func (h *Handler) SendMessage(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid user id")
		return
	}
	taskId, err := strconv.Atoi(c.Param("id"))
	var msg models.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	req := h.service.Chat.AddMessage(taskId, userId, msg) //Добавляем вопрос пользователя
	if req != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "send message failed "+req.Error())
		return
	}
	logrus.Infof("send message to user %d", userId)

	history, err := h.service.Chat.Chat(taskId, userId)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "failed get history: "+err.Error())
		return

	}
	//fmt.Println(history)
	ask, err := h.service.AskLLM(history)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "ask fail: "+err.Error())
		return
	}
	logrus.Infof("ask comlite")
	if ask == nil {
		NewErrorResponse(c, http.StatusInternalServerError, "ask llm failed")
		return
	}
	if err = h.service.Chat.AddMessage(taskId, userId, *ask); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "add msg failed "+err.Error())
		return
	}
	logrus.Infof("add message to user %d", userId)
	final, err := h.service.Chat.Chat(taskId, userId)
	logrus.Infof("final")
	c.JSON(http.StatusOK, gin.H{"result": final})
	//fmt.Println(final)
}

func (h *Handler) ClearContext(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid user id")
		return
	}
	taskId, err := strconv.Atoi(c.Param("id"))
	del := h.service.Chat.ClearContext(taskId, userId)

	theory, _ := h.service.Theory.SendTheory(c.Param("id"), true)
	msg := models.Message{
		Role:    "system",
		Content: initPrompt + theory,
	}
	req := h.service.Chat.AddMessage(taskId, userId, msg) //Добавляем вопрос пользователя
	if req != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "send message failed "+req.Error())
		return
	}

	if del != nil {
		c.JSON(http.StatusOK, gin.H{"result": del})
	}
	c.JSON(http.StatusOK, gin.H{"result": "ok"})
}
