package main

import (
	"fmt"
)

func main() {
	NinjaLevel6_exercise1(foo, bar)
	NinjaLevel6_exercise2()
	NinjaLevel6_exercise3()
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
