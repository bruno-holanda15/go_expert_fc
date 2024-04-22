package main

import "fmt"

// goroutine 1
func main() {
	ch := make(chan int)

	go publisher(ch)
	consumer(ch)
}

func publisher(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
}

func consumer(ch chan int) {
	for x := range ch {
		fmt.Printf("Consumed ch and received ==> %d\n", x)
	}
}
