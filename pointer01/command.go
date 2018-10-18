package main

import (
	"fmt"
	"go-blackchain/pointer01/demo2"
)

func main()  {
	var ptr *int
	var a int = 20
	ptr = &a
	fmt.Printf("ptr 的地址为 : %x\n", ptr)
	if (ptr != nil) {
		fmt.Printf("ptr 的地址为 : %d\n", *ptr)
	}
	ret := demo2.Say()
	fmt.Println(ret)
}