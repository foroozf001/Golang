package main

import (
	"fmt"

	"github.com/foroozf001/Golang/010_application/ex1"
	"github.com/foroozf001/Golang/010_application/ex2"
	"github.com/foroozf001/Golang/010_application/ex3"
	"github.com/foroozf001/Golang/010_application/ex4"
	"github.com/foroozf001/Golang/010_application/ex5"
)

func main() {
	// Exercise1()
	// Exercise2()
	// Exercise3()
	// Exercise4()
	Exercise5()
}

func Exercise1() {
	// slice of slice of byte
	var exs [][]byte
	exs = append(exs, ex1.Ex1())
	fmt.Printf("%s\n", string(exs[0]))
}

func Exercise2() {
	var users []ex2.User
	users = ex2.Ex2()
	fmt.Printf("%T, %v\n", users, users)
}

func Exercise3() {
	ex3.Ex3()
}

func Exercise4() {
	ex4.Ex4()
}

func Exercise5() {
	ex5.Ex5()
}
