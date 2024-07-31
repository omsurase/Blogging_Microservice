package main

import (
	"os"

	"github.com/gin-gonic/gin"
	routes "github.com/omsurase/Blogging_Microservice/auth-service/routes"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-1"})
	})

	router.Run(":" + port)
}
