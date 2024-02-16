package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, v := range expressions {
		ans, err := eval(v)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ans)
		}
	}
}

func eval(expression []string) (int, error) {
	m := map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	if len(expression) != 3 {
		return 0, errors.New("invalid expression: " + fmt.Sprint(expression))
	}

	operator := expression[1]
	if _, ok := m[operator]; !ok {
		return 0, errors.New("invalid operator: " + operator)
	}

	operand1, err := strconv.Atoi(expression[0])
	if err != nil {
		return 0, errors.New("invalid operand: " + expression[0])
	}

	operand2, err := strconv.Atoi(expression[2])
	if err != nil {
		return 0, errors.New("invalid operand: " + expression[2])
	}

	return m[operator](operand1, operand2), nil
}

func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }
func mul(a, b int) int { return a * b }
func div(a, b int) int { return a / b }
