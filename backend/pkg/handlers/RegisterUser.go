package handlers

import (
	"net/http"

	"github.com/bzhtux/my-go-gallery/backend/models"
	"github.com/bzhtux/my-go-gallery/backend/pkg/tokens"
	"github.com/gin-gonic/gin"
)

func (h Handler) RegisterUser(c *gin.Context) {
	userEmail := c.Params.ByName("email")
	userRegToken := c.Params.ByName("registrationtoken")
	var user = models.User{}
	result := h.DB.Where("Email = ?", userEmail).First(&user)
	if result != nil {
		var token = models.RegistrationToken{}
		h.DB.Where("user_id = ?", user.ID).First(&token)
		if tokens.TokenIsValid(userRegToken, token.Value) {
			// token is valid, let's update user table to activate user
			if h.ActivateUser(user.ID) {
				if h.DeleteRegistrationToken(user.ID) {
					c.JSON(http.StatusOK, gin.H{"Status": "User " + userEmail + "is now valid"})
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"Status": "Can not register user with token"})
				}
				// c.JSON(http.StatusOK, gin.H{"Status": "User " + userEmail + "is now valid"})
			}
		} else {
			// token is not valid, nothing to do ...
			c.JSON(http.StatusBadRequest, gin.H{"Status": "Can not register user with token"})
		}
	}
}
