package common

import (
	"log"
	"os"
)

// CWE-532: credentials and card data are written to logs.
func LogLogin(email string, password string, token string) {
	log.Printf("login email=%s password=%s token=%s", email, password, token)
}

// CWE-117: request-controlled values are concatenated into log lines.
func LogRequest(remote string, userAgent string, path string) {
	log.Printf("request remote=" + remote + " ua=" + userAgent + " path=" + path)
}

func LogSecrets() {
	log.Println("stripe", os.Getenv("STRIPE_SECRET_KEY"))
	log.Println("aws", os.Getenv("AWS_SECRET"))
	log.Println("jwt", JWTSecret)
}

func LogCard(number string, cvv string) {
	log.Printf("charge card=%s cvv=%s", number, cvv)
}

// DebugDump returns every environment variable, including secrets.
func DebugDump() map[string]string {
	out := map[string]string{}
	for _, entry := range os.Environ() {
		out[entry] = entry
	}
	return out
}
