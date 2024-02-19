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

type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Social
	Advertisements
	Spam
)

type Employee struct {
	Id   int
	Name string
}

func (e Employee) Describe() string {
	return fmt.Sprintf("%s (%d)", e.Name, e.Id)
}

type Manager struct {
	Employee // Embedded field
	Reports  []Employee
}

func (m Manager) getReportsWithName(name string) []Employee {
	// business logic
	return nil
}

type Inner struct {
	X int
}

type Outer struct {
	Inner
	X int
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

	/*
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
	*/

	// var mc MailCategory = Uncategorized
	// fmt.Println(mc)
	// fmt.Println(Uncategorized)
	// fmt.Println(Spam)

	// m := Manager{
	// 	Employee: Employee{
	// 		Id:   1,
	// 		Name: "Dheerubhai Ambani",
	// 	},
	// 	Reports: []Employee{
	// 		{Id: 2, Name: "Mukesh Ambani"},
	// 		{Id: 3, Name: "Anil Ambani"},
	// 	},
	// }
	// fmt.Println(m.Id, m.Name) // Employee fields "promoted" to Manager type
	// fmt.Println(m.Describe()) // Employee methods "promoted" to Manager type
	// fmt.Println(m.getReportsWithName("Dummy Employee"))

	// o := Outer{
	// 	Inner: Inner{X: 10},
	// 	X:     20,
	// }
	// fmt.Println(o.X, o.Inner.X) // Explicitly specify the "conflicting" fields

	// Embedding is NOT inheritance!
	m := Manager{}
	// var e1 Employee = m // compiler error: Manager "is not" an Employee
	var e2 Employee = m.Employee // fine: Manager "has" an Empoyee
	fmt.Println(e2)
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
