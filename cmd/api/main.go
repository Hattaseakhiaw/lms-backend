// main.go

package main

import (
	"fmt"

	"github.com/Hattaseakhiaw/lms-backend/config"
	"github.com/Hattaseakhiaw/lms-backend/internal/models"
	"github.com/Hattaseakhiaw/lms-backend/internal/routes"
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	config.ConnectDB()

	// ให้ GORM สร้างตารางอัตโนมัติ
	config.DB.AutoMigrate(
		&models.User{},
		&models.Leave{},
		&models.LeaveHistory{},
		&models.LeaveBalance{},
	)

	// เริ่ม API Server
	r := routes.SetupRouter()
	fmt.Println("🚀 LMS API Server is running on port 8080...")
	r.Run(":8080")
}
