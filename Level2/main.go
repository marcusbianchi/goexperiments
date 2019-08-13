package main

import "fmt"

func main() {

	//Exercise 1
	x := 10
	fmt.Printf("%d\n", x)
	fmt.Printf("%b\n", x)
	fmt.Printf("%X\n", x)
	fmt.Println("---")

	//Exercise 2
	fmt.Println(1 == 1)
	fmt.Println(1 <= 1)
	fmt.Println(1 >= 1)
	fmt.Println(1 != 1)
	fmt.Println(1 < 1)
	fmt.Println(1 > 1)
	fmt.Println("---")

	//Exercise 3
	const a = 1
	const b int = 5
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("---")

	//Exercise 4
	c := x << 1
	fmt.Printf("%d\n", c)
	fmt.Printf("%b\n", c)
	fmt.Printf("%X\n", c)
	fmt.Println("---")

	//Exercise 5
	d := `	String
	with
	linebreak`
	fmt.Println(d)
	fmt.Println("---")

	//Exercise 6
	const (
		e = iota
		f = iota
		g = iota
		h = iota
	)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println("---")
}
