package Exercise2

import "fmt"

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func (p *person) Speak() {
	fmt.Printf("Hi, my name is %s %s, and I'm %d years old.\n", (*p).First, (*p).Last, (*p).Age)
}

type human interface {
	Speak()
}

// Now this method can call the Speak function on anything that implements the Human interface, not only type Person
func SaySomething(h human) {
	h.Speak()
}

func Ex2() {
	person1 := person{"Bob", "Ross", 71}
	SaySomething(&person1)
}
