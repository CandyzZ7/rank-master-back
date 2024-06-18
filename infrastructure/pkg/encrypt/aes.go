package encrypt

import (
	"github.com/zeromicro/go-zero/core/codec"
)

// EncryptASEBase64ByECB ECB模式
func EncryptASEBase64ByECB(key, encrypt string) (string, error) {
	data, err := codec.EcbEncryptBase64(encrypt, key)
	if err != nil {
		return "", err
	}
	return data, nil
}

// DecryptASEBase64ByECB ECB模式
func DecryptASEBase64ByECB(encrypt, key string) (string, error) {
	data, err := codec.EcbDecryptBase64(key, encrypt)
	if err != nil {
		return "", err
	}
	return data, nil
}
