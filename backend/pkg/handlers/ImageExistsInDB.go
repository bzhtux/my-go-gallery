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
		log.Println("False (expected 0) => result.RowsAffected: ", result.RowsAffected)
		return false
	} else {
		log.Println("True (image already exists) => result.RowsAffected: ", result.RowsAffected)
		return true
	}
}
