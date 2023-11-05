package encrypt

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"strings"
)

const RandomNumberLen = 10

func RandomString(len int) (string, error) {
	bytes := make([]byte, len)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func Encryption(target, cryptSalt string) string {
	return Md5Sum([]byte(strings.TrimSpace(target + cryptSalt)))
}

// EqualsEncryption 对比密码是否正确
func EqualsEncryption(target, encryptTarget string) bool {
	md5Target := Md5Sum([]byte(strings.TrimSpace(target)))
	if md5Target == encryptTarget {
		return true
	}
	return false
}

func Md5Sum(data []byte) string {
	return hex.EncodeToString(byte16ToBytes(md5.Sum(data)))
}

func byte16ToBytes(in [16]byte) []byte {
	tmp := make([]byte, 16)
	for _, value := range in {
		tmp = append(tmp, value)
	}
	return tmp[16:]
}
