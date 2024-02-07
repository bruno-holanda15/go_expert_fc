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
	fmt.Println("Ola")
	fmt.Printf("O tipo de F é %T\n", f)
	fmt.Printf("O tipo de E é %T", e)
}
