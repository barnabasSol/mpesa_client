package stkpush

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
)

func generatePassword(shortCode string, timestamp string) string {
	shortCodePasskeyMap := map[string]string{
		"4646": "c737b0e916cd216d85a8d79053105cbbf59adf69534d02b94e269712629d41bc",
	}
	passkey, ok := shortCodePasskeyMap[shortCode]

	if !ok {
		return ""
	}

	data := shortCode + passkey + timestamp

	hashBytes := sha256.Sum256([]byte(data))
	hexString := fmt.Sprintf("%x", hashBytes)
	base64String := base64.StdEncoding.EncodeToString([]byte(hexString))

	log.Println("the base64")
	log.Println(base64String)
	return base64String
}
