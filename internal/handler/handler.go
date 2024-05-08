package handler

import (
	"github.com/Jereyji/canban-board/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		user := api.Group("/user") 
		{
			user.GET("/", h.getUserInfo)
			user.PUT("/", h.updateUser)
		}
		boards := api.Group("/boards")
		{
			boards.POST("/", h.createBoard)
			boards.GET("/", h.getAllBoards)
			boards.POST("/:board_id", h.addPermission)
			boards.GET("/:board_id", h.getBoardById)
			boards.PUT("/:board_id", h.updateBoard)
			boards.DELETE("/:board_id", h.deleteBoard)

			cards := boards.Group("/:board_id/cards")
			{
				cards.POST("/", h.createCard)
				cards.GET("/", h.getAllCards)
			}
		}
		cards := api.Group("/cards") 
		{
			cards.GET("/:card_id", h.getCardById)
			cards.PUT("/:card_id", h.updateCard)
			cards.DELETE("/:card_id", h.deleteCard)
		}
	}
	return router
}
