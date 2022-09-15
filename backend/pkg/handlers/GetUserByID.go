package handlers

import (
	"net/http"
	"strconv"

	"github.com/bzhtux/my-go-gallery/backend/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetUserByID(c *gin.Context) {
	userID := c.Params.ByName("uid")
	intVal, _ := strconv.Atoi(userID)
	// db := db.OpenDB()
	// db := db.Conn{}
	var user = models.User{}
	result := h.DB.Where("ID = ?", intVal).First(&user)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No Username with id " + userID})
	} else {
		c.JSON(http.StatusOK, gin.H{"firstname": user.FirstName, "lastname": user.LastName, "email": user.Email, "nickname": user.NickName})
	}
}
