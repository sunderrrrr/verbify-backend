package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) SendTheory(c *gin.Context) {
	n := c.Param("id")
	if n == "" {
		NewErrorResponse(c, http.StatusBadRequest, "task number is required")
		return
	}

	theory, err := h.service.Theory.SendTheory(n)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"theory": theory})
}
