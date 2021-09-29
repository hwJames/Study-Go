package v1

import (
	"github.com/gin-gonic/gin"
)

func InitV1Routes(api *gin.RouterGroup)  {
	v1Route := api.Group("v1")
	setTestRoutes(v1Route)
}