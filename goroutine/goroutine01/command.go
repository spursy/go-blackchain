package main

import (
	"fmt"
	// "time"
	// "sync/atomic"
)

func main()  {
	num := uint32(10)
	sign := make(chan struct{})

	for i := uint32(0); i < num; i ++ {
		go func(i uint32)  {
			fmt.Println(i)
			sign <- struct{}{}
		}(i)
	}
	// time.Sleep(time.Millisecond * 500)
	for j := uint32(0); j < num; j ++ {
		<-sign
	}
}

