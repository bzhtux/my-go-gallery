package handlers

import "github.com/bzhtux/my-go-gallery/backend/models"

func (h Handler) DeleteRegistrationToken(userid uint) bool {
	var token = models.RegistrationToken{}
	h.DB.Where("user_id = ?", userid).First(&token)
	err := h.DB.Delete(&token, token.ID)
	if err != nil {
		return true
	} else {
		return false
	}
}
