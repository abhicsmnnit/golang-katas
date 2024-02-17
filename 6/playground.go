package main

func main() {
	// x := 10
	// pointerToX := &x
	// fmt.Println(pointerToX)  // prints a memory address
	// fmt.Println(*pointerToX) // prints 10
	// z := 5 + *pointerToX
	// fmt.Println(z) // prints 15

	// var x *int
	// fmt.Println(x == nil) // prints true
	// fmt.Println(x)
	// fmt.Println(*x) // panics

	// // The built-in function new creates a pointer variable. It returns a pointer to a zero-value instance of the provided type:
	// var x = new(int)
	// fmt.Println(x == nil) // prints false
	// fmt.Println(x)
	// fmt.Println(*x) // prints 0

	/*
		type person struct {
			firstName  string
			middleName *string
			lastName   string
		}

		p := person{
			firstName:  "Narendra",
			middleName: makePtr("Damodardas"),
			lastName:   "Modi",
		}

		fmt.Println(p)
		fmt.Println(p.firstName, *p.middleName, p.lastName)
	*/

}

func makePtr[T any](t T) *T {
	return &t
}
