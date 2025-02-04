package publish

import (
	"ScheduleApiGo/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PublishRepositoryImpl struct {
	db *gorm.DB
}

func NewRepositoryImpl(db *gorm.DB) *PublishRepositoryImpl {
	return &PublishRepositoryImpl{db: db}
}

func (r *PublishRepositoryImpl) GetjobById(c *gin.Context) (*model.Job, error) {
	var job model.Job
	id := c.Param("id")
	if err := r.db.First(&job, id).Error; err != nil {
		return nil, err
	}
	return &job, nil
}
