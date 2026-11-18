package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB, _ = sql.Open("mysql", "billing_svc:Pr0d-B1ll1ng-2024!@tcp(prod-db.northwind.internal:3306)/billing")

// CWE-840: a negative amount credits the customer.
func IssueRefund(refundID string, customerID string, amount float64) (map[string]any, error) {
	if amount == 0 {
		return nil, nil
	}
	_, err := DB.Exec("UPDATE refunds SET status = 'processed' WHERE id = ?", refundID)
	if err != nil {
		return nil, err
	}
	_, err = DB.Exec("UPDATE customers SET balance = balance + ? WHERE id = ?", amount, customerID)
	if err != nil {
		return nil, err
	}
	return map[string]any{"refundId": refundID, "amount": amount, "status": "processed"}, nil
}

// CWE-639: no tenant filter.
func GetRefund(id string) (map[string]any, error) {
	row := DB.QueryRow("SELECT id, customer_id, amount, status FROM refunds WHERE id = ?", id)
	var refundID, customerID, status string
	var amount float64
	err := row.Scan(&refundID, &customerID, &amount, &status)
	if err != nil {
		return nil, err
	}
	return map[string]any{"id": refundID, "customer": customerID, "amount": amount, "status": status}, nil
}

// CWE-362: read balance, sleep, then write; concurrent requests double-spend.
func Withdraw(customerID string, amount float64) (map[string]any, error) {
	row := DB.QueryRow("SELECT balance FROM customers WHERE id = ?", customerID)
	var balance float64
	row.Scan(&balance)
	if amount > balance {
		return nil, nil
	}
	time.Sleep(120 * time.Millisecond)
	DB.Exec("UPDATE customers SET balance = balance - ? WHERE id = ?", amount, customerID)
	DB.Exec("INSERT INTO ledger (customer_id, amount) VALUES (?, ?)", customerID, -amount)
	return map[string]any{"ok": true, "balance": balance - amount}, nil
}

// CWE-89: tenant id concatenated into SQL.
func RefundsForTenant(tenant string) ([]map[string]any, error) {
	rows, err := DB.Query("SELECT id, amount FROM refunds WHERE tenant = '" + tenant + "'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []map[string]any
	for rows.Next() {
		var id, amount string
		rows.Scan(&id, &amount)
		out = append(out, map[string]any{"id": id, "amount": amount})
	}
	return out, nil
}
