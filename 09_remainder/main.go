package main

import "fmt"

func main() {

	x := 13 / 3
	y := 13 % 3
	fmt.Println(x)
	fmt.Println(y)

	if y == 1 {
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

}
