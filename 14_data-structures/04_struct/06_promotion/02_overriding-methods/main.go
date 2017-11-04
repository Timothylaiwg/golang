package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("I'm just a normal person.")
}

func (dz DoubleZero) Greeting() {
	fmt.Println("Miss Elliot, so good to see you.")
}

func main() {
	p1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   20,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Missy",
			Last:  "Elliot",
			Age:   42,
		},
		LicenseToKill: false,
	}

	p1.Greeting()
	p2.Greeting()
	p2.Person.Greeting()
}
