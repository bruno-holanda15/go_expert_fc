package main

import "fmt"

func main() {
	canal := make(chan string) // Criando canal

	go func ()  {
		canal <- "Olá amigo!" // PREENCHE canal
		canal <- "Olá amigo novo!" // Não irá preencher POIS JÁ ESTÁ CHEIO
	}()

	msg := <-canal // ESVAZIA canal
	fmt.Println(msg)
}