package main

import "fmt"

var x int

func increment() int {
	x++
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

//x is in the package scope
