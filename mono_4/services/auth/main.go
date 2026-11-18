package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", LoginHandler)
	mux.HandleFunc("/register", RegisterHandler)
	mux.HandleFunc("/forgot", ForgotHandler)
	mux.HandleFunc("/verify", VerifyHandler)
	mux.HandleFunc("/users/search", SearchUsersHandler)
	mux.HandleFunc("/users", UserHandler)
	mux.HandleFunc("/ldap", LDAPHandler)
	mux.HandleFunc("/debug/jwt", DebugJWT)
	mux.HandleFunc("/reset", ResetHandler)
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// No TLS, no timeouts, no max header size.
	http.ListenAndServe("0.0.0.0:8080", mux)
}
