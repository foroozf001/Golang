package main

import "fmt"

func main() {
	// NinjaLevel5_exercise1()
	// NinjaLevel5_exercise2()
	// NinjaLevel5_exercise3()
	NinjaLevel5_exercise4()
}

type person struct {
	first              string
	last               string
	favorite_icecreams []string
}

type ninja struct {
	person
	license_to_kill bool
}

func NinjaLevel5_exercise1() {
	person1 := person{
		first: "bob",
		last:  "ross",
		favorite_icecreams: []string{
			"mocha",
			"vanilla",
			"chocolate",
		},
	}

	for _, v := range person1.favorite_icecreams {
		fmt.Printf("%s ", v)
	}
	println()
}

func NinjaLevel5_exercise2() {
	person1 := person{
		first: "bob",
		last:  "ross",
		favorite_icecreams: []string{
			"mocha",
			"vanilla",
			"chocolate",
		},
	}

	person2 := person{
		first: "bob",
		last:  "marley",
		favorite_icecreams: []string{
			"strawberry",
			"caramel",
			"mint",
		},
	}

	m := map[string]person{
		person1.last: person1,
		person2.last: person2,
	}

	for k, v := range m {
		fmt.Printf("%s\n", k)
		for _, f := range v.favorite_icecreams {
			fmt.Printf("\t%s\n", f)
		}
	}
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func NinjaLevel5_exercise3() {
	truck1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "red",
		},
		fourWheel: true,
	}
	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		luxury: true,
	}
	fmt.Printf("%v\n", truck1)
	fmt.Printf("%v\n", sedan1)
}

func NinjaLevel5_exercise4() {
	anon1 := struct {
		prop1 string
		prop2 string
		prop3 int
		prop4 bool
		prop5 map[string][]string
	}{
		prop1: "one",
		prop2: "green",
		prop3: 24,
		prop4: true,
		prop5: map[string][]string{
			"key1": {"pair1", "pair2"},
			"key2": {"string12", "string22"},
		},
	}
	fmt.Printf("%v\n", anon1)
}
