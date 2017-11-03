package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "Buenos dias!"
	greeting = append(greeting, "Suprebadham!")

	fmt.Println(greeting[3])

	//if append beyond capacity, the capacity will double
}
