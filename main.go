package main

import (
	"ScheduleApiGo/config"
	"ScheduleApiGo/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	configs, err := configSet()
	if err != nil {
		logger.Log.Error("Error when trying to load configuration file: " + err.Error())
		return
	}
	logger.Log.Info("Start application " + configs.App.Name)
	r := gin.Default()
	r.GET("/job", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Teste",
		})
	})
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
