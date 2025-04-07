package routes

import (
	"github.com/Hattaseakhiaw/lms-backend/internal/handlers"
	"github.com/Hattaseakhiaw/lms-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Group สำหรับ route ที่ต้องใช้ JWT Auth
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/leaves", handlers.CreateLeave)
		auth.GET("/leaves/me", handlers.GetMyLeaves)
	}

	// Public route ไม่ต้องมี JWT
	r.POST("/login", handlers.Login)
	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)

	return r
}
