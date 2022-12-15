package main

import (
	"fmt"
	"math/rand"
)

func longTimeOperation() <-chan int32 {
	ch := make(chan int32)
	func run() {
		defer close(ch)
		time.Sleep(time.Second * 5)
		ch <- rand.Int31n(300)
	}
	go run()
	return ch
}

func main() {
	fmt.Println("Soli Deo Gloria")

	ch := longTimeOperation()
	fmt.Println(ch)

}
