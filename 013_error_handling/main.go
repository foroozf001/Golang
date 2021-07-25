package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// Ex1()
	// Ex2()
	// Ex3()
	Ex4()
}

type person struct {
	First   string
	Last    string
	Sayings []string
}

func Ex1() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("Error: %e", err)
	}
	return bs, err
}

func Ex2() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)

	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(string(bs))
}

func Ex3() {
	var c1 customError
	c1 = customError{
		info: "This is my new error type",
	}
	foo(c1)
}

type customError struct {
	info string
}

func (ce customError) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func foo(err error) {
	log.Println(err)
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func Ex4() {
	// lat "50.2289 N"
	// long "99.4656 W"
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := fmt.Errorf("more coffee needed - value was %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}
