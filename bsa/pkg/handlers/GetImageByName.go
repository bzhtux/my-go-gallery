package handlers

import (
	"net/http"

	"github.com/bzhtux/my-go-gallery/bsa/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetImageByName(c *gin.Context) {
	filename, _ := c.FormFile("file")
	var i models.Image
	result := h.DB.Where("Name = ?", filename.Filename).First(&i)
	if result.RowsAffected == 0 {
		// No record found with this ID
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not found",
			"message": "Image " + filename.Filename + " does not exist in DB",
		})
	} else {
		// Found record here
		c.JSON(http.StatusFound, gin.H{
			"status":  "Found",
			"message": "Image with Name " + filename.Filename + " was found",
			"data": gin.H{
				"ID":          i.ID,
				"Uploaded at": i.CreatedAt,
			},
		})
	}
}
