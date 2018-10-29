package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main()  {
	forAndCAS1()
	fmt.Println()
}

func forAndCAS1()  {
	sign := make(chan struct{}, 2)
	num := int32(0)
	fmt.Printf("The number: %d\n", num)
	go func ()  {
		defer func() {
			sign <- struct{}{}
		}()
		for {
			time.Sleep(time.Millisecond * 500)
			newNum := atomic.AddInt32(&num, 2)
			fmt.Printf("The number: %d\n", newNum)
			if newNum == 10 {
				break
			}
		}
	}()
	go func ()  {
		defer func ()  {
			sign <- struct{}{}
		}()
		for {
			if atomic.CompareAndSwapInt32(&num, 10, 0) {
				fmt.Println("The number has gone to zero.")
				break
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	<- sign
	<- sign
}