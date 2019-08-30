package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Exercise 1
type person struct {
	First   string
	Last    string
	Sayings []string
}

//Exercise 3
type customErr struct {
	description string
	code        int
}

//Exercise 4
type sqrtError struct {
	lat  string
	long string
	err  error
}

func main() {
	//Exercise 1
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	fmt.Println("---")

	//Exercise 2
	bs, err = toJSON(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	fmt.Println("---")

	//Exercise 3
	custError := customErr{
		description: "Custom error",
		code:        042,
	}
	printError(custError)
	fmt.Println("---")

	//Exercise 4
	_, err = sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("---")
}

//Exercise 2
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return nil, fmt.Errorf("new error")
	}
	return bs, nil
}

//Exercise 3
func (err customErr) Error() string {
	return fmt.Sprint("Erro: ", err.description, " Code: ", err.code)
}

func printError(err error) {
	fmt.Println(err)
}

//Exercise 4
func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := fmt.Errorf("more coffee needed - value was %v", f)
		return 0, sqrtError{"50.2289 N", "99.4656 W", e}
	}
	return 42, nil
}
