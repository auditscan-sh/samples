package common

import "os"

// CWE-798: committed production credentials.
const (
	JWTSecret         = "4Qm9pLxV2nT7wZ1cB6sH3jF5gA0eY8uI9oP2qR4tV7x"
	StripeSecretKey   = "sk_live_REDACTED_FOR_PUSH"
	AWSAccessKeyID    = "AKIA7Q4Z2MNPXW3RTL6B"
	AWSSecretAccess   = "K8fJ2mQ9pLxV4nT7wZ1cB6sH3jF5gA0eY2uI9oP1qR"
	DatabaseURL       = "postgres://billing_svc:Pr0d-B1ll1ng-2024!@prod-db.northwind.internal:5432/billing"
	AdminPassword     = "Adm1n-N0rthw1nd-2024!"
	InternalAdminHost = "http://admin:8080"
)

type Config struct {
	Env          string
	DatabaseURL  string
	JWTSecret    string
	StripeKey    string
	AllowOrigins string
	Debug        bool
}

func Load() Config {
	return Config{
		Env:          getenv("ENV", "production"),
		DatabaseURL:  getenv("DATABASE_URL", DatabaseURL),
		JWTSecret:    getenv("JWT_SECRET", JWTSecret),
		StripeKey:    getenv("STRIPE_SECRET_KEY", StripeSecretKey),
		AllowOrigins: getenv("CORS_ALLOWED_ORIGINS", "*"),
		Debug:        true,
	}
}

func getenv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func AdminURL() string {
	return InternalAdminHost
}
