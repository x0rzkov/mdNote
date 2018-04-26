package handler

import (
	"crypto/rand"
	"io"

	"github.com/jinzhu/gorm"
)

type Handler struct {
	DB        *gorm.DB
	SecretKey []byte
}

func GenerateRandomKey(length int) []byte {
	k := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return nil
	}
	return k
}
