package main

import "fmt"

func main() {
	// NinjaLevel3_exercise1()
	// NinjaLevel3_exercise2()
	// NinjaLevel3_exercise3()
	// NinjaLevel3_exercise4()
	// NinjaLevel3_exercise5()
	// NinjaLevel3_exercise6()
	// NinjaLevel3_exercise7()
	// NinjaLevel3_exercise8()
	NinjaLevel3_exercise9()
}

func NinjaLevel3_exercise1() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d\t\n", i)
	}
}

func NinjaLevel3_exercise2() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d:\n", i)
		for h := 0; h < 3; h++ {
			fmt.Printf("\t%U\n", i)
		}
	}
}

func NinjaLevel3_exercise3() {
	year := 1991
	for year <= 2021 {
		fmt.Printf("%d\n", year)
		year++
	}
}

func NinjaLevel3_exercise4() {
	year := 1991
	for {
		if year > 2021 {
			break
		} else {
			fmt.Printf("%d\n", year)
			year++
		}
	}
}

func NinjaLevel3_exercise5() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%d\n", i%4)
	}
}

func NinjaLevel3_exercise6() {
	x := "James Bond"

	if x == "James Bond" {
		fmt.Println(x)
	}
}

func NinjaLevel3_exercise7() {
	x := "Bugs Bunny"

	if x == "James Bond" {
		fmt.Println(x)
	} else if x == "Bugs Bunny" {
		fmt.Println(x + "!")
	} else {
		fmt.Println("None")
	}
}

func NinjaLevel3_exercise8() {
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}

func NinjaLevel3_exercise9() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}
