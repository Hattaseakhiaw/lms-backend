package models

// User represents the users table
type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`                       // Primary Key (id)
	Name  string `gorm:"not null"`                                       // Name of the user
	Email string `gorm:"unique;not null"`                                // Email of the user (unique)
	Role  string `gorm:"check:role in ('employee', 'manager', 'admin')"` // Role of the user (check constraint)
}
