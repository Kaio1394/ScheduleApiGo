package routes

import (
	"ScheduleApiGo/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	r.POST("/token", controllers.Auth)
}
