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

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

type Score int
type HighScore Score

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

	// var c Counter
	// doUpdateWrong(c)
	// fmt.Println("in main:", c.String())
	// doUpdateRight(&c)
	// fmt.Println("in main:", c.String())

	// a := Adder{start: 10}
	// f1 := a.AddTo // Method value
	// fmt.Println(f1(5))
	// f2 := Adder.AddTo // Method expression
	// fmt.Println(f2(a, 10))

	var i int = 100
	var s Score = 200
	var h HighScore = 300
	fmt.Println(i, s, h)
	// s = i // compiler error
	// h = s // compiler error
	s = Score(i)
	h = HighScore(i)
	h = HighScore(s)
	fmt.Println(i, s, h)
	newScore := s + 500 // + operator on Score
	fmt.Println(newScore)
	printScore(10) // Implicit conversion: 10 is a Score literal
	printScore(Score(10))
	printHighScore(10) // Implicit conversion: 10 is a HighScore literal
	// printHighScore(Score(10)) // compiler error
	printHighScore(HighScore(10))
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

func printScore(s Score) {
	fmt.Println(s)
}

func printHighScore(h HighScore) {
	fmt.Println(h)
}
