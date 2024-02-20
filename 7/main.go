package main

import (
	"fmt"
	"io"
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

// ////////////////////////
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

type Incrementer interface {
	Increment()
}

// ////////////////////////
type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

// ////////////////////////
type Score int
type HighScore Score

// ////////////////////////
type MailCategory int

const (
	Uncategorized MailCategory = iota
	Personal
	Social
	Advertisements
	Spam
)

// ////////////////////////
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

// ////////////////////////
type Inner struct {
	X int
}

type Outer struct {
	Inner
	X int
}

// ////////////////////////
type Reader interface {
	Read(p []byte) (count int, err error)
}

type Closer interface {
	Close() error
}

type ReadCloser interface {
	Reader
	Closer
}

// ////////////////////////
type Doubler interface {
	Double()
}

type DoubleInt int

func (di *DoubleInt) Double() {
	*di = *di * 2
}

type DoubleIntSlice []int

func (dis DoubleIntSlice) Double() {
	for i := range dis {
		dis[i] = 2 * dis[i]
	}
}

func DoubleCompare(d1, d2 Doubler) bool {
	return d1 == d2
}

// ////////////////////////
type MyInt int

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

	// // Embedding is NOT inheritance!
	// m := Manager{}
	// // var e1 Employee = m // compiler error: Manager "is not" an Employee
	// var e2 Employee = m.Employee // fine: Manager "has" an Empoyee
	// fmt.Println(e2)

	// var incrementer Incrementer
	// var counterPtr *Counter
	// fmt.Println(incrementer == nil) // true
	// fmt.Println(counterPtr == nil)  // true
	// incrementer = counterPtr
	// fmt.Println(incrementer == nil) // false. Why? Short answer: incrementer is now pointing to a concrete type (which is nil as well). But since it needs to know what type it is pointing to, it can't stay nil. An instance of an interface is nil when both its type and value are nil. For more details, read book.

	// var di1 DoubleInt = 10
	// var di2 DoubleInt = 10
	// var dis1 DoubleIntSlice = []int{1, 2, 3}
	// var dis2 DoubleIntSlice = []int{1, 2, 3}
	// fmt.Println(DoubleCompare(&di1, &di2))
	// fmt.Println(DoubleCompare(&di1, dis1))
	// fmt.Println(DoubleCompare(dis1, dis2)) // panic: comparing incomparable types

	// var i interface{} // `i` can hold values of any type now
	// i = 10
	// i = "Hello"
	// i = struct {
	// 	FirstName string
	// 	LastName  string
	// }{"Abhinav", "Verma"}

	// var i2 any // `i2` can hold values of any type as well
	// i2 = 10
	// i2 = "Hello"
	// i2 = struct {
	// 	FirstName string
	// 	LastName  string
	// }{"Abhinav", "Verma"}
	// fmt.Println(i, i2)

	// var i any
	// var myInt MyInt = 10
	// i = myInt
	// fmt.Println(5 + i.(MyInt)) // type assertion in action
	// // fmt.Println("5" + i.(string)) // panic: interface conversion: interface {} is main.MyInt, not string
	// // fmt.Println(5 + i.(int))      // This panics as well even though MyInt is an int! panic: interface conversion: interface {} is main.MyInt, not int
	// i2, ok := i.(int) // comma-ok idiom
	// if !ok {
	// 	fmt.Println("Fishy type assertion")
	// 	return
	// }
	// fmt.Println(i2 + 1) // Never reached

	var i any
	doThings(5)
	doThings("5")
	doThings(i)
	i = MyInt(5)
	doThings(i)
	doThings(MyInt(10))
	doThings(false)
	doThings([]int{})
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

func doThings(i any) {
	// Type switch
	switch j := i.(type) {
	case nil:
		// i is nil, type of j is any
		fmt.Println("nil", j)
	case int:
		// j is of type int
		j++
		fmt.Println("int", j)
	case MyInt:
		// j is of type MyInt
		j = j + 10
		fmt.Println("MyInt", j)
	case io.Reader:
		// j is of type io.Reader
		j.Read(nil)
		fmt.Println("io.Reader", j)
	case string:
		// j is a string
		j = j + "string"
		fmt.Println("string", j)
	case bool, rune:
		// i is either a bool or rune, so j is of type any
		fmt.Println("bool, or rune", j)
	default:
		// no idea what i is, so j is of type any
		fmt.Println("Nothing matched", j)
	}
}
