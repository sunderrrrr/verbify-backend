package handler

import (
	"WhyAi/pkg/utils/responser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SendTheory(c *gin.Context) {
	n := c.Param("id")
	if n == "" {
		responser.NewErrorResponse(c, http.StatusBadRequest, "task number is required")
		return
	}

	theory, err := h.service.Theory.SendTheory(n, false)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"theory": theory})
}
