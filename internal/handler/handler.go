package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		boards := api.Group("/boards")
		{
			boards.POST("/", h.createBoard)
			boards.GET("/", h.getAllBoards)
			boards.GET("/:id", h.getBoardById)
			boards.PUT("/:id", h.updateBoard)
			boards.DELETE("/:id", h.deleteBoard)

			cards := boards.Group(":id/cards")
			{
				cards.POST("/", h.createCard)
				cards.GET("/", h.getAllCards)
			}
		}
		cards := boards.Group("/cards")
		{
			cards.GET("/:id", h.getCardById)
			cards.PUT("/:id", h.updateCard)
			cards.DELETE("/:id", h.deleteCard)
		}
	}
	return router
}
