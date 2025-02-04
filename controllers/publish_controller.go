package controllers

import (
	"ScheduleApiGo/config"
	"ScheduleApiGo/enums"
	"ScheduleApiGo/helper"
	"ScheduleApiGo/logger"
	"ScheduleApiGo/model"
	"ScheduleApiGo/service"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	configs config.Config
)

type PublishController struct {
	sp *service.PublishService
	js *service.JobService
}

func NewPublishController(service *service.PublishService, js *service.JobService) *PublishController {
	return &PublishController{sp: service, js: js}
}

func (s *PublishController) Publish(c *gin.Context) {

	jobId, err := strconv.Atoi(c.GetHeader("JobId"))
	if err != nil {
		logger.Log.Error(err)
		return
	}
	job, err := s.js.GetJobById(context.Background(), jobId)
	if err != nil {
		logger.Log.Errorf("get job by id err: %v", err)
		return
	}

	server := c.GetHeader("server")
	portStr := c.GetHeader("port")
	user := c.GetHeader("user")
	password := c.GetHeader("password")
	queue := c.DefaultQuery("queue", "")

	if server == "" || portStr == "" || user == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing required parameters (host, user, password, port).",
		})
		logger.Log.Error("Missing required parameters (host, user, password, port).")
		return
	}

	port, err := strconv.ParseUint(portStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid port parameter.",
		})
		logger.Log.Error("Invalid port parameter.")
		return
	}

	var jobHistory = model.HistoryExecution{
		JobId:    job.Id,
		Status:   enums.Running.String(),
		ServerId: job.ServerId,
		CreateAt: time.Now(),
	}

	err = s.js.SaveToTableHistoryExecution(context.Background(), jobHistory)
	if err != nil {
		logger.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var rabbit_config helper.Rabbit

	rabbit_config = helper.Rabbit{}

	rabbit_config = helper.Rabbit{
		Host:     server,
		Port:     uint32(port),
		User:     user,
		Password: password,
	}

	if rabbit_config.HasEmptyParams() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing required parameters (host, user, password, port).",
		})
		logger.Log.Error("Missing required parameters (host, user, password, port).")
		return
	}

	con, err := rabbit_config.Connection()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":            "Connection error.",
			"ParamsConnection": rabbit_config,
		})
		logger.Log.Error("Connection error.")
		return
	}

	rabbit_config.SendMessage(job, queue, con)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Publish Job.",
		"job":     job,
	})
	logger.Log.Info("Sended message to queue Job.Schedule.Test")

	jobHistory.Status = enums.Completed.String()
	jobHistory.CreateAt = time.Now()

	err = s.js.SaveToTableHistoryExecution(context.Background(), jobHistory)
	if err != nil {
		logger.Log.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Publish Job.",
		"job":     job,
	})
}
