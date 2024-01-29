package controllers

import (
	"configservice/service/pingpong"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingPongController handles HTTP requests related to ping pong.
type PingPongController struct {
	service *pingpong.Service
}

// NewPingPongController creates a new instance of PingPongController.
func NewPingPongController(service *pingpong.Service) *PingPongController {
	return &PingPongController{service: service}
}

func (c *PingPongController) SetupRouter(router *gin.Engine) {
	pingPongGroup := router.Group("/ping")
	{
		pingPongGroup.GET("/", c.ping)
	}
}

func (c *PingPongController) ping(ctx *gin.Context) {
	response, err := c.service.Ping(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}
