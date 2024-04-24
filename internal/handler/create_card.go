package handler

import (
	"net/http"

	todo "github.com/Jereyji/canban-board"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCard(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	boardId := c.Param("board_id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	
	err = h.services.CheckPermissionToCard(userId, boardId)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input todo.Card
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	
	if input.UserId == "" {
		input.UserId = userId
	}
	
	id, err := h.services.Card.Create(userId, boardId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
