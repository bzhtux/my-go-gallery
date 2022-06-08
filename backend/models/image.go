package models

import (
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	UserID    uint
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
