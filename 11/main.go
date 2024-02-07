package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
}

// Não pode passar propriedades, apenas métodos
type Pessoa interface {
	Desativar()
}

// Cliente implementa automaticamente a interface Pessoa
type Cliente struct {
	Nome    string
	Idade   int
	Ativo   bool
	Address Endereco
}

func (c *Cliente) Desativar() {
	c.Ativo = false
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {

	brunin := Cliente{
		Nome:  "Bruno",
		Idade: 28,
		Ativo: true,
		Address: Endereco{
			"Avenida Sucesso",
			33333333,
		},
	}

	// brunin.Ativo = false
	// brunin.Desativar()
	Desativacao(&brunin)

	fmt.Printf("Nome: %s - Idade: %d - Ativo: %t - Address: %s nº %d\n",
		brunin.Nome,
		brunin.Idade, 
		brunin.Ativo, 
		brunin.Address.Rua, 
		brunin.Address.Numero,
	)
	fmt.Println(brunin)
}
