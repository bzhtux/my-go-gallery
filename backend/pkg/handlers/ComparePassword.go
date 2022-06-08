package handlers

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(plain, hash string) bool {
	p_pass := []byte(plain)
	h_pass := []byte(hash)
	err := bcrypt.CompareHashAndPassword(h_pass, p_pass)
	if err != nil {
		log.Println(err)
		return false
	} else {
		return true
	}
}
