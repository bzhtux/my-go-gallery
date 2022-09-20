package handlers

import (
	"fmt"

	"github.com/bzhtux/my-go-gallery/bsa/models"
)

func (h Handler) AddDefaultUser() uint {
	var user = models.User{}
	anon := models.User{FirstName: "Anonymous", LastName: "RuleZ", Password: HashPassword("nimda"), Email: "ano@nymous.org", NickName: "Anon", Valid: true}
	result := h.DB.Where("Email = ?", "ano@nymous.org").First(&user)

	if result.RowsAffected == 0 {
		h.DB.Create(&anon)
		return anon.ID
	}

	h.DB.Where("Email = ?", "ano@nymous.org").First(&user)
	fmt.Printf("User %s already exists\n", anon.FirstName)
	return user.ID

}
