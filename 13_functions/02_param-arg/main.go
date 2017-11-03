package main

import "fmt"

func main() {
	greet("John")
	greet("Jane")
}

func greet(name string) {
	fmt.Println("Hello,", name)
}
