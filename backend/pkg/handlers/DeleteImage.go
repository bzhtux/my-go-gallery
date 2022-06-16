package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/bzhtux/my-go-gallery/backend/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) DeleteImage(c *gin.Context) {
	id := c.Params.ByName("id")
	var i = models.Image{}
	result := h.DB.Where("ID = ?", id).First(&i)
	img_name := i.Name
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Status": "Image with ID " + id + " does not exist in DB"})
	} else {
		result := h.DB.Delete(&i, id)
		if result.RowsAffected != 0 {
			err := os.Remove(dest + "/" + img_name)
			if err != nil {
				log.Println(err)
			} else {
				c.JSON(http.StatusOK, gin.H{"Status": "Image ID " + id + " was successfully deleted"})
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": "Image ID " + id + " was not deleted"})
		}
	}

}
