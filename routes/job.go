package routes

import (
	"ScheduleApiGo/controllers"
	"ScheduleApiGo/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterJobRoutes(r *gin.Engine, db *gorm.DB) {
	jobRepo := repository.NewJobRepository(db)
	jobController := controllers.NewJobController(jobRepo)
	r.POST("/job", jobController.CreateJob)
}

func RegisterListJobRoutes(r *gin.Engine) {
	r.GET("/list/job", controllers.ListJob)
}
