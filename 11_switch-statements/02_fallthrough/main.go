package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Hi Tim")
	case "Jenny":
		fmt.Println("Hi Jenny")
	case "Marcus":
		fmt.Println("Hi Marcus")
		fallthrough
	case "Daniel":
		fmt.Println("Hi Daniel")
		fallthrough
	case "Medhi":
		fmt.Println("Hi Medhi")
	case "Julian":
		fmt.Println("Hi Julian")
	default:
		fmt.Println("Have you no friends?")
	}
}
