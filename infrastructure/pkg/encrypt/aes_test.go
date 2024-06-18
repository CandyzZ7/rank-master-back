package encrypt

import (
	"fmt"
	"testing"
)

func TestEncryptASEBase64ByECB(t *testing.T) {
	encryptedData, err := EncryptASEByECB(MobileAesKey, "13110708227")
	if err != nil {
		t.Error(err)
	}
	data, err := DecryptASEByECB(MobileAesKey, encryptedData)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(data)
}
