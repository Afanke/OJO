package ctrl

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/afanke/OJO/utils/log"
	"strconv"
)

func SHA256(str string) (result string) {
	return fmt.Sprintf("%x", sha256.Sum256([]byte("ojo"+str+"ojo")))
}

func EqualIfSHA256(before, after string) bool {
	return SHA256(before) == after
}

func DESEncrypt(p []byte, key string) ([]byte, error) {
	k := []byte(key)
	block, err := des.NewCipher(k)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	iv := bytes.Repeat([]byte("a"), blockSize)
	mode := cipher.NewCBCEncrypter(block, iv)
	p, err = paddingNumber(p, blockSize)
	if err != nil {
		return nil, err
	}
	mode.CryptBlocks(p, p)
	return p, nil
}

func DESDecrypt(d []byte, key string) ([]byte, error) {
	k := []byte(key)
	block, err := des.NewCipher(k)
	if err != nil {
		return nil, err
	}
	iv := bytes.Repeat([]byte("a"), block.BlockSize())
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(d, d)
	d, err = unPaddingNumber(d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

//填充数据
func paddingNumber(src []byte, blockSize int) ([]byte, error) {
	if src == nil {
		return nil, errors.New("src length can less than 0")
	}
	leftNumber := len(src) % blockSize
	needNumber := blockSize - leftNumber
	b := byte(needNumber)
	newSlice := bytes.Repeat([]byte{b}, needNumber)
	src = append(src, newSlice...)
	return src, nil
}

//解密后去除填充数据
func unPaddingNumber(src []byte) ([]byte, error) {
	lastChar := src[len(src)-1]
	num := int(lastChar)
	return src[:len(src)-num], nil
}

func EncryptId(id int64) (string, error) {
	src := strconv.FormatInt(id, 10) //明文
	key := "lovelove"                //秘钥
	data, err := DESEncrypt([]byte(src), key)
	if err != nil {
		fmt.Println("err:", err)
		return "", err
	}
	res := fmt.Sprintf("%x", data)
	return res, nil
}

func DecryptId(s string) (int64, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		fmt.Println("err:", err)
		return 0, err
	}
	key := "lovelove" //秘钥
	//调用解密函数
	res, err := DESDecrypt(b, key)
	if err != nil {
		fmt.Println("err:", err)
		return 0, err
	}
	i, err := strconv.ParseInt(string(res), 10, 64)
	if err != nil {
		fmt.Println("err:", err)
		return 0, err
	}
	return i, nil
}

func BatchDES(lens int, getInt func(i int) *int64, retStr func(i int) *string) error {
	for i := 0; i < lens; i++ {
		id := getInt(i)
		if id == nil {
			log.Error("nil pointer error")
			return errors.New("nil pointer error")
		}
		encryptedId, err := EncryptId(*id)
		if err != nil {
			return err
		}
		str := retStr(i)
		if str == nil {
			log.Error("nil pointer error")
			return errors.New("nil pointer error")
		}
		*str = encryptedId
		*id = 0
	}
	return nil
}
