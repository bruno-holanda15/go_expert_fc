package main

// recebendo o endereço na memória, ponteiro
func soma(a, b *int) int {

	// mudando o valor na memória
	*a = 50
	*b = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	// passando o endereço na memória, ponteiro
	soma(&minhaVar1, &minhaVar2)

	println(minhaVar1)
	println(minhaVar2)
}