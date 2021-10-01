package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hwjames/Study-Go/models"
	v1 "github.com/hwjames/Study-Go/routers/v1"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Seoul"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
	  panic("failed to connect database")
	}

    db.AutoMigrate(&models.User{})

	// user := models.User{ Nick: "Test", ID: "test@test.com", Pwd: "1234", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	// db.Create(&user)

	r.Run(":" + port)
}
