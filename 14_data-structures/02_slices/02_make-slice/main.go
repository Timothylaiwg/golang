package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println("----------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("----------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "Capacity: ", cap(mySlice), "Value: ", mySlice[i])
	}
	//golang doubles the size of the slice everytime the slice exceeds the Capacity
	//but it is good to set capacity to increase performance
	//if you try to add an index >range, golang does not allow it
	//must use append method
}
