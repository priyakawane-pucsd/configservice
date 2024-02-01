package controllers

import (
	"configservice/models"
	"configservice/models/dto"
	"configservice/service/pingpong"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
		pingPongGroup.POST("/", c.createPing)
		pingPongGroup.GET("/all", c.getAllPings)
		pingPongGroup.GET("/:id", c.getPingByID)
		pingPongGroup.PUT("/:id", c.updatePingByID)
		pingPongGroup.DELETE("/:id", c.deletePingByID)

	}
}

// @Summary Get ping pong health
// @Description Get ping pong health api
// @Tags pingpong
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.PingResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /ping [get]
func (c *PingPongController) ping(ctx *gin.Context) {
	response, err := c.service.Ping(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// @Summary Create a new Ping Pong entry
// @Description Create a new Ping Pong entry
// @Tags pingpong
// @Accept json
// @Produce json
// @Param text body dto.PingRequest true "Ping Pong text"
// @Success 201 {object} dto.PingResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /ping [post]
func (c *PingPongController) createPing(ctx *gin.Context) {
	var request struct {
		Text string `json:"text" binding:"required"`
		//ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{Error: err.Error()})
		return
	}
	response, err := c.service.CreatePing(request.Text)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// get all pings
// @Summary Get all Ping Pong entries
// @Description Get all Ping Pong entries
// @Tags pingpong
// @Accept json
// @Produce json
// @Success 200 {array} models.PingPong[]
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /ping/all [get]
func (c *PingPongController) getAllPings(ctx *gin.Context) {
	pings, err := c.service.GetAllPings(ctx)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(404, dto.ErrorResponse{Error: "No ping records found"})
			return
		}
		ctx.JSON(500, dto.ErrorResponse{Error: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, pings)
}

// @Summary Get ping by ID
// @Description Get ping by ID
// @Tags pingpong
// @Accept  json
// @Produce  json
// @Param id path string true "Ping ID"
// @Success 200 {object} models.PingPong
// @Failure 500 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /ping/{id} [get]
func (c *PingPongController) getPingByID(ctx *gin.Context) {
	id := ctx.Param("id")
	ping, err := c.service.GetPingByID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, dto.ErrorResponse{Error: "No ping record found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, ping)
}

// @Summary Update ping by ID
// @Description Update ping by ID
// @Tags pingpong
// @Accept json
// @Produce json
// @Param id path string true "Ping ID"
// @Param input body dto.PingRequest true "Ping data"
// @Success 200 {object} dto.PingResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /ping/{id} [put]
func (c *PingPongController) updatePingByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var update models.PingPong
	if err := ctx.BindJSON(&update); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	err := c.service.UpdatePingByID(ctx, id, &update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "No ping record found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Ping record updated successfully"})
}

// @Summary Delete ping by ID
// @Description Delete ping by ID
// @Tags pingpong
// @Accept json
// @Produce json
// @Param id path string true "Ping ID"
// @Success 200 {object} dto.PingResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Router /ping/{id} [delete]
func (c *PingPongController) deletePingByID(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.DeletePingByID(ctx, id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "No ping record found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Ping record deleted successfully"})
}
