package auth

import (
	"ScheduleApiGo/model"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{}
}

func (a AuthRepositoryImpl) FindByUsername(username string) (*model.User, error) {
	var user model.User
	if err := a.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
