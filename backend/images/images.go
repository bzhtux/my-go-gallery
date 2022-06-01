package images

import (
	"errors"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bzhtux/my-go-gallery/backend/db"
	"github.com/bzhtux/my-go-gallery/backend/users"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	dest = "/Users/yfoeillet/go/src/github.com/bzhtux/my-go-gallery/backend/uploaded_files"
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

func RecordImage(img string) bool {
	db := db.OpenDB()
	image := Image{Name: img, UserID: 1}
	result := db.Create(&image)
	if result.RowsAffected != 0 {
		return true
	} else {
		return false
	}

}

func ImageExistsInDir(filename string) bool {
	if _, err := os.Stat(dest + "/" + filename); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}

func GetImageByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var i Image
	db := db.OpenDB()
	result := db.Where("ID = ?", id).First(&i)
	if result.RowsAffected == 0 {
		// No record found with this ID
		c.JSON(http.StatusNotFound, gin.H{"Status": "Image with ID " + id + " was not found"})
	} else {
		// Found record here
		c.JSON(http.StatusOK, gin.H{"Name": i.Name})
	}
}

func ImageExistsInDB(filename string) bool {
	db := db.OpenDB()
	var i = Image{}
	result := db.Where("Name = ?", filename).First(&i)
	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}

func UploadImage(c *gin.Context) {
	// Upload image first and if upload is successful record image in DB
	filename, _ := c.FormFile("file")
	if !ImageExistsInDB(filename.Filename) {
		if RecordImage(filename.Filename) {
			dst := dest + "/" + filename.Filename
			err := c.SaveUploadedFile(filename, dst)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Status": "images " + filename.Filename + " was NOT uploaded"})
			} else {
				c.JSON(http.StatusOK, gin.H{"Status": "images " + filename.Filename + " was uploaded"})
			}
		}
	} else {
		c.JSON(http.StatusConflict, gin.H{"Status": "Image " + filename.Filename + " already exists"})
	}

}

func DeleteImage(c *gin.Context) {
	id := c.Params.ByName("id")
	var i Image
	db := db.OpenDB()
	result := db.Where("ID = ?", id).First(&i)
	img_name := i.Name
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Status": "Image with ID " + id + " does not exist in DB"})
	} else {
		result := db.Delete(&i, id)
		if result.RowsAffected != 0 {
			err := os.Remove(dest + "/" + img_name)
			if err != nil {
				log.Println(err)
			}
			c.JSON(http.StatusOK, gin.H{"Status": "Image ID " + id + " was successfully deleted"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": "Image ID " + id + " was not deleted"})
		}
	}

}
