package main

import "fmt"

func main() {

	// // Shadowing variables
	// x := 10
	// if x > 5 {
	// 	fmt.Println(x)
	// 	x := 5
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)

	// fmt.Println(10)
	// fmt := "oops"
	// fmt.Println(fmt) // Compile-time error

	// fmt.Println(true)
	// true := 10
	// fmt.Println(true)

	// n := rand.Intn(10)
	// if n == 0 {
	// 	fmt.Println("That's too low")
	// } else if n > 5 {
	// 	fmt.Println("That's too big:", n)
	// } else {
	// 	fmt.Println("That's a good number:", n)
	// }

	// if n := rand.Intn(10); n == 0 {
	// 	fmt.Println("That's too low")
	// } else if n > 5 {
	// 	fmt.Println("That's too big:", n)
	// } else {
	// 	fmt.Println("That's a good number:", n)
	// }
	// // fmt.Println(n) // n is undefined here

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// i := 1
	// for i < 100 { // Acts like the while loop
	// 	fmt.Println(i)
	// 	i = i * 2
	// }

	// // Acts like the do-while loop
	// i := 100
	// for {
	// 	fmt.Println(i)
	// 	if i > 10 {
	// 		break
	// 	}
	// }

	// evenInts := []int{0, 2, 4, 6, 8, 10}
	// for i, v := range evenInts {
	// 	fmt.Println(i, v)
	// }

	// evenInts := []int{0, 2, 4, 6, 8, 10}
	// for _, v := range evenInts {
	// 	fmt.Println(v)
	// }

	// m := map[string]int{"a": 1, "b": 2}
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }
	// for _, v := range m {
	// 	fmt.Println(v)
	// }
	// for k := range m {
	// 	fmt.Println(k)
	// }

	// m := map[string]int{
	// 	"a": 1,
	// 	"c": 3,
	// 	"b": 2,
	// }
	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Loop", i)
	// 	for k, v := range m {
	// 		fmt.Println(k, v)
	// 	}
	// }

	// samples := []string{"hello", "apple_π!"}
	// for _, sample := range samples {
	// 	for i, r := range sample { // Iterates over runes, not bytes!
	// 		fmt.Println(i, r, string(r))
	// 	}
	// 	fmt.Println()
	// }

	// // Modifying the value doesn’t modify the source
	// evenVals := []int{1, 2, 3, 4, 5}
	// for _, v := range evenVals {
	// 	v *= 2 // Source value not modified
	// }
	// fmt.Println(evenVals)

	// 	samples := []string{"hello", "apple_π!"}
	// outer:
	// 	for _, sample := range samples {
	// 	inner:
	// 		for i, r := range sample {
	// 			if r == 'e' {
	// 				continue inner
	// 			}
	// 			fmt.Println(i, r, string(r))
	// 			if r == 'l' {
	// 				continue outer
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}

	// words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	// for _, word := range words {
	// 	switch size := len(word); size {
	// 	case 1, 2, 3, 4:
	// 		fmt.Println(word, "is a short word!")
	// 	case 5:
	// 		wordLen := len(word)
	// 		fmt.Println(word, "is exactly the right length:", wordLen)
	// 	case 6, 7, 8, 9:
	// 	default:
	// 		fmt.Println(word, "is a long word!")
	// 	}
	// }

	// words := []string{"hi", "salutations", "hello"}
	// for _, word := range words {
	// 	switch wordLen := len(word); { // Blank switch
	// 	case wordLen < 5:
	// 		fmt.Println(word, "is a short word!")
	// 	case wordLen > 10:
	// 		fmt.Println(word, "is a long word!")
	// 	default:
	// 		fmt.Println(word, "is exactly the right length.")
	// 	}
	// }

	// FizzBuzz is so clean with switch in golang
	for i := 1; i <= 20; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
