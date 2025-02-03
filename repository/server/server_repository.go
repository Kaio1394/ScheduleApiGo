package server

import (
	"ScheduleApiGo/model"
	"context"
)

type ServerRepository interface {
	Create(ctx context.Context, tag string, ip string) (int, error)
	GetServers(ctx context.Context) ([]model.Server, error)
}
