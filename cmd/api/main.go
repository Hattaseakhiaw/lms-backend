package main

import (
	"d:/lms-backend/config" // Import Config Database
	"fmt"
	"net/http"

	_ "github.com/lib/pq" // Import Driver PostgreSQL
)

func main() {
	// เชื่อมต่อฐานข้อมูล
	config.ConnectDB()

	// เริ่ม API Server
	fmt.Println("🚀 LMS API Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
