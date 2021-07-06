package main

import "fmt"

func main() {
	NinjaLevel5_exercise1()
}

func NinjaLevel5_exercise1() {
	// slice
	x := []int{1, 2, 6, 10}
	for i, v := range x {
		fmt.Printf("%d %d\n", i, v)
	}

	// append elements
	x = append(x, 12, 13)
	fmt.Printf("%v %d\n", x, len(x))

	// join two slices
	y := []int{-1, -2}
	x = append(x, y...)
	fmt.Printf("%v %d\n", x, len(x))

	// delete elements
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	// make
	z := make([]int, 4, 10)
	// there's 4 elements in the slice and there's 10 spots in total in the underlying array
	fmt.Printf("slice: %v\t length: %d\t capacity: %d\n", z, len(z), cap(z))
}
