package handler

import (
	"WhyAi/pkg/utils/responser"
	"github.com/gin-gonic/gin"
)

func (h *Handler) adminIdentity(c *gin.Context) {
	r, err := getRoleId(c)
	if err != nil {
		responser.NewErrorResponse(c, 500, "get role id failed")
	}
	if r != 0 {
		responser.NewErrorResponse(c, 403, "you are not admin")
		return
	}
	c.Next()
}

func (h *Handler) GetAllUsersList(c *gin.Context) {}

func (h *Handler) SetSubscription(c *gin.Context) {}

func (h *Handler) RemoveSubscription(c *gin.Context) {}

func (h *Handler) DeleteUser(c *gin.Context) {}
