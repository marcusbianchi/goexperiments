package main

import (
	"fmt"
)

func main() {
	//Exercise 1
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
	fmt.Println("---")
	//Exercise 2.1
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)
	fmt.Printf("cs\t%T\n", cs)
	//Exercise 2.2
	cr := make(chan int)

	go func() {
		cr <- 42
	}()
	fmt.Println(<-cr)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Println("---")
	//Exercise 3
	c1 := gen()
	receive(c1)

	fmt.Println("---")

	//Exercise 4
	q := make(chan int)
	c2 := gen2(q)

	receive2(c2, q)
	fmt.Println("---")

	//Exercise 5
	c3 := make(chan int)
	go func() {
		c3 <- 1
	}()

	v, ok := <-c3
	fmt.Println(v, ok)

	close(c3)

	v, ok = <-c3
	fmt.Println(v, ok)
	fmt.Println("---")

	//Exercise 6
	genChannel := generator()
	receiver(genChannel)
	fmt.Println("---")

	//Exercise 7
	x := 10
	y := 10
	c7 := generatorX10(x, y)
	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c7)
	}

}

//Exercise 3
func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}

		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println((v))
	}
}

//Exercise 4
func gen2(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()
	return c
}

func receive2(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)

		case <-q:
			return
		}
	}
}

//Exercise 6
func generator() <-chan int {
	channel := make(chan int)
	go func() {
		for index := 0; index < 10; index++ {
			channel <- index
		}
		close(channel)
	}()
	return channel
}

func receiver(c <-chan int) {
	for {
		select {
		case v, ok := <-c:
			if ok {
				fmt.Println(v)
			} else {
				return
			}
		}
	}
}

//Exercise 7
func generatorX10(maxI, maxJ int) <-chan string {
	channel := make(chan string)

	for i := 0; i < maxI; i++ {
		go func(i int) {
			for j := 0; j < maxJ; j++ {
				channel <- fmt.Sprint(j, i)
			}
		}(i)
	}

	return channel
}
