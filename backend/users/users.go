package users

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"not null" json:"firstname"`
	LastName  string `gorm:"not null" json:"lastname"`
	Password  string `gorm:"not null" json:"password"`
	Email     string `gorm:"unique;index" json:"email"`
	NickName  string `json:"nickname"`
	Avatar    string
	Valid     bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func HashPassword(userpassword string) string {
	encrypt_pass, err := bcrypt.GenerateFromPassword([]byte(userpassword), 14)
	if err != nil {
		fmt.Println(err)
	}
	return string(encrypt_pass)
}

func GetPasswordFromEmail(db *gorm.DB, email string) string {
	var u = User{}
	result := db.Where("Email = ?", email).First(&u)
	if result.RowsAffected == 0 {
		return ""
	}
	return u.Password
}

func ComparePassword(plain, hash string) bool {
	p_pass := []byte(plain)
	h_pass := []byte(hash)
	err := bcrypt.CompareHashAndPassword(h_pass, p_pass)
	if err != nil {
		// log.Println("ComparePasswordError: ", err)
		return false
	}
	return true
}

func UserExists(db *gorm.DB, email string) bool {
	var u = User{}
	result := db.Where("Email = ?", email).First(&u)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func AddDefaultUser(db *gorm.DB) uint {
	var user = User{}
	anon := User{FirstName: "Anonymous", LastName: "RuleZ", Password: HashPassword("nimda"), Email: "ano@nymous.org", NickName: "Anon", Valid: true}
	result := db.Where("Email = ?", "ano@nymous.org").First(&user)

	if result.RowsAffected == 0 {
		db.Create(&anon)
		return anon.ID
	}

	db.Where("Email = ?", "ano@nymous.org").First(&user)
	fmt.Printf("User %s already exists\n", anon.FirstName)
	return user.ID

}

func AddNewUser(db *gorm.DB, uFName, uLName, uPass, uEmail, uNick string) uint {
	var user = User{}
	result := db.Where("Email = ?", uEmail).First(&user)

	user.FirstName = uFName
	user.LastName = uLName
	user.Password = HashPassword(uPass)
	user.Email = uEmail
	user.NickName = uNick

	if result.RowsAffected == 0 {
		db.Create(&user)
		return user.ID
	}

	fmt.Printf("User %s already exists\n", user.FirstName)
	return user.ID
}

func GetUserByID(db *gorm.DB, id int) *User {
	var user = User{}
	result := db.Where("ID = ?", id).First(&user)

	if result.RowsAffected == 0 {
		return nil
	}

	return &user
}
