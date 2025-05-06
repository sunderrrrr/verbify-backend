package handler

import (
	"WhyAi/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}
	signUp, err := h.service.Auth.CreateUser(input)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "sign up failed")
		logrus.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": signUp})
	fmt.Println(signUp)
}

// TODO Сделать обработчик входа
func (h *Handler) signIn(c *gin.Context) {
	var input models.AuthUser
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input")
		return
	}
	signIn, err := h.service.Auth.GenerateToken(input)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "sign in failed")
		logrus.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": signIn})
	fmt.Println(signIn)
}
