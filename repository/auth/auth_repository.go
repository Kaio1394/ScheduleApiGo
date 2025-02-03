package auth

import "ScheduleApiGo/model"

type AuthRepository interface {
	FindByUsername(username string) (*model.User, error)
}
