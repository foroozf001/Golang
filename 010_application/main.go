package main

import (
	"fmt"

	"github.com/foroozf001/Golang/010_application/ex1"
	"github.com/foroozf001/Golang/010_application/ex2"
)

func main() {
	// Exercise1()
	Exercise2()
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
