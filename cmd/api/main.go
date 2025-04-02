package main

import (
	// Import Config Database
	"fmt"
	"net/http"

	"github.com/Hattaseakhiaw/lms-backend/config"
	_ "github.com/lib/pq" // Import Driver PostgreSQL
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	config.ConnectDB()

	// เริ่ม API Server
	fmt.Println("🚀 LMS API Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
