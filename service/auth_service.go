package service

import (
	"ScheduleApiGo/logger"
	"ScheduleApiGo/model"
	"ScheduleApiGo/repository/auth"
	"fmt"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	repo *auth.AuthRepositoryImpl
}

func NewAuthService(repo *auth.AuthRepositoryImpl) *AuthService {
	return &AuthService{repo: repo}
}

func getSecretKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		logger.Log.Error("JWT_SECRET wasn't configurate.")
		return nil
	}
	return []byte(secret)
}

func (service *AuthService) GenerateJWT(user *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Username,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(getSecretKey())
}

func (service *AuthService) ValidateJWT(tokenString string) (*jwt.Token, error) {
	secretKey := getSecretKey()
	if secretKey == nil {
		return nil, fmt.Errorf("missing secret key")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	return token, err
}

func (service *AuthService) VerifyPassword(providedPassword, storedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}

func (service *AuthService) Authenticate(username, password string) (string, error) {
	user, err := service.repo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if password != user.Password {
		return "", fmt.Errorf("Password invalid.")
	}

	token, err := service.GenerateJWT(user)
	if err != nil {
		return "", fmt.Errorf("Error to generate token: %v", err)
	}

	return token, nil
}
