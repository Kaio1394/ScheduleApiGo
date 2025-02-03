package service

import (
	"ScheduleApiGo/model"
	"ScheduleApiGo/repository/server"
	"context"
)

type ServerService struct {
	repo *server.ServerRepositoryimpl
}

func NewServerService(repo *server.ServerRepositoryimpl) *ServerService {
	return &ServerService{repo: repo}
}

func (s *ServerService) CreateServer(ctx context.Context, tag string, ip string) (int, error) {
	return s.repo.Create(ctx, tag, ip)
}
func (s *ServerService) GetServer(ctx context.Context) ([]model.Server, error) {
	return s.repo.GetServers(ctx)
}
