package main

import (
	"ScheduleApiGo/config"
	"ScheduleApiGo/logger"
	"ScheduleApiGo/routes"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ScheduleApiGo
// @version 1.0
// @description This is a sample API to demonstrate Swagger with Gin.
// @host localhost:8080
// @BasePath /
func main() {
	configs, err := configSet()
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

	r.Run(":" + configs.Port)
}

func configSet() (config.Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Error(err)
		return config.Config{}, err
	}

	var configs config.Config
	if err := viper.Unmarshal(&configs); err != nil {
		logger.Log.Error(err)
		return config.Config{}, err
	}

	return configs, nil
}
