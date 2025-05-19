package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	usernameCtx         = "username"
)

// Промежуточкая аунтентификация пользователя
func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "middleware.go: no authorization header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		NewErrorResponse(c, http.StatusUnauthorized, "middleware.go: invalid authorization header")
		return
	}

	// parse token
	//fmt.Println("middleware: headerParts", headerParts[1])
	userId, err := h.service.Auth.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(userCtx, userId.Id)
	c.Set(usernameCtx, userId.Name)
	//fmt.Println("middleware.go: userId:", userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusUnauthorized, "middleware.go: no user id")
		return 0, errors.New("middleware.go: no user id")
	}
	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(c, http.StatusUnauthorized, "middleware.go: invalid user id")
		return 0, errors.New("middleware.go: invalid user id")
	}
	return idInt, nil
}

func getUsername(c *gin.Context) (string, error) {
	username, ok := c.Get(usernameCtx)
	if !ok {
		NewErrorResponse(c, http.StatusUnauthorized, "middleware.go: no user id")
		return "", errors.New("middleware.go: no user id")
	}
	return username.(string), nil
}
