package server

import (
	"ScheduleApiGo/logger"
	"ScheduleApiGo/model"
	"context"

	"gorm.io/gorm"
)

type ServerRepositoryimpl struct {
	db *gorm.DB
}

func NewServerRepository(db *gorm.DB) *ServerRepositoryimpl {
	return &ServerRepositoryimpl{db: db}
}

func (r *ServerRepositoryimpl) Create(ctx context.Context, tag string, ip string) (int, error) {
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

func (r *ServerRepositoryimpl) GetServers(ctx context.Context) ([]model.Server, error) {
	var servers []model.Server
	if err := r.db.WithContext(ctx).Find(&servers).Error; err != nil {
		logger.Log.Error("Error fetching jobs: " + err.Error())
		return nil, err
	}
	return servers, nil
}
