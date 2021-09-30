package v1

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/hwjames/Study-Go/controllers/v1"
)

func setTestRoutes(v1Route *gin.RouterGroup)  {
	testRoute := v1Route.Group("test")

	testRoute.GET("/ping", v1.Ping)
}