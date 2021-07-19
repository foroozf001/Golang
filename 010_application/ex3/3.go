package ex3

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string   `json:"First"`
	Last    string   `json:"Surname"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func Ex3() {
	var users []user
	var u1, u2, u3 user
	var err error

	u1 = user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 = user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 = user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users = []user{u1, u2, u3}

	// https://pkg.go.dev/encoding/json#Encoder.Encode
	err = json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err.Error())
	}
}
