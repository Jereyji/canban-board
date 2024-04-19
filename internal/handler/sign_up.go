package handler

import (
	"net/http"

	todo "github.com/Jereyji/canban-board"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	h.services.CreateUser(input)
}
