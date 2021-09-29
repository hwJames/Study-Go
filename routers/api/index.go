package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/hwjames/Study-Go/routers/api/v1"
)

func InitAPIRoutes(api *gin.Engine)  {
	apiRoute := api.Group("api")
	
	v1.InitV1Routes(apiRoute)
}