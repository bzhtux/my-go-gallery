package handlers

import (
	"net/http"

	"github.com/bzhtux/my-go-gallery/backend/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetImageByName(c *gin.Context) {
	filename, _ := c.FormFile("file")
	var i models.Image
	result := h.DB.Where("Name = ?", filename.Filename).First(&i)
	if result.RowsAffected == 0 {
		// No record found with this ID
		c.JSON(http.StatusNotFound, gin.H{"Status": "Image " + filename.Filename + " does not exist in DB"})
	} else {
		// Found record here
		c.JSON(http.StatusConflict, gin.H{"Status": "Image with Name " + filename.Filename + " already exists", "ID": i.ID, "Uploaded at": i.CreatedAt})
	}
}
