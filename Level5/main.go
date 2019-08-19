package main

import "fmt"

func main() {
	//Exercise 1
	type person struct {
		firstName      string
		lastName       string
		favoriteFlavor []string
	}

	person1 := person{
		firstName:      "Marcus",
		lastName:       "Santos",
		favoriteFlavor: []string{"vanilla", "cream"},
	}
	person2 := person{
		firstName:      "Marcus",
		lastName:       "Silva",
		favoriteFlavor: []string{"banana"},
	}

	fmt.Println(person1)
	for _, x := range person1.favoriteFlavor {
		fmt.Println(x)
	}

	fmt.Println(person2)
	for _, x := range person2.favoriteFlavor {
		fmt.Println(x)
	}
	fmt.Println("---")

	//Exercise 2
	mapValue := make(map[string]person)
	mapValue["person1"] = person1
	mapValue["person2"] = person2
	for _, value := range mapValue {
		fmt.Println(value)
		for _, x := range value.favoriteFlavor {
			fmt.Println(x)
		}
	}
	fmt.Println("---")

	//Exercise 3
	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	truck1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white"},
		fourWheel: true,
	}
	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(sedan1)
	fmt.Println(sedan1.color)
	fmt.Println(truck1)
	fmt.Println(truck1.color)
	fmt.Println("---")
	//Exercise 4
	anon := struct {
		name     string
		function string
	}{
		name:     "marcus",
		function: "teste",
	}
	fmt.Println(anon)
	fmt.Println("---")

}
