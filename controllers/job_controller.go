package controllers

import (
	"ScheduleApiGo/model"
	"ScheduleApiGo/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobController struct {
	service *service.JobService
}

func NewJobController(service *service.JobService) *JobController {
	return &JobController{service: service}
}

func (jc *JobController) CreateJob(c *gin.Context) {
	var request struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
		CmdExecute  bool   `json:"cmdExecute"`
		Script      string `json:"script"`
		Date        string `json:"date"`
		Hold        bool   `json:"hold"`
		Priority    int    `json:"priority"`
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
		Script:      request.Script,
		Date:        request.Date,
		Hold:        request.Hold,
		ServerId:    request.ServerId,
		Priority:    request.Priority,
	}

	id, err := jc.service.CreateJob(context.Background(), &job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fail to insert data.", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func (jc *JobController) GetJobs(c *gin.Context) {
	jobs, err := jc.service.GetJobs(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch jobs", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jobs": jobs})
}
