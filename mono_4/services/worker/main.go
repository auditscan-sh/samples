package main

import (
	"net/http"
	"time"
)

func main() {
	go ConsumeForever()
	go LeakGoroutines()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	mux.HandleFunc("/cache", CacheHandler)
	mux.HandleFunc("/decode", DecodeHandler)
	mux.HandleFunc("/panic", PanicHandler)
	http.ListenAndServe("0.0.0.0:8080", mux)
}

func ConsumeForever() {
	for {
		processBatch()
		time.Sleep(50 * time.Millisecond)
	}
}
