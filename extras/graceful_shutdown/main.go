package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := &http.Server{Addr: ":3000"}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(4 * time.Second)
		w.Write([]byte("Olá carinha"))
	})

	// criado uma goroutine/"thread" acionando a concorrência no sistema
	go func() {
		fmt.Println("Server is running at http://localhost:3000")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed { // quando acionamos o shutdown no sistema, ele basicamente é um ErrServerClose e não cai nesse if 
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	// criado channel do tipo os.Signal
	stop := make(chan os.Signal, 1)

	// enviando/PREENCHENDO para p channel stop uma mensagem caso ocorra algum dos casos abaixo
	// Interrupt pelo os ou syscall e Terminated pelo syscall
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// bloqueando programa para receber/ESVAZIAR msg no channel
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	}

	fmt.Println("Server stopped")
}
