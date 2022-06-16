package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null" json:"firstname"`
	LastName  string `gorm:"not null" json:"lastname"`
	Password  string `gorm:"not null" json:"password"`
	Email     string `gorm:"not null;unique;index" json:"email" binding:"required"`
	NickName  string `json:"nickname"`
	Avatar    string
	Valid     bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
