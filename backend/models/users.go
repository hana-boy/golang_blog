package models

import "time"

type User struct {
	ID        uint       `gorm:"primary_key"`
	Name      string     `gorm:"size:255"`
	Email     string     `gorm:"size:255"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"index"`
}
