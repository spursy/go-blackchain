package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"sync/atomic"	
)

func main()  {
	// demo01
	var box atomic.Value
	fmt.Println("Copy box to box2.")
	box2 := box
	v1 := [...]int{1, 2, 3}
	fmt.Printf("Store %v to box.\n", v1)
	box.Store(v1)
	fmt.Printf("The value load from box is %v.\n", box.Load())
	fmt.Printf("The value load from box2 is %v.\n", box2.Load())
	fmt.Println()

	// demo02
	v2 := "123"
	fmt.Printf("Store %q to box2.\n", v2)
	box2.Store(v2)
	fmt.Printf("The value load from box is %v.\n", box.Load())
	fmt.Printf("The value load from box is %q.\n", box2.Load())
	fmt.Println()

	// demo03
	fmt.Println("Copy box to box3.")
	box3 := box
	fmt.Printf("The value load from box3 is %v.\n", box3.Load())
	// v3 := 123
	// box3.Store(v3)
	_ = box3
	fmt.Println()

	// demo04
	var box4 atomic.Value
	v4 := errors.New("Something wrong")
	fmt.Printf("Store an error with message %q to box4.\n", v4)	
	box4.Store(v4)
	v41 := io.EOF
	fmt.Println("Strore a value of the same type to box4.")
	box4.Store(v41)
	v42, ok := interface{}(&os.PathError{}).(error)
	if ok {
		fmt.Printf("Store a value of type %T that implements error interface to box4.\n", v42)
		// box4.Store(v42)
	}
	fmt.Println()

	// demo05
	box5, err := NewAtomicValue(v4)
	if err != nil {
		fmt.Printf("errors: %s\n", err)
	}
	fmt.Printf("The legal type in box5 is %s.\n", box5.TypeOfValue())
	fmt.Printf("Store a value of the same type to box5.\n")
	err = box5.Strore(v42)
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Println()

	// demo06
	var box6 atomic.Value
	v6 := []int{1, 2, 3}
	fmt.Printf("Store %v to box6.\n", v6)
	box6.Store(v6)
	v6[1] = 4
	fmt.Printf("The value load from box6 is %v.\n", box6.Load())
	fmt.Println()

	// demo07
	var box7 atomic.Value
	v7 := []int{1,2,3}
	stroe := func (v []int)  {
		replica := make([]int, len(v))
		copy(replica, v)
		box7.Store(replica)
	}
	fmt.Printf("Store %v to box7.\n", v7)
	stroe(v7)
	v7[2] = 5
	fmt.Printf("The value load from box7 is %v.\n", box7.Load())
}

type atomicValue struct {
	v atomic.Value
	t reflect.Type
}

func NewAtomicValue(example interface{}) (*atomicValue, error)  {
	if example == nil {
		return nil, errors.New("atomic value: nil example")
	}
	return &atomicValue{
		t: reflect.TypeOf(example),
	}, nil
}

func (av *atomicValue) Strore(v interface{}) error {
	if v == nil {
		return errors.New("atomic value: nil value")
	}
	t := reflect.TypeOf(v)
	if t != av.t {
		return fmt.Errorf("atomic value: wrong type: %s", t)
	}
	av.v.Store(v)
	return nil
}

func (av *atomicValue) Load() interface{} {
	return av.v.Load()
}

func (av *atomicValue) TypeOfValue() reflect.Type {
	return av.t
}