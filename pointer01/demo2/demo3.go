package demo2

import(
	"fmt"
)

func Say() string {
	var ptr *int
	var a int = 20
	ptr = &a
	if (ptr != nil) {
		fmt.Printf("值是%d\n", *ptr)
	}
	return "demo2"
}