package users

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Password  string
	Email     string `gorm:"unique"`
	NickName  string
	Avatar    string
	Valid     bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func AutoMigrate(db *gorm.DB, database interface{}) {

	db.AutoMigrate(database)

}

func AddDefaultUser(db *gorm.DB) uint {
	var user = User{}
	anon := User{Name: "Anonymous", Password: "nimda", Email: "ano@nymous.org", NickName: "Anon", Valid: true}
	result := db.Where("Name = ?", "Anonymous").First(&user)

	if result.RowsAffected == 0 {
		db.Create(&anon)
		return anon.ID
	}

	db.Where("Name = ?", "Anonymous").First(&user)
	fmt.Printf("User %s already exists\n", anon.Name)
	return user.ID

}

func AddNewUser(db *gorm.DB, usr *User) uint {
	var user = User{}
	result := db.Where("Name = ?", usr.Name).First(&user)

	if result.RowsAffected == 0 {
		db.Create(&usr)
		return usr.ID
	}

	fmt.Printf("User %s already exists\n", usr.Name)
	return usr.ID
}
