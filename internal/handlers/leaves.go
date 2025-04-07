package handlers

import (
	"net/http"

	"github.com/Hattaseakhiaw/lms-backend/config"
	"github.com/Hattaseakhiaw/lms-backend/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateLeave(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ไม่พบ user"})
		return
	}

	var leave models.Leave
	if err := c.ShouldBindJSON(&leave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ข้อมูลไม่ถูกต้อง"})
		return
	}

	leave.UserID = userID.(uint)

	if err := config.DB.Create(&leave).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถบันทึกวันลาได้"})
		return
	}

	c.JSON(http.StatusOK, leave)
}
