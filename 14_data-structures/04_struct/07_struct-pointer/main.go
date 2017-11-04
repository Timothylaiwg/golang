package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	//create a pointer
	p1 := &person{"James", 20}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	//Go gets value from address
	fmt.Println(p1.name)
	fmt.Println(p1.age)
}
