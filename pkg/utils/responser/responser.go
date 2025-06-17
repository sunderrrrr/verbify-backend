package responser

import (
	"WhyAi/pkg/utils/logger"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Msg string `json:"result:"`
}

func NewErrorResponse(c *gin.Context, statusCode int, msg string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{Msg: msg})
	logger.Log.Error("Error while binding input: %v", msg)
}
