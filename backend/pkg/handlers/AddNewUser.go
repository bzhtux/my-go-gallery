package handlers

import (
	"log"
	"net/http"

	"github.com/bzhtux/my-go-gallery/backend/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) AddNewUser(c *gin.Context) {
	var user = models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Status": "Can not register new user:" + err.Error()})
	} else {
		// log.Println(user.Email)
		result := h.DB.Where("Email = ?", user.Email).First(&user)
		if result.RowsAffected == 0 {
			// hpass := HashPassword(user.Password)
			user.Password = HashPassword(user.Password)
			h.DB.Create(&user)
			newRegToken := h.AddNewRegistrationToken(user.ID)
			// body := email.PrepareRegistrationEmail(newRegToken, user.FirstName, user.LastName, user.Email)
			// email.SendRegistrationEmail(user.FirstName, user.LastName, user.Email, newRegToken)
			log.Println(newRegToken)
			c.JSON(http.StatusOK, gin.H{"Status": user.Email + " was successfully created.", "userID": user.ID})
		} else {
			c.JSON(http.StatusConflict, gin.H{"Status": user.Email + " already exists.", "userID": user.ID})
		}
	}
}
