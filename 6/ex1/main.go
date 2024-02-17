package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p1 := MakePerson("Abhinav", "Verma", 30)
	fmt.Println(p1)
	// Surprisingly, p1 escapes to the heap.
	// The reason for this is that it is passed into fmt.Println. This is because the parameter to fmt.Println are ...any.
	// The current Go compiler moves to the heap any value that is passed in to a function via a parameter that is of an interface type.

	p2 := MakePersonPointer("Abhinav", "Verma", 30)
	fmt.Println(p2)
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}
