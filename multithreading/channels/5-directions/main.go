package main

import "fmt"

// Ou um channel recebe dados
// Ou entrega dados para alguma variável

func main()  {
	ch := make(chan string)
	
	go read(ch)

	for i := 0; i < 10; i++ {
		fillChannel(fmt.Sprintf("Hello baby do id: %d", i), ch)
	}

}

// chan apenas recebe chan<- // enche o canal
func fillChannel(name string, ch chan<- string) {
	ch <- name
}

// chan apenas entrega <-chan // esvazia o canal
func read(ch <-chan string)  {
	for x := range ch {
		fmt.Println(x)
	}
}