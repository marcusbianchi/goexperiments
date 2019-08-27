package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

//Exercise 2
type person struct {
	name string
}

func main() {
	//Exercise 1
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		fmt.Println("1")
		waitGroup.Done()

	}()
	go func() {
		fmt.Println("2")
		waitGroup.Done()
	}()

	fmt.Println("Main Ended")
	waitGroup.Wait()
	fmt.Println("---")

	//Exercise 2
	p := person{
		name: "Marcus",
	}
	//does not work
	//saySomething(p)

	//work
	saySomething(&p)

	//work
	p.speak()
	fmt.Println("---")

	//Exercise 3

	var newWaitGroup sync.WaitGroup

	//Contains Race conditions
	/* 	newWaitGroup.Add(100)
	count := 0

	for v := 0; v < 100; v++ {
		go func() {
			localCount := count
			runtime.Gosched()
			localCount++
			count = localCount
			mt.Println("Counter", count)
			newWaitGroup.Done()
		}()
	}

	newWaitGroup.Wait()
	fmt.Println("Final Counter", count)
	fmt.Println("---") */

	//Exercise 4
	count := 0
	var mutex sync.Mutex
	newWaitGroup.Add(100)
	for v := 0; v < 100; v++ {
		go func() {
			mutex.Lock()
			count++
			fmt.Println("Counter", count)
			mutex.Unlock()
			newWaitGroup.Done()
		}()
	}

	newWaitGroup.Wait()
	fmt.Println("Final Counter", count)
	fmt.Println("---")

	//Exercise 5
	var counter int64
	newWaitGroup.Add(100)

	for v := 0; v < 100; v++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter", atomic.LoadInt64(&counter))
			newWaitGroup.Done()
		}()
	}
	newWaitGroup.Wait()
	fmt.Println("Final Counter", counter)
	fmt.Println("---")

	//Exercise 6
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("ARCH", runtime.GOARCH)
}

//Exercise 2

func (p *person) speak() {
	fmt.Println(p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}
