package controllers

import (
	"ScheduleApiGo/model"
	"ScheduleApiGo/repository"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobController struct {
	repo *repository.JobRepository
}

func NewJobController(repo *repository.JobRepository) *JobController {
	return &JobController{repo: repo}
}

func (jc *JobController) CreateJob(c *gin.Context) {
	var request struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		CmdExecute  bool   `json:"cmdExecute"`
		Script      string `json:"script"` // Alterado para string
		Date        string `json:"date"`
		Hold        bool   `json:"hold"`
		ServerId    int    `json:"serverId" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}

	job := model.Job{
		Name:        request.Name,
		Description: request.Description,
		CmdExecute:  request.CmdExecute,
		Script:      request.Script, // Agora pode ser uma string
		Date:        request.Date,
		Hold:        request.Hold,
		ServerId:    request.ServerId,
	}

	id, err := jc.repo.Create(context.Background(), &job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to insert data.", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func ListJob(c *gin.Context) {

}
