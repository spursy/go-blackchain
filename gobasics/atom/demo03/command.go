package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main()  {
	forAndCAS2()
	fmt.Println()
}

func forAndCAS2()  {
	sign := make (chan struct{}, 2)
	num := int32(0)
	fmt.Printf("The number: %d\n", num)
	max := int32(20)

	go func (id int, max int32)  {
		defer func ()  {
			sign <- struct{}{}
		}()
		for i := 0; ; i ++ {
			currNum := atomic.LoadInt32(&num)
			if currNum >= max {
				break
			}
			newNum := currNum + 2
			time.Sleep(time.Millisecond * 200)
			if  atomic.CompareAndSwapInt32(&num, currNum, newNum) {
				fmt.Printf("The number: %d [%d-%d]\n", newNum, id, i)
			} else {
				fmt.Printf("The CAS operation failed. [%d-%d]\n", id, i)
			}
		}
	}(1, max)

	go func (id int, max int32)  {
		defer func ()  {
			sign <- struct{}{}
		}()
		for j := 0; ; j ++ {
			currNum := atomic.LoadInt32(&num)
			if currNum >= max {
				break
			}
			newNum := currNum + 2
			time.Sleep(time.Millisecond * 200)
			if atomic.CompareAndSwapInt32(&num, currNum, newNum) {
				fmt.Printf("The number: %d [%d-%d]\n", newNum, id, j)
			} else {
				fmt.Printf("The CAS operation failed. [%d-%d]\n", id, j)
			}
		}
	}(2, max)

	<- sign
	<- sign
}