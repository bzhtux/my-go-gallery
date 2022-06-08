package handlers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(userpassword string) string {
	encrypt_pass, err := bcrypt.GenerateFromPassword([]byte(userpassword), 14)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(encrypt_pass)
}
