package middleware

import (
	"net/http"
	"strings"

	"github.com/Hattaseakhiaw/lms-backend/pkg/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "กรุณา login ก่อน"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token ไม่ถูกต้อง"})
			c.Abort()
			return
		}

		// เก็บ userID ใน context
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
