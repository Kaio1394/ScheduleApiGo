package auth

import (
	"ScheduleApiGo/logger"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getSecretKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		logger.Log.Error()
		return nil
	}
	return []byte(secret)
}

func GenerateJWT(userID string, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID,
		"name": username,
		"exp":  time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString(getSecretKey())
}
