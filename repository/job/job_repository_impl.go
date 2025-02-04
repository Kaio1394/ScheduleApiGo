package job

import (
	"ScheduleApiGo/logger"
	"ScheduleApiGo/model"
	"context"
	"gorm.io/gorm"
)

type JobRepositoryImpl struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepositoryImpl {
	return &JobRepositoryImpl{db: db}
}

func (r *JobRepositoryImpl) Create(ctx context.Context, job *model.Job) (int, error) {
	if err := r.db.WithContext(ctx).Create(&job).Error; err != nil {
		logger.Log.Error("Error to insert a new server: " + err.Error())
		return 0, err
	}
	return job.Id, nil
}

func (r *JobRepositoryImpl) GetJobs(ctx context.Context) ([]model.Job, error) {
	var jobs []model.Job
	if err := r.db.WithContext(ctx).Find(&jobs).Error; err != nil {
		logger.Log.Error("Error fetching jobs: " + err.Error())
		return nil, err
	}
	return jobs, nil
}

func (r *JobRepositoryImpl) GetJobById(ctx context.Context, id int) (*model.Job, error) {
	var job model.Job
	if err := r.db.WithContext(ctx).First(&job, id).Error; err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *JobRepositoryImpl) SendJobToTableHistory(ctx context.Context, jobHistoryExecution model.HistoryExecution) error {
	if err := r.db.WithContext(ctx).Create(&jobHistoryExecution).Error; err != nil {
		logger.Log.Error("Error to insert to history execution: " + err.Error())
		return err
	}
	return nil
}
