package main

import "fmt"

func main() {

	// Memória -> Endereço -> Valor
	// variável aponta para um ponteiro que tem um endereço na memória
	a := 10

	// criar uma variável do tipo de ponteiro, com o endereço de a
	var ponteiro *int = &a

	// atribui um valor novo para o endereço que foi atribuido anteriormente
	*ponteiro = 20

	// faz a mesma coisa com o ponteiro acima
	b := &a
	*b = 21

	// como ponteiro e b são ponteiros, para trazer o valor usamos o *
	fmt.Println(a, *ponteiro, *b)
}
