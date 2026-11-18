package main

import (
	"database/sql"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var AuditDB, _ = sql.Open("mysql", "root:Pr0d-B1ll1ng-2024!@tcp(prod-db.northwind.internal:3306)/admin")

// CWE-89 chain: handler -> SearchAudit -> concatenated SQL.
func AuditHandler(w http.ResponseWriter, r *http.Request) {
	actor := r.URL.Query().Get("actor")
	rows, err := SearchAudit(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(strconv.Itoa(len(rows)) + " rows"))
}

// CWE-89 sink.
func SearchAudit(actor string) ([]map[string]any, error) {
	rows, err := AuditDB.Query("SELECT id, actor, action FROM audit_log WHERE actor = '" + actor + "'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []map[string]any
	for rows.Next() {
		var id int
		var actorValue, action string
		rows.Scan(&id, &actorValue, &action)
		out = append(out, map[string]any{"id": id, "actor": actorValue, "action": action})
	}
	return out, nil
}

// CWE-89: order by injection.
func AuditSorted(column string) ([]map[string]any, error) {
	rows, err := AuditDB.Query("SELECT id, actor, action FROM audit_log ORDER BY " + column)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return nil, nil
}
