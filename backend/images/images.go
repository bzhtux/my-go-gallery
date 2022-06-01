package images

type Image struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	UserID    uint
	User      users.User
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
