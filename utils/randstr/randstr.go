package randstr

import (
	"math/rand"
)

var BigSam = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

//var Sam =[]byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func BigRandN(i int) string {
	var b = make([]byte, i, i)
	for j := 0; j < i; j++ {
		b[j] = BigSam[rand.Intn(36)]
	}
	return string(b)
}
