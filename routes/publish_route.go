package routes

import (
	"ScheduleApiGo/controllers"
	"ScheduleApiGo/helper"
	"ScheduleApiGo/repository/job"
	"ScheduleApiGo/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Consume jobs of RabbitMQ
// @Summary Consume a job
// @Description Publishes a job message to RabbitMQ with connection parameters
// @Tags Consume
// @Accept json
// @Produce json
// @Param host header string true "RabbitMQ host"
// @Param port header string true "RabbitMQ port"
// @Param user header string true "RabbitMQ user"
// @Param password header string true "RabbitMQ password"
// @Param queue query string true "RabbitMQ queue"
// @Success 201 {object} map[string]string "Consumer success"
// @Failure 400 {object} map[string]interface{} "Error response"
// @Router /consumer/start [post]
func RegisterPublishJobRoute(r *gin.Engine, db *gorm.DB) {
	//configs, err := viper.ConfigSet()
	//if err != nil {
	//	logger.Log.Error(err.Error())
	//	return
	//}
	rabbit := helper.Rabbit{}
	publishService := service.NewPublishService(&rabbit)
	repoJob := job.NewJobRepository(db)
	jobService := service.NewJobService(repoJob)
	publishController := controllers.NewPublishController(publishService, jobService)
	r.POST("/publish/job", publishController.Publish)
}
