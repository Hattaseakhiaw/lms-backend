package handlers

import (
	"net/http"

	"github.com/Hattaseakhiaw/lms-backend/config"
	"github.com/Hattaseakhiaw/lms-backend/internal/models"
	"github.com/Hattaseakhiaw/lms-backend/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email ไม่ถูกต้อง"})
		return
	}

	var user models.User
	if err := config.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ไม่พบผู้ใช้งาน"})
		return
	}

	// สร้าง JWT token
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถสร้าง token ได้"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}
