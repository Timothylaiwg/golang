package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Hi Tim, or Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Both your names start with M")
	case "Julian":
		fmt.Println("Hi Julian")
	default:
		fmt.Println("Have you no friends?")
	}
}
