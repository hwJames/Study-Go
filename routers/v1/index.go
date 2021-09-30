package v1

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.RouterGroup)  {
	v1Route := route.Group("v1")
	setTestRoutes(v1Route)
}