package handlers

import (
	"net/http"

	"github.com/bzhtux/my-go-gallery/backend/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetImageByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var i models.Image
	result := h.DB.Where("ID = ?", id).First(&i)
	if result.RowsAffected == 0 {
		// No record found with this ID
		c.JSON(http.StatusNotFound, gin.H{"Status": "Image with ID " + id + " was not found"})
	} else {
		// Found record here
		var u = models.User{}
		h.DB.Where("Id = ?", i.UserID).First(&u)
		c.JSON(http.StatusOK, gin.H{"ID": i.ID, "Name": i.Name, "Uploaded by": u.NickName, "Uploaded": i.CreatedAt})
	}
}
