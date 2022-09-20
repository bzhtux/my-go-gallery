package handlers

import (
	"net/http"

	"github.com/bzhtux/my-go-gallery/backend/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetImageByName(filename string) {
	var i models.Image
	result := h.DB.Where("Name = ?", filename).First(&i)
	if result.RowsAffected == 0 {
		// No record found with this ID
		c.JSON(http.StatusOK, gin.H{"Status": "Image " + filename + " does not exist in DB"})
	} else {
		// Found record here
		c.JSON(http.StatusNotFound, gin.H{"Status": "Image with Name " + filename + " already exists", "ID": i.ID, "Uploaded at": i.CreatedAt})
	}
}
