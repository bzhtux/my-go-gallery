package handlers

import (
	"net/http"

	"github.com/bzhtux/my-go-gallery/backend/models"
	"github.com/gin-gonic/gin"
)

var (
	dest = "/Users/yfoeillet/go/src/github.com/bzhtux/my-go-gallery/backend/uploaded_files"
)

func (h Handler) UploadImage(c *gin.Context) {
	// Upload image first and if upload is successful record image in DB
	filename, _ := c.FormFile("file")
	if !h.ImageExistsInDB(filename.Filename) {
		if h.RecordImage(filename.Filename) {
			dst := dest + "/" + filename.Filename
			err := c.SaveUploadedFile(filename, dst)
			var i = models.Image{}
			h.DB.Where("Name = ?", filename.Filename).First(&i)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Status": "images " + filename.Filename + " was NOT uploaded", "ID": 0})
			} else {
				c.JSON(http.StatusOK, gin.H{"Status": "images " + filename.Filename + " was uploaded", "ID": i.ID})
			}
		}
	} else {
		c.JSON(http.StatusConflict, gin.H{"Status": "Image " + filename.Filename + " already exists"})
	}

}
