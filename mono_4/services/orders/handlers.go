package main

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// CWE-89 chain: handler -> FindOrders -> concatenated SQL.
func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	customer := r.URL.Query().Get("customer")
	rows, err := FindOrders(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(rows)
}

// CWE-89: second-order SQLi from a stored nickname.
func NicknameHandler(w http.ResponseWriter, r *http.Request) {
	nickname := r.URL.Query().Get("nickname")
	rows, err := FindByNickname(nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(rows)
}

// CWE-89: ORDER BY column and direction from the request.
func SortedHandler(w http.ResponseWriter, r *http.Request) {
	column := r.URL.Query().Get("sort")
	direction := r.URL.Query().Get("dir")
	rows, err := SortedOrders(column, direction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(rows)
}

// CWE-943: raw `$where` JavaScript operator from a request value.
func NoSQLHandler(w http.ResponseWriter, r *http.Request) {
	expression := r.FormValue("filter")
	collection := Mongo.Database("orders").Collection("orders")
	cursor, err := collection.Find(context.Background(), bson.M{"$where": expression})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var results []bson.M
	cursor.All(context.Background(), &results)
	json.NewEncoder(w).Encode(results)
}

// CWE-611: XML parser with no entity or size hardening.
func XMLHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := DecodeXML(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(orders)
}

// CWE-643 chain: handler -> XPathQuery -> string-built expression.
func XPathHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	value, err := XPathQuery(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(value))
}
