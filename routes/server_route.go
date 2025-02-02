package routes

import (
	"ScheduleApiGo/controllers"
	"ScheduleApiGo/repository"
	"ScheduleApiGo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterServerRoute(r *gin.Engine, db *gorm.DB) {
	serverRepo := repository.NewServerRepository(db)
	serverService := service.NewServerService(serverRepo)
	serverController := controllers.NewServerController(serverService)

	r.POST("/server", serverController.CreateServer)
}
