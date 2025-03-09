package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ripu2/blahblah/internal/handlers"
	"github.com/ripu2/blahblah/internal/middleware"
)

func SetupRoutes(server *gin.Engine) {
	// server.GET("/api/v1/users")
	authServer := server.Group("/")
	authServer.Use(middleware.CheckForAuthentication)

	server.POST("/api/v1/signUp", handlers.CreateUserHandler)
	server.GET("/api/v1/login", handlers.LoginUserHandler)
	server.GET("/api/v1/channels", handlers.GetAllChannelsHandler)

	authServer.POST("/api/v1/createChanel", handlers.CreateChanelHandler)
	authServer.GET("/api/v1/myChannels", handlers.GetOwnChannelsHandler)
}
