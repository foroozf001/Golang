package ex1

import (
	"encoding/json"
	"fmt"
)

// Type is user with underlying type of struct
type user struct {
	First string `json:"First"` //The json tag describes the field name in json
	Age   int    `json:"Age"`
}

// Marshal documentation: https://pkg.go.dev/encoding/json#Marshal
func Ex1() []byte {
	// Initialization
	var err error
	var users []user
	var u4 user

	// Initialization + declaration
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

	// Declaration with one-liner
	u4 = user{"Faraz", 29}

	// Declaration using either statement is fine
	users = []user{u1, u2, u3, u4} // users = append(users, u1, u2, u3, u4)

	// Type of users is []exercises.user
	// fmt.Printf("type: %T, value: %v\n", users, users)

	// Marshal returns a slice of byte and error
	var users_sb []byte
	users_sb, err = json.Marshal(users)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Type of users_sb is []uint8, that's a positive byte. The values in the slice are hexadecimal UTF-8 characters
	// One string character equals a uint8/byte
	// fmt.Printf("type: %T\n, ascii: %v\n, string: %s\n", users_sb, users_sb, string(users_sb))
	return users_sb
}
