package handlers

import (
	"net/http"
	"time"

	"github.com/bzhtux/my-go-gallery/bsa/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) DeleteUser(c *gin.Context) {
	userID := c.Params.ByName("uid")
	var user = models.User{}
	result := h.DB.Where("ID = ?", userID).First(&user)
	user_email := user.Email
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "User " + user_email + " does not exist!",
		})
	} else {
		now := time.Now()
		new_email := user.Email + "." + now.String()
		h.DB.Model(&models.User{}).Where("ID = ? ", userID).Update("Email", new_email)
		h.DB.Delete(&user, userID)
		c.JSON(http.StatusOK, gin.H{
			"status":  "Deleted",
			"message": "User with ID " + userID + " was successfully deleted",
		})
	}
}
