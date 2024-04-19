package handler

import (
	"net/http"

	todo "github.com/Jereyji/canban-board"
	"github.com/gin-gonic/gin"
)

type getAllBoardResponse struct {
	Data []todo.Board `json:"data"`
}

func (h *Handler) getAllBoards(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return 
	}

	boards, err := h.services.Board.GetAll(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllBoardResponse{
		Data: boards,
	})
}
