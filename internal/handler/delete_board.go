package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) deleteBoard(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id := c.Param("board_id")

	err = h.services.CheckPermissionToBoard(userId, id, adminLevel)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = h.services.Board.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Activity: "Delete board",
		Status: "ok",
	})
}
