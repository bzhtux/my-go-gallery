package handlers

import "github.com/bzhtux/my-go-gallery/backend/models"

func (h Handler) RecordImage(img string) bool {
	image := models.Image{Name: img, UserID: 1}
	result := h.DB.Create(&image)
	if result.RowsAffected != 0 {
		return true
	} else {
		return false
	}

}
