package service

import (
	"ScheduleApiGo/model"
	"ScheduleApiGo/repository"
	"context"
	"errors"
)

type JobService struct {
	repo *repository.JobRepository
}

func NewJobService(repo *repository.JobRepository) *JobService {
	return &JobService{repo: repo}
}

func (s *JobService) CreateJob(ctx context.Context, job *model.Job) (int, error) {
	if job.Name == "" || job.ServerId == 0 {
		return 0, errors.New("Job name and ServerId are required")
	}

	jobID, err := s.repo.Create(ctx, job)
	if err != nil {
		return 0, err
	}

	return jobID, nil
}
