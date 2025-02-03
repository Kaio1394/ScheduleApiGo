package routes

import (
	"ScheduleApiGo/controllers"
	"ScheduleApiGo/repository/server"
	"ScheduleApiGo/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterServerRoute(r *gin.Engine, db *gorm.DB) {
	serverRepo := server.NewServerRepository(db)
	serverService := service.NewServerService(serverRepo)
	serverController := controllers.NewServerController(serverService)

	r.POST("/server", serverController.CreateServer)
	r.GET("/server", serverController.GetServers)
}
