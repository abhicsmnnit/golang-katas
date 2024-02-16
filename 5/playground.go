package main

import (
	"errors"
	"fmt"
	"os"
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

	result, remainer, error := divAndRemainderNamedReturnVars(10, 3)
	if error != nil {
		fmt.Println(error)
		os.Exit(1)
	}
	fmt.Println(result, remainer)
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
