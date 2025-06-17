package handler

import (
	"WhyAi/pkg/utils/responser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetFact(c *gin.Context) {
	req, err := h.service.Facts.GetFacts()
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "facts get failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{"facts": req})
}
