package main

import "fmt"

func main() {
	NinjaLevel5_exercise1()
}

type person struct {
	first string
	last  string
}

func NinjaLevel5_exercise1() {
	person1 := person{
		first: "bob",
		last:  "ross",
	}

	fmt.Printf("%v\n", person1)
}
