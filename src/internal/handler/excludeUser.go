package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserExcludeRequest struct {
    Username string `json:"username"`
}

func (h *Handler) excludeUser(c *gin.Context) {
    userId, err := getUserId(c)
    if err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    boardId := c.Param("board_id")

    err = h.services.CheckPermissionToBoard(userId, boardId, adminLevel)
    if err != nil {
        newErrorResponse(c, http.StatusUnauthorized, err.Error())
        return
    }

    var userExcludeReq UserExcludeRequest
    if err := c.BindJSON(&userExcludeReq); err != nil {
        newErrorResponse(c, http.StatusBadRequest, err.Error())
        return
    }

    if err := h.services.Authorization.ExcludeUser(userExcludeReq.Username, boardId); err != nil {
        newErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    c.JSON(http.StatusOK, map[string]interface{}{
        "user": "excluded",
    })
}
