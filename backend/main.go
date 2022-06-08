package main

import (
	"fmt"
	"os"

	"github.com/bzhtux/my-go-gallery/backend/db"
	"github.com/bzhtux/my-go-gallery/backend/images"
	"github.com/bzhtux/my-go-gallery/backend/users"
	"github.com/gin-gonic/gin"
)

const (
	version = "0.0.2"
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
)

func main() {
	fmt.Println("\033[32m***********************************")
	fmt.Println("*** Starting with version " + version + " ***")
	fmt.Println("***********************************")

	dbConn := db.OpenDB()
	dbConn.AutoMigrate(&images.Image{})
	dbConn.AutoMigrate(&users.User{})
	uid := users.AddDefaultUser(dbConn)
	fmt.Println("User Anonymous has ID:", uid)
	fmt.Println("***********************************\033[0m")

	router := gin.Default()
	router.MaxMultipartMemory = 16 << 32 // 16 MiB

	router.GET("/user/:uid", users.GetUserByID)
	router.POST("/user", users.AddNewUser)
	router.POST("/user/auth", users.AuthUser)
	router.POST("/image/upload", images.UploadImage)
	router.GET("image/:id", images.GetImageByID)
	router.DELETE("/image/delete/:id", images.DeleteImage)
	router.DELETE("/user/delete/:uid", users.DeleteUser)

	router.Run(":" + os.Getenv("APP_PORT"))

}
