package main

import (
	"fmt"
	"time"
)


func main() {
	data := make(chan int)
	QtdWorkers := 100

	// inicializa os workers
	for i := 0; i < QtdWorkers; i++ {
		go worker(i, data)
	}


	for i := 0; i < 1000; i++ {
		data <- i // enviando para o canal o i
	}
}

func worker(workedId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workedId, x)
		time.Sleep(time.Second)
	}
}