package service

import (
	"ScheduleApiGo/repository"
	"context"
	"errors"
)

type ServerService struct {
	repo *repository.ServerRepository
}

func NewServerService(repo *repository.ServerRepository) *ServerService {
	return &ServerService{repo: repo}
}

func (s *ServerService) CreateServer(ctx context.Context, tag string, ip string) (int, error) {
	if tag == "" || ip == "" {
		return 0, errors.New("Tag and IP are required")
	}

	serverID, err := s.repo.Create(ctx, tag, ip)
	if err != nil {
		return 0, err
	}

	return serverID, nil
}
