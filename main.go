package main

import (
	"ScheduleApiGo/database"
	"ScheduleApiGo/logger"
	"ScheduleApiGo/routes"
	"ScheduleApiGo/viper"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ScheduleApiGo
// @version 1.0
// @description This is a sample API to demonstrate Swagger with Gin.
// @host localhost:8080
// @BasePath /
func main() {

	db, err := database.ConnectDatabase()
	if err != nil {
		return
	}
	var currentDB string
	db.Raw("SELECT current_database()").Scan(&currentDB)
	logger.Log.Info("Connected to database: " + currentDB)

	configs, err := viper.ConfigSet()
	if err != nil {
		logger.Log.Error("Error when trying to load configuration file: " + err.Error())
		return
	}
	logger.Log.Info("Start application " + configs.App.Name)
	r := gin.Default()
	logger.Log.Info("Start Swagger.")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.RegisterPublishJobRoute(r)
	routes.RegisterAuthRoutes(r)
	routes.RegisterServerRoute(r, db)
	routes.RegisterJobRoutes(r, db)
	//routes.RegisterPublishJobRoute(r)
	r.Run(":" + configs.Server.Port)
}
