package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	adminLevel          = "admin"
	editorLevel         = "editor"
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	var token string
	cookie, err := c.Request.Cookie("JTV")

	if err == nil {
		token = cookie.Value
	} else {
		header := c.GetHeader(authorizationHeader)
		if header == "" {
			newErrorResponse(c, http.StatusUnauthorized, "missing session cookie and empty auth header")
			return
		}
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			newErrorResponse(c, http.StatusUnauthorized, header + " invalid auth header")
			return
		}
		token = headerParts[1]
	}

	userId, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (string, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "user id not found")
		return "", errors.New("user id not found")
	}

	idInt, ok := id.(string)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return "", errors.New("user id not found")
	}

	return idInt, nil
}
