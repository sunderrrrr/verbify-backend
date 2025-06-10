package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Msg string `json:"result:"`
}

func NewErrorResponse(c *gin.Context, statusCode int, msg string) {
	logrus.Println(msg)
	c.AbortWithStatusJSON(statusCode, errorResponse{Msg: msg})
}
