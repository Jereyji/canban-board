package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getCardById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	cardId := c.Param("card_id")

	// Check permission to board?

	card, err := h.services.Card.GetById(userId, cardId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, card)
}
