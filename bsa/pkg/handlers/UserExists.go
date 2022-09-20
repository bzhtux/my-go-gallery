package handlers

import "github.com/bzhtux/my-go-gallery/bsa/models"

func (h Handler) UserExists(email string) bool {
	var u = models.User{}
	result := h.DB.Where("Email = ?", email).First(&u)
	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
