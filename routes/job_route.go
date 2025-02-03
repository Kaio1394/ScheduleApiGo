package routes

import (
	"ScheduleApiGo/controllers"
	"ScheduleApiGo/repository/job"
	"ScheduleApiGo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterJobRoutes(r *gin.Engine, db *gorm.DB) {
	jobRepo := job.NewJobRepository(db)
	jobService := service.NewJobService(jobRepo)
	jobController := controllers.NewJobController(jobService)
	r.POST("/job", jobController.CreateJob)
	r.GET("/job/list", jobController.GetJobs)
}
