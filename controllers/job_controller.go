package controllers

import (
	"ScheduleApiGo/logger"
	"ScheduleApiGo/model"
	"ScheduleApiGo/service"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type JobController struct {
	service *service.JobService
}

func NewJobController(service *service.JobService) *JobController {
	return &JobController{service: service}
}

// CreateJob godoc
// @Summary Create a new job
// @Description Add a new job to the database
// @Tags job
// @Accept json
// @Produce json
// @Param job body model.Job true "Job data"
// @Success 201 {object} map[string]int
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /job [post]
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

// GetJobs godoc
// @Summary List of jobs
// @Description Return a list of jobs
// @Tags job
// @Produce json
// @Success 200 {array} model.Job
// @Failure 500 {object} map[string]string
// @Router /job/list [get]
func (jc *JobController) GetJobs(c *gin.Context) {
	jobs, err := jc.service.GetJobs(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch jobs", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jobs": jobs})
}

// GetJobById godoc
// @Summary Get job by ID
// @Description Retrieve a job by its ID from the database
// @Tags job
// @Accept json
// @Produce json
// @Param id header string true "Job ID"
// @Success 200 {object} model.Job
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /job [get]
func (jc *JobController) GetJobById(c *gin.Context) {
	id := c.GetHeader("id")
	id_int, err := strconv.Atoi(id)
	if err != nil {
		logger.Log.Errorf("Fail to parse job id %s", id)
	}

	job, err := jc.service.GetJobById(context.Background(), id_int)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch job", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, job)
}
