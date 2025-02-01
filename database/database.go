package database

import (
	"ScheduleApiGo/viper"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() *gorm.DB {
	configs, err := viper.ConfigSet()
	if err != nil {
		log.Fatalf("Fail to load configurations: %v", err)
	}

	connString := configs.DataBase.StringConnection
	typeDatabase := configs.DataBase.TypeDatabase

	var db *gorm.DB
	switch typeDatabase {
	case "postgres":
		db, err = gorm.Open(postgres.Open(connString), &gorm.Config{})
	default:
		log.Fatalf("Database not supported: %s", typeDatabase)
	}

	if err != nil {
		log.Fatalf("Fail to connect to database: %v", err)
	}

	fmt.Println("Connection with database successful!")

	return db
}
