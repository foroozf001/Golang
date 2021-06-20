package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	x := quote.Hello()
	fmt.Printf("%v\n", x)
}
