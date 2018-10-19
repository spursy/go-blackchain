package main

import (
	"fmt"
	"time"
	"sync/atomic"
)

func main()  {
	num := uint32(10)
	var count uint32

	trigger := func(i uint32, fn func())  {
		for {
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				atomic.AddUint32(&count, 1)
				break
			}
			time.Sleep(time.Nanosecond)
		}
	}

	for i := uint32(0); i < num; i ++ {
		go func(i uint32)  {
			// fmt.Println(i)
			// sign <- struct{}{}
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	// time.Sleep(time.Millisecond * 500)
	// for j := 0; j < num; j ++ {
	// 	<-sign
	// }
	

	trigger(10, func() {})
}

