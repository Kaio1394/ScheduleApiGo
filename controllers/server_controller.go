package controllers

import (
	"ScheduleApiGo/repository"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerController struct {
	repo *repository.ServerRepository
}

func NewServerController(repo *repository.ServerRepository) *ServerController {
	return &ServerController{repo: repo}
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

	id, err := sc.repo.Create(context.Background(), request.Tag, request.IP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to insert data.", "message": "" + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
