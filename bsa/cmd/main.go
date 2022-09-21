package main

import (
	"fmt"
	"os"

	"github.com/bzhtux/my-go-gallery/bsa/models"
	"github.com/bzhtux/my-go-gallery/bsa/pkg/db"
	"github.com/bzhtux/my-go-gallery/bsa/pkg/handlers"

	"github.com/gin-gonic/gin"
)

const (
	version = "v0.0.8"
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
	h := handlers.New(dbConn)
	dbConn.AutoMigrate(&models.Image{})
	dbConn.AutoMigrate(&models.User{})
	uid := h.AddDefaultUser()
	fmt.Println("User Anonymous has ID:", uid)
	fmt.Println("***********************************\033[0m")

	router := gin.Default()
	router.MaxMultipartMemory = 16 << 32 // 16 MiB

	router.GET("/user/:uid", h.GetUserByID)
	router.POST("/user", h.AddNewUser)
	router.POST("/user/auth", h.AuthUser)
	router.POST("/image/upload", h.UploadImage)
	router.GET("image/:id", h.GetImageByID)
	router.DELETE("/image/delete/:id", h.DeleteImage)
	router.DELETE("/user/delete/:uid", h.DeleteUser)

	router.Run(":" + os.Getenv("APP_PORT"))

}
