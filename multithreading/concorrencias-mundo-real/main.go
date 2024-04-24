package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync/atomic"

	// "sync"
	"time"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		atomic.AddUint64(&number, 1)
		// m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Olá, você é o visitante número %d - Número de processadores %d", number, runtime.NumCPU())))
	})

	http.ListenAndServe(":3000", nil)
}
