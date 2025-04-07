package models

// Leave represents the leaves table
type Leave struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`                                              // Primary Key (id)
	UserID    uint   `gorm:"not null"`                                                              // Foreign Key to User (user_id)
	StartDate string `gorm:"not null"`                                                              // Start date of the leave
	EndDate   string `gorm:"not null"`                                                              // End date of the leave
	Type      string `gorm:"check:type in ('vacation', 'sick', 'personal')"`                        // Type of leave
	Status    string `gorm:"default:'pending';check:status in ('pending', 'approved', 'rejected')"` // Status of leave
}
