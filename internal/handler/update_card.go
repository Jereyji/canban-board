package handler

import (
	"net/http"

	todo "github.com/Jereyji/canban-board"
	"github.com/gin-gonic/gin"
)

func (h *Handler) updateCard(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	board_id := c.Param("board_id")
	card_id := c.Param("card_id")

	var input todo.UpdateCardInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.CheckPermissionToCard(userId, board_id)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if err := h.services.Card.Update(userId, card_id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, statusResponse {
		Activity: "Update board",
		Status: "ok",
	})
}
