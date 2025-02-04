package controllers

import (
	"ScheduleApiGo/service"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServerController struct {
	service *service.ServerService
}

func NewServerController(service *service.ServerService) *ServerController {
	return &ServerController{service: service}
}

// CreateServer godoc
// @Summary Create a new server
// @Description Add a new Server to database
// @Tags server
// @Accept json
// @Produce json
// @Param server body model.Server true "Object Data"
// @Success 201 {object} model.Server
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /server [post]
func (sc *ServerController) CreateServer(c *gin.Context) {
	var request struct {
		Tag string `json:"tag" binding:"required"`
		IP  string `json:"ip" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}

	id, err := sc.service.CreateServer(context.Background(), request.Tag, request.IP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to insert data.", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetServers godoc
// @Summary list of servers
// @DescriptionReturn a list of servers
// @Tags server
// @Produce json
// @Success 200 {array} model.Server
// @Failure 500 {object} map[string]string
// @Router /server [get]
func (sc *ServerController) GetServers(c *gin.Context) {
	jobs, err := sc.service.GetServer(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch jobs", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jobs": jobs})
}
