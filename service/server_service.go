package service

import (
	"ScheduleApiGo/model"
	"ScheduleApiGo/repository"
	"context"
)

type ServerService struct {
	repo *repository.ServerRepository
}

func NewServerService(repo *repository.ServerRepository) *ServerService {
	return &ServerService{repo: repo}
}

func (s *ServerService) CreateServer(ctx context.Context, tag string, ip string) (int, error) {
	return s.repo.Create(ctx, tag, ip)
}
func (s *ServerService) GetServer(ctx context.Context) ([]model.Server, error) {
	return s.repo.GetServers(ctx)
}
