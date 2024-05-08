package handler

import (
	"net/http"

	todo "github.com/Jereyji/canban-board"
	"github.com/gin-gonic/gin"
)

func (h *Handler) updateUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input todo.UpdateUserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Authorization.UpdateUser(userId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse {
		Activity: "Update User",
		Status: "ok",
	})
}