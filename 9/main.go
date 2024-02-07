package main

import (
	"fmt"
)

func main() {
	valor := sum(2, 43, 1, 3, 2, 12)

	fmt.Println(valor)
}

func sum(numeros ...int) int {
	var total int
	for _, numero := range numeros {
		total += numero
	}

	return total
}
