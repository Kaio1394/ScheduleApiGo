package controllers

import (
	"ScheduleApiGo/auth"
	"ScheduleApiGo/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	user := c.GetHeader("user")
	userId := c.GetHeader("userId")

	token, err := auth.GenerateJWT(userId, user)
	if err != nil {
		logger.Log.Error("Erro to generate token: " + err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"jwt_token": token,
	})

}
