package service

import (
	"ScheduleApiGo/helper"
	"ScheduleApiGo/logger"
	"errors"
)

type PublishService struct {
	rabbit *helper.Rabbit
	//repo   publish.PublishRepositoryImpl
}

func NewPublishService(rabbit *helper.Rabbit) *PublishService {
	return &PublishService{rabbit: rabbit}
}

func (ps *PublishService) Publish(topic string, payload string) error {
	if ps.rabbit == nil {
		logger.Log.Error("erro: service.rabbit its nil")
		return errors.New("erro: service.rabbit its nil")
	}
	con, err := ps.rabbit.Connection()
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	err = ps.rabbit.SendMessage(payload, topic, con)
	if err != nil {
		logger.Log.Error(err.Error())
		return err
	}
	return nil
}
