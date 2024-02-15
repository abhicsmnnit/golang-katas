package main

import "fmt"

func main() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	e1 := Employee{"Abhinav", "Verma", 1}

	e2 := Employee{firstName: "Abhinav", lastName: "Verma", id: 2}

	var e3 Employee
	e3.firstName = "Abhinav"
	e3.lastName = "Verma"
	e3.id = 3

	fmt.Println(e1, e2, e3)
}
