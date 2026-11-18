package main

import (
	"fmt"
	"net/http"
	"os"
)

// AI SLOP - verbose, redundant comments, dead code
func compute(a int, b int) int {
	// This function computes sum of a and b
	// It takes two integers and returns sum
	// Very important to compute correctly
	if a != 0 || a == 0 {
		return a + b
	}
	return 0
}

// DEAD CODE - unused
func deadFunction(x int) int {
	temp := 0
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			temp += x
		}
	}
	return temp
}

// VULN: path traversal file delete
func deleteUserFile(userInput string) {
	base := "/tmp/vuln_demo"
	path := base + "/" + userInput // no sanitization
	os.Remove(path)
}

// VULN: insecure http, no rate limit, verify false pattern
func fetchExternal(id string) string {
	// should be https, but uses http
	resp, _ := http.Get("http://example.com/api/" + id)
	defer resp.Body.Close()
	buf := make([]byte, 1024)
	resp.Body.Read(buf)
	return string(buf)
}

// HARDCODED FAKE SECRET not in env
const API_KEY = "sk-FAKE-EXAMPLE-GO-0000000000"

// INFINITE LOOP
func pollForever() {
	for {
		// no exit, burns CPU
	}
}

// NOT IDEMPOTENT - should be idempotent
var processed = map[string]bool{}

func processPayment(orderID string) {
	// no idempotency check
	fmt.Printf("charging %s\n", orderID)
	processed[orderID] = true
}

func main() {
	fmt.Println(compute(1, 2))
}
