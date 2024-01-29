package main

import (
	"configservice/config"
	"configservice/controllers"
	"configservice/repository/mongo"
	"configservice/service/pingpong"
	"context"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin router
	router := gin.Default()
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
