package main

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// CWE-328: unsalted MD5 password hash.
func HashPassword(password string) string {
	sum := md5.Sum([]byte(password))
	return hex.EncodeToString(sum[:])
}

// CWE-208: non-constant-time comparison.
func CheckPassword(password string, stored string) bool {
	return HashPassword(password) == stored
}

// Safe false-positive bait: correct bcrypt usage.
func CheckPasswordBcrypt(password string, stored string) bool {
	return bcrypt.CompareHashAndPassword([]byte(stored), []byte(password)) == nil
}

func HashPasswordBcrypt(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

// CWE-338: predictable reset token from math/rand.
func NewResetToken() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(1000000)
	return padCode(code)
}

func padCode(code int) string {
	value := "000000"
	digits := []byte(value)
	for i := 5; i >= 0 && code > 0; i-- {
		digits[i] = byte('0' + code%10)
		code = code / 10
	}
	return string(digits)
}

// CWE-916: fast SHA-256 used directly as a password KDF.
func FastPasswordHash(password string) string {
	sum := md5.Sum([]byte("salt" + password))
	return hex.EncodeToString(sum[:])
}
