package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetFact(c *gin.Context) {
	req, err := h.service.Facts.GetFacts()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "facts get failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{"facts": req})
}
