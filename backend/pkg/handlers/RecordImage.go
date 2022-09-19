package handlers

import "github.com/bzhtux/my-go-gallery/backend/models"

func (h Handler) RecordImage(img string) bool {
	// Anonymous context so all images belong to anonumous user that should have id 0
	image := models.Image{Name: img, UserID: 0}
	result := h.DB.Create(&image)
	if result.RowsAffected != 0 {
		return true
	} else {
		return false
	}

}
