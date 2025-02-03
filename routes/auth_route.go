package routes

import (
	"ScheduleApiGo/controllers"
	"ScheduleApiGo/repository/auth"
	"ScheduleApiGo/service"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, db *gorm.DB) {
	authRepository := auth.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authController := controllers.NewAuthController(authService)
	r.POST("/infra/auth", authController.Authenticate)
}
