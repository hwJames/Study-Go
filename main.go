package main

import (
	"os"

	"github.com/gin-gonic/gin"
	v1 "github.com/hwjames/Study-Go/routers/v1"
	"github.com/joho/godotenv"
)

var r *gin.Engine

func init() {
	godotenv.Load()

	r = gin.Default()
	api := r.Group("/api")
	v1.InitRoutes(api)
}

func main() {
	port := os.Getenv("PORT") 
	if "" == port {
		port = "8080"
	}

	r.Run(":" + port)
}
