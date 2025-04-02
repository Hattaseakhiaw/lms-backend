package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// อ่านค่าจาก ENV (ใช้ os.Getenv() เพื่อความยืดหยุ่น)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "lms_user"),
		getEnv("DB_PASSWORD", "Doo08092538"),
		getEnv("DB_NAME", "lms_db"),
		getEnv("DB_PORT", "5432"),
	)

	// เชื่อมต่อ Database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("❌ ไม่สามารถเชื่อมต่อ Database ได้: " + err.Error())
	}

	DB = db
	fmt.Println("✅ เชื่อมต่อฐานข้อมูลสำเร็จ!")
}

// ฟังก์ชันช่วยอ่านค่า ENV และใช้ค่า Default ถ้าไม่มี
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
