package handler

import (
	"WhyAi/models"
	"WhyAi/pkg/utils/logger"
	"WhyAi/pkg/utils/responser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		responser.NewErrorResponse(c, http.StatusBadRequest, "invalid input")
		logger.Log.Errorf("Error while binding input: %v", err)
		return
	}
	signUp, err := h.service.Auth.CreateUser(input)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "sign up failed")
		logger.Log.Error("Error while creating user: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": signUp})

}

// TODO Сделать обработчик входа
func (h *Handler) signIn(c *gin.Context) {
	var input models.AuthUser
	if err := c.BindJSON(&input); err != nil {
		responser.NewErrorResponse(c, http.StatusBadRequest, "invalid input")
		logger.Log.Error("Error while binding input: %v", err)
		return
	}
	signIn, err := h.service.Auth.GenerateToken(input)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "sign in failed")
		logger.Log.Error("Error while generating token: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": signIn})
}
