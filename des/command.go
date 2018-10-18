package main

import (
	"crypto/des"
	"bytes"
	"crypto/cipher"
	"fmt"
	"encoding/base64"
)

func MyDESEncrypt(origData, key []byte) {
	block, _ := des.NewCipher(key)
	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([] byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	fmt.Println(base64.StdEncoding.EncodeToString(crypted))
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length - 1])
	return origData[:(length - unpadding)]
}

func MyDESDecrypt(data string, key []byte)  {
	crypted, _ := base64.StdEncoding.DecodeString(data)
	block, _ := des.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	fmt.Println(string(origData))
}

func main() {
	data := []byte("hello workld")
	key := []byte("12345678")

	MyDESEncrypt(data, key)
	MyDESDecrypt("CyqS6B+0nOGkMmaqyup7gQ==",key)
}