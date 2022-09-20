package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null" json:"firstname" binding:"required"`
	LastName  string `gorm:"not null" json:"lastname" binding:"required"`
	Password  string `gorm:"not null" json:"password" binding:"required"`
	Email     string `gorm:"not null,unique;index" json:"email" binding:"required"`
	NickName  string `json:"nickname"`
	Avatar    string
	Valid     bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
