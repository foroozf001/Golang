package main

import "fmt"

func main() {
	NinjaLevel1_exercise1()
	NinjaLevel1_exercise2()
	NinjaLevel1_exercise3()
	NinjaLevel1_exercise4()
}

func NinjaLevel1_exercise1() {
	println("Exercise 1:")
	// Short-hand operator with basic types (primitives)
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("Value x: %d,\tvalue y: %s,\tvalue z: %v\n", x, y, z)
	fmt.Printf("Value x: %d is type: %T\n", x, x)
	fmt.Printf("Value y: %s is type: %T\n", y, y)
	fmt.Printf("Value z: %v is type: %T\n", z, z)
	println()
}

var x int
var y string
var z bool

func NinjaLevel1_exercise2() {
	println("Exercise 2:")
	// Print zero values for different basic types
	fmt.Printf("Value x: %d,\tvalue y: %s,\tvalue z: %v\n", x, y, z)
	fmt.Printf("Value x: %d is type: %T\n", x, x)
	fmt.Printf("Value y: %s is type: %T\n", y, y)
	fmt.Printf("Value z: %v is type: %T\n", z, z)
	println()
}

var i int = 42
var j string = "James Bond"
var k bool = true

func NinjaLevel1_exercise3() {
	println("Exercise 3:")
	// Assign return value to variable
	s := fmt.Sprintf("Value i: %d,\tvalue j: %s,\tvalue k: %v\n", i, j, k)
	println(s)
}

type hotdog int

var u hotdog
var v int

func NinjaLevel1_exercise4() {
	println("Exercise 4:")
	// Assign return value to variable
	fmt.Printf("Value u: %v,\t type u: %T\n", u, u)

	u = 42

	fmt.Printf("Value u: %v,\t type u: %T\n", u, u)

	// Type conversion
	v = int(u)

	fmt.Printf("Value u: %v,\t type u: %T\n", v, v)
}
