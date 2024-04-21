package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var number uint64 = 0

func main() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		w.Write([]byte(fmt.Sprintf("Olá, você é o visitante número %d", number)))
		time.Sleep(300 * time.Millisecond)
	})

	http.ListenAndServe(":3000", nil)
}