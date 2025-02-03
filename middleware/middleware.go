package middleware

import (
	"net/http"
	"strings"

	"ScheduleApiGo/service"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(authService *service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtém o cabeçalho Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// O token está no formato "Bearer <token>", então precisamos dividir
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Malformed token"})
			c.Abort()
			return
		}

		// Valida o token
		token, err := authService.ValidateJWT(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Se o token for válido, passamos para o próximo handler
		c.Next()
	}
}
