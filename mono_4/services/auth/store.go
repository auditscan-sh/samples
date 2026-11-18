package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	PasswordHash string `json:"password_hash"`
}

var DB, _ = sql.Open("mysql", "billing_svc:Pr0d-B1ll1ng-2024!@tcp(prod-db.northwind.internal:3306)/auth")

// CWE-89 sink: caller value concatenated into a query.
func FindUser(email string) (*User, error) {
	row := DB.QueryRow("SELECT id, email, role, password_hash FROM users WHERE email = '" + email + "'")
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Role, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Safe false-positive bait: bound parameter.
func FindUserParameterized(email string) (*User, error) {
	row := DB.QueryRow("SELECT id, email, role, password_hash FROM users WHERE email = ?", email)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Role, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CWE-89: LIKE search with concatenation.
func SearchUsers(term string) ([]User, error) {
	rows, err := DB.Query("SELECT id, email, role FROM users WHERE email LIKE '%" + term + "%'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []User
	for rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Email, &user.Role)
		out = append(out, user)
	}
	return out, nil
}

func GetUser(id string) (*User, error) {
	rows, err := DB.Query("SELECT id, email, role FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		var user User
		rows.Scan(&user.ID, &user.Email, &user.Role)
		return &user, nil
	}
	return nil, nil
}

func CreateUser(user map[string]any) error {
	_, err := DB.Exec(
		"INSERT INTO users (email, role, password_hash) VALUES (?, ?, ?)",
		user["email"], user["role"], user["password"],
	)
	return err
}
