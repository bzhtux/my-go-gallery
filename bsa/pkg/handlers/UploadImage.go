package handlers

import (
	"net/http"
	"os"

	"github.com/bzhtux/my-go-gallery/bsa/models"
	"github.com/gin-gonic/gin"
)

var (
	// dest = "/Users/yfoeillet/go/src/github.com/bzhtux/my-go-gallery/bsa/uploaded_files"
	// dest = "/uploaded_files"
	dest = os.Getenv("UPLOAD_DIR")
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
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "Failed",
					"message": "images " + filename.Filename + " was NOT uploaded",
					"data": gin.H{
						"ID":    0,
						"Error": err,
					},
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status":  "Uploaded",
					"message": "images " + filename.Filename + " was successfuly uploaded",
					"data": gin.H{
						"ID": i.ID,
					},
				})
			}
		}
	} else {
		h.GetImageByName(c)
	}

}
