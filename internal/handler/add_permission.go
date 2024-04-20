package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userPermission struct {
	Email string `json:"email" db:"email"`
	AccessLevel string `json:"access_level"`
}

func (h *Handler) addPermission(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	board_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id parametr")
		return
	}

	var input userPermission
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user_id, err := h.services.CheckUser(input.Email) 
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Board.AddPermission(board_id, user_id, input.AccessLevel)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Activity: "Give permission",
		Status: "ok",
	})
}