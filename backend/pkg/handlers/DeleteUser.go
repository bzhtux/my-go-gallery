package handlers

import (
	"net/http"
	"time"

	"github.com/bzhtux/my-go-gallery/backend/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) DeleteUser(c *gin.Context) {
	userID := c.Params.ByName("uid")
	var user = models.User{}
	var token = models.RegistrationToken{}
	result := h.DB.Where("ID = ?", userID).First(&user)
	user_email := user.Email
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Status": "User " + user_email + " does not exist!"})
	} else {
		now := time.Now()
		// old_email := user.Email
		new_email := user.Email + "." + now.String()
		h.DB.Model(&models.User{}).Where("ID = ? ", userID).Update("Email", new_email)
		h.DB.Delete(&user, userID)
		h.DB.Where("user_id = ?", userID).First(&token)
		h.DB.Delete(&token, token.ID)
		c.JSON(http.StatusOK, gin.H{"Status": "User with ID " + userID + " was successfully deleted!"})
	}
}
