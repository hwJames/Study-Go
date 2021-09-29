package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hwjames/Study-Go/routers/api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT") 
	if "" == port {
		port = "8080"
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	api.InitAPIRoutes(r)

	r.Run(":" + port)
}
