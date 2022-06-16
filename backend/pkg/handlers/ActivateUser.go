package handlers

import "github.com/bzhtux/my-go-gallery/backend/models"

func (h Handler) ActivateUser(userid uint) bool {

	err := h.DB.Model(&models.User{}).Where("ID = ? ", userid).Update("Valid", true)
	if err != nil {
		return true
	} else {
		return false
	}

}
