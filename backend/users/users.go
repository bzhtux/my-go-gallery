package users

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/bzhtux/my-go-gallery/backend/db"
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

func GetPasswordFromEmail(db *gorm.DB, email string) string {
	var u = User{}
	result := db.Where("Email = ?", email).First(&u)
	if result.RowsAffected == 0 {
		return ""
	} else {
		return u.Password
	}
}

func UserExists(db *gorm.DB, email string) bool {
	var u = User{}
	result := db.Where("Email = ?", email).First(&u)
	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

func HashPassword(userpassword string) string {
	encrypt_pass, err := bcrypt.GenerateFromPassword([]byte(userpassword), 14)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(encrypt_pass)
}

func ComparePassword(plain, hash string) bool {
	p_pass := []byte(plain)
	h_pass := []byte(hash)
	err := bcrypt.CompareHashAndPassword(h_pass, p_pass)
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
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

func AddNewUser(c *gin.Context) {
	db := db.OpenDB()
	var user = User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Can not register nes user:" + err.Error()})
	} else {
		result := db.Where("Email = ?", user.Email).First(&user)
		if result.RowsAffected == 0 {
			hpass := HashPassword(user.Password)
			user.Password = hpass
			db.Create(&user)
			c.JSON(http.StatusOK, gin.H{"userID": user.ID})
		} else {
			c.JSON(http.StatusConflict, gin.H{"Status": user.Email + " already exists"})
		}
	}
}

func GetUserByID(c *gin.Context) {
	userID := c.Params.ByName("uid")
	intVal, _ := strconv.Atoi(userID)
	db := db.OpenDB()
	// db := db.Conn{}
	var user = User{}
	result := db.Where("ID = ?", intVal).First(&user)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No Username with id " + userID})
	} else {
		c.JSON(http.StatusOK, gin.H{"firstname": user.FirstName, "lastname": user.LastName, "email": user.Email, "nickname": user.NickName})
	}
}

func AuthUser(c *gin.Context) {
	var user = User{}
	db := db.OpenDB()
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
	}
	dbPass := GetPasswordFromEmail(db, user.Email)
	if !ComparePassword(user.Password, dbPass) {
		c.JSON(http.StatusForbidden, gin.H{"Status": "Unauthorized"})
	} else {
		c.JSON(http.StatusOK, gin.H{"Status": "Authorized - creating JWT"})
	}
}

func DeleteUser(c *gin.Context) {
	userID := c.Params.ByName("uid")
	db := db.OpenDB()
	var user = User{}
	result := db.Where("ID = ?", userID).First(&user)
	user_email := user.Email
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Status": "User " + user_email + " does not exist!"})
	} else {
		now := time.Now()
		new_email := user.Email + "." + now.String()
		db.Model(&User{}).Where("ID = ? ", userID).Update("Email", new_email)
		db.Delete(&user, userID)
		c.JSON(http.StatusOK, gin.H{"Status": "User with ID " + userID + " was successfully deleted!"})
	}
}
