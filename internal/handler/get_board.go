package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getBoardById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id := c.Param("board_id")

	board, err := h.services.Board.GetById(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, board)
}
