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

	twoMult := makeMult(2)
	threeMult := makeMult(3)
	for i := 0; i < 5; i++ {
		fmt.Println(twoMult(i), threeMult(i))
	}
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
