package repository

import (
	"ScheduleApiGo/logger"
	"ScheduleApiGo/model"
	"context"

	"gorm.io/gorm"
)

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{db: db}
}

func (r *JobRepository) Create(ctx context.Context, job *model.Job) (int, error) {
	if err := r.db.WithContext(ctx).Create(&job).Error; err != nil {
		logger.Log.Error("Error to insert a new server: " + err.Error())
		return 0, err
	}
	return job.Id, nil
}
