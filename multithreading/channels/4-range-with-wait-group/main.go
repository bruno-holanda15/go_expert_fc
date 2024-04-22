package main

import (
	"fmt"
	"sync"
)

// goroutine 1
func main() {
	ch := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(10)

	go publisher(ch)
	go consumer(ch, &wg)

	wg.Wait()
}

func publisher(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
}

func consumer(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Consumed ch and received ==> %d\n", x)
		wg.Done()
	}
}
