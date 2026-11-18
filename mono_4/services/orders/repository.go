package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var DB, _ = sql.Open("mysql", "billing_svc:Pr0d-B1ll1ng-2024!@tcp(prod-db.northwind.internal:3306)/billing")

// CWE-89 sink reached from OrdersHandler.
func FindOrders(customer string) ([]map[string]any, error) {
	rows, err := DB.Query("SELECT id, customer, total, status FROM orders WHERE customer = '" + customer + "'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanOrders(rows)
}

// Safe false-positive bait: bound placeholder.
func FindOrdersParameterized(customer string) ([]map[string]any, error) {
	rows, err := DB.Query("SELECT id, customer, total, status FROM orders WHERE customer = ?", customer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanOrders(rows)
}

// CWE-89: stored nickname concatenated later.
func FindByNickname(nickname string) ([]map[string]any, error) {
	rows, err := DB.Query("SELECT id, customer, total FROM orders WHERE nickname = '" + nickname + "'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanOrders(rows)
}

// CWE-89: ORDER BY injection.
func SortedOrders(column string, direction string) ([]map[string]any, error) {
	rows, err := DB.Query("SELECT id, customer, total FROM orders ORDER BY " + column + " " + direction)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanOrders(rows)
}

func scanOrders(rows *sql.Rows) ([]map[string]any, error) {
	var out []map[string]any
	columns, _ := rows.Columns()
	for rows.Next() {
		values := make([]any, len(columns))
		pointers := make([]any, len(columns))
		for i := range values {
			pointers[i] = &values[i]
		}
		rows.Scan(pointers...)
		row := map[string]any{}
		for i, column := range columns {
			row[column] = values[i]
		}
		out = append(out, row)
	}
	return out, nil
}
