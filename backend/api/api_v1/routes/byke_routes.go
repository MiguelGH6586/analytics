package routes

import (
	"api/api_v1/controllers"

	"github.com/gin-gonic/gin"
)

func BykeRouter(router *gin.Engine) {
	router.GET("/analytics/bikes", controllers.GetInfoBikes())
}
