package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ripu2/blahblah/internal/config/db"
	RedisClient "github.com/ripu2/blahblah/internal/config/redis"
	"github.com/ripu2/blahblah/internal/routes"
)

func main() {
	db.InitDB()
	RedisClient.InitRedisClient()
	_ = godotenv.Load()
	server := gin.Default()
	routes.SetupRoutes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Server is up an running on http://localhost:8080")
	err := server.Run(":" + port) //localhost:8080
	fmt.Printf("Server Failed to start: %v\n", err.Error())

}
