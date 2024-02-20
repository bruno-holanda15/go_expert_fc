package main

import (
	"errors"
	"fmt"
)

func main() {
	valor, err := sum(2, 43)

	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(valor)
}

func sum(a, b int) (int, error) {
	if a + b >= 50 {
		return a + b, nil
	}
	return a + b, errors.New("A soma é menor que 50")
}