package main

import "fmt"

type Mission struct {
	qtd int
}

// método de criação de uma nova struct Mission
func NewMission() *Mission {
	return &Mission{qtd: 0}
}

// soma o valor da qtd da struct pois não está pegando uma cópia
// *Mission pega o endereço da memória da struct e altera nela
func (m *Mission) add(n int) {
	m.qtd += n
}

func main()  {
	mission := NewMission()

	mission.add(100)
	mission.add(200)

	fmt.Printf("Quantidade de missões: %d\n", mission.qtd)
}
