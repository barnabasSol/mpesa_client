package stkpush

import (
	"crypto/sha256"
	"encoding/base64"
)

func generatePassword(shortCode string, timestamp string) string {
	var shortCodePasskeyMap = map[string]string{
		"2060": "5ab0ecb13d56a1818f182cbe463b84370c3768a5f3e355aa1dd706043d722dee",
		"4646": "c737b0e916cd216d85a8d79053105cbbf59adf69534d02b94e269712629d41bc",
		"8088": "cf9f7a9887e2aee5115c2f617629d523b0bbed6b6824bb11dca24aa97f682fd0",
		"6768": "43bfce8765c660777374d0961248fe1be877a882c4b49783a9e161922e393ab9",
		"6564": "d708c5edb61fd71261a42972dadeb6fdd162cd6918930260cf63ebcb60ff3998",
	}
	passkey, ok := shortCodePasskeyMap[shortCode]
	if !ok {
		return ""
	}

	data := shortCode + passkey + timestamp

	hash := sha256.Sum256([]byte(data))

	password := base64.StdEncoding.EncodeToString(hash[:])

	return password
}
