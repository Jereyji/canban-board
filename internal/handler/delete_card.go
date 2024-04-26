package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) deleteCard(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	cardId := c.Param("card_id")
	boardId, err := h.services.GetBoardIdByCard(cardId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.CheckPermissionToCard(userId, boardId)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = h.services.Card.Delete(userId, cardId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Activity: "Delete card",
		Status: "ok",
	})
}
