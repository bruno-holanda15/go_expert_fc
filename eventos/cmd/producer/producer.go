package main

import (
	"fmt"
	"go_eventos/pkg/rabbitmq"
)

func main() {
	ch := rabbitmq.OpenChannel()
	defer ch.Close()

	err := rabbitmq.Publish(ch, "Olá RabbitMQ", "fila_teste", "amq.direct")
	if err != nil {
		fmt.Printf("Error Publishing %s", err)
	}
}