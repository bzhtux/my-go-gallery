package handlers

import (
	"log"

	"github.com/bzhtux/my-go-gallery/backend/models"
)

func (h Handler) ImageExistsInDB(filename string) bool {
	var i = models.Image{}
	result := h.DB.Where("Name = ?", filename).First(&i)
	log.Println("Results: ", result)
	if result.RowsAffected == 0 {
		return false
	} else {
		return true
	}
}
