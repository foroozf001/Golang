package main

import "fmt"

func main() {
	// NinjaLevel4_exercise1()
	// NinjaLevel4_exercise2()
	// NinjaLevel4_exercise3()
	// NinjaLevel4_exercise4()
	// NinjaLevel4_exercise5()
	// NinjaLevel4_exercise6()
	NinjaLevel4_exercise7()
}

func NinjaLevel4_exercise0() {
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

// array
func NinjaLevel4_exercise1() {
	var my_array [5]int
	my_array[0] = 41
	my_array[1] = 42
	my_array[2] = 43
	my_array[3] = 44
	my_array[4] = 45
	// x := [5]int{41, 42, 43, 44, 45}
	for i, v := range my_array {
		fmt.Printf("i: %d, v: %d\n", i, v)
	}
	fmt.Printf("%T\n", my_array)
}

// slice
func NinjaLevel4_exercise2() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}

func NinjaLevel4_exercise3() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x1 := x[:5]
	fmt.Printf("%v\n", x1)
	x2 := x[5:]
	fmt.Printf("%v\n", x2)
	x3 := x[2:7]
	fmt.Printf("%v\n", x3)
	x4 := x[1:6]
	fmt.Printf("%v\n", x4)
}

// appending slices
func NinjaLevel4_exercise4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Printf("%v\n", y)
}

func NinjaLevel4_exercise5() {
	states := []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon",
		"Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming"}
	fmt.Printf("len: %d, cap: %d\n", len(states), cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Printf("index: %d, state: %s\n", i, states[i])
	}
}

func NinjaLevel4_exercise6() {
	i1 := []string{"James", "Bond", "Shaken, not stirred"}
	i2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	ii := [][]string{i1, i2}

	for _, i := range ii {
		for j := range i {
			fmt.Printf("%s\t", i[j])
		}
		println()
	}
}

func NinjaLevel4_exercise7() {
	m := map[string][]string{
		"ross_bob":  []string{"painting", "brush", "paint"},
		"dylan_bob": []string{"music", "singing", "guitar"},
	}
	for _, item := range m {
		fmt.Printf("%v\n", item)
	}
}
