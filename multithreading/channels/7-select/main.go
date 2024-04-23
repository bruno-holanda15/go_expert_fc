package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// A ideia do select para os channels, é atuar como um switch case de canais

type Message struct {
	Id int64
	Msg string
}

var i int64 = 0

func main()  {
	c1 := make(chan Message)
	c2 := make(chan Message)
	
	// RabbitMQ
	go func ()  {
		// c1 <- 0 com isso ele atribui primeiro e não cai no default
		for {
			atomic.AddInt64(&i, 1) // evitar problema de race condition
			msg := Message{i, "RabbitMQ"}
			time.Sleep(time.Second)
			c1 <- msg
		}
	}()

	// Kafka
	go func ()  {
		for {
			atomic.AddInt64(&i, 1) // evitar problema de race condition
			msg := Message{i, "Kafka"}
			time.Sleep(time.Second * 2)
			c1 <- msg
		}
	}()

	for {
	// for i := 0; i < 3; i++ {
		select {
		case msg := <-c1: // rabbitmq
			fmt.Printf("received from %s - id %d\n", msg.Msg, msg.Id)
		case msg := <-c2: // kafka
			fmt.Printf("received from %s - id %d\n", msg.Msg, msg.Id)
		case <-time.After(time.Second *5):
			println("timeout")
		// default:
		// 	println("default")
		}
	}
}


