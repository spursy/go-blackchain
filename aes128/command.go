package main
import (
	"crypto/cipher"
	"crypto/aes"
	"bytes"
	"fmt"
)

func main() {
	key := []byte("1234567890123456")
	origData := []byte("hello world")
	en := AESEncrypt(origData, key)
	de := AESDecrypt(en, key)
	fmt.Println(string(de))
}

func AESDecrypt(crypted, key []byte) []byte{
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData
}

func PKCS7UnPadding(origData []byte) []byte{
	length := len(origData)
	unpadding := int(origData[length - 1])
	return origData[:length - unpadding]
}

func AESEncrypt(origData, key []byte) []byte {
	block,_ := aes.NewCipher(key)
	origData = PKCS7Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])

	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted
}

func PKCS7Padding(origData []byte,blockSize int)[]byte {
	padding := blockSize-len(origData)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)},padding)
	return append(origData,padtext...)
}

