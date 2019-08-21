package main

import "fmt"

//Exercise 2
type person struct {
	first string
}

func main() {
	//Exercise 1
	x := 1
	fmt.Println(&x)
	fmt.Printf("%T\n", &x)
	fmt.Println("---")

	//Exercise 2
	p1 := person{
		first: "Marcus",
	}
	fmt.Println(p1.first)
	changeMe(&p1)
	fmt.Println(p1.first)
	fmt.Println("---")

}

//Exercise 2
func changeMe(p *person) {
	p.first = "New First"
}
