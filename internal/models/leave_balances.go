package models

// LeaveBalance represents the leave_balances table
type LeaveBalance struct {
	ID        uint `gorm:"primaryKey;autoIncrement"` // Primary Key (id)
	UserID    uint `gorm:"not null"`                 // Foreign Key to User (user_id)
	Year      int  `gorm:"not null"`                 // Year for leave balance
	TotalDays int  `gorm:"default:10;not null"`      // Total leave days available for the user
	UsedDays  int  `gorm:"default:0;not null"`       // Total leave days used by the user
}
