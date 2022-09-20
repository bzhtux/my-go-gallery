package models

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null" json:"name" binding:"required"`
	UserID    uint
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
