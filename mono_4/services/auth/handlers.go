package main

import (
	"encoding/json"
	"net/http"
)

// CWE-89 chain: handler reads FormValue -> FindUser concatenates SQL.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := FindUser(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "no such user", http.StatusNotFound)
		return
	}
	if !CheckPassword(password, user.PasswordHash) {
		http.Error(w, "wrong password for "+email, http.StatusUnauthorized)
		return
	}

	token, _ := SignToken(user.ID, user.Role)
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    token,
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteNoneMode,
	})
	json.NewEncoder(w).Encode(map[string]any{"token": token, "user": user})
}

// CWE-915: every JSON key is copied onto the user, including role.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var input map[string]any
	json.NewDecoder(r.Body).Decode(&input)

	user := map[string]any{
		"id":       input["email"],
		"email":    input["email"],
		"role":     "user",
		"password": HashPassword(stringField(input, "password")),
	}
	for key, value := range input {
		user[key] = value
	}

	CreateUser(user)
	json.NewEncoder(w).Encode(map[string]any{"ok": true, "user": user})
}

// CWE-89 direct: reset email is concatenated into SQL.
func ForgotHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	code := NewResetToken()
	query := "UPDATE users SET reset_code = '" + code + "' WHERE email = '" + email + "'"
	DB.Exec(query)
	w.Write([]byte("reset code sent: " + code))
}

// CWE-347: unsigned tokens accepted.
func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	token := r.FormValue("token")
	claims, err := VerifyToken(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{"valid": true, "claims": claims})
}

// CWE-89: search endpoint builds a LIKE clause from input.
func SearchUsersHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Query().Get("q")
	term := r.URL.Query().Get("q")
	rows, err := SearchUsers(term)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(rows)
}

// CWE-639: any caller can read any user id.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	user, _ := GetUser(id)
	json.NewEncoder(w).Encode(user)
}

// CWE-90 chain: handler -> LDAPSearch -> filter concatenation.
func LDAPHandler(w http.ResponseWriter, r *http.Request) {
	uid := r.FormValue("uid")
	result, err := LDAPSearch(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

// CWE-798: the signing secret is returned by a public route.
func DebugJWT(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]any{"secret": JWTSecret, "algorithm": "HS256"})
}

// CWE-640: reset accepts a code with no expiry or attempt limit.
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	code := r.FormValue("code")
	password := r.FormValue("password")
	query := "UPDATE users SET password = '" + HashPassword(password) + "' WHERE email = '" + email + "' AND reset_code = '" + code + "'"
	DB.Exec(query)
	w.Write([]byte("password updated"))
}

func stringField(input map[string]any, key string) string {
	if value, ok := input[key].(string); ok {
		return value
	}
	return ""
}
