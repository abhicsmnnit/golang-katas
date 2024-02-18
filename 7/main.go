package main

import (
	"fmt"
	"time"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p.firstName, p.lastName, p.age)
}

type Counter struct {
	val      int
	updateAt time.Time
}

func (c *Counter) Increment() {
	c.val++
	c.updateAt = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("counter: %d, updated at: %v", c.val, c.updateAt)
}

func main() {
	// p := Person{"Abhinav", "Verma", 30}
	// fmt.Println(p.String())

	// c := Counter{}
	// fmt.Println(c.String())
	// c.Increment() // Go automatically converts to (&c).Increment()
	// fmt.Println(c.String())

	// c1 := &Counter{}
	// fmt.Println(c1.String()) // Go automatically converts to (*c1).String()
	// c1.Increment()
	// fmt.Println(c1.String()) // Go automatically converts to (*c1).String()

	// var c2 *Counter
	// fmt.Println(c2.String()) // panic: nil pointer dereference
	// c2.Increment()           // panic: nil pointer dereference
	// fmt.Println(c2.String()) // panic: nil pointer dereference

	var c Counter
	doUpdateWrong(c)
	fmt.Println("in main:", c.String())
	doUpdateRight(&c)
	fmt.Println("in main:", c.String())
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}
