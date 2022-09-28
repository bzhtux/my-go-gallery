package handlers

import (
	"net/http"
	"strconv"

	"github.com/bzhtux/my-go-gallery/bsa/models"
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
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not found",
			"message": "No User found with this ID ",
			"data": gin.H{
				"ID": userID,
			},
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  "OK",
			"message": "User with ID " + userID + " was found",
			"data": gin.H{
				"Firstname": user.FirstName,
				"Lastname":  user.LastName,
				"Email":     user.Email,
				"Nickname":  user.NickName,
			},
		})
	}
}
