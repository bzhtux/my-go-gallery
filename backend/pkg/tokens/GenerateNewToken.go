package tokens

import (
	"crypto/rand"
	"encoding/hex"
)

var (
	length = 36
)

func GenerateNewtoken() string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	} else {
		return hex.EncodeToString(b)
	}
}
