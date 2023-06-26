package main

import (
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	//m := sync.RWMutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		//number++
		//	m.RUnlock()
		atomic.AddInt64(&number, 1)

		time.Sleep(3000 * time.Millisecond)
		w.Write([]byte("Tivemos um total de " + string(number) + "vizitas"))

	})
	http.ListenAndServe(":3000", nil)

}
