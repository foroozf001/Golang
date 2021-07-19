package ex4

import (
	"fmt"
	"sort" // sort package is what you're looking for, https://pkg.go.dev/sort#pkg-overview
)

func Ex4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	// sort integers: https://pkg.go.dev/sort#Ints
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	// sort strings: https://pkg.go.dev/sort#Strings
	sort.Strings(xs)
	fmt.Println(xs)
}
