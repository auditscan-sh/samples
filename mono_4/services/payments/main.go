package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/refunds", RefundHandler)
	mux.HandleFunc("/refunds/get", GetRefundHandler)
	mux.HandleFunc("/withdraw", WithdrawHandler)
	mux.HandleFunc("/webhooks/stripe", WebhookHandler)
	mux.HandleFunc("/receipt", ReceiptHandler)
	mux.HandleFunc("/charge", ChargeHandler)
	mux.HandleFunc("/debug/config", ConfigHandler)
	http.ListenAndServe("0.0.0.0:8080", mux)
}
