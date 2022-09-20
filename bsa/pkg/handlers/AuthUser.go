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
		c.JSON(http.StatusBadRequest, gin.H{"Error": err})
	} else {
		dbPass := h.GetPasswordFromEmail(user.Email)
		if !ComparePassword(user.Password, dbPass) {
			c.JSON(http.StatusForbidden, gin.H{"Status": "Unauthorized"})
		} else {
			c.JSON(http.StatusOK, gin.H{"Status": "Authorized - creating JWT"})
		}
	}
}
