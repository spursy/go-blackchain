package main

import (
	"fmt"
	"errors"
)

func main()  {
	fmt.Println("Enter function main")
	defer func ()  {
		fmt.Println("Enter defer func call")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Printf("Exit defer function")
	}()
	panic(errors.New("Something wrong"))
	fmt.Println("Exit function main")
}