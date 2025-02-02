package routes

import (
	"ScheduleApiGo/controllers"
	"ScheduleApiGo/repository"
	"ScheduleApiGo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterJobRoutes(r *gin.Engine, db *gorm.DB) {
	jobRepo := repository.NewJobRepository(db)
	jobService := service.NewJobService(jobRepo)
	jobController := controllers.NewJobController(jobService)
	r.POST("/job", jobController.CreateJob)
}

func RegisterListJobRoutes(r *gin.Engine) {
	r.GET("/list/job", controllers.ListJob)
}
