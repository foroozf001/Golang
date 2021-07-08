package main

import (
	"fmt"
)

func main() {
	NinjaLevel7_exercise1()
}

// Ninja Level 7 exercise 1
func NinjaLevel7_exercise1() {
	value := 42
	fmt.Println(value)
	fmt.Println(&value) //address
	fmt.Printf("%T\n", value)
	fmt.Printf("%T\n", &value) //type is pointer to int

	address := &value
	fmt.Println(*address) //dereference address gives value
	fmt.Println(*&value)

	*address = 43 // reassign value to whatever's stored at address
	fmt.Println(value)
}
