package models

// LeaveHistory represents the leave_history table
type LeaveHistory struct {
	ID       uint `gorm:"primaryKey;autoIncrement"` // Primary Key (id)
	UserID   uint `gorm:"not null"`                 // Foreign Key to User (user_id)
	LeaveID  uint `gorm:"not null"`                 // Foreign Key to Leave (leave_id)
	Year     int  `gorm:"not null"`                 // Year for leave history
	DaysUsed int  `gorm:"not null"`                 // Number of days used in the leave
}
