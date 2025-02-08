package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// Generate ID format YYMMDDHmmss + random 3 digit
func GenerateID() string {
	now := time.Now()
	timestamp := now.Format("060102150405") // YYMMDDHmmss

	// Tambahkan angka/huruf random untuk menghindari duplikasi
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	randomPart := make([]byte, 3)
	for i := range randomPart {
		randomPart[i] = charset[rand.Intn(len(charset))]
	}

	return fmt.Sprintf("%s%s", timestamp, string(randomPart))
}
