package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/ping", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})
		}
	}

	r.Run(":4000")
}
