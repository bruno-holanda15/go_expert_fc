package main

import (
	"fmt"
	"packaging/1/math"
)

func main() {
	fmt.Println("Hello baby!")

	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}