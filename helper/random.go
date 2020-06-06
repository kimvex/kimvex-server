package helper

import (
	"crypto/rand"
	"encoding/hex"
)

//RandomCode helper
func RandomCode(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
