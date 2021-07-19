package ex5

import (
	"fmt"
	"sort"
)

type user struct {
	First   string   `json:"First"`
	Last    string   `json:"Surname"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func (u user) String() string {
	return fmt.Sprintf("%s %s: %d %v", u.First, u.Last, u.Age, u.Sayings)
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func Ex5() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	sort.Sort(ByAge(users))
	for _, u := range users {
		sort.Strings(u.Sayings)
	}

	fmt.Println(users)

	// your code goes here

}
