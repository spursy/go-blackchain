package main

import (
	"fmt"
)

var p *int

func foo() (*int, error)  {
	var i int = 5
	return &i, nil
}

func bar()  {
	fmt.Printf("%p, %T\n", p, p)
	fmt.Println(*p)
}

func main()  {
	// p, err := foo()
	var err error
	p, err = foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%p, %T\n", p, p)
	bar()
	
	fmt.Println(*p)
}



