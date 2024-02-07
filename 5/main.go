package main

import (
	"fmt"
)

const a = "Hello Brunin"

type ID int

var (
	b bool = true
	c int
	d string
	e float64
	f ID
)

func main() {
	// array determina seu tamanho e não muda
	var meuArray [3]int

	meuArray[0] = 1
	meuArray[1] = 21
	meuArray[2] = 33333333

	fmt.Println(len(meuArray))
	fmt.Println(meuArray)

	// range percorre os itens do meuArray
	for i, v := range meuArray {
		fmt.Printf("index: %d - valor: %d\n", i, v)
	}

}
