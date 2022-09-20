package handlers

import (
	"errors"
	"os"
)

func ImageExistsInDir(filename string) bool {
	if _, err := os.Stat(dest + "/" + filename); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}
