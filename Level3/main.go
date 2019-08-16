package main

import "fmt"

func main() {
	//Exercise 1
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
	fmt.Println("---")

	//Exercise 2
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%U %q\n", i, i)
		}
	}
	fmt.Println("---")

	//Exercise 3
	year := 1988
	for year <= 2019 {
		fmt.Println(year)
		year++
	}
	fmt.Println("---")

	//Exercise 4
	year = 1988
	for {
		fmt.Println(year)
		if year == 2019 {
			break
		}
		year++
	}
	fmt.Println("---")

	//Exercise 5
	for index := 10; index <= 100; index++ {
		fmt.Println(index % 4)
	}
	fmt.Println("---")

	//Exercise 6
	if year == 2019 {
		fmt.Println("Current year")
	}
	fmt.Println("---")

	//Exercise 7
	if year < 2000 {
		fmt.Println("Last century")
	} else {
		fmt.Println("Current century")
	}
	fmt.Println("---")

	//Exercise 8
	switch {
	case year == 2019:
		fmt.Println("Current year")
	default:
		fmt.Println("Other year")
	}
	fmt.Println("---")

	//Exercise 9
	favSport := "kenjutsu"
	switch favSport {
	case "kenjutsu":
		fmt.Println("Fav Sport")
	default:
		fmt.Println("Other sport")
	}
	fmt.Println("---")

	//Exercise 10
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
