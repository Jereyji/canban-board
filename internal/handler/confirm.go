package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type confirmationInput struct {
	Email string `json:"email" binding:"required"`
	Code  string `json:"code" binding:"required"`
}

func (h *Handler) confirm(c *gin.Context) {
	var input confirmationInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	expectedCode, err := h.services.Authorization.GetCode(input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if input.Code != expectedCode {
		newErrorResponse(c, http.StatusUnauthorized, "Invalid confirmation code")
		return
	}

	token, err := h.services.GenerateToken(input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed autorization")
		return
	}

	cookie := http.Cookie{
		Name:     "JTV",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(c.Writer, &cookie)

	c.JSON(http.StatusOK, "Authentication successful")
}
