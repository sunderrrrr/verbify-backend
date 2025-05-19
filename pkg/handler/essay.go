package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetEssayTasks(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	themes, err := h.service.Essay.GetEssayThemes()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": themes})
}

func (h *Handler) SendEssay(c *gin.Context) {

}
