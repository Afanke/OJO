package ctrl

import (
	"crypto/sha256"
	"fmt"
)

func SHA256(str string, salt string) (result string) {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(salt+str+salt)))
}

func Encrypt(str string) (result string) {
	return SHA256(str, "ojo")
}

func EqualIfEncrypt(before, after string) bool {
	return Encrypt(before) == after
}
