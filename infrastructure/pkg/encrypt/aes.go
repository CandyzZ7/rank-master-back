package encrypt

import (
	"encoding/base64"

	"github.com/zeromicro/go-zero/core/codec"
)

func EncryptASEByECB(key, encrypt string) (string, error) {
	data, err := codec.EcbEncrypt([]byte(key), []byte(encrypt))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func DecryptASEByECB(key, encrypt string) (string, error) {
	originalData, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", err
	}
	data, err := codec.EcbDecrypt([]byte(key), originalData)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func EncryptASEBase64ByECB(key, encrypt string) (string, error) {
	data, err := codec.EcbEncryptBase64(key, encrypt)
	if err != nil {
		return "", err
	}
	return data, nil
}

func DecryptASEBase64ByECB(key, encrypt string) (string, error) {
	data, err := codec.EcbDecryptBase64(key, encrypt)
	if err != nil {
		return "", err
	}
	return data, nil
}
