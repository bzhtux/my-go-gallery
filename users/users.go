package users

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(userpassword string) (string, error) {
	encrypt_pass, err := bcrypt.GenerateFromPassword([]byte(userpassword), 14)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(encrypt_pass), nil
}

func ComparePassword(plain, hash string) bool {
	p_pass := []byte(plain)
	h_pass := []byte(hash)
	err := bcrypt.CompareHashAndPassword(h_pass, p_pass)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
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

func GetUserByID(db *gorm.DB, id int) *User {
	var user = User{}
	result := db.Where("ID = ?", id).First(&user)

	if result.RowsAffected == 0 {
		return nil
	}

	return &user
}
