package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type errorResponse struct {
	Msg string `json:"response:"`
}

func NewErrorResponse(c *gin.Context, statusCode int, msg string) {
	log.Println("response.go: " + msg)
	c.AbortWithStatusJSON(statusCode, errorResponse{Msg: msg})
}
