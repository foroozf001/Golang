package ex2

import (
	"encoding/json"
	"fmt"
)

// Only members with Capitals are exported to other packages
type User struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

// Unmarshal documentation: https://pkg.go.dev/encoding/json#Unmarshal
func Ex2() []User {
	var s string
	var sb []byte
	var err error
	var users []User
	s = `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	sb = []byte(s)
	err = json.Unmarshal(sb, &users)
	if err != nil {
		fmt.Println(err.Error())
	}
	return users
}
