package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

// goroutine 1 -> main já é uma go routine/thread
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(15)

	// goroutine 2
	go task("Deploy", &waitGroup)

	// goroutine 3
	go task("Build", &waitGroup)

	// goroutine 4
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "Func Anônima")
			time.Sleep(1 * time.Second)
			waitGroup.Done()
		}
	}()

	waitGroup.Wait()
	// Caso não tenha algo que "prenda" a função main e só rodem as goroutines em backgroung
	// o código irá "finalizar" e sair
	// Com os waitGroups, conseguimos mapear quantas goroutines estão em execução,
	// criadas e esperar essas operações serem finalizadas
}
