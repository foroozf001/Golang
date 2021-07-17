package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type user struct {
	First string
	Age   int
}

func main() {
	// NinjaLevel8_exercise1()
	// NinjaLevel8_exercise2()
	// NinjaLevel8_exercise3()
	// NinjaLevel8_exercise4()
	// NinjaLevel8_exercise5()
	test()
}

func test() {
	var people []user
	bobby := user{
		First: "Bobby",
		Age:   32,
	}

	jessy := user{
		First: "Jessy",
		Age:   21,
	}
	people = append(people, bobby, jessy)
	fmt.Printf("%T\t%v", people, people)
}

func NinjaLevel8_exercise1() {
	u1 := user{
		First: "James",
		Age:   32,
	}
	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}
	u3 := user{
		First: "M",
		Age:   54,
	}
	users := []user{u1, u2, u3}
	out, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(users)
	fmt.Println(string(out))
}

type spy struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func NinjaLevel8_exercise2() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	// fmt.Println(s)
	var spies []spy
	err := json.Unmarshal([]byte(s), &spies)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(spies)
}

func NinjaLevel8_exercise3() {
	u1 := spy{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := spy{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := spy{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []spy{u1, u2, u3}

	// fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}

func NinjaLevel8_exercise4() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

func NinjaLevel8_exercise5() {
	u1 := spy{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := spy{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := spy{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	spies := []spy{u1, u2, u3}

	sort.Sort(ByAge(spies)) //convert to a type that implements a custom interface
	for _, spy := range spies {
		spy.speak()
	}
	sort.Sort(ByLast(spies))
	for _, spy := range spies {
		spy.speak()
	}

}

func (s spy) speak() {
	fmt.Println(s.First, s.Last, s.Age)
	for _, saying := range s.Sayings {
		fmt.Println("\t", saying)
	}
}

type ByAge []spy
type ByLast []spy

func (b ByAge) Len() int           { return len(b) }
func (b ByAge) Less(i, j int) bool { return b[i].Age < b[j].Age }
func (b ByAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func (b ByLast) Len() int           { return len(b) }
func (b ByLast) Less(i, j int) bool { return b[i].Last < b[j].Last }
func (b ByLast) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
