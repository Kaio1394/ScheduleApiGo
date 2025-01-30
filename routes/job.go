package routes

import (
	"ScheduleApiGo/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterJobRoutes(r *gin.Engine) {
	r.GET("/job", controllers.Job)

}

func RegisterListJobRoutes(r *gin.Engine) {
	r.GET("/list/job", controllers.ListJob)
}
