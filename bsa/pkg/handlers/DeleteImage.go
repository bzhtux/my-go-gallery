package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/bzhtux/my-go-gallery/bsa/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) DeleteImage(c *gin.Context) {
	id := c.Params.ByName("id")
	var i = models.Image{}
	result := h.DB.Where("ID = ?", id).First(&i)
	img_name := i.Name
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not found",
			"message": "Image with ID " + id + " was not found, ensure ID is correct",
		})
	} else {
		result := h.DB.Delete(&i, id)
		if result.RowsAffected != 0 {
			err := os.Remove(dest + "/" + img_name)
			if err != nil {
				log.Println(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  "deleted",
				"message": "Image with ID " + id + " was successfully deleted",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Internal Error",
				"message": "Image with ID " + id + " was not deleted",
			})
		}
	}

}
