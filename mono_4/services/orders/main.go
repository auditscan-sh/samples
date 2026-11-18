package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", OrdersHandler)
	mux.HandleFunc("/orders/nickname", NicknameHandler)
	mux.HandleFunc("/orders/sorted", SortedHandler)
	mux.HandleFunc("/orders/nosql", NoSQLHandler)
	mux.HandleFunc("/orders/xml", XMLHandler)
	mux.HandleFunc("/orders/xpath", XPathHandler)
	mux.HandleFunc("/orders/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	http.ListenAndServe("0.0.0.0:8080", mux)
}
