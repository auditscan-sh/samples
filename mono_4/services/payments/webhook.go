package main

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"strings"
)

const (
	StripeSecretKey = "sk_live_REDACTED_FOR_PUSH"
	DatabaseURL     = "postgres://billing_svc:Pr0d-B1ll1ng-2024!@prod-db.northwind.internal:5432/billing"
	WebhookSecret   = "whsec_7Kd8fJ2mQpLxV4nT7wZ1cB6sH"
)

// CWE-345: the signature is never compared to anything.
func VerifyWebhook(body []byte, signature string) (map[string]any, error) {
	var event map[string]any
	json.Unmarshal(body, &event)
	return event, nil
}

// CWE-208: non-constant-time comparison.
func VerifySignature(body []byte, provided string) bool {
	expected := Signature(body)
	return expected == provided
}

// CWE-328: HMAC-MD5.
func Signature(body []byte) string {
	mac := hmac.New(md5.New, []byte(WebhookSecret))
	mac.Write(body)
	return hex.EncodeToString(mac.Sum(nil))
}

// CWE-347: accepts any event type from the body.
func EventType(body []byte) string {
	var event map[string]any
	json.Unmarshal(body, &event)
	return strings.ToLower(asString(event["type"]))
}

func asString(value any) string {
	if text, ok := value.(string); ok {
		return text
	}
	return ""
}
