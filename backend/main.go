package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/bzhtux/my-go-gallery/backend/db"
	"github.com/bzhtux/my-go-gallery/backend/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	version = "0.0.1"
)

var (
	DBUser      = os.Getenv("DB_USER")
	DBName      = os.Getenv("DB_NAME")
	DBHost      = os.Getenv("DB_HOST")
	DBPort      = os.Getenv("DB_PORT")
	DBPassword  = os.Getenv("DB_PASSWORD")
	record_user = true
)

type Image struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	UserID    uint
	User      users.User
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// type User struct {
// 	ID        uint `gorm:"primaryKey"`
// 	Name      string
// 	Password  string
// 	Email     string `gorm:"unique"`
// 	NickName  string
// 	Avatar    string
// 	Valid     bool
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

func main() {
	fmt.Println("Starting with version " + version)
	// encrypt_pass, err := users.HashPassword("tuxpasss")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println("Encrypted Password:", encrypt_pass)
	// if !users.ComparePassword("tuxpasss", encrypt_pass) {
	// 	// fmt.Println("Password mismatch !")
	// 	log.Println("Password mismatch, authentication failed")
	// 	record_user = false
	// 	os.Exit(1)
	// }
	// if record_user {
	// 	fmt.Println("Password match !")
	// 	os.Exit(0)
	// }

	// var img = Image{}
	// var user = User{}

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)
	// if !db.HealthCheck(dsn) {
	// 	msg := "Could not connect to DB"
	// }
	dbConn := db.OpenDB(dsn)
	// fmt.Println(dbConn)
	dbConn.AutoMigrate(&Image{})
	dbConn.AutoMigrate(&users.User{})
	uid := users.AddDefaultUser(dbConn)
	fmt.Println(uid)

	router := gin.Default()

	router.GET("/user/:uid", func(c *gin.Context) {
		uid := c.Params.ByName("uid")
		intVal, _ := strconv.Atoi(uid)
		user := users.GetUserByID(dbConn, intVal)
		if user != nil {
			uFName := user.FirstName
			uLName := user.LastName
			uEmail := user.Email
			uNick := user.NickName
			c.JSON(http.StatusOK, gin.H{"username": uFName, "lastname": uLName, "email": uEmail, "nickname": uNick})
		}
		if user == nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "No Username with id " + uid})
		}

	})

	router.POST("/user", func(c *gin.Context) {
		var nu users.User
		err := c.BindJSON(&nu)
		if err != nil {
			log.Println(err)
			record_user = false
		}
		if !record_user {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Can not register new user: " + nu.FirstName})
		}
		if record_user {
			newUserID := users.AddNewUser(dbConn, nu.FirstName, nu.LastName, nu.Password, nu.Email, nu.NickName)
			fmt.Println("New User ID:", newUserID)
			c.JSON(http.StatusOK, gin.H{"userID": newUserID})
		}

	})

	router.POST("/user/auth", func(c *gin.Context) {
		var nu users.User
		err := c.BindJSON(&nu)
		if err != nil {
			// log.Println("BindJSONError: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		}
		dbPass := users.GetPasswordFromEmail(dbConn, nu.Email)
		if !users.ComparePassword(nu.Password, dbPass) {
			c.JSON(http.StatusForbidden, gin.H{"Status": "Unauthorized"})
		}
		if users.ComparePassword(nu.Password, dbPass) {

			c.JSON(http.StatusOK, gin.H{"Status": "Authorized - creating JWT"})
		}

	})
	router.Run(":" + os.Getenv("APP_PORT"))

	// demo := Image{Name: "demo.jpg", UserID: yann.ID}

	// _ = db.Create(&demo)

	// fmt.Println(demo.ID)

	// db.Find(&img)

	// fmt.Println("All Images: ", img.ID)

}
