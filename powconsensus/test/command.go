package main

import (
	"fmt"
)

func SayHello(data string) {
	fmt.Println(data)
}

func Say(data *string)  {
	SayHello(*data)
}

func main()  {
	value := "cool"
	Say(&value)
}