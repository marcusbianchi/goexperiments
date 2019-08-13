package main

import "fmt"

var x_1 int
var y_1 string
var z_1 bool

type new_type int

var n new_type

func main() {

	//Exercise 1
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("---")

	//Exercise 2

	fmt.Println(x_1, y_1, z_1)
	fmt.Println(x_1)
	fmt.Println(y_1)
	fmt.Println(z_1)
	fmt.Println("---")

	//Exercise 3
	x_1 = 42
	y_1 = "James Bond"
	z_1 = true

	fmt.Println(x_1, y_1, z_1)
	fmt.Println(x_1)
	fmt.Println(y_1)
	fmt.Println(z_1)
	fmt.Println("---")

	//Exercise 4
	fmt.Println(n)
	fmt.Printf("%T \n", n)
	n = 42
	fmt.Println(n)
	fmt.Println("---")

	//Exercise 5
	x = 5
	x = int(n)
	fmt.Println(n)
	fmt.Println("---")

}
