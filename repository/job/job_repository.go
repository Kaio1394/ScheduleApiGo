package job

import (
	"ScheduleApiGo/model"
	"context"
)

type IJobRepository interface {
	GetJobs(ctx context.Context) ([]model.Job, error)
	Create(ctx context.Context, job *model.Job) (int, error)
	GetJobById(ctx context.Context, id int) (*model.Job, error)
}
