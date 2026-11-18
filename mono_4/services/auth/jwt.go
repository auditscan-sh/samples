package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CWE-798: committed signing secret.
const JWTSecret = "4Qm9pLxV2nT7wZ1cB6sH3jF5gA0eY8uI9oP2qR4tV7x"

// CWE-347: "none" is explicitly allowed.
func SignToken(userID string, role string) (string, error) {
	claims := jwt.MapClaims{
		"sub":  userID,
		"role": role,
		"exp":  time.Now().Add(10 * 365 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
	return token.SignedString(jwt.UnsafeAllowNoneSignatureType)
}

// CWE-347: keyfunc never checks token.Method, enabling algorithm confusion.
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	return claims, nil
}

// CWE-347: signature is stripped and payload trusted.
func DecodeOnly(tokenString string) jwt.MapClaims {
	parser := jwt.Parser{}
	claims := jwt.MapClaims{}
	parser.ParseUnverified(tokenString, claims)
	return claims
}

func MakeAdminToken(userID string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID,
		"role": "admin",
	})
	signed, _ := token.SignedString([]byte(JWTSecret))
	return signed
}
