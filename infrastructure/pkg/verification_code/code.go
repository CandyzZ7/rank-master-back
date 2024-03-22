package verification_code

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	Six              = 6
	CodeValidityTime = 3
)

func GetRand(codeLen int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var s string
	for i := 0; i < codeLen; i++ {
		s += strconv.Itoa(r.Intn(10))
	}
	return s
}
