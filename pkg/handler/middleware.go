package handler

import (
	"WhyAi/pkg/utils/responser"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	roleCtx             = "roleId"
)

// Промежуточкая аунтентификация пользователя
func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "middleware.go: no authorization header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "middleware.go: invalid authorization header")
		return
	}
	userId, err := h.service.Auth.ParseToken(headerParts[1])
	if err != nil {
		responser.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	roleId, err := h.service.User.GetRoleById(userId.Id)
	if err != nil {
		responser.NewErrorResponse(c, http.StatusInternalServerError, "middleware.go: get role by id failed")
		return
	}
	c.Set(roleCtx, roleId)
	c.Set(userCtx, userId.Id)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "middleware.go: no user id")
		return 0, errors.New("middleware.go: no user id")
	}
	idInt, ok := id.(int)
	if !ok {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "middleware.go: invalid user id")
		return 0, errors.New("middleware.go: invalid user id")
	}
	return idInt, nil
}

func getRoleId(c *gin.Context) (int, error) {
	id, ok := c.Get(roleCtx)
	if !ok {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "middleware.go: no role id")
		return 0, errors.New("middleware.go: no role id")
	}
	idInt, ok := id.(int)
	if !ok {
		responser.NewErrorResponse(c, http.StatusUnauthorized, "middleware.go: invalid role id")
		return 0, errors.New("middleware.go: invalid role id")
	}
	return idInt, nil
}
