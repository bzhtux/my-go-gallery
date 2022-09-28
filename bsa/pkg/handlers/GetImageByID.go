package handlers

import (
	"net/http"

	"github.com/bzhtux/my-go-gallery/bsa/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetImageByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var i models.Image
	result := h.DB.Where("ID = ?", id).First(&i)
	if result.RowsAffected == 0 {
		// No record found with this ID
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not found",
			"message": "Image with ID " + id + " was not found",
		})
	} else {
		// Found record here
		c.JSON(http.StatusOK, gin.H{
			"status":  "Found",
			"message": "Image with ID " + id + " was found",
			"data": gin.H{
				"ID":          id,
				"Name":        i.Name,
				"Uploaded at": i.CreatedAt,
			},
		})
	}
}
