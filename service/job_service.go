package service

import (
	"ScheduleApiGo/model"
	"ScheduleApiGo/repository"
	"context"
)

type JobService struct {
	repo *repository.JobRepository
}

func NewJobService(repo *repository.JobRepository) *JobService {
	return &JobService{repo: repo}
}

func (s *JobService) CreateJob(ctx context.Context, job *model.Job) (int, error) {
	return s.repo.Create(ctx, job)
}

func (s *JobService) GetJobs(ctx context.Context) ([]model.Job, error) {
	return s.repo.GetJobs(ctx)
}
