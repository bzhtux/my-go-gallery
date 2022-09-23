package handlers

import "github.com/bzhtux/my-go-gallery/bsa/models"

func (h Handler) RecordImage(img string) bool {
	// Anonymous context so all images belong to anonumous user that should have id 1
	image := models.Image{Name: img, UserID: 1}
	result := h.DB.Create(&image)
	if result.RowsAffected != 0 {
		return true
	} else {
		return false
	}

}
