package main

import "fmt"

func main() {
	records := make([][]string, 0)
	//student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.0"
	student1[3] = "74.0"
	//store the record
	records = append(records, student1)
	//student 2
	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.0"
	student2[3] = "96.0"
	//store the record
	records = append(records, student2)
	//print
	fmt.Println(records)
}
