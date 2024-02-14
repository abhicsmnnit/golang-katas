package main

import "fmt"

func main() {
	// var x = []int{1, 2, 3, 4}
	// var y = []int{1, 2, 3, 4}
	// var z = []int{1, 2, 3, 4, 5}

	// var a []int

	// fmt.Println(slices.Equal(x, y))
	// fmt.Println(slices.Equal(x, z))
	// slices.Reverse(z)
	// fmt.Println(z)

	// fmt.Println(len(z))

	// fmt.Println(len(a), z[1])

	// x = append(x, 5)

	// fmt.Println(x)

	// slices.Reverse(z)

	// fmt.Println(slices.Equal(x, z))

	// fmt.Println(append(x, y...))

	// fmt.Println(cap(z))
	// z = append(z, 6)
	// fmt.Println(cap(z))

	// b := make([]int, 10)
	// b = append(b, 10)
	// fmt.Println("b: ", b, len(b), cap(b))
	// clear(b)
	// fmt.Println("b: ", b, len(b), cap(b))

	// c := make([]string, 10)
	// c = append(c, "10")
	// fmt.Println("c: ", c, len(c), cap(c))
	// clear(c)
	// fmt.Println("c: ", c, len(c), cap(c))

	// x := []string{"a", "b", "c", "d"}
	// y := x[:2]
	// z := x[2:]
	// fmt.Println(cap(x), cap(y), cap(z))
	// y = append(y, "z")
	// fmt.Println("x:", x, len(x), cap(x))
	// fmt.Println("y:", y, len(y), cap(y))
	// fmt.Println("z:", z, len(z), cap(z))

	// [a, b, i, j, y]
	// x := make([]string, 0, 5)
	// x = append(x, "a", "b", "c", "d")
	// y := x[:2]
	// z := x[2:]
	// fmt.Println(cap(x), cap(y), cap(z))
	// y = append(y, "i", "j", "k")
	// x = append(x, "x")
	// z = append(z, "y")
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// x := make([]string, 0, 5)
	// x = append(x, "a", "b", "c", "d")
	// y := x[:2:2]
	// z := x[2:4:4]
	// fmt.Println(cap(x), cap(y), cap(z))
	// x[0] = "0"
	// y = append(y, "i", "j", "k")
	// x = append(x, "x")
	// x[2] = "2"
	// z = append(z, "y")
	// x[3] = "3"
	// x[1] = "1"
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	x := [4]int{5, 6, 7, 8}
	y := x[:2]
	z := x[2:]
	x[0] = 10
	y = append(y, 101, 1001, 10001)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

}
