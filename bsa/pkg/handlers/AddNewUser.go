package handlers

import (
	"net/http"
	"strconv"

	"github.com/bzhtux/my-go-gallery/bsa/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) AddNewUser(c *gin.Context) {
	var user = models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad request",
			"message": "Can not register new user:" + err.Error(),
		})
	} else {
		// log.Println("AddNewUser Error: ", err.Error())
		result := h.DB.Where("Email = ?", user.Email).First(&user)
		if result.RowsAffected == 0 {
			hpass := HashPassword(user.Password)
			user.Password = hpass
			h.DB.Create(&user)
			c.JSON(http.StatusOK, gin.H{
				"status":  "created",
				"message": "New user with email " + user.Email + " was successfuly created",
				"data": gin.H{
					"ID":        strconv.FormatUint(uint64(user.ID), 10),
					"FirstName": user.FirstName,
					"LastName":  user.LastName,
					"Email":     user.Email,
					"NickName":  user.NickName,
				},
			})
		} else {
			c.JSON(http.StatusConflict, gin.H{
				"status":  "Conflict",
				"message": "User with email " + user.Email + " already exists",
				"data": gin.H{
					"ID": strconv.FormatUint(uint64(user.ID), 10),
				},
			})
		}
	}
}
