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

	authServer.POST("/api/v1/createChanel", handlers.CreateChanelHandler)
}
