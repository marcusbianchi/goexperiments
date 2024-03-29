package main

import (
	"fmt"
	"math"
)

//Exercise 3
type person struct {
	first string
	last  string
	age   int
}

//Exercise 5
type circle struct {
	r float64
}

type square struct {
	l float64
	w float64
}

type shape interface {
	area() float64
}

func main() {
	//Exercise 1
	fmt.Println(foo())
	fmt.Println(bar())
	fmt.Println("---")

	//Excercise 2
	array := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(foo1(array...))
	fmt.Println(bar1(array))
	fmt.Println("---")

	//Exercise 3
	defer fmt.Println("------ Defer ------")
	fmt.Println("Posterior")

	//Exercise 4
	p := person{
		first: "Marcus",
		last:  "Santos",
		age:   31,
	}
	p.speak()

	//Exercise 5
	c := circle{
		r: 1.5,
	}
	sq := square{
		l: 3.2,
		w: 3,
	}
	info(c)
	info(sq)
	fmt.Println("---")

	//Exercise 6
	func(s string) {
		fmt.Println(s)
	}("We are anonimous")
	fmt.Println("---")

	//Exercise 7
	fn := func(s string) {
		fmt.Println(s)
	}
	fn("New Annonimous")
	fmt.Println("---")
	//Exercises 8
	returnfn := returnfunc()

	fmt.Println("Callback", returnfn())
	fmt.Println("---")

	//Exercise 9
	fn2 := func(i int) {
		fmt.Println("callback", i)
	}
	receivefn(fn2, 1894)
	fmt.Println("---")

	//Exercise 10
	cl := closure()

	cl(2)
	cl(10)
	fmt.Println("---")

}

//Exercise 1
func foo() int {
	return 1984
}
func bar() (int, string) {
	return 1984, "Brazil 2019"
}

//Exercise 2
func foo1(array ...int) int {
	total := 0
	for _, i := range array {
		total += i
	}
	return total
}

func bar1(array []int) int {
	total := 0
	for _, i := range array {
		total += i
	}
	return total
}

//Exercise 3
func (p person) speak() {
	fmt.Println("I'm a person!", p.age)
}

//Exercise 5
func (c circle) area() float64 {
	return math.Pow(c.r, 2) * math.Pi
}

func (s square) area() float64 {
	return s.l * s.w
}

func info(s shape) {
	fmt.Println("The Area is:", s.area())
}

//Exercise 8
func returnfunc() func() int {
	return func() int {
		return 42
	}
}

//Exercise 9
func receivefn(fn func(i int), i int) {
	fn(i)
}

//Exercise 10
func closure() func(i int) {
	x := 0
	fmt.Println("Original x", x)
	return func(i int) {
		x += i
		fmt.Println("New x", x)
	}
}
