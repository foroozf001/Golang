package main

import (
	"fmt"
	"math"
)

func main() {
	// NinjaLevel6_exercise1(foo, bar)
	// NinjaLevel6_exercise2()
	// NinjaLevel6_exercise3()
	// NinjaLevel6_exercise4()
	// NinjaLevel6_exercise5()
	// NinjaLevel6_exercise678()
	// NinjaLevel6_exercise9()
	// NinjaLevel6_exercise10()
	NinjaLevel6_exercise11()
}

// Ninja Level 6 exercise 1
func NinjaLevel6_exercise1(f1 func() int, f2 func() (string, int)) {
	r0 := f1()
	r1, r2 := f2()
	fmt.Println(r0, r1, r2)
}

func foo() int {
	return 7
}

func bar() (string, int) {
	return "bar", 8
}

// Ninja Level 6 exercise 2
func NinjaLevel6_exercise2() {
	var my_slice []int
	my_slice = []int{1, 2, 7, 31}
	fmt.Printf("%d\n", sum(even, my_slice...))
}

func even(in int) int {
	if in%2 == 0 {
		return in
	} else {
		return 0
	}
}

func sum(f func(in int) int, values ...int) int {
	ii := []int{}
	sum := 0
	// Add all even numbers to new slice
	for _, val := range values {
		ii = append(ii, f(val))
	}
	// Sum all values in slice containing even numbers
	for _, val := range ii {
		sum += val
	}
	return sum
}

// Ninja Level 6 exercise 3
func NinjaLevel6_exercise3() {
	defer fmt.Println("Deferred")
	fmt.Println(foo())
}

// Ninja Level 6 exercise 4
func NinjaLevel6_exercise4() {
	person1 := person{
		first: "Bob",
		last:  "Ross",
		age:   42,
	}
	person2 := person{
		first: "Bob",
		last:  "Marley",
		age:   38,
	}
	person1.speak()
	person2.speak()
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hi, my name is %s %s and I'm %d years old.\n", p.first, p.last, p.age)
}

// Ninja Level 6 exercise 5
func NinjaLevel6_exercise5() {
	circle1 := circle{
		id:     0,
		radius: 2.2,
	}
	square1 := square{
		id:     1,
		length: 4.2,
	}
	info(circle1)
	info(square1)
}

type shape interface {
	getid() int
	area() float32
}

type circle struct {
	id     int
	radius float32
}

type square struct {
	id     int
	length float32
}

func (s square) area() float32 {
	return s.length * s.length
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

func (c circle) getid() int {
	return c.id
}

func (s square) getid() int {
	return s.id
}

func info(s shape) {
	fmt.Printf("id: %d, Type: %T, Area: %f\n", s.getid(), s, s.area())
}

// Ninja Level 6 exercise 6,7,8
func NinjaLevel6_exercise678() {
	anonymous := func() func(i int) int {
		return func(i int) int {
			return i
		}
	}()
	fmt.Println(anonymous(-1))
}

// Ninja Level 6 exercise 9
func NinjaLevel6_exercise9() {
	measured_lengths := []float32{1.6, 1.0, 3.2, 0.3, 2.0, 9.7, 5.1, 5.2, 6.0}
	fmt.Println(filterDecimals(filterLengths, measured_lengths, 2))
}

func filterLengths(lengths []float32, threshold float32) []float32 {
	ret := []float32{}
	for _, length := range lengths {
		if length >= threshold {
			ret = append(ret, length)
		}
	}
	return ret
}

func filterDecimals(f func([]float32, float32) []float32, lengths []float32, threshold float32) []float32 {
	ret := []float32{}
	for _, length := range lengths {
		//has decimal?
		if float64(length) == math.Trunc(float64(length)) {
			ret = append(ret, length)
		}
	}
	return f(ret, threshold)
}

// Ninja Level 6 exercise 10
func NinjaLevel6_exercise10() {
	f := re()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func re() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

// Ninja Level 6 exercise 11
func NinjaLevel6_exercise11() {
	fmt.Println(factorial(4))
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
