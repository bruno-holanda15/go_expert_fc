package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	select {
	case <-time.After(3 * time.Second):
		// log impreme no comand line -> stdout
		log.Println("Request processada com sucesso")
		w.Write([]byte("Request processada com sucesso"))

	case <-ctx.Done():
		// impreme no command line
		// caso o cliente cancele a requesição
		log.Println("Request cancelada pelo client")
	}

}