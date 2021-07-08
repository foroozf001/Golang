package main

import (
	"fmt"
)

func main() {
	// NinjaLevel7_exercise1()
	NinjaLevel7_exercise2()
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

	x := 0
	foo(&x)
	fmt.Println(x)
}

func foo(in *int) {
	fmt.Println(*in)
	*in = 73
	fmt.Println(*in)
}

// Ninja Level 7 exercise 2
func NinjaLevel7_exercise2() {
	person1 := person{
		first: "Bob",
		last:  "Ross",
		age:   42,
	}
	fmt.Printf("%v\n", person1)
	fmt.Printf("%v\n", person1.speak())
	//passing in a pointer or a value does not matter
	info(&person1)
	info(person1)
}

type creature interface {
	speak() string
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() string {
	return fmt.Sprintf("Hi I'm %s", p.first)
}

func info(c creature) {
	fmt.Println(c.speak())
}

// func (p *person) changeMe() {
// 	(*p).first = "Steve"
// 	(*p).last = "Banner"
// 	(*p).age *= 2
// }
