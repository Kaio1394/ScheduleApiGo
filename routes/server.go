package routes

import (
	"ScheduleApiGo/controllers"
	"ScheduleApiGo/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterServerRoute(r *gin.Engine, db *gorm.DB) {
	serverRepo := repository.NewServerRepository(db)
	serverController := controllers.NewServerController(serverRepo)

	r.POST("/server", serverController.CreateServer)
}
