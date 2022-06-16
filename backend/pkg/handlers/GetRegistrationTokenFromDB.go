package handlers

import (
	"github.com/bzhtux/my-go-gallery/backend/models"
)

func (h Handler) GetRegistrationTokenFromDB(email string) string {
	var token = models.RegistrationToken{}
	result := h.DB.Where("Email = ?", email).First(&token)
	if result.RowsAffected == 0 {
		return ""
	} else {
		return token.Value
	}
}
