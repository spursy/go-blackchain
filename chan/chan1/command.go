package main

import (
	"fmt"
)

func main()  {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	elem1 := <- ch1
	fmt.Printf("the value is %v\n", elem1)
}