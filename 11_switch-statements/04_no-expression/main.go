package main

import "fmt"

func main() {

	myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Hi my friend with name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Hey Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Hey Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("LOL")
	case myFriendsName == "Julian":
		fmt.Println("Hey Julian")
	default:
		fmt.Println("nothing matched")
	}
}
