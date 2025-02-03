package repository

import (
	"ScheduleApiGo/logger"
	"ScheduleApiGo/model"
	"context"

	"gorm.io/gorm"
)

type ServerRepository struct {
	db *gorm.DB
}

func NewServerRepository(db *gorm.DB) *ServerRepository {
	return &ServerRepository{db: db}
}

func (r *ServerRepository) Create(ctx context.Context, tag string, ip string) (int, error) {
	server := model.Server{
		Tag: tag,
		Ip:  ip,
	}

	if err := r.db.WithContext(ctx).Create(&server).Error; err != nil {
		logger.Log.Error("Error to insert a new server: " + err.Error())
		return 0, err
	}

	return server.Id, nil
}

func (r *ServerRepository) GetServers(ctx context.Context) ([]model.Server, error) {
	var servers []model.Server
	if err := r.db.WithContext(ctx).Find(&servers).Error; err != nil {
		logger.Log.Error("Error fetching jobs: " + err.Error())
		return nil, err
	}
	return servers, nil
}
