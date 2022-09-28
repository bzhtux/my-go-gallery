package handlers

import (
	"net/http"

	"github.com/bzhtux/my-go-gallery/bsa/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) AuthUser(c *gin.Context) {
	var user = models.User{}
	// db := db.OpenDB()
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": err,
		})
	} else {
		dbPass := h.GetPasswordFromEmail(user.Email)
		if !ComparePassword(user.Password, dbPass) {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  "Unauthorized",
				"message": "User " + user.Email + " is not authorized, please ensure credentials",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "Authorized",
				"Status": "User " + user.Email + " successfuly authenticated",
			})
		}
	}
}
