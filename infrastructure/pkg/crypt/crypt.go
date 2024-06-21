package crypt

import (
	"crypto/rand"
	"encoding/hex"
)

const RandomNumberLen = 10

const MobileAesKey = "5A2E746B08D846502F37A6E2D85D583B"

func RandomString(len int) (string, error) {
	bytes := make([]byte, len)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func byte16ToBytes(in [16]byte) []byte {
	tmp := make([]byte, 0, 16)
	for _, value := range in {
		tmp = append(tmp, value)
	}
	return tmp
}
