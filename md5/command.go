package main

import (
	"fmt"
	"crypto/md5"
	"encoding/hex"
)

func main()  {
	MyMD5()
}

func MyMD5()  {
	// 方法一
	data := []byte("hello world")
	// fmt.Printf("%v", data)
	s := fmt.Sprintf("%x\n", md5.Sum(data))
	fmt.Printf(s)

	// 方法二
	m := md5.New()
	m.Write([]byte("hello world"))
	s = hex.EncodeToString(m.Sum(nil))
	fmt.Println(s)
}