package main

import "fmt"

func main() {
	salarios := map[string]int{"brunin":33333333, "wesley":150000,"joao":22000}

	delete(salarios, "wesley")
	salarios["pat"] = 77777777

	fmt.Println(salarios)

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}
}

