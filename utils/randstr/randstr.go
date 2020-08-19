package randstr

import (
	"math/rand"
	"time"
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

func RandInt(num int) string {
	bs := make([]byte, num, num)
	rand.Seed(time.Now().Unix())
	rand.Intn(10)
	for i, j := 0, len(bs); i < j; i++ {
		bs[i] = byte(48 + rand.Intn(10))
	}
	return string(bs)
}
