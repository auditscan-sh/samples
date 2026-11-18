package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// CWE-306 + CWE-840: unauthenticated refund with a trusted amount.
func RefundHandler(w http.ResponseWriter, r *http.Request) {
	refundID := r.FormValue("refundId")
	customerID := r.FormValue("customerId")
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
	refund, err := IssueRefund(refundID, customerID, amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(refund)
}

// CWE-639: refund read has no ownership or tenant check.
func GetRefundHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	refund, err := GetRefund(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(refund)
}

// CWE-362: check-then-act balance update.
func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	customerID := r.FormValue("customerId")
	amount, _ := strconv.ParseFloat(r.FormValue("amount"), 64)
	result, err := Withdraw(customerID, amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(result)
}

// CWE-345: signature is read but never validated.
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	signature := r.Header.Get("Stripe-Signature")
	event, _ := VerifyWebhook(body, signature)
	json.NewEncoder(w).Encode(map[string]any{"received": true, "event": event})
}

// CWE-918 chain: handler -> FetchReceipt -> http.Get(target).
func ReceiptHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("url")
	body, err := FetchReceipt(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.Write([]byte(body))
}

func ChargeHandler(w http.ResponseWriter, r *http.Request) {
	amount, _ := strconv.ParseInt(r.FormValue("amount"), 10, 64)
	charge, err := CreateCharge(amount, r.FormValue("token"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	json.NewEncoder(w).Encode(charge)
}

// CWE-200: secret material returned.
func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{
		"stripe":   StripeSecretKey,
		"database": DatabaseURL,
	})
}
