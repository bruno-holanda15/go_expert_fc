package main

import "fmt"

func main() {
	var nome interface{} = "Percy"

	// println(nome)
	// println(nome.(string))

	// tenta transformar o nome em int, se der certo o ok é true e o res diferente de 0
	res, ok := nome.(int)
	fmt.Printf("O valor de res é %v e o valor de ok é %v\n", res, ok)

	res2, ok2 := nome.(string)
	fmt.Printf("O valor de res é %v e o valor de ok é %v\n", res2, ok2)

}
