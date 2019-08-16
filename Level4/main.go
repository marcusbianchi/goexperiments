package main

import (
	"fmt"
)

func main() {
	//Exercise 1
	array := [5]int{1, 2, 3, 4, 5}
	array[0] = 6
	array[1] = 7
	array[2] = 8
	array[3] = 9
	array[4] = 10
	for _, v := range array {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", array)
	fmt.Println("---")

	//Exercise 2
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for _, v := range slice {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", slice)
	fmt.Println("---")

	//Exercise 3
	sliceSub := slice[:5]
	fmt.Println(sliceSub)
	sliceSub = slice[5:]
	fmt.Println(sliceSub)
	sliceSub = slice[2:7]
	fmt.Println(sliceSub)
	sliceSub = slice[1:6]
	fmt.Println(sliceSub)
	fmt.Println("---")

	//Exercise 4
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("---")

	//Exercise 5
	x = []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
	fmt.Println("---")

	//Exercise 6
	z := []string{" Alabama", " Alaska", " Arizona", " Arkansas", " California", " Colorado", " Connecticut", " Delaware", " Florida", " Georgia", " Hawaii", " Idaho", " Illinois", " Indiana", " Iowa", " Kansas", " Kentucky", " Louisiana", " Maine", " Maryland", " Massachusetts", " Michigan", " Minnesota", " Mississippi", " Missouri", " Montana", " Nebraska", " Nevada", " New Hampshire", " New Jersey", " New Mexico", " New York", " North Carolina", " North Dakota", " Ohio", " Oklahoma", " Oregon", " Pennsylvania", " Rhode Island", " South Carolina", " South Dakota", " Tennessee", " Texas", " Utah", " Vermont", " Virginia", " Washington", " West Virginia", " Wisconsin", " Wyoming"}
	fmt.Println(len(z))
	fmt.Println(cap(z))

	for index := 0; index < len(z); index++ {
		fmt.Println(z[index])
	}
	fmt.Println("---")

	//Exercise 7
	k := [][]string{{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."}}

	for _, v1 := range k {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	fmt.Println("---")

	//Exercise 8
	mapValue := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}
	for k, v := range mapValue {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	fmt.Println("---")

	//Exercise 9
	mapValue["Gerald"] = []string{"Yennifer", "Triss"}
	for k, v := range mapValue {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	fmt.Println("---")

	//Exercise 10
	delete(mapValue, `no_dr`)

	for k, v := range mapValue {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
	fmt.Println("---")

}
