package main

import (
	"api/api_v1/configs"
	"api/api_v1/cors"
	"api/api_v1/routes"

	//_ "api/api_v1/docs"

	"github.com/gin-gonic/gin"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Bikes API
// @version 1.0
// @description Bikes API in Go using Gin framework

// @host	localhost:6060
// @BasePath /
func main() {
	router := gin.Default()

	//run database
	configs.ConnectDB()

	//routes
	//routes.UserRouter(router)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// add swagger
	//router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.BykeRouter(router)

	router.Run(":6060")
}
