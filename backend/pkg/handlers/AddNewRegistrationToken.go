package handlers

import (
	"github.com/bzhtux/my-go-gallery/backend/models"
	"github.com/bzhtux/my-go-gallery/backend/pkg/tokens"
)

func (h Handler) AddNewRegistrationToken(userID uint) string {
	var token = models.RegistrationToken{}
	result := h.DB.Where("user_id = ?", userID).First(&token)
	if result.RowsAffected == 0 {
		token.Value = tokens.GenerateNewtoken()
		token.UserID = userID
		err := h.DB.Create(&token)
		if err != nil {
			return token.Value
		} else {
			return ""
		}
	} else {
		return token.Value
	}
}
