package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	f, err := os.Create("arquivo.txt")

	if err != nil {
		panic(err)
	}

	// tamanho, err := f.WriteString("Adicionando texto ao arquivo")
	tamanho, err := f.Write([]byte("Escrevendo texto ao arquivo"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("O tamanho do arquivo é %d bytes\n", tamanho)

	f.Close()

	// leitura
	arquivo, err := os.ReadFile("arquivo.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(arquivo))

	// leitura de pouco em pouco, abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		// n é a posição de onde ele está fazendo a leitura
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivo.txt")
	if  err != nil {
		panic(err)
	}
}