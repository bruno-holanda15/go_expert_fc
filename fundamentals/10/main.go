package main

import (
	"fmt"
)

func main() {
	valor := sum(2, 43, 1, 3, 2, 12)

	funcAnon := func() int {
		return sum(1,2,3,4)
	}()

	fmt.Println(valor)
	fmt.Println(funcAnon)
}

func sum(numeros ...int) int {
	var total int
	for _, numero := range numeros {
		total += numero
	}

	return total
}
