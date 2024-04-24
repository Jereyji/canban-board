package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllCards(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	boardId := c.Param("board_id")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	cards, err := h.services.Card.GetAll(userId, boardId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, cards)
}
