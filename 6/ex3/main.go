package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var people []Person
	// people := make([]Person, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		people = append(people, Person{"Abhinav", "Verma", 30})
	}
}
