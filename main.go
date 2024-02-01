package main

import (
	"configservice/config"
	"configservice/controllers"
	"configservice/repository/mongo"
	"configservice/service/pingpong"
	"context"

	_ "configservice/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	conf := config.LoadConfig()
	//Initializing repository
	pingPongRepository := mongo.NewRepository(context.Background(), &conf.Mongo)

	//Initiating service
	pingPongService := pingpong.NewService(pingPongRepository)

	// Initialize controller
	pingPongController := controllers.NewPingPongController(pingPongService)

	// Setup router with the controller
	pingPongController.SetupRouter(router)

	// Start the server
	router.Run(":8080")
}
