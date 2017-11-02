package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 42
	fmt.Println(a)

	// we can pass a memory address instead of a bunch of values
	// we can change the value of whatever is stored at that memory
	// this makes out programs more performant
	// we don't have to pass around large amounts of data
}
