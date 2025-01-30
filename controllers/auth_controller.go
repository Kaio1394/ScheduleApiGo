package controllers

import (
	"ScheduleApiGo/logger"
	"ScheduleApiGo/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	user := c.GetHeader("user")
	userId := c.GetHeader("userId")

	if user == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "User is required"})
		return
	}
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "UserId is required"})
		return
	}
	token, err := service.GenerateJWT(userId, user)
	if err != nil {
		logger.Log.Error("Erro to generate token: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"user":      user,
		"jwt_token": token,
		"createAt":  time.Now().Format(time.RFC3339),
	})
}
