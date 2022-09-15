package handlers

import (
	"github.com/bzhtux/my-go-gallery/backend/models"
)

func (h Handler) GetPasswordFromEmail(email string) string {
	var u = models.User{}
	result := h.DB.Where("Email = ?", email).First(&u)
	if result.RowsAffected == 0 {
		return ""
	} else {
		return u.Password
	}
}
