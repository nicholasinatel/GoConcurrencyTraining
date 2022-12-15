package main

import (
	"fmt"
	"sync"
)

// func Example1() {
// 	go printSomething("this is the first thing to be printed")
// 	time.Sleep(time.Second)
// 	printSomething("this is the second thing to be printed")
// }

func Example2() {

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()
}

func main() {
	fmt.Println("Soli Deo Gloria")
	// Example1()
	Example2()
}

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
