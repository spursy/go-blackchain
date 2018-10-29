package main

import (
	"log"
	"sync"
	"time"
)

type counter struct {
	num uint
	mu sync.RWMutex
}

func (c *counter) number() uint {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.num
}

func (c *counter) add(increment uint) uint {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.num += increment
	return c.num
}

func count(c *counter)  {
	sign := make(chan struct{}, 3)
	go func() {
		defer func() {
			sign <- struct{}{}
		}()
		for i := 1; i <= 10; i ++ {
			time.Sleep(time.Millisecond * 500)
			c.add(1)
		}
	}()

	go func ()  {
		defer func ()  {
			sign <- struct{}{}
		}()
		for j := 1; j <= 20; j ++ {
			time.Sleep(time.Millisecond * 200)
			log.Printf("The number in counter: %d [%d-%d]", c.number(), 1, j)
		}
	}()

	go func () {
		defer func () {
			sign <- struct{}{}
		}()
		for k := 1; k <= 20; k ++ {
			time.Sleep(time.Microsecond * 300)
			log.Printf("The number in counter: %d [%d-%d]", c.number(), 2, k)
		}
	}()

	<- sign
	<- sign
	<- sign
}

func redundantUnlock()  {
	var rwMu sync.RWMutex

	// demo 1 引发panic
	// rwMu.Unlock()

	// demo 2 引发panic
	// rwMu.RUnlock()

	// demo 3 
	// rwMu.RLock()
	// rwMu.RUnlock()

	// demo 4
	rwMu.Lock()
	rwMu.Unlock()
}

func main()  {
	c := counter{}
	count(&c)
	redundantUnlock()
}