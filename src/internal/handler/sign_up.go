package handler

import (
	"net/http"
	"unicode"

	todo "github.com/Jereyji/canban-board"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.ValidateUserInput(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	if isWeakPassword(input.Password) {
		newErrorResponse(c, http.StatusBadRequest, "Password is too weak")
		return
	}	

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func isWeakPassword(password string) bool {
	hasOnlyDigits := true
	hasOnlyLetters := true
	for _, char := range password {
		if unicode.IsDigit(char) {
			hasOnlyLetters = false
		} else if unicode.IsLetter(char) {
			hasOnlyDigits = false
		} else {
			return false
		}
	}

	return hasOnlyDigits || hasOnlyLetters
}