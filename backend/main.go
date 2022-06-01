package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	// "time"

	"github.com/bzhtux/my-go-gallery/backend/db"
	"github.com/bzhtux/my-go-gallery/backend/images"
	"github.com/bzhtux/my-go-gallery/backend/users"
	"github.com/gin-gonic/gin"
)

const (
	version = "0.0.1"
)

var (
	DBUser       = os.Getenv("DB_USER")
	DBName       = os.Getenv("DB_NAME")
	DBHost       = os.Getenv("DB_HOST")
	DBPort       = os.Getenv("DB_PORT")
	DBPassword   = os.Getenv("DB_PASSWORD")
	SMTPUSer     = os.Getenv("SMTP_USER")
	SMTPPassword = os.Getenv("SMTP_PASSWORD")
	SMTPHost     = os.Getenv("SMTP_HOST")
	SMTPPort     = os.Getenv("SMTP_PORT")
	record_user  = true
)

func main() {
	fmt.Println("Starting with version " + version)

	dbConn := db.OpenDB()
	dbConn.AutoMigrate(&images.Image{})
	dbConn.AutoMigrate(&users.User{})
	uid := users.AddDefaultUser(dbConn)
	fmt.Println(uid)

	router := gin.Default()
	router.MaxMultipartMemory = 16 << 32 // 16 MiB

	router.GET("/user/:uid", func(c *gin.Context) {
		uid := c.Params.ByName("uid")
		intVal, _ := strconv.Atoi(uid)
		user := users.GetUserByID(dbConn, intVal)
		if user != nil {
			uFName := user.FirstName
			uLName := user.LastName
			uEmail := user.Email
			uNick := user.NickName
			c.JSON(http.StatusOK, gin.H{"firstname": uFName, "lastname": uLName, "email": uEmail, "nickname": uNick})
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

	router.POST("/image/upload", images.UploadImage)
	router.GET("image/:id", images.GetImageByID)
	router.DELETE("/image/delete/:id", images.DeleteImage)

	router.Run(":" + os.Getenv("APP_PORT"))

}
