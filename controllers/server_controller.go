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
