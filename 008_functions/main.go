package main

import (
	"fmt"
	"math"
)

func main() {
	// Ninja Level 6 exercise 1
	f := foo1()
	a, b := bar1()
	fmt.Println(f, a, b)

	// Ninja Level 6 exercise 2
	ninja2 := []int{1, 8, 3}
	c := foo2(ninja2...) //extract values
	fmt.Println(c)

	d := bar2(ninja2)
	fmt.Println(d)

	// Ninja Level 6 exercise 3
	defer foo3() //deferred function
	fmt.Println("This function is placed after foo3")

	// Ninja Level 6 exercise 4
	p1 := person{
		first: "Bob",
		last:  "Ross",
		age:   50,
	}
	p1.speak()
	// Ninja Level 6 exercise 5
	circ := circle{
		radius: 2.0,
	}
	squa := square{
		length: 2.0,
	}
	info(circ)
	info(squa)

	// Ninja Level 6 exercise 6
	func() {
		fmt.Println("This is a an anonymous function")
	}()

	// Ninja Level 6 exercise 7
	fcc := firstClassCitizen
	fmt.Println(fcc())

	// Ninja Level 6 exercise 8
	fret := fret()
	fmt.Println(fret())

	// Ninja Level 6 exercise 9
	mff := func(f func() int) int {
		return f() * -42
	}
	fmt.Println(mff(callb))
}

// Ninja Level 6 exercise 1
func foo1() int {
	return 123
}

func bar1() (int, string) {
	return 456, "bar"
}

// Ninja Level 6 exercise 2
func foo2(sum ...int) int { //input is variadic parameter
	total := 0
	for _, i := range sum {
		total += i
	}
	return total
}

func bar2(sum []int) int { //input is slice of int
	total := 0
	for _, i := range sum {
		total += i
	}
	return total
}

// Ninja Level 6 exercise 3
func foo3() {
	fmt.Println("Foo3 is deferred")
}

// Ninja Level 6 exercise 4
type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hi, my name is", p.first, p.last, "and I'm", p.age, "years old")
}

// Ninja Level 6 exercise 5
type shape interface {
	area() float64
}

type circle struct {
	shape
	radius float64
}

type square struct {
	shape
	length float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

func info(s shape) {
	fmt.Println(s.area())
}

// Ninja Level 6 exercise 7
func firstClassCitizen() string {
	return "This is a function stored in a variable"
}

// Ninja Level 6 exercise 8
func fret() func() int {
	return func() int {
		return 404
	}
}

// Ninja Level 6 exercise 9
func callb() int {
	return -3
}
