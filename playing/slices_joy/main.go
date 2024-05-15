package main

import "fmt"

func main() {
	eventos := []string{"teste", "teste2", "teste3", "teste4", "teste5"}

	// possibilitando um shift
	fmt.Println(eventos[:2])
	fmt.Println(eventos[3:])
}
