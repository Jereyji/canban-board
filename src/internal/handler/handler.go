package handler

import (
    "log"
    "time"

    "github.com/Jereyji/canban-board/internal/service"
    "github.com/gin-contrib/cors"
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
    router.Use(func(c *gin.Context) {
        log.Printf("Received %s request from %s", c.Request.Method, c.Request.Header.Get("Origin"))
        c.Next()
    })

    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8000/"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type", "Authorization", "Accept", "X-Requested-With"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    router.OPTIONS("/*path", func(c *gin.Context) {
        c.AbortWithStatus(204)
    })


    auth := router.Group("/auth")
    {
        auth.POST("/sign-up", h.signUp)
        auth.POST("/sign-in", h.signIn)
        auth.POST("/send-code", h.sendCode)
        auth.POST("/confirm", h.confirm)
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
            boards.GET("/:board_id", h.getBoardById)
            boards.PUT("/:board_id", h.updateBoard)
            boards.DELETE("/:board_id", h.deleteBoard)

            cards := boards.Group("/:board_id/cards")
            {
                cards.POST("/", h.createCard)
                cards.GET("/", h.getAllCards)
            }
            users := boards.Group(":board_id/users")
            {
                users.GET("/", h.getAllUsers)
                users.POST("/", h.addPermission)
                users.DELETE("/", h.excludeUser)
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
