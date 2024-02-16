package main

import (
	"errors"
	"fmt"
)

func main() {
	// fmt.Println(addTo(3))
	// fmt.Println(addTo(3, 5))
	// fmt.Println(addTo(3, 5, 10, 12, 15))
	// vals := []int{5, 10, 12, 15}
	// fmt.Println(addTo(3, vals...))
	// fmt.Println(addTo(3, []int{5, 10, 12, 15}...))

	// fmt.Println(divAndRemainder(10, 3))
	// fmt.Println(divAndRemainder(10, 5))
	// fmt.Println(divAndRemainder(10, 0))
	// fmt.Println(divAndRemainder(0, 0))
	// result, remainer, error := divAndRemainder(10, 3)
	// if error != nil {
	// 	fmt.Println(error)
	// 	os.Exit(1)
	// }
	// fmt.Println(result, remainer)

	// result, _, error := divAndRemainder(10, 3)
	// if error != nil {
	// 	fmt.Println(error)
	// 	os.Exit(1)
	// }
	// fmt.Println(result)

	// result, remainer, error := divAndRemainderNamedReturnVars(10, 3)
	// if error != nil {
	// 	fmt.Println(error)
	// 	os.Exit(1)
	// }
	// fmt.Println(result, remainer)

	// result, remainer, error := divAndRemainderNamedReturnVars(10, 3)
	// if error != nil {
	// 	fmt.Println(error)
	// 	os.Exit(1)
	// }
	// fmt.Println(result, remainer)

	/*
		type Person struct {
			FirstName string
			LastName  string
			Age       int
		}

		people := []Person{
			{"Abhinav", "Verma", 30},
			{"Prachi", "Jain", 31},
			{"Siddhanta", "Das", 35},
		}

		fmt.Println(people)

		sort.Slice(people, func(i, j int) bool {
			return people[i].LastName < people[j].LastName
		})
		fmt.Println(people)

		sort.Slice(people, func(i, j int) bool {
			return people[i].Age < people[j].Age
		})
		fmt.Println(people)
	*/

	// twoMult := makeMult(2)
	// threeMult := makeMult(3)
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(twoMult(i), threeMult(i))
	// }

	// deferExample()

	// goIsCallByValue()

	testModForMapsAndSlices()
}

// Variadic Input Parameters
func addTo(base int, vals ...int) []int {
	result := make([]int, 0, len(vals))
	for _, v := range vals {
		result = append(result, base+v)
	}
	return result
}

func divAndRemainder(num int, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("can't divide by zero")
	}

	return num / denom, num % denom, nil
}

func divAndRemainderNamedReturnVars(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func divAndRemainderBlankReturns(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return
	}
	result, remainder = num/denom, num%denom
	return
}

func makeMult(base int) func(int) int {
	return func(num int) int {
		return base * num
	}
}

// Order of defers and param evaluation logic of defer functions
func deferExample() int {
	a := 10
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)
	a = 20
	defer func(val int) {
		fmt.Println("second:", val)
	}(a)
	a = 30
	fmt.Println("exiting:", a)
	return a
}

type person struct {
	name string
}

func goIsCallByValue() {
	i := 10
	s := "hello"
	p := person{}

	fmt.Println(i, s, p)
	modFails(i, s, p)
	fmt.Println(i, s, p)
}

func modFails(i int, s string, p person) {
	i = i * 2
	s = s + "modified"
	p.name = "Abhinav"
}

func testModForMapsAndSlices() {
	m := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
	}

	s := []int{0, 1, 2, 3, 4}

	fmt.Println(m, s)
	mod(m, s)
	fmt.Println(m, s)
}

func mod(m map[string]int, s []int) {
	m["a"] = 100
	m["d"] = 1000
	delete(m, "b")

	for i, v := range s {
		s[i] = 2 * v
	}
	s = append(s, 100)
}
