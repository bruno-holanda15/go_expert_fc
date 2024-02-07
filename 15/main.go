package main

import "fmt"

// essa interface "implementa" todo mundo
// pois não tem nenhuma especificidade
type hercules interface{}

func main() {
	var x interface{} = "Percy"
	var y interface{} = 28

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", t, t)
}