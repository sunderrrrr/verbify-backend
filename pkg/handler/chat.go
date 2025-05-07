package handler

import (
	"WhyAi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
	c.JSON(http.StatusOK, gin.H{"data": req})
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
	req := h.service.Chat.AddMessage(taskId, userId, msg)
	if req != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "send message failed")
	}
	c.JSON(http.StatusOK, gin.H{"result": req})
}
