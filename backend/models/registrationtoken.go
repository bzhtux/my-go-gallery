package models

import (
	"time"

	"gorm.io/gorm"
)

type RegistrationToken struct {
	ID        uint   `gorm:"primaryKey"`
	Value     string `gorm:"not null;unique" json:"token" binding:"required"`
	UserID    uint
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
