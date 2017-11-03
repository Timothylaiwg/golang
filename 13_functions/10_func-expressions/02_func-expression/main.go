package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Println("Hello, world")
	}

	greeting()

	//assign a func to a variable
	//only way to assign a function within a function
}
