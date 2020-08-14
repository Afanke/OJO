package ctrl

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
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

//输入明文，得到密文
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

//输入密文，得到明文
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
		return nil, errors.New("src长度不能小于0")
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
	lastChar := src[len(src)-1] //byte(3)
	num := int(lastChar)        //int(3)
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

func BatchEncrypt(lens int, getInt func(i int) *int64, retStr func(i int) *string) error {
	for i := 0; i < lens; i++ {
		encryptedId, err := EncryptId(*getInt(i))
		if err != nil {
			return err
		}
		*retStr(i) = encryptedId
		*getInt(i) = 0
	}
	return nil
}
