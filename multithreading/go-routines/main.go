package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1 -> main já é uma go routine/thread
func main() {
	// Thread 2
	go task("Deploy")

	// Thread 3
	go task("Build")

	go func(){
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "Func Anônima")
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(8 * time.Second)

	// Caso não tenha algo que "prenda" a função main e só rodem as threads em backgroung
	// o código irá "finalizar" e sair
}