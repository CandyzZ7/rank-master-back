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

// Md5Sum MD5加密多用于数据完整性校验。
// 例如，下载文件时，网站可能会提供该文件的MD5值，用户下载后通过计算文件的MD5值并与提供的值进行对比，就可以判断文件是否在传输过程中发生了改变。
// 此外，MD5也经常被用于存储用户的密码，例如在数据库中存储用户的MD5密码，即使数据库被泄露，攻击者也无法直接得知用户的真实密码。
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
