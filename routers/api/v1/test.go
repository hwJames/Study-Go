package v1

import (
	"github.com/gin-gonic/gin"
)

func setTestRoutes(v1 *gin.RouterGroup)  {
	testRoute := v1.Group("test")

	testRoute.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}