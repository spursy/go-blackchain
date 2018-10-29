package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main()  {
	var mailbox uint8
	var lock sync.Mutex  

	secondCond := sync.NewCond(&lock)
	recvCond := sync.NewCond(&lock)

	send := func (id, index int)  {
		lock.Lock()
		for mailbox == 1 {
			secondCond.Wait()
		}
		log.Printf("sender [%d-%d]: the mailbox is empty.", id, index)
		mailbox = 1
		log.Printf("sender [%d-%d]: the letter has been sent.", id, index)
		lock.Unlock()
		recvCond.Broadcast()
	}

	recv := func (id, index int)  {
		lock.Lock()
		for mailbox == 0 {
			recvCond.Wait()
		}
		mailbox = 0
		log.Printf("receiver [%d-%d]; the mailbox is full.", id, index)
		lock.Unlock()
		secondCond.Signal()
	}

	sign := make(chan struct{}, 3)
	max := 6
	go func (id, max int)  {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= max; i ++ {
			time.Sleep(time.Microsecond * 500)
			send(id, i)
		}
	}(0, max)
	go func (id, max int)  {
		defer func() {
			sign <- struct{}{}
		}()
		for j := 1; j <= max; j ++ {
			time.Sleep(time.Millisecond * 200)
			recv(id, j)
		}
	}(1, max/2)
	go func (id, max int)  {
		defer func ()  {
			sign <- struct{}{}	
		}()
		for k := 1; k <= max; k ++ {
			time.Sleep(time.Microsecond * 200)
			recv(id, k)
		}
	}(2, max/2)

	<- sign
	fmt.Println(111)
	<- sign
	fmt.Println(222)
	<- sign
	fmt.Println(333)
}