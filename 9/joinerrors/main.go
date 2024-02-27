package main

import (
	"errors"
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func ValidatePerson(p Person) error {
	var errs []error

	if len(p.firstName) < 2 {
		errs = append(errs, errors.New("field firstName must have at least 2 characters"))
	}

	if len(p.lastName) < 2 {
		errs = append(errs, errors.New("field lastName must have at least 2 characters"))
	}

	if p.age < 0 {
		errs = append(errs, errors.New("field age must be positive"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func main() {
	p := Person{
		firstName: "B",
		lastName:  "Baggins",
		age:       111,
	}

	err := ValidatePerson(p)

	if err != nil {
		fmt.Println(err)
		fmt.Println("--")
	}

	p1 := Person{
		firstName: "B",
		lastName:  "B",
		age:       111,
	}

	err = ValidatePerson(p1)

	if err != nil {
		fmt.Println(err)
		fmt.Println("--")
	}

	p2 := Person{
		firstName: "B",
		lastName:  "B",
		age:       -1,
	}

	err = ValidatePerson(p2)

	if err != nil {
		fmt.Println(err)
		fmt.Println("--")
	}
}
