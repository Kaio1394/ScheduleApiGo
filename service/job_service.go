package service

import (
	"ScheduleApiGo/model"
	"ScheduleApiGo/repository/job"
	"context"
)

type JobService struct {
	repo *job.JobRepositoryImpl
}

func NewJobService(repo *job.JobRepositoryImpl) *JobService {
	return &JobService{repo: repo}
}

func (s *JobService) CreateJob(ctx context.Context, job *model.Job) (int, error) {
	return s.repo.Create(ctx, job)
}

func (s *JobService) GetJobs(ctx context.Context) ([]model.Job, error) {
	return s.repo.GetJobs(ctx)
}

func (s *JobService) GetJobById(ctx context.Context, id int) (*model.Job, error) {
	return s.repo.GetJobById(ctx, id)
}
func (s *JobService) SaveToTableHistoryExecution(ctx context.Context, jobHistory model.HistoryExecution) error {
	return s.repo.SendJobToTableHistory(ctx, jobHistory)
}
