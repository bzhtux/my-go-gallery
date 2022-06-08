package handlers

import "github.com/bzhtux/my-go-gallery/backend/models"

func (h Handler) ImageExistsInDB(filename string) bool {
	var i = models.Image{}
	result := h.DB.Where("Name = ?", filename).First(&i)
	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
