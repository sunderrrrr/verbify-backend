package handler

import (
	"WhyAi/models"
	"WhyAi/pkg/utils/responser"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetUserInfo(c *gin.Context) {
}

func (h *Handler) UpdateUserInfo(c *gin.Context) {
}

func (h *Handler) DeleteUserInfo(c *gin.Context) {
}

func (h *Handler) SendResetRequest(c *gin.Context) {
	var input models.ResetRequest
	err := c.BindJSON(&input)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println(input)
	err = h.service.User.ResetPasswordRequest(input)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "request sent"})

}

func (h *Handler) UpdatePassword(c *gin.Context) {
	var input models.UserReset
	err := c.BindJSON(&input)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.service.User.ResetPassword(input)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "password reset confirmed"})
	}

}
