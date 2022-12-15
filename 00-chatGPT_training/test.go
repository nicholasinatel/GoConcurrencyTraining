package main

import "fmt"

func fillChannel(functions ...func() int) chan int {
	ch := make(chan int, len(functions))
	go func(ch chan int) {
		for i := len(functions) - 1; i >= 0; i-- {
			ch <- functions[i]()
		}
		close(ch)
	}(ch)

	return ch
}

func exampleFunction(counter int) int {
	sum := 0
	for i := 0; i < counter; i++ {
		sum += 1
	}
	return sum
}

func main() {
	expensiveFunction := func() int { return exampleFunction(20000000000) }
	cheapFunction := func() int { return exampleFunction(1000000000) }

	ch := fillChannel(expensiveFunction, cheapFunction)

	if ch != nil {
		for i := 0; i < 2; i++ {
			fmt.Printf("Result: %d\n", <-ch)
		}
	}
}
