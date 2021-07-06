package main

import (
	"fmt"
)

func main() {
	NinjaLevel2_exercise1()
	NinjaLevel2_exercise2()
	NinjaLevel2_exercise3()
	NinjaLevel2_exercise4()
	NinjaLevel2_exercise5()
	NinjaLevel2_exercise6()
}

func NinjaLevel2_exercise1() {
	println("Exercise 1:")
	i := 24
	fmt.Printf("Decimal: %d\t binary: %b\t hexadecimal: %#x\n\n", i, i, i)
}

func NinjaLevel2_exercise2() {
	println("Exercise 2:")
	g := 24 == 32
	h := 24 <= 32
	i := 24 >= 32
	j := 24 != 32
	k := 24 > 32
	l := 23 < 32

	fmt.Printf("g: %t, h: %t, i: %t, j: %t, k: %t, l: %t\n\n", g, h, i, j, k, l)
}

func NinjaLevel2_exercise3() {
	println("Exercise 3:")
	// Typed and untyped constants
	const (
		a     = 1
		b int = 2
	)
	fmt.Printf("a: %d\tb: %d\n\n", a, b)
}

func NinjaLevel2_exercise4() {
	println("Exercise 4:")
	// Bit shift
	var v int8 = 19
	var w int8 = v << 1
	var x int8 = w >> 2
	fmt.Printf("Decimal: %d\t binary: %08b\t hexadecimal: %#x\n\n", v, v, v)
	fmt.Printf("Decimal: %d\t binary: %08b\t hexadecimal: %#x\n\n", w, w, w)
	fmt.Printf("Decimal: %d\t binary: %08b\t hexadecimal: %#x\n\n", x, x, x)
}

func NinjaLevel2_exercise5() {
	println("Exercise 5:")
	// Raw string literal
	var k string = `This is	a raw
	string        'f'        "literal"`

	fmt.Printf("%s\n\n", k)
}

func NinjaLevel2_exercise6() {
	println("Exercise 6:")
	// iota
	const (
		a = 2021 + iota
		b = 2021 + iota
		c = 2021 + iota
		d = 2021 + iota
	)

	fmt.Printf("Year 1: %d\tyear 2: %d\tyear 3: %d\t year 4: %d\n\n", a, b, c, d)
}
