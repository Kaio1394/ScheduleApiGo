package publish

import (
	"ScheduleApiGo/model"
	"github.com/gin-gonic/gin"
)

type PublishRepository interface {
	GetjobById(c *gin.Context) (*model.Job, error)
}
