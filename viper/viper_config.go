package viper

import (
	"ScheduleApiGo/config"
	"ScheduleApiGo/logger"

	"github.com/spf13/viper"
)

func ConfigSet() (config.Config, error) {
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
