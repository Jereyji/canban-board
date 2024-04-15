package handler

import "github.com/gin-gonic/gin"

type Handler struct {

}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth"); {
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}

	api := router.Group("/api"); {
		boards := api.Group("/boards"); {
			boards.POST("/")
			boards.GET("/")
			boards.GET("/:id")
			boards.PUT("/:id")
			boards.DELETE("/:id")

			cards := boards.Group(":id/cards"); {
				cards.POST("/")
				cards.GET("/")
			}
		}
		cards := boards.Group("/cards"); {
			cards.GET("/:id")
			cards.PUT("/:id")
			cards.DELETE("/:id")
		}
	}
	return router
}